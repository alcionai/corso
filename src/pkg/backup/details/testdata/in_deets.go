package testdata

import (
	"context"
	"fmt"
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
)

// InDeets is a helper for comparing details state in tests
// across backup instances.
type InDeets struct {
	Entries map[string]details.DetailsEntry
	Deleted map[string]struct{}
}

func NewInDeets() *InDeets {
	id := &InDeets{}
	id.init()

	return id
}

const clip = 82

func (id *InDeets) init() {
	id.Entries = map[string]details.DetailsEntry{}
	id.Deleted = map[string]struct{}{}
}

func (id *InDeets) Add(deets details.Details) {
	if id.Entries == nil {
		id.init()
	}

	for _, ent := range deets.Entries {
		id.Entries[ent.RepoRef] = ent
	}
}

func (id *InDeets) AddEntry(rr string) {
	fmt.Printf("-----\nadding %+v\n", rr[clip:])
	addEnt(id, rr, details.DetailsEntry{RepoRef: rr})
}

func addent(id *InDeets, rr string, ent details.DetailsEntry) {
	id.Entries[rr] = ent
	delete(id.Deleted, rr)
}

func (id *InDeets) DeleteEntry(rr string) {
	fmt.Printf("-----\ndeleting %+v\n", rr[clip:])
	delete(id.Entries, rr)
	id.Deleted[rr] = struct{}{}
}

func (id *InDeets) DeletePrefix(rr string) {
	fmt.Printf("-----\ndelpfx %+v\n", rr[clip:])
	delete(id.Entries, rr)
	id.Deleted[rr] = struct{}{}

	for k := range id.Entries {
		if strings.HasPrefix(k, rr) {
			id.DeleteEntry(k)
		}
	}
}

func (id *InDeets) Subset(prefixRR string) *InDeets {
	nd := NewInDeets()

	for rr, ent := range id.Entries {
		if strings.HasPrefix(rr, prefixRR) {
			nd.Entries[rr] = ent
		}
	}

	for rr, ent := range id.Deleted {
		if strings.HasPrefix(rr, prefixRR) {
			nd.Deleted[rr] = ent
		}
	}

	return nd
}

// QoL func that makes a couple assumptions about how item migration
// works in details.  Spedifically that movement is the act of trimming
// the old prefix, and appending the remaining suffix to the new prefix.
// eg: from = pfx/foo, to = newpfx/bar, result = newpfx/bar/foo
// This implies that item IDs remain constant across moves, which may
// not be guaranteed by the underlying service, so use this helper
// with a measure of caution.
func (id *InDeets) MovePrefix(t *testing.T, fromRR, toRR string) {
	fmt.Printf("-----\nmove from %+v\n", fromRR[clip:])
	ss := id.Subset(fromRR)

	for rr := range ss.Entries {
		rr = strings.TrimPrefix(rr, fromRR)
		if len(rr) == 0 {
			continue
		}

		rr = strings.TrimSuffix(toRR, "/") + "/" + strings.TrimPrefix(rr, "/")
		id.AddEntry(rr)
	}

	id.DeletePrefix(fromRR)
	id.AddEntry(toRR)
}

// ---------------------------------------------------------------------------
//
// ---------------------------------------------------------------------------

// // InDeets is a helper for comparing details state in tests
// // across backup instances.
// type InDeets struct {
// 	Items          map[string]details.DetailsEntry
// 	Folders        map[string]details.DetailsEntry
// 	DeletedItems   map[string]struct{}
// 	DeletedFolders map[string]struct{}
// }

// func NewInDeets() *InDeets {
// 	id := &InDeets{}
// 	id.init()

// 	return id
// }

// func (id *InDeets) init() {
// 	id.Items = map[string]details.DetailsEntry{}
// 	id.Folders = map[string]details.DetailsEntry{}
// 	id.DeletedItems = map[string]struct{}{}
// 	id.DeletedFolders = map[string]struct{}{}
// }

// func (id *InDeets) Add(deets details.Details) {
// 	if id.Items == nil {
// 		id.init()
// 	}

// 	for _, ent := range deets.Entries {
// 		if ent.Folder == nil {
// 			id.Items[ent.RepoRef] = ent
// 		} else {
// 			id.Folders[ent.RepoRef] = ent
// 		}
// 	}
// }

// func (id *InDeets) AddItem(parentRR, i string) {
// 	rr := parentRR + "/" + i

// 	id.Items[rr] = details.DetailsEntry{RepoRef: rr}
// 	delete(id.DeletedItems, rr)
// }

// func (id *InDeets) RemoveItem(rr string) {
// 	fmt.Printf("\n-----\ndeleting %+v\n-----\n", rr)
// 	delete(id.Items, rr)
// 	id.DeletedItems[rr] = struct{}{}
// }

// func (id *InDeets) AddFolder(rr string) {
// 	id.Folders[rr] = details.DetailsEntry{RepoRef: rr}
// 	delete(id.DeletedFolders, rr)
// }

// func (id *InDeets) RemoveFolder(rr string) {
// 	fmt.Printf("\n-----\ndeleting %+v\n-----\n", rr)
// 	delete(id.Folders, rr)
// 	id.DeletedFolders[rr] = struct{}{}

// 	for k := range id.Items {
// 		if strings.HasPrefix(k, rr) {
// 			id.RemoveItem(k)
// 		}
// 	}
// }

// func (id *InDeets) Subset(folderRR string) *InDeets {
// 	nd := NewInDeets()

// 	fEnt, ok := id.Folders[folderRR]
// 	if ok {
// 		nd.Folders[folderRR] = fEnt
// 	}

// 	dfEnt, ok := id.DeletedFolders[folderRR]
// 	if ok {
// 		nd.DeletedFolders[folderRR] = dfEnt
// 	}

// 	for rr, ent := range id.Items {
// 		if strings.HasPrefix(rr, folderRR) {
// 			nd.Items[rr] = ent
// 		}
// 	}

// 	for rr, ent := range id.DeletedItems {
// 		if strings.HasPrefix(rr, folderRR) {
// 			nd.DeletedItems[rr] = ent
// 		}
// 	}

// 	return nd
// }

// // assumes destination does not already exist
// // also assumes item IDs don't change.  if item IDs change
// // as a result of the move, you should instead chaing:
// // RemoveFolder(), AddFolder(), AddItems()
// func (id *InDeets) MoveFolder(t *testing.T, fromRR, toRR string) {
// 	fmt.Printf("\n-----\nmove to %+v\n-----\n", toRR)
// 	id.AddFolder(toRR)

// 	moved := map[string]struct{}{}

// 	for rr := range id.Items {
// 		if strings.HasPrefix(rr, fromRR) {
// 			moved[rr] = struct{}{}
// 		}
// 	}

// 	for rr := range moved {
// 		delete(id.Items, rr)

// 		_, i := path.Split(rr)
// 		id.AddItem(toRR, i)
// 	}

// 	id.RemoveFolder(fromRR)

// 	fmt.Printf("\n-----\nmove from %+v\n-----\n", fromRR)
// }

// ---------------------------------------------------------------------------
//
// ---------------------------------------------------------------------------

// // InDeets is a helper for comparing details state in tests
// // across backup instances.
// type InDeets struct {
// 	Items          map[string]struct{}
// 	Folders        map[string]struct{}
// 	ItemsInFolders map[string][]string
// 	DeletedItems   map[string]struct{}
// 	DeletedFolders map[string]struct{}
// }

// func NewInDeets() *InDeets {
// 	id := &InDeets{}
// 	id.init()

// 	return id
// }

// func (id *InDeets) init() {
// 	id.Items = map[string]struct{}{}
// 	id.ItemsInFolders = map[string][]string{}
// 	id.Folders = map[string]struct{}{}
// 	id.DeletedItems = map[string]struct{}{}
// 	id.DeletedFolders = map[string]struct{}{}
// }

// func (id *InDeets) Add(deets details.Details) {
// 	if id.Items == nil {
// 		id.init()
// 	}

// 	for _, ent := range deets.Entries {
// 		if ent.Folder == nil {
// 			id.AddItem(ent.ItemRef, ent.ParentRef)
// 		} else {
// 			id.AddFolder(ent.ParentRef, ent.Folder.DisplayName)
// 		}
// 	}
// }

// func (id *InDeets) AddItem(i, f string) {
// 	id.Items[i] = struct{}{}

// 	is := id.ItemsInFolders[f]
// 	id.ItemsInFolders[f] = append(is, i)

// 	delete(id.DeletedItems, i)
// }

// func (id *InDeets) RemoveItem(i string) {
// 	delete(id.Items, i)
// 	id.DeletedItems[i] = struct{}{}
// }

// func (id *InDeets) AddFolder(parent, f string) {
// 	f = parent + "/" + f

// 	id.Folders[f] = struct{}{}
// 	delete(id.DeletedFolders, f)
// }

// func (id *InDeets) RemoveFolder(parent, f string) {
// 	f = parent + "/" + f

// 	for _, i := range id.ItemsInFolders[f] {
// 		id.RemoveItem(i)
// 	}

// 	delete(id.Folders, f)
// 	id.DeletedFolders[f] = struct{}{}
// }

// // assumes destination does not already exist
// func (id *InDeets) Movefolder(fromParent, fromF, toParent, toF string) {
// 	from := fromParent + "/" + fromF
// 	to := toParent + "/" + toF
// 	is := id.ItemsInFolders[from]

// 	delete(id.ItemsInFolders, from)
// 	delete(id.Folders, from)

// 	id.ItemsInFolders[to] = is
// 	id.Folders[to] = struct{}{}
// }

func CheckBackupDetails(
	t *testing.T,
	ctx context.Context, //revive:disable:context-as-argument
	backupID model.StableID,
	// kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	ssr streamstore.Reader,
	expect *InDeets,
) {
	_, result := GetDeetsInBackup(t, ctx, backupID, ms, ssr)

	t.Log("details entries in result")

	for rr := range result.Entries {
		t.Log(rr)
	}

	assert.Subset(t, maps.Keys(result.Entries), maps.Keys(expect.Entries), "result missing expected entry")

	for rr := range expect.Deleted {
		_, ok := result.Entries[rr]
		assert.False(t, ok, "deleted entry found in result: %s", rr)
	}
}

// func CheckBackupDetails(
// 	t *testing.T,
// 	ctx context.Context, //revive:disable:context-as-argument
// 	backupID model.StableID,
// 	// kw *kopia.Wrapper,
// 	ms *kopia.ModelStore,
// 	ssr streamstore.Reader,
// 	expect *InDeets,
// ) {
// 	_, result := GetDeetsInBackup(t, ctx, backupID, ms, ssr)

// 	assert.Subset(t, maps.Keys(result.Items), maps.Keys(expect.Items), "result should contain expected items")
// 	assert.Subset(t, maps.Keys(result.Folders), maps.Keys(expect.Folders), "result should contain expected folders")

// 	for i := range expect.DeletedItems {
// 		assert.Nil(t, result.Items[i], "item was deleted and should not exist: %s", i)
// 	}

// 	for f := range expect.DeletedFolders {
// 		assert.Nil(t, result.Folders[f], "folder was deleted and should not exist: %s", f)
// 	}
// }

func GetDeetsInBackup(
	t *testing.T,
	ctx context.Context, //revive:disable:context-as-argument
	backupID model.StableID,
	// kw *kopia.Wrapper,
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

	id := NewInDeets()
	id.Add(deets)

	return deets, id
}
