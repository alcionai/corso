package path

import (
	"strings"

	"github.com/alcionai/corso/src/internal/path"
)

type mockDataLayerResourcePath struct {
	elems    []string
	category path.CategoryType
	service  path.ServiceType
	hasItem  bool
}

func MockEmptyPath() path.Path {
	return mockDataLayerResourcePath{
		category: path.UnknownCategory,
		service:  path.UnknownService,
	}
}

func MockPath(
	vs []string,
	item bool,
	srv path.ServiceType,
	cat path.CategoryType,
) path.Path {
	return mockDataLayerResourcePath{
		elems:    append([]string{"tenant_id", srv.String(), "resource_owner", cat.String()}, vs...),
		category: cat,
		service:  srv,
		hasItem:  item,
	}
}

// Tenant returns the tenant ID embedded in the dataLayerResourcePath.
func (rp mockDataLayerResourcePath) Tenant() string {
	if len(rp.elems) == 0 {
		return ""
	}

	return rp.elems[0]
}

// Service returns the ServiceType embedded in the dataLayerResourcePath.
func (rp mockDataLayerResourcePath) Service() path.ServiceType {
	return rp.service
}

// Category returns the CategoryType embedded in the dataLayerResourcePath.
func (rp mockDataLayerResourcePath) Category() path.CategoryType {
	return rp.category
}

// ResourceOwner returns the user ID or group ID embedded in the
// dataLayerResourcePath.
func (rp mockDataLayerResourcePath) ResourceOwner() string {
	if len(rp.elems) == 0 {
		return ""
	}

	return rp.elems[2]
}

// Folder returns the folder segment embedded in the dataLayerResourcePath.
func (rp mockDataLayerResourcePath) Folder() string {
	if len(rp.elems) == 0 {
		return ""
	}

	endIdx := len(rp.elems)

	if rp.hasItem {
		endIdx--
	}

	return rp.join(4, endIdx)
}

func (rp mockDataLayerResourcePath) join(start, end int) string {
	return strings.Join(rp.elems[start:end], "/")
}

// Item returns the item embedded in the dataLayerResourcePath if the path
// refers to an item.
func (rp mockDataLayerResourcePath) Item() string {
	if len(rp.elems) == 0 {
		return ""
	}

	if rp.hasItem {
		return rp.elems[len(rp.elems)-1]
	}

	return ""
}

func (rp mockDataLayerResourcePath) String() string {
	return rp.join(0, len(rp.elems))
}
