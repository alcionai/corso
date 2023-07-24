package api

import (
	"context"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Teams() Teams {
	return Teams{c}
}

// Teams is an interface-compliant provider of the client.
type Teams struct {
	Client
}

// GetAllTeams retrieves all teams.
func (c Teams) GetAll(
	ctx context.Context,
	errs *fault.Bus,
) ([]models.Teamable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	resp, err := service.Client().Teams().Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all teams")
	}

	iter, err := msgraphgocore.NewPageIterator[models.Teamable](
		resp,
		service.Adapter(),
		models.CreateTeamCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating teams iterator")
	}

	var (
		teams = make([]models.Teamable, 0)
		el    = errs.Local()
	)

	iterator := func(item models.Teamable) bool {
		if el.Failure() != nil {
			return false
		}

		err := ValidateTeams(item)
		if err != nil {
			el.AddRecoverable(ctx, graph.Wrap(ctx, err, "validating teams"))
		} else {
			teams = append(teams, item)
		}

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, graph.Wrap(ctx, err, "iterating all teams")
	}

	return teams, el.Failure()
}

// GetID retrieves team by teamID.
func (c Teams) GetByID(
	ctx context.Context,
	identifier string,
) (models.Teamable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	resp, err := service.Client().Teams().ByTeamId(identifier).Get(ctx, nil)
	if err != nil {
		err := graph.Wrap(ctx, err, "getting team by id")

		return nil, err
	}

	return resp, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// ValidateTeams ensures the item is a Teamable, and contains the necessary
// identifiers that we handle with all teams.
func ValidateTeams(item models.Teamable) error {
	if item.GetId() == nil {
		return clues.New("missing ID")
	}

	if item.GetDisplayName() == nil {
		return clues.New("missing display name")
	}

	return nil
}
