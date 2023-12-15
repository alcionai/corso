package site

import (
	"context"

	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// DeleteList removes a list object from a site.
// deletes require unique http clients
// https://github.com/alcionai/corso/issues/2707
func DeleteList(
	ctx context.Context,
	gs graph.Servicer,
	siteID, listID string,
) error {
	err := gs.Client().Sites().BySiteId(siteID).Lists().ByListId(listID).Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting list")
	}

	return nil
}
