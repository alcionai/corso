package selectors

import (
	"golang.org/x/exp/maps"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// reasoner interface compliance
// ---------------------------------------------------------------------------

var _ identity.Reasoner = &backupReason{}

type backupReason struct {
	category path.CategoryType
	resource string
	service  path.ServiceType
	tenant   string
}

func (br backupReason) Tenant() string {
	return br.tenant
}

func (br backupReason) ProtectedResource() string {
	return br.resource
}

func (br backupReason) Service() path.ServiceType {
	return br.service
}

func (br backupReason) Category() path.CategoryType {
	return br.category
}

func (br backupReason) SubtreePath() (path.Path, error) {
	srs, err := path.NewServiceResources(br.service, br.resource)
	if err != nil {
		return nil, clues.Wrap(err, "building path prefix services")
	}

	return path.BuildPrefix(br.tenant, srs, br.category)
}

func (br backupReason) key() string {
	return br.category.String() + br.resource + br.service.String() + br.tenant
}

// ---------------------------------------------------------------------------
// common transformer
// ---------------------------------------------------------------------------

type servicerCategorizerProvider interface {
	pathServicer
	pathCategorier
	idname.Provider
}

// produces the Reasoner basis described by the selector.
// In cases of reasons with subservices (ie, multiple
// services described by a backup or path), the selector
// will only ever generate a ServiceResource for the first
// service+resource pair in the set.
//
// TODO: it may be possible, if necessary, to add subservice
// recognition to the service via additional scopes.
func reasonsFor(
	sel servicerCategorizerProvider,
	tenantID string,
	useOwnerNameForID bool,
) []identity.Reasoner {
	service := sel.PathService()
	reasons := map[string]identity.Reasoner{}

	resource := sel.ID()
	if useOwnerNameForID {
		resource = sel.Name()
	}

	pc := sel.PathCategories()

	for _, sl := range [][]path.CategoryType{pc.Includes, pc.Filters} {
		for _, cat := range sl {
			br := backupReason{
				category: cat,
				resource: resource,
				service:  service,
				tenant:   tenantID,
			}

			reasons[br.key()] = br
		}
	}

	return maps.Values(reasons)
}
