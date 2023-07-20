package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Lists() Lists {
	return Lists{c}
}

// Lists is an interface-compliant provider of the client.
type Lists struct {
	Client
}

// PostDrive creates a new list of type drive.  Specifically used to create
// documentLibraries for SharePoint Sites.
func (c Lists) PostDrive(
	ctx context.Context,
	siteID, driveName string,
) (models.Driveable, error) {
	list := models.NewList()
	list.SetDisplayName(&driveName)
	list.SetDescription(ptr.To("corso auto-generated restore destination"))

	li := models.NewListInfo()
	li.SetTemplate(ptr.To("documentLibrary"))
	list.SetList(li)

	// creating a list of type documentLibrary will result in the creation
	// of a new drive owned by the given site.
	builder := c.Stable.
		Client().
		Sites().
		BySiteId(siteID).
		Lists()

	newList, err := builder.Post(ctx, list, nil)
	if graph.IsErrItemAlreadyExistsConflict(err) {
		return nil, clues.Stack(graph.ErrItemAlreadyExistsConflict, err).WithClues(ctx)
	}

	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating documentLibrary list")
	}

	// drive information is not returned by the list creation.
	drive, err := builder.
		ByListId(ptr.Val(newList.GetId())).
		Drive().
		Get(ctx, nil)

	return drive, graph.Wrap(ctx, err, "fetching created documentLibrary").OrNil()
}
