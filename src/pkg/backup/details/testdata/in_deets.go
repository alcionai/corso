package testdata

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
//	location set handling
// ---------------------------------------------------------------------------

var exists = struct{}{}

type locSet struct {
	// map [locationRef] map [itemRef] {}
	// refs may be either the canonical ent refs, or something else,
	// so long as they are consistent for the test in question
	Locations map[string]map[string]struct{}
	Deleted   map[string]map[string]struct{}
}

func newLocSet() *locSet {
	return &locSet{
		Locations: map[string]map[string]struct{}{},
		Deleted:   map[string]map[string]struct{}{},
	}
}

func (ls *locSet) AddItem(locationRef, itemRef string) {
	ls.AddLocation(locationRef)

	ls.Locations[locationRef][itemRef] = exists
	delete(ls.Deleted[locationRef], itemRef)
}

func (ls *locSet) RemoveItem(locationRef, itemRef string) {
	delete(ls.Locations[locationRef], itemRef)

	if _, ok := ls.Deleted[locationRef]; !ok {
		ls.Deleted[locationRef] = map[string]struct{}{}
	}

	ls.Deleted[locationRef][itemRef] = exists
}

func (ls *locSet) MoveItem(fromLocation, toLocation, ir string) {
	ls.RemoveItem(fromLocation, ir)
	ls.AddItem(toLocation, ir)
}

func (ls *locSet) AddLocation(locationRef string) {
	if _, ok := ls.Locations[locationRef]; !ok {
		ls.Locations[locationRef] = map[string]struct{}{}
	}
	// don't purge previously deleted items, or child locations.
	// Assumption is that their itemRef is unique, and still deleted.
	delete(ls.Deleted, locationRef)
}

func (ls *locSet) RemoveLocation(locationRef string) {
	ss := ls.Subset(locationRef)

	for lr := range ss.Locations {
		items := ls.Locations[lr]

		delete(ls.Locations, lr)

		if _, ok := ls.Deleted[lr]; !ok {
			ls.Deleted[lr] = map[string]struct{}{}
		}

		for ir := range items {
			ls.Deleted[lr][ir] = exists
		}
	}
}

// MoveLocation takes the LAST elemet in the fromLocation (and all)
// children matching the prefix, and relocates it as a child of toLocation.
// ex: MoveLocation("/a/b/c", "/d") will move all entries with the prefix
// "/a/b/c" into "/d/c".  This also deletes all "/a/b/c" entries and children.
// assumes item IDs don't change across the migration.  If item IDs do change,
// that difference will need to be handled manually by the caller.
// returns the base folder's new location (ex: /d/c)
func (ls *locSet) MoveLocation(fromLocation, toLocation string) string {
	fromBuilder := path.Builder{}.Append(path.Split(fromLocation)...)
	toBuilder := path.Builder{}.Append(path.Split(toLocation)...).Append(fromBuilder.LastElem())

	ls.RenameLocation(fromBuilder.String(), toBuilder.String())

	return toBuilder.String()
}

func (ls *locSet) RenameLocation(fromLocation, toLocation string) {
	ss := ls.Subset(fromLocation)
	fromBuilder := path.Builder{}.Append(path.Split(fromLocation)...)
	toBuilder := path.Builder{}.Append(path.Split(toLocation)...)

	for lr, items := range ss.Locations {
		lrBuilder := path.Builder{}.Append(path.Split(lr)...)
		lrBuilder.UpdateParent(fromBuilder, toBuilder)

		newLoc := lrBuilder.String()

		for ir := range items {
			ls.RemoveItem(lr, ir)
			ls.AddItem(newLoc, ir)
		}

		ls.RemoveLocation(lr)
		ls.AddLocation(newLoc)
	}
}

// Subset produces a new locSet containing only Items and Locations
// whose location matches the locationPfx
func (ls *locSet) Subset(locationPfx string) *locSet {
	ss := newLocSet()

	for lr, items := range ls.Locations {
		if strings.HasPrefix(lr, locationPfx) {
			ss.AddLocation(lr)

			for ir := range items {
				ss.AddItem(lr, ir)
			}
		}
	}

	return ss
}

// ---------------------------------------------------------------------------
// The goal of InDeets is to provide a struct and interface which allows
// tests to predict not just the elements within a set of details entries,
// but also their changes (relocation, renaming, etc) in a way that consolidates
// building an "expected set" of details entries that can be compared against
// the details results after a backup.
// ---------------------------------------------------------------------------

// InDeets is a helper for comparing details state in tests
// across backup instances.
type InDeets struct {
	// only: tenantID/service/resourceOwnerID
	RRPrefix string
	// map of container setting the uniqueness boundary for location
	// ref entries (eg, data type like email, contacts, etc, or
	// drive id) to the unique entries in that set.
	Sets map[string]*locSet
}

func NewInDeets(repoRefPrefix string) *InDeets {
	return &InDeets{
		RRPrefix: repoRefPrefix,
		Sets:     map[string]*locSet{},
	}
}

func (id *InDeets) getSet(set string) *locSet {
	s, ok := id.Sets[set]
	if ok {
		return s
	}

	return newLocSet()
}

func (id *InDeets) AddAll(deets details.Details, ws whatSet) {
	if id.Sets == nil {
		id.Sets = map[string]*locSet{}
	}

	for _, ent := range deets.Entries {
		set, err := ws(ent)
		if err != nil {
			set = err.Error()
		}

		dir := ent.LocationRef

		if ent.Folder != nil {
			dir = dir + ent.Folder.DisplayName
			id.AddLocation(set, dir)
		} else {
			id.AddItem(set, ent.LocationRef, ent.ItemRef)
		}
	}
}

func (id *InDeets) AddItem(set, locationRef, itemRef string) {
	id.getSet(set).AddItem(locationRef, itemRef)
}

func (id *InDeets) RemoveItem(set, locationRef, itemRef string) {
	id.getSet(set).RemoveItem(locationRef, itemRef)
}

func (id *InDeets) MoveItem(set, fromLocation, toLocation, ir string) {
	id.getSet(set).MoveItem(fromLocation, toLocation, ir)
}

func (id *InDeets) AddLocation(set, locationRef string) {
	id.getSet(set).AddLocation(locationRef)
}

// RemoveLocation removes the provided location, and all children
// of that location.
func (id *InDeets) RemoveLocation(set, locationRef string) {
	id.getSet(set).RemoveLocation(locationRef)
}

// MoveLocation takes the LAST elemet in the fromLocation (and all)
// children matching the prefix, and relocates it as a child of toLocation.
// ex: MoveLocation("/a/b/c", "/d") will move all entries with the prefix
// "/a/b/c" into "/d/c".  This also deletes all "/a/b/c" entries and children.
// assumes item IDs don't change across the migration.  If item IDs do change,
// that difference will need to be handled manually by the caller.
// returns the base folder's new location (ex: /d/c)
func (id *InDeets) MoveLocation(set, fromLocation, toLocation string) string {
	return id.getSet(set).MoveLocation(fromLocation, toLocation)
}

func (id *InDeets) RenameLocation(set, fromLocation, toLocation string) {
	id.getSet(set).RenameLocation(fromLocation, toLocation)
}

// Subset produces a new locSet containing only Items and Locations
// whose location matches the locationPfx
func (id *InDeets) Subset(set, locationPfx string) *locSet {
	return id.getSet(set).Subset(locationPfx)
}

// ---------------------------------------------------------------------------
// whatSet helpers for extracting a set identifier from an arbitrary repoRef
// ---------------------------------------------------------------------------

type whatSet func(details.Entry) (string, error)

// common whatSet parser that extracts the service category from
// a repoRef.
func CategoryFromRepoRef(ent details.Entry) (string, error) {
	p, err := path.FromDataLayerPath(ent.RepoRef, false)
	if err != nil {
		return "", err
	}

	return p.Category().String(), nil
}

// common whatSet parser that extracts the driveID from a repoRef.
func DriveIDFromRepoRef(ent details.Entry) (string, error) {
	p, err := path.FromDataLayerPath(ent.RepoRef, false)
	if err != nil {
		return "", err
	}

	odp, err := path.ToDrivePath(p)
	if err != nil {
		return "", err
	}

	return odp.DriveID, nil
}

// ---------------------------------------------------------------------------
// helpers and comparators
// ---------------------------------------------------------------------------

func CheckBackupDetails(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	backupID model.StableID,
	ws whatSet,
	ms *kopia.ModelStore,
	ssr streamstore.Reader,
	expect *InDeets,
	// standard check is assert.Subset due to issues of external data cross-
	// pollination.  This should be true if the backup contains a unique directory
	// of data.
	mustEqualFolders bool,
) {
	deets, result := GetDeetsInBackup(t, ctx, backupID, "", "", path.UnknownService, ws, ms, ssr)

	t.Log("details entries in result")

	for _, ent := range deets.Entries {
		if ent.Folder == nil {
			t.Log(ent.LocationRef)
			t.Log(ent.ItemRef)
		}

		assert.Truef(
			t,
			strings.HasPrefix(ent.RepoRef, expect.RRPrefix),
			"all details should begin with the expected prefix\nwant: %s\ngot:  %s",
			expect.RRPrefix, ent.RepoRef)
	}

	for set := range expect.Sets {
		check := assert.Subsetf

		if mustEqualFolders {
			check = assert.ElementsMatchf
		}

		check(
			t,
			maps.Keys(result.Sets[set].Locations),
			maps.Keys(expect.Sets[set].Locations),
			"results in %s missing expected location", set)

		for lr, items := range expect.Sets[set].Deleted {
			_, ok := result.Sets[set].Locations[lr]
			assert.Falsef(t, ok, "deleted location in %s found in result: %s", set, lr)

			for ir := range items {
				_, ok := result.Sets[set].Locations[lr][ir]
				assert.Falsef(t, ok, "deleted item in %s found in result: %s", set, lr)
			}
		}
	}
}

func GetDeetsInBackup(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	backupID model.StableID,
	tid, resourceOwner string,
	service path.ServiceType,
	ws whatSet,
	ms *kopia.ModelStore,
	ssr streamstore.Reader,
) (details.Details, *InDeets) {
	bup := backup.Backup{}

	err := ms.Get(ctx, model.BackupSchema, backupID, &bup)
	require.NoError(t, err, clues.ToCore(err))

	ssid := bup.StreamStoreID
	require.NotEmpty(t, ssid, "stream store ID")

	var deets details.Details
	err = ssr.Read(
		ctx,
		ssid,
		streamstore.DetailsReader(details.UnmarshalTo(&deets)),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	id := NewInDeets(path.Builder{}.Append(tid, service.String(), resourceOwner).String())
	id.AddAll(deets, ws)

	return deets, id
}
