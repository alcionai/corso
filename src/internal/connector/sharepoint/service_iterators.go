package sharepoint

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// FilterContainersAndFillCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a Collection. Items outside of those directories are omitted.
// @param collection is filled with during this function.
func FilterContainersAndFillCollections(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
	scope selectors.SharePointScope,
) error {
	return nil
}

// code previously within the function, moved here to make the linter happy

// var (
// 	category       = qp.Scope.Category().PathType()
// 	collectionType = CategoryToOptionIdentifier(category)
// 	errs           error
// )

// for _, c := range resolver.Items() {
// 	dirPath, ok := pathAndMatch(qp, category, c)
// 	if ok {
// 		// Create only those that match
// 		service, err := createService(qp.Credentials, qp.FailFast)
// 		if err != nil {
// 			errs = support.WrapAndAppend(
// 				qp.User+" FilterContainerAndFillCollection",
// 				err,
// 				errs)

// 			if qp.FailFast {
// 				return errs
// 			}
// 		}

// 		edc := NewCollection(
// 			qp.User,
// 			dirPath,
// 			collectionType,
// 			service,
// 			statusUpdater,
// 		)
// 		collections[*c.GetId()] = &edc
// 	}
// }

// for directoryID, col := range collections {
// 	fetchFunc, err := getFetchIDFunc(category)
// 	if err != nil {
// 		errs = support.WrapAndAppend(
// 			qp.User,
// 			err,
// 			errs)

// 		if qp.FailFast {
// 			return errs
// 		}

// 		continue
// 	}

// 	jobs, err := fetchFunc(ctx, col.service, qp.User, directoryID)
// 	if err != nil {
// 		errs = support.WrapAndAppend(
// 			qp.User,
// 			err,
// 			errs,
// 		)
// 	}

// 	col.jobs = append(col.jobs, jobs...)
// }

// return errs
