package its

import "github.com/alcionai/corso/src/pkg/path"

type ResourceServicer interface {
	Resource() IDs
	Service() path.ServiceType
}

type resourceAndService struct {
	ProtectedResource IDs
	ServiceType       path.ServiceType
}

func (ras resourceAndService) Resource() IDs {
	return ras.ProtectedResource
}

func (ras resourceAndService) Service() path.ServiceType {
	return ras.ServiceType
}

func NewResourceService(r IDs, s path.ServiceType) ResourceServicer {
	return &resourceAndService{r, s}
}
