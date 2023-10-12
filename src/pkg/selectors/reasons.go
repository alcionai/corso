package selectors

import (
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

func key(br identity.Reasoner) string {
	return br.Category().String() +
		br.ProtectedResource() +
		br.Service().String() +
		br.Tenant()
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
			br := identity.NewReason(tenantID, resource, service, cat)
			reasons[key(br)] = br
		}
	}

	return maps.Values(reasons)
}
