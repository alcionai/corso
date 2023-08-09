package selectors

import (
	"golang.org/x/exp/maps"

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
	return path.ServicePrefix(
		br.tenant,
		br.resource,
		br.service,
		br.category)
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
