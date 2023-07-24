package m365

import (
	"context"
	"runtime/trace"
	"sync"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// must comply with BackupProducer and RestoreConsumer
var (
	_ inject.BackupProducer  = &Controller{}
	_ inject.RestoreConsumer = &Controller{}
)

// Controller is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type Controller struct {
	AC api.Client

	tenant      string
	credentials account.M365Config

	ownerLookup getOwnerIDAndNamer
	// maps of resource owner ids to names, and names to ids.
	// not guaranteed to be populated, only here as a post-population
	// reference for processes that choose to populate the values.
	IDNameLookup idname.Cacher

	// wg is used to track completion of tasks
	wg     *sync.WaitGroup
	region *trace.Region

	// mutex used to synchronize updates to `status`
	mu     sync.Mutex
	status support.ControllerOperationStatus // contains the status of the last run status

	// backupDriveIDNames is populated on restore.  It maps the backup's
	// drive names to their id. Primarily for use when creating or looking
	// up a new drive.
	backupDriveIDNames idname.CacheBuilder
}

func NewController(
	ctx context.Context,
	acct account.Account,
	rc resource.Category,
	pst path.ServiceType,
	co control.Options,
) (*Controller, error) {
	graph.InitializeConcurrencyLimiter(ctx, pst == path.ExchangeService, co.Parallelism.ItemFetch)

	creds, err := acct.M365Config()
	if err != nil {
		return nil, clues.Wrap(err, "retrieving m365 account configuration").WithClues(ctx)
	}

	ac, err := api.NewClient(creds)
	if err != nil {
		return nil, clues.Wrap(err, "creating api client").WithClues(ctx)
	}

	rCli, err := getResourceClient(rc, ac)
	if err != nil {
		return nil, clues.Wrap(err, "creating resource client").WithClues(ctx)
	}

	ctrl := Controller{
		AC:           ac,
		IDNameLookup: idname.NewCache(nil),

		credentials:        creds,
		ownerLookup:        rCli,
		tenant:             acct.ID(),
		wg:                 &sync.WaitGroup{},
		backupDriveIDNames: idname.NewCache(nil),
	}

	return &ctrl, nil
}

// ---------------------------------------------------------------------------
// Processing Status
// ---------------------------------------------------------------------------

// AwaitStatus waits for all tasks to complete and then returns status
func (ctrl *Controller) Wait() *data.CollectionStats {
	defer func() {
		if ctrl.region != nil {
			ctrl.region.End()
			ctrl.region = nil
		}
	}()
	ctrl.wg.Wait()

	// clean up and reset statefulness
	dcs := data.CollectionStats{
		Folders:   ctrl.status.Folders,
		Objects:   ctrl.status.Metrics.Objects,
		Successes: ctrl.status.Metrics.Successes,
		Bytes:     ctrl.status.Metrics.Bytes,
		Details:   ctrl.status.String(),
	}

	ctrl.wg = &sync.WaitGroup{}
	ctrl.status = support.ControllerOperationStatus{}

	return &dcs
}

// UpdateStatus is used by initiated tasks to indicate completion
func (ctrl *Controller) UpdateStatus(status *support.ControllerOperationStatus) {
	defer ctrl.wg.Done()

	if status == nil {
		return
	}

	ctrl.mu.Lock()
	defer ctrl.mu.Unlock()
	ctrl.status = support.MergeStatus(ctrl.status, *status)
}

// Status returns the current status of the controller process.
func (ctrl *Controller) Status() support.ControllerOperationStatus {
	return ctrl.status
}

// PrintableStatus returns a string formatted version of the status.
func (ctrl *Controller) PrintableStatus() string {
	return ctrl.status.String()
}

func (ctrl *Controller) incrementAwaitingMessages() {
	ctrl.wg.Add(1)
}

func (ctrl *Controller) CacheItemInfo(dii details.ItemInfo) {
	if dii.SharePoint != nil {
		ctrl.backupDriveIDNames.Add(dii.SharePoint.DriveID, dii.SharePoint.DriveName)
	}

	if dii.OneDrive != nil {
		ctrl.backupDriveIDNames.Add(dii.OneDrive.DriveID, dii.OneDrive.DriveName)
	}
}

// ---------------------------------------------------------------------------
// Resource Lookup Handling
// ---------------------------------------------------------------------------

func getResourceClient(rc resource.Category, ac api.Client) (*resourceClient, error) {
	switch rc {
	case resource.Users:
		return &resourceClient{enum: rc, getter: ac.Users()}, nil
	case resource.Sites:
		return &resourceClient{enum: rc, getter: ac.Sites()}, nil
	default:
		return nil, clues.New("unrecognized owner resource enum").With("resource_enum", rc)
	}
}

type resourceClient struct {
	enum   resource.Category
	getter getIDAndNamer
}

type getIDAndNamer interface {
	GetIDAndName(ctx context.Context, owner string) (
		ownerID string,
		ownerName string,
		err error,
	)
}

var _ getOwnerIDAndNamer = &resourceClient{}

type getOwnerIDAndNamer interface {
	getOwnerIDAndNameFrom(
		ctx context.Context,
		discovery api.Client,
		owner string,
		ins idname.Cacher,
	) (
		ownerID string,
		ownerName string,
		err error,
	)
}

// getOwnerIDAndNameFrom looks up the owner's canonical id and display name.
// If the owner is present in the idNameSwapper, then that interface's id and
// name values are returned.  As a fallback, the resource calls the discovery
// api to fetch the user or site using the owner value. This fallback assumes
// that the owner is a well formed ID or display name of appropriate design
// (PrincipalName for users, WebURL for sites).
func (r resourceClient) getOwnerIDAndNameFrom(
	ctx context.Context,
	discovery api.Client,
	owner string,
	ins idname.Cacher,
) (string, string, error) {
	if ins != nil {
		if n, ok := ins.NameOf(owner); ok {
			return owner, n, nil
		} else if i, ok := ins.IDOf(owner); ok {
			return i, owner, nil
		}
	}

	ctx = clues.Add(ctx, "owner_identifier", owner)

	var (
		id, name string
		err      error
	)

	id, name, err = r.getter.GetIDAndName(ctx, owner)
	if err != nil {
		if graph.IsErrUserNotFound(err) {
			return "", "", clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		return "", "", err
	}

	if len(id) == 0 || len(name) == 0 {
		return "", "", clues.Stack(graph.ErrResourceOwnerNotFound)
	}

	return id, name, nil
}

// PopulateProtectedResourceIDAndName takes the provided owner identifier and produces
// the owner's name and ID from that value.  Returns an error if the owner is
// not recognized by the current tenant.
//
// The id-name cacher is optional.  Some processes will look up all owners in
// the tenant before reaching this step.  In that case, the data gets handed
// down for this func to consume instead of performing further queries.  The
// data gets stored inside the controller instance for later re-use.
func (ctrl *Controller) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	owner string, // input value, can be either id or name
	ins idname.Cacher,
) (string, string, error) {
	id, name, err := ctrl.ownerLookup.getOwnerIDAndNameFrom(ctx, ctrl.AC, owner, ins)
	if err != nil {
		return "", "", clues.Wrap(err, "identifying resource owner")
	}

	ctrl.IDNameLookup = idname.NewCache(map[string]string{id: name})

	return id, name, nil
}
