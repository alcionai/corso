package repository

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/errs"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// BackupGetter deals with retrieving metadata about backups from the
// repository.
type BackupGetter interface {
	Backup(ctx context.Context, id string) (*backup.Backup, error)
	Backups(ctx context.Context, ids []string) ([]*backup.Backup, *fault.Bus)
	BackupsByTag(ctx context.Context, fs ...store.FilterOption) ([]*backup.Backup, error)
	GetBackupDetails(
		ctx context.Context,
		backupID string,
	) (*details.Details, *backup.Backup, *fault.Bus)
	GetBackupErrors(
		ctx context.Context,
		backupID string,
	) (*fault.Errors, *backup.Backup, *fault.Bus)
}

type Backuper interface {
	NewBackup(
		ctx context.Context,
		self selectors.Selector,
	) (operations.BackupOperation, error)
	NewBackupWithLookup(
		ctx context.Context,
		self selectors.Selector,
		ins idname.Cacher,
	) (operations.BackupOperation, error)
	DeleteBackups(
		ctx context.Context,
		failOnMissing bool,
		ids ...string,
	) error
}

// NewBackup generates a BackupOperation runner.
func (r repository) NewBackup(
	ctx context.Context,
	sel selectors.Selector,
) (operations.BackupOperation, error) {
	return r.NewBackupWithLookup(ctx, sel, nil)
}

// NewBackupWithLookup generates a BackupOperation runner.
// ownerIDToName and ownerNameToID are optional populations, in case the caller has
// already generated those values.
func (r repository) NewBackupWithLookup(
	ctx context.Context,
	sel selectors.Selector,
	ins idname.Cacher,
) (operations.BackupOperation, error) {
	err := r.ConnectDataProvider(ctx, sel.PathService())
	if err != nil {
		return operations.BackupOperation{}, clues.Wrap(err, "connecting to m365")
	}

	ownerID, ownerName, err := r.Provider.PopulateProtectedResourceIDAndName(ctx, sel.DiscreteOwner, ins)
	if err != nil {
		return operations.BackupOperation{}, clues.Wrap(err, "resolving resource owner details")
	}

	// TODO: retrieve display name from gc
	sel = sel.SetDiscreteOwnerIDName(ownerID, ownerName)

	return operations.NewBackupOperation(
		ctx,
		r.Opts,
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		r.Provider,
		r.Account,
		sel,
		sel, // the selector acts as an IDNamer for its discrete resource owner.
		r.Bus)
}

// Backup retrieves a backup by id.
func (r repository) Backup(ctx context.Context, id string) (*backup.Backup, error) {
	return getBackup(ctx, id, store.NewWrapper(r.modelStore))
}

// getBackup handles the processing for Backup.
func getBackup(
	ctx context.Context,
	id string,
	sw store.BackupGetter,
) (*backup.Backup, error) {
	b, err := sw.GetBackup(ctx, model.StableID(id))
	if err != nil {
		return nil, errWrapper(err)
	}

	return b, nil
}

// Backups lists backups by ID. Returns as many backups as possible with
// errors for the backups it was unable to retrieve.
func (r repository) Backups(ctx context.Context, ids []string) ([]*backup.Backup, *fault.Bus) {
	var (
		bups []*backup.Backup
		bus  = fault.New(false)
		sw   = store.NewWrapper(r.modelStore)
	)

	for _, id := range ids {
		ictx := clues.Add(ctx, "backup_id", id)

		b, err := sw.GetBackup(ictx, model.StableID(id))
		if err != nil {
			bus.AddRecoverable(ctx, errWrapper(err))
		}

		bups = append(bups, b)
	}

	return bups, bus
}

func addBackup(
	bup *lineageNode,
	seen map[model.StableID]*lineageNode,
	allNodes map[model.StableID]*lineageNode,
) {
	if bup == nil {
		return
	}

	if _, ok := seen[bup.ID]; ok {
		// We've already traversed this node.
		return
	}

	for baseID := range bup.MergeBases {
		addBackup(allNodes[baseID], seen, allNodes)
	}

	for baseID := range bup.AssistBases {
		addBackup(allNodes[baseID], seen, allNodes)
	}

	seen[bup.ID] = bup

	for _, descendent := range bup.children {
		addBackup(allNodes[descendent.ID], seen, allNodes)
	}
}

func filterLineages(
	bups map[model.StableID]*lineageNode,
	backupIDs ...string,
) map[model.StableID]*lineageNode {
	if len(backupIDs) == 0 {
		return bups
	}

	res := map[model.StableID]*lineageNode{}

	// For each backup we're interested in, traverse up and down the hierarchy.
	// Going down the hierarchy is more difficult because backups only have
	// backpointers to their ancestors.
	for _, id := range backupIDs {
		addBackup(bups[model.StableID(id)], res, bups)
	}

	return res
}

func addBase(
	baseID model.StableID,
	baseReasons []identity.Reasoner,
	current *BackupNode,
	allNodes map[model.StableID]*BackupNode,
	bups map[model.StableID]*lineageNode,
) {
	parent, parentOK := allNodes[baseID]
	if !parentOK {
		parent = &BackupNode{}
		allNodes[baseID] = parent
	}

	parent.Label = string(baseID)

	// If the parent isn't in the set of backups passed in it must have been
	// deleted.
	if p, ok := bups[baseID]; !ok || p.deleted {
		parent.Deleted = true

		// If the backup was deleted we should also attempt to recreate the
		// set of Reasons which it encompassed. We can get partial info on this
		// by collecting all the Reasons it was a base.
		for _, reason := range baseReasons {
			if !slices.ContainsFunc(
				parent.Reasons,
				func(other identity.Reasoner) bool {
					return other.Service() == reason.Service() &&
						other.Category() == reason.Category()
				}) {
				parent.Reasons = append(parent.Reasons, reason)
			}
		}
	}

	parent.Children = append(
		parent.Children,
		&BackupEdge{
			BackupNode: current,
			Reasons:    baseReasons,
		})
}

func buildOutput(bups map[model.StableID]*lineageNode) ([]*BackupNode, error) {
	var roots []*BackupNode

	allNodes := map[model.StableID]*BackupNode{}

	for _, bup := range bups {
		node := allNodes[bup.ID]
		if node == nil {
			node = &BackupNode{}
			allNodes[bup.ID] = node
		}

		node.Label = string(bup.ID)
		node.Type = MergeNode
		node.Created = bup.CreationTime

		if bup.Tags[model.BackupTypeTag] == model.AssistBackup {
			node.Type = AssistNode
		}

		topLevel := true

		if !bup.deleted {
			reasons, err := bup.Reasons()
			if err != nil {
				return nil, clues.Wrap(err, "getting reasons").With("backup_id", bup.ID)
			}

			node.Reasons = reasons

			bases, err := bup.Bases()
			if err != nil {
				return nil, clues.Wrap(err, "getting bases").With("backup_id", bup.ID)
			}

			for baseID, baseReasons := range bases.Merge {
				topLevel = false

				addBase(baseID, baseReasons, node, allNodes, bups)
			}

			for baseID, baseReasons := range bases.Assist {
				topLevel = false

				addBase(baseID, baseReasons, node, allNodes, bups)
			}
		}

		// If this node has no ancestors then add it directly to the root.
		if bup.deleted || topLevel {
			roots = append(roots, node)
		}
	}

	return roots, nil
}

// lineageNode is a small in-memory wrapper around *backup.Backup that provides
// information about children. This just makes it easier to traverse lineages
// during filtering.
type lineageNode struct {
	*backup.Backup
	children []*lineageNode
	deleted  bool
}

func (r repository) BackupLineage(
	ctx context.Context,
	tenantID string,
	protectedResourceID string,
	service path.ServiceType,
	category path.CategoryType,
	backupIDs ...string,
) ([]*BackupNode, error) {
	sw := store.NewWrapper(r.modelStore)

	fs := []store.FilterOption{
		store.Tenant(tenantID),
		//store.ProtectedResource(protectedResourceID),
		//store.Reason(service, category),
	}

	bs, err := sw.GetBackups(ctx, fs...)
	if err != nil {
		return nil, clues.Stack(err)
	}

	if len(bs) == 0 {
		return nil, clues.Stack(errs.NotFound)
	}

	// Put all the backups in a map so we can access them easier when building the
	// graph.
	bups := make(map[model.StableID]*lineageNode, len(bs))

	for _, b := range bs {
		current := bups[b.ID]

		if current == nil {
			current = &lineageNode{}
		}

		current.Backup = b
		current.deleted = false
		bups[b.ID] = current

		for id := range b.MergeBases {
			parent := bups[id]
			if parent == nil {
				// Populate the ID so we don't NPE on it when building the tree if
				// something was deleted.
				parent = &lineageNode{
					Backup: &backup.Backup{
						BaseModel: model.BaseModel{
							ID: id,
						},
					},
					deleted: true,
				}
			}

			parent.children = append(parent.children, current)
			bups[id] = parent
		}

		for id := range b.AssistBases {
			parent := bups[id]
			if parent == nil {
				// Populate the ID so we don't NPE on it when building the tree if
				// something was deleted.
				parent = &lineageNode{
					Backup: &backup.Backup{
						BaseModel: model.BaseModel{
							ID: id,
						},
					},
					deleted: true,
				}
			}

			parent.children = append(parent.children, current)
			bups[id] = parent
		}
	}

	// Filter the map of backups to just those in the lineages we're interested
	// about.
	filtered := filterLineages(bups, backupIDs...)

	// Build the output graph.
	res, err := buildOutput(filtered)

	return res, clues.Stack(err).OrNil()
}

// BackupsByTag lists all backups in a repository that contain all the tags
// specified.
func (r repository) BackupsByTag(ctx context.Context, fs ...store.FilterOption) ([]*backup.Backup, error) {
	sw := store.NewWrapper(r.modelStore)
	return backupsByTag(ctx, sw, fs)
}

// backupsByTag returns all backups matching all provided tags.
//
// TODO(ashmrtn): This exists mostly for testing, but we could restructure the
// code in this file so there's a more elegant mocking solution.
func backupsByTag(
	ctx context.Context,
	sw store.BackupWrapper,
	fs []store.FilterOption,
) ([]*backup.Backup, error) {
	bs, err := sw.GetBackups(ctx, fs...)
	if err != nil {
		return nil, clues.Stack(err)
	}

	// Filter out assist backup bases as they're considered incomplete and we
	// haven't been displaying them before now.
	res := make([]*backup.Backup, 0, len(bs))

	for _, b := range bs {
		if t := b.Tags[model.BackupTypeTag]; t != model.AssistBackup {
			res = append(res, b)
		}
	}

	return res, nil
}

// BackupDetails returns the specified backup.Details
func (r repository) GetBackupDetails(
	ctx context.Context,
	backupID string,
) (*details.Details, *backup.Backup, *fault.Bus) {
	bus := fault.New(false)

	deets, bup, err := getBackupDetails(
		ctx,
		backupID,
		r.Account.ID(),
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		bus)

	return deets, bup, bus.Fail(err)
}

// getBackupDetails handles the processing for GetBackupDetails.
func getBackupDetails(
	ctx context.Context,
	backupID, tenantID string,
	kw *kopia.Wrapper,
	sw store.BackupGetter,
	bus *fault.Bus,
) (*details.Details, *backup.Backup, error) {
	b, err := sw.GetBackup(ctx, model.StableID(backupID))
	if err != nil {
		return nil, nil, errWrapper(err)
	}

	ssid := b.StreamStoreID
	if len(ssid) == 0 {
		ssid = b.DetailsID
	}

	if len(ssid) == 0 {
		return nil, b, clues.New("no streamstore id in backup").WithClues(ctx)
	}

	var (
		sstore = streamstore.NewStreamer(kw, tenantID, b.Selector.PathService())
		deets  details.Details
	)

	err = sstore.Read(
		ctx,
		ssid,
		streamstore.DetailsReader(details.UnmarshalTo(&deets)),
		bus)
	if err != nil {
		return nil, nil, err
	}

	// Retroactively fill in isMeta information for items in older
	// backup versions without that info
	// version.Restore2 introduces the IsMeta flag, so only v1 needs a check.
	if b.Version >= version.OneDrive1DataAndMetaFiles && b.Version < version.OneDrive3IsMetaMarker {
		for _, d := range deets.Entries {
			if d.OneDrive != nil {
				d.OneDrive.IsMeta = metadata.HasMetaSuffix(d.RepoRef)
			}
		}
	}

	deets.DetailsModel = deets.FilterMetaFiles()

	return &deets, b, nil
}

// BackupErrors returns the specified backup's fault.Errors
func (r repository) GetBackupErrors(
	ctx context.Context,
	backupID string,
) (*fault.Errors, *backup.Backup, *fault.Bus) {
	bus := fault.New(false)

	fe, bup, err := getBackupErrors(
		ctx,
		backupID,
		r.Account.ID(),
		r.dataLayer,
		store.NewWrapper(r.modelStore),
		bus)

	return fe, bup, bus.Fail(err)
}

// getBackupErrors handles the processing for GetBackupErrors.
func getBackupErrors(
	ctx context.Context,
	backupID, tenantID string,
	kw *kopia.Wrapper,
	sw store.BackupGetter,
	bus *fault.Bus,
) (*fault.Errors, *backup.Backup, error) {
	b, err := sw.GetBackup(ctx, model.StableID(backupID))
	if err != nil {
		return nil, nil, errWrapper(err)
	}

	ssid := b.StreamStoreID
	if len(ssid) == 0 {
		return nil, b, clues.New("missing streamstore id in backup").WithClues(ctx)
	}

	var (
		sstore = streamstore.NewStreamer(kw, tenantID, b.Selector.PathService())
		fe     fault.Errors
	)

	err = sstore.Read(
		ctx,
		ssid,
		streamstore.FaultErrorsReader(fault.UnmarshalErrorsTo(&fe)),
		bus)
	if err != nil {
		return nil, nil, err
	}

	return &fe, b, nil
}

// DeleteBackups removes the backups from both the model store and the backup
// storage.
//
// If failOnMissing is true then returns an error if a backup model can't be
// found. Otherwise ignores missing backup models.
//
// Missing models or snapshots during the actual deletion do not cause errors.
//
// All backups are delete as an atomic unit so any failures will result in no
// deletions.
func (r repository) DeleteBackups(
	ctx context.Context,
	failOnMissing bool,
	ids ...string,
) error {
	return deleteBackups(ctx, store.NewWrapper(r.modelStore), failOnMissing, ids...)
}

// deleteBackup handles the processing for backup deletion.
func deleteBackups(
	ctx context.Context,
	sw store.BackupGetterModelDeleter,
	failOnMissing bool,
	ids ...string,
) error {
	// Although we haven't explicitly stated it, snapshots are technically
	// manifests in kopia. This means we can use the same delete API to remove
	// them and backup models. Deleting all of them together gives us both
	// atomicity guarantees (around when data will be flushed) and helps reduce
	// the number of manifest blobs that kopia will create.
	var toDelete []manifest.ID

	for _, id := range ids {
		b, err := sw.GetBackup(ctx, model.StableID(id))
		if err != nil {
			if !failOnMissing && errors.Is(err, errs.NotFound) {
				continue
			}

			return clues.Stack(errWrapper(err)).
				WithClues(ctx).
				With("delete_backup_id", id)
		}

		toDelete = append(toDelete, b.ModelStoreID)

		if len(b.SnapshotID) > 0 {
			toDelete = append(toDelete, manifest.ID(b.SnapshotID))
		}

		ssid := b.StreamStoreID
		if len(ssid) == 0 {
			ssid = b.DetailsID
		}

		if len(ssid) > 0 {
			toDelete = append(toDelete, manifest.ID(ssid))
		}
	}

	return sw.DeleteWithModelStoreIDs(ctx, toDelete...)
}
