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
	category         CategoryType
	serviceResources []ServiceResource
	hasItem          bool
}

// performs no validation, assumes the caller has validated the inputs.
func newDataLayerResourcePath(
	pb Builder,
	tenant string,
	srs []ServiceResource,
	cat CategoryType,
	isItem bool,
) dataLayerResourcePath {
	pfx := append([]string{tenant}, ServiceResourcesToElements(srs)...)
	pfx = append(pfx, cat.String())

	return dataLayerResourcePath{
		Builder:          *pb.withPrefix(pfx...),
		serviceResources: srs,
		category:         cat,
		hasItem:          isItem,
	}
}

// Tenant returns the tenant ID embedded in the dataLayerResourcePath.
func (rp dataLayerResourcePath) Tenant() string {
	return rp.Builder.elements[0]
}

func (rp dataLayerResourcePath) ServiceResources() []ServiceResource {
	return rp.serviceResources
}

func (rp dataLayerResourcePath) PrimaryService() ServiceType {
	srs := rp.serviceResources

	if len(srs) == 0 {
		return UnknownService
	}

	return srs[0].Service
}

func (rp dataLayerResourcePath) PrimaryProtectedResource() string {
	srs := rp.serviceResources

	if len(srs) == 0 {
		return ""
	}

	return srs[0].ProtectedResource
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

func (rp dataLayerResourcePath) prefixLen() int {
	return 2 + 2*len(rp.serviceResources)
}

// Folder returns the folder segment embedded in the dataLayerResourcePath.
func (rp dataLayerResourcePath) Folder(escape bool) string {
	endIdx := rp.lastFolderIdx()
	pfxLen := rp.prefixLen()

	if endIdx == pfxLen {
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
	pfxLen := rp.prefixLen()

	// if endIdx == prefix length, there are no folders
	if endIdx == pfxLen {
		return nil
	}

	return append([]string{}, rp.elements[pfxLen:endIdx]...)
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
	// Dir is not allowed to slice off any prefix values.
	// The prefix len is determined by the length of the number of
	// service+resource tuples, plus 2 (tenant and category).
	if len(rp.elements) <= 2+(2*len(rp.serviceResources)) {
		return nil, clues.New("unable to shorten path").With("path", rp)
	}

	return &dataLayerResourcePath{
		Builder:          *rp.Builder.Dir(),
		serviceResources: rp.serviceResources,
		category:         rp.category,
		hasItem:          false,
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
		Builder:          *rp.Builder.Append(elems...),
		serviceResources: rp.serviceResources,
		category:         rp.category,
		hasItem:          isItem,
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

func (rp *dataLayerResourcePath) Halves() (*Builder, Elements) {
	pfx, sfx := &Builder{}, Elements{}

	b := rp.Builder
	if len(b.elements) > 0 {
		lenPfx := 2 + (len(rp.serviceResources) * 2)

		pfx = &Builder{elements: append(Elements{}, b.elements[:lenPfx]...)}
		sfx = append(Elements{}, b.elements[lenPfx-1:]...)
	}

	return pfx, sfx
}
