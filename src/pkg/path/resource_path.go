package path

import (
	"github.com/alcionai/clues"
)

// dataLayerResourcePath allows callers to extract information from a
// resource-specific path. This struct is unexported so that callers are
// forced to use the pre-defined constructors, making it impossible to create a
// dataLayerResourcePath with invalid service/category combinations.
//
// All dataLayerResourcePaths start with the same prefix:
// <tenant ID>/<service>/<resource owner ID>/<category>
// which allows extracting high-level information from the path. The path
// elements after this prefix represent zero or more folders and, if the path
// refers to a file or item, an item ID. A valid dataLayerResourcePath must have
// at least one folder or an item so that the resulting path has at least one
// element after the prefix.
type dataLayerResourcePath struct {
	Builder
	category CategoryType
	service  ServiceType
	hasItem  bool
}

// Tenant returns the tenant ID embedded in the dataLayerResourcePath.
func (rp dataLayerResourcePath) Tenant() string {
	return rp.Builder.elements[0]
}

// Service returns the ServiceType embedded in the dataLayerResourcePath.
func (rp dataLayerResourcePath) Service() ServiceType {
	return rp.service
}

// Category returns the CategoryType embedded in the dataLayerResourcePath.
func (rp dataLayerResourcePath) Category() CategoryType {
	return rp.category
}

// ResourceOwner returns the user ID or group ID embedded in the
// dataLayerResourcePath.
func (rp dataLayerResourcePath) ResourceOwner() string {
	return rp.Builder.elements[2]
}

func (rp dataLayerResourcePath) lastFolderIdx() int {
	endIdx := len(rp.Builder.elements)

	if rp.hasItem {
		endIdx--
	}

	return endIdx
}

// Folder returns the folder segment embedded in the dataLayerResourcePath.
func (rp dataLayerResourcePath) Folder(escape bool) string {
	endIdx := rp.lastFolderIdx()
	if endIdx == 4 {
		return ""
	}

	fs := rp.Folders()

	if !escape {
		return join(fs)
	}

	// builder.String() will escape all individual elements.
	return Builder{}.Append(fs...).String()
}

// Folders returns the individual folder elements embedded in the
// dataLayerResourcePath.
func (rp dataLayerResourcePath) Folders() Elements {
	endIdx := rp.lastFolderIdx()
	if endIdx == 4 {
		return nil
	}

	return append([]string{}, rp.elements[4:endIdx]...)
}

// Item returns the item embedded in the dataLayerResourcePath if the path
// refers to an item.
func (rp dataLayerResourcePath) Item() string {
	if rp.hasItem {
		return rp.Builder.elements[len(rp.Builder.elements)-1]
	}

	return ""
}

// Dir removes the last element from the path.  If this would remove a
// value that is part of the standard prefix structure, an error is returned.
func (rp dataLayerResourcePath) Dir() (Path, error) {
	if len(rp.elements) <= 4 {
		return nil, clues.New("unable to shorten path").With("path", rp)
	}

	return &dataLayerResourcePath{
		Builder:  *rp.Builder.Dir(),
		service:  rp.service,
		category: rp.category,
		hasItem:  false,
	}, nil
}

func (rp dataLayerResourcePath) Append(
	isItem bool,
	elems ...string,
) (Path, error) {
	if rp.hasItem {
		return nil, clues.New("appending to an item path")
	}

	return &dataLayerResourcePath{
		Builder:  *rp.Builder.Append(elems...),
		service:  rp.service,
		category: rp.category,
		hasItem:  isItem,
	}, nil
}

func (rp dataLayerResourcePath) AppendItem(item string) (Path, error) {
	return rp.Append(true, item)
}

func (rp dataLayerResourcePath) ToBuilder() *Builder {
	// Safe to directly return the Builder because Builders are immutable.
	return &rp.Builder
}

func (rp *dataLayerResourcePath) UpdateParent(prev, cur Path) bool {
	return rp.Builder.UpdateParent(prev.ToBuilder(), cur.ToBuilder())
}
