package groups

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/pkg/filters"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

func IsServiceEnabled(
	ctx context.Context,
	gbi api.GetByIDer[models.Groupable],
	resource string,
) (bool, error) {
	resp, err := gbi.GetByID(ctx, resource, api.CallConfig{})
	if err != nil {
		return false, clues.WrapWC(ctx, err, "getting group")
	}

	// according to graph api docs: https://learn.microsoft.com/en-us/graph/api/resources/group?view=graph-rest-1.0
	// "If the collection contains Unified, the group is a Microsoft 365 group;
	// otherwise, it's either a security group or distribution group."
	//
	// Basically, if it's "unified", then we actually have data to back up.
	// If it's not unified, then its purely a mailing list, and has no backing data.
	isUnified := filters.
		Equal(resp.GetGroupTypes()).
		Compare("unified")

	return isUnified, nil
}
