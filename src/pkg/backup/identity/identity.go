package identity

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/path"
)

// Reasoner describes the parts of the backup that make up its
// data identity: the tenant, protected resources, services, and
// categories which are held within the backup.
type Reasoner interface {
	Tenant() string
	ProtectedResource() string
	Service() path.ServiceType
	Category() path.CategoryType
	// SubtreePath returns the path prefix for data in existing backups that have
	// parameters (tenant, protected resourced, etc) that match this Reasoner.
	SubtreePath() (path.Path, error)
	// ToMetadata returns the corresponding metadata reason for this reason.
	ToMetadata() Reasoner
}

func NewReason(
	tenantID, resourceID string,
	service path.ServiceType,
	category path.CategoryType,
) Reasoner {
	return reason{
		tenant:   tenantID,
		resource: resourceID,
		service:  service,
		category: category,
	}
}

type reason struct {
	// tenant appears here so that when this is moved to an inject package nothing
	// needs changed. However, kopia itself is blind to the fields in the reason
	// struct and relies on helper functions to get the information it needs.
	tenant   string
	resource string
	service  path.ServiceType
	category path.CategoryType
}

func (r reason) Tenant() string {
	return r.tenant
}

func (r reason) ProtectedResource() string {
	return r.resource
}

func (r reason) Service() path.ServiceType {
	return r.service
}

func (r reason) Category() path.CategoryType {
	return r.category
}

func (r reason) SubtreePath() (path.Path, error) {
	p, err := path.BuildPrefix(
		r.Tenant(),
		r.ProtectedResource(),
		r.Service(),
		r.Category())

	return p, clues.Wrap(err, "building path").OrNil()
}

func (r reason) ToMetadata() Reasoner {
	return NewReason(
		r.Tenant(),
		r.ProtectedResource(),
		r.Service().ToMetadata(),
		r.Category())
}
