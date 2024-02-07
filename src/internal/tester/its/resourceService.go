package its

import "github.com/alcionai/canario/src/pkg/path"

type ResourceServicer interface {
	Resource() IDs
	Service() path.ServiceType
}

var _ ResourceServicer = resourceAndService{}

type resourceAndService struct {
	protectedResource IDs
	serviceType       path.ServiceType
}

func (ras resourceAndService) Resource() IDs {
	return ras.protectedResource
}

func (ras resourceAndService) Service() path.ServiceType {
	return ras.serviceType
}

func NewResourceService(r IDs, s path.ServiceType) ResourceServicer {
	return &resourceAndService{
		protectedResource: r,
		serviceType:       s,
	}
}
