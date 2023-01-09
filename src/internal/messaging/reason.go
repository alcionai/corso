package messaging

import (
	"github.com/alcionai/corso/src/pkg/path"
)

type Reason struct {
	ResourceOwner string
	Service       path.ServiceType
	Category      path.CategoryType
}
