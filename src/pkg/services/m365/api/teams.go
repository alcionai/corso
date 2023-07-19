package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/common/tform"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

const teamService = "Team"

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Groups() Groups {
	return Groups{c}
}

// On creation of each Microsoft Teams a corrsponding group gets created from them.
// Most of the information like events, drive and mail info will be fetched directly
// from groups. So we pull in group and process only the once which are associated with
// a team for further proccessing of teams.

// Teams is an interface-compliant provider of the client.
type Groups struct {
	Client
}

// GetAll retrieves all groups.
func (c Groups) GetAll(
	ctx context.Context,
	filterTeams bool,
	errs *fault.Bus,
) ([]models.Groupable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	return getGroups(ctx, filterTeams, errs, service)
}

// GetAll retrieves all groups.
func getGroups(
	ctx context.Context,
	filterTeams bool,
	errs *fault.Bus,
	service graph.Servicer,
) ([]models.Groupable, error) {

	resp, err := service.Client().Groups().Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all groups")
	}

	iter, err := msgraphgocore.NewPageIterator[models.Groupable](
		resp,
		service.Adapter(),
		models.CreateTeamCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating groups iterator")
	}

	var (
		groups = make([]models.Groupable, 0)
		el     = errs.Local()
	)

	iterator := func(item models.Groupable) bool {
		if el.Failure() != nil {
			return false
		}

		err := ValidateGroup(item)
		if err != nil {
			el.AddRecoverable(ctx, graph.Wrap(ctx, err, "validating groups"))
		} else {
			isTeam := IsTeam(item)
			if !filterTeams || isTeam {
				groups = append(groups, item)
			}
		}

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, graph.Wrap(ctx, err, "iterating all groups")
	}

	return groups, el.Failure()
}

func IsTeam(g models.Groupable) bool {
	if g.GetAdditionalData()["resourceProvisioningOptions"] != nil {
		val, _ := tform.AnyValueToT[[]any]("resourceProvisioningOptions", g.GetAdditionalData())
		for _, v := range val {
			s, err := str.AnyToString(v)
			if err != nil {
				return false
			}
			if s == teamService {
				return true
			}
		}
	}
	return false
}

// GetID retrieves team by groupID/teamID.
func (c Groups) GetByID(
	ctx context.Context,
	identifier string,
) (models.Groupable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	resp, err := service.Client().Groups().ByGroupId(identifier).Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting group by ID")
	}

	if err != nil {
		err := graph.Wrap(ctx, err, "getting teams by id")

		// TODO: check if its applicable here
		if graph.IsErrItemNotFound(err) {
			err = clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		return nil, err
	}

	return resp, err
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// ValidateGroup ensures the item is a Groupable, and contains the necessary
// identifiers that we handle with all groups.
func ValidateGroup(item models.Groupable) error {
	if item.GetId() == nil {
		return clues.New("missing ID")
	}

	if item.GetDisplayName() == nil {
		return clues.New("missing principalName")
	}

	return nil
}
