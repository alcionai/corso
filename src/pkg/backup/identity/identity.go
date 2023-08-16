package identity

import "github.com/alcionai/corso/src/pkg/path"

type SubtreeReasoner interface {
	Reasoner
	SubtreePather
}

// Reasoner describes the parts of the backup that make up its
// data identity: the tenant, protected resources, services, and
// categories which are held within the backup.
//
// Reasoner only recognizes the "primary" protected resource and
// service. IE: subservice resources and services are not recognized
// as part of the backup Reason.
type Reasoner interface {
	Tenant() string
	// ProtectedResource represents the Primary protected resource.
	// IE: if a path or backup supports subservices, this value
	// should only provide the first service's resource, and not the
	// resource for any subservice.
	ProtectedResource() string
	// Service represents the Primary service.
	// IE: if a path or backup supports subservices, this value
	// should only provide the first service; not a subservice.
	Service() path.ServiceType
	Category() path.CategoryType
}

type SubtreePather interface {
	// SubtreePath returns the path prefix for data in existing backups that have
	// parameters (tenant, protected resourced, etc) that match this Reasoner.
	SubtreePath() (path.Path, error)
}
