package identity

import "github.com/alcionai/corso/src/pkg/path"

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
}
