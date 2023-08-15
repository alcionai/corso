package mock

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/path"
)

type Reason struct {
	TenantID   string
	Cat        path.CategoryType
	Svc        path.ServiceType
	Resource   string
	SubtreeErr error
}

func (r Reason) Tenant() string {
	return r.TenantID
}

func (r Reason) Category() path.CategoryType {
	return r.Cat
}

func (r Reason) Service() path.ServiceType {
	return r.Svc
}

func (r Reason) ProtectedResource() string {
	return r.Resource
}

func (r Reason) SubtreePath() (path.Path, error) {
	if r.SubtreeErr != nil {
		return nil, r.SubtreeErr
	}

	p, err := path.BuildPrefix(
		r.Tenant(),
		[]path.ServiceResource{{
			ProtectedResource: r.Resource,
			Service:           r.Svc,
		}},
		r.Category())

	return p, clues.Wrap(err, "building path").OrNil()
}
