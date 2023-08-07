package details

import (
	"context"

	"github.com/alcionai/corso/src/cli/print"
)

// DetailsModel describes what was stored in a Backup
type DetailsModel struct {
	Entries []Entry `json:"entries"`
}

// Print writes the DetailModel Entries to StdOut, in the format
// requested by the caller.
func (dm DetailsModel) PrintEntries(ctx context.Context) {
	printEntries(ctx, dm.Entries)
}

type infoer interface {
	Entry | *Entry
	// Need this here so we can access the infoType function without a type
	// assertion. See https://stackoverflow.com/a/71378366 for more details.
	infoType() ItemType
}

func printEntries[T infoer](ctx context.Context, entries []T) {
	if print.DisplayJSONFormat() {
		printJSON(ctx, entries)
	} else {
		printTable(ctx, entries)
	}
}

func printTable[T infoer](ctx context.Context, entries []T) {
	perType := map[ItemType][]print.Printable{}

	for _, ent := range entries {
		it := ent.infoType()
		ps, ok := perType[it]

		if !ok {
			ps = []print.Printable{}
		}

		perType[it] = append(ps, print.Printable(ent))
	}

	for _, ps := range perType {
		print.All(ctx, ps...)
	}
}

func printJSON[T infoer](ctx context.Context, entries []T) {
	ents := []print.Printable{}

	for _, ent := range entries {
		ents = append(ents, print.Printable(ent))
	}

	print.All(ctx, ents...)
}

// Paths returns the list of Paths for non-folder and non-meta items extracted
// from the Entries slice.
func (dm DetailsModel) Paths() []string {
	r := make([]string, 0, len(dm.Entries))

	for _, ent := range dm.Entries {
		if ent.Folder != nil || ent.isMetaFile() {
			continue
		}

		r = append(r, ent.RepoRef)
	}

	return r
}

// Items returns a slice of *ItemInfo that does not contain any FolderInfo
// entries. Required because not all folders in the details are valid resource
// paths, and we want to slice out metadata.
func (dm DetailsModel) Items() entrySet {
	res := make([]*Entry, 0, len(dm.Entries))

	for i := 0; i < len(dm.Entries); i++ {
		ent := dm.Entries[i]
		if ent.Folder != nil || ent.isMetaFile() {
			continue
		}

		res = append(res, &ent)
	}

	return res
}

// FilterMetaFiles returns a copy of the Details with all of the
// .meta files removed from the entries.
func (dm DetailsModel) FilterMetaFiles() DetailsModel {
	d2 := DetailsModel{
		Entries: []Entry{},
	}

	for _, ent := range dm.Entries {
		if !ent.isMetaFile() {
			d2.Entries = append(d2.Entries, ent)
		}
	}

	return d2
}

// SumNonMetaFileSizes returns the total size of items excluding all the
// .meta files from the items.
func (dm DetailsModel) SumNonMetaFileSizes() int64 {
	var size int64

	// Items will provide only files and filter out folders
	for _, ent := range dm.FilterMetaFiles().Items() {
		size += ent.size()
	}

	return size
}
