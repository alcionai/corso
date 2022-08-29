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
// resource-specific path. This struct is package-private so that callers are
// forced to use the pre-defined constructors, making it impossible to create a
// dataLayerResourcePath with invalid service/category combinations.
//
// All dataLayerResourcePaths start with the same prefix:
// <tenant ID>/<service>/<owner ID> which allows extracting high-level
// information from the path. The prefix may additionally contain a <category>
// after the owner ID if the service has multiple disjoint data sets (e.x.
// Exchange has email and contacts). The path elements after this prefix
// represent zero or more folders the path refers to and possibly an item ID if
// the path refers to an individual item. A valid dataLayerResourcePath must
// have at least one folder or an item so that the resulting path has at least
// on element after the prefix.
type dataLayerResourcePath struct {
	Builder
	category CategoryType
	service  ServiceType
	hasItem  bool
}

// Tenant returns the tenant ID embedded in the dataLayerResourcePath.
func (dlrp dataLayerResourcePath) Tenant() string {
	return dlrp.Builder.elements[0]
}

// Service returns the ServiceType embedded in the dataLayerResourcePath.
func (dlrp dataLayerResourcePath) Service() ServiceType {
	return dlrp.service
}

// Category returns the CategoryType embedded in the dataLayerResourcePath.
func (dlrp dataLayerResourcePath) Category() CategoryType {
	return dlrp.category
}

// ResourceOwner returns the user ID or group ID embedded in the
// dataLayerResourcePath.
func (dlrp dataLayerResourcePath) ResourceOwner() string {
	return dlrp.Builder.elements[2]
}

// Folder returns the folder segment embedded in the dataLayerResourcePath.
func (dlrp dataLayerResourcePath) Folder() string {
	startIdx := 3

	if dlrp.category != UnknownCategory {
		startIdx++
	}

	endIdx := len(dlrp.Builder.elements)

	if dlrp.hasItem {
		endIdx--
	}

	return dlrp.Builder.join(startIdx, endIdx)
}

// Item returns the item embedded in the dataLayerResourcePath if the path
// refers to an item.
func (dlrp dataLayerResourcePath) Item() string {
	if dlrp.hasItem {
		return dlrp.Builder.elements[len(dlrp.Builder.elements)-1]
	}

	return ""
}
