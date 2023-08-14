package details

import (
	"sync"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/pkg/path"
)

// Builder should be used to create a details model.
type Builder struct {
	d            Details
	mu           sync.Mutex       `json:"-"`
	knownFolders map[string]Entry `json:"-"`
}

func (b *Builder) Empty() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.d.Entries) == 0
}

func (b *Builder) Add(
	repoRef path.Path,
	locationRef *path.Builder,
	info ItemInfo,
) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	entry, err := b.d.add(
		repoRef,
		locationRef,
		info)
	if err != nil {
		return clues.Wrap(err, "adding entry to details")
	}

	if err := b.addFolderEntries(
		repoRef.ToBuilder().Dir(),
		locationRef,
		entry,
	); err != nil {
		return clues.Wrap(err, "adding folder entries")
	}

	return nil
}

func (b *Builder) addFolderEntries(
	repoRef, locationRef *path.Builder,
	entry Entry,
) error {
	if len(repoRef.Elements()) < len(locationRef.Elements()) {
		return clues.New("RepoRef shorter than LocationRef").
			With("repo_ref", repoRef, "location_ref", locationRef)
	}

	if b.knownFolders == nil {
		b.knownFolders = map[string]Entry{}
	}

	// Need a unique location because we want to have separate folders for
	// different drives and categories even if there's duplicate folder names in
	// them.
	uniqueLoc, err := entry.uniqueLocation(locationRef)
	if err != nil {
		return clues.Wrap(err, "getting LocationIDer")
	}

	for uniqueLoc.elementCount() > 0 {
		mapKey := uniqueLoc.ID().ShortRef()

		name := uniqueLoc.lastElem()
		if len(name) == 0 {
			return clues.New("folder with no display name").
				With("repo_ref", repoRef, "location_ref", uniqueLoc.InDetails())
		}

		shortRef := repoRef.ShortRef()
		rr := repoRef.String()

		// Get the parent of this entry to add as the LocationRef for the folder.
		uniqueLoc.dir()

		repoRef = repoRef.Dir()
		parentRef := repoRef.ShortRef()

		folder, ok := b.knownFolders[mapKey]
		if !ok {
			loc := uniqueLoc.InDetails().String()

			folder = Entry{
				RepoRef:     rr,
				ShortRef:    shortRef,
				ParentRef:   parentRef,
				LocationRef: loc,
				ItemInfo: ItemInfo{
					Folder: &FolderInfo{
						ItemType: FolderItem,
						// TODO(ashmrtn): Use the item type returned by the entry once
						// SharePoint properly sets it.
						DisplayName: name,
					},
				},
			}

			if err := entry.updateFolder(folder.Folder); err != nil {
				return clues.Wrap(err, "adding folder").
					With("parent_repo_ref", repoRef, "location_ref", loc)
			}
		}

		folder.Folder.Size += entry.size()

		itemModified := entry.Modified()
		if folder.Folder.Modified.Before(itemModified) {
			folder.Folder.Modified = itemModified
		}

		// Always update the map because we're storing structs not pointers to
		// structs.
		b.knownFolders[mapKey] = folder
	}

	return nil
}

func (b *Builder) Details() *Details {
	b.mu.Lock()
	defer b.mu.Unlock()

	ents := make([]Entry, len(b.d.Entries))
	copy(ents, b.d.Entries)

	// Write the cached folder entries to details
	details := &Details{
		DetailsModel{
			Entries: append(ents, maps.Values(b.knownFolders)...),
		},
	}

	return details
}
