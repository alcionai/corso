package source

import "github.com/pkg/errors"

type service int

//go:generate stringer -type=service -linecomment
const (
	ServiceUnknown  service = iota // Unknown Service
	ServiceExchange                // Exchange
)

var ErrorBadSourceCast = errors.New("wrong source service type")

const (
	scopeKeyGranularity = "granularity"
	scopeKeyFullPath    = "fullPath"
)

// The core source.  Has no api for setting or retrieving data.
// Is only used to pass along more specific source instances.
type Source struct {
	TenantID string  // The tenant making the request.
	service  service // The service scope of the data.  Exchange, Teams, Sharepoint, etc.
	scopes   []any   // A slice of scopes, held as 'any' in the source.
}

// helper for specific source instance constructors.
func newSource(tenantID string, s service) Source {
	return Source{
		TenantID: tenantID,
		service:  s,
		scopes:   []any{},
	}
}

// Service return the service enum for the source.
func (s Source) Service() service {
	return s.service
}

func BadCastErr(cast, is service) error {
	return errors.Wrapf(ErrorBadSourceCast, "%s service is not %s", cast, is)
}
