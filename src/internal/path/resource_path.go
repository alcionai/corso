package path

type ServiceType int

//go:generate stringer -type=ServiceType -linecomment
const (
	UnknownService  ServiceType = iota
	ExchangeService             // exchange
)

type CategoryType int

//go:generate stringer -type=CategoryType -linecomment
const (
	UnknownCategory CategoryType = iota
	EmailCategory                // email
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

// Folder returns the folder segment embedded in the dataLayerResourcePath.
func (rp dataLayerResourcePath) Folder() string {
	endIdx := len(rp.Builder.elements)

	if rp.hasItem {
		endIdx--
	}

	return rp.Builder.join(4, endIdx)
}

// Item returns the item embedded in the dataLayerResourcePath if the path
// refers to an item.
func (rp dataLayerResourcePath) Item() string {
	if rp.hasItem {
		return rp.Builder.elements[len(rp.Builder.elements)-1]
	}

	return ""
}
