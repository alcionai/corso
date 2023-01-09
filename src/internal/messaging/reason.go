package messaging

import (
	"github.com/alcionai/corso/src/pkg/path"
)

type Reason struct {
	ResourceOwner string
	Service       path.ServiceType
	Category      path.CategoryType
}

func (r Reason) TagKeys() []string {
	return []string{
		r.ResourceOwner,
		serviceCatString(r.Service, r.Category),
	}
}

func serviceCatString(s path.ServiceType, c path.CategoryType) string {
	return s.String() + c.String()
}
