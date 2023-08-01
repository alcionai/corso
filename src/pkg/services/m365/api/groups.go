package api

import (
	"context"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/common/tform"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	teamsAdditionalDataLabel    = "Team"
	ResourceProvisioningOptions = "resourceProvisioningOptions"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Groups() Groups {
	return Groups{c}
}

// On creation of each Teams team a corrsponding group gets created.
// The group acts as the protected resource, and all teams data like events,
// drive and mail messages are owned by that group.

// Groups is an interface-compliant provider of the client.
type Groups struct {
	Client
}

// GetAllGroups retrieves all groups.
func (c Groups) GetAll(
	ctx context.Context,
	errs *fault.Bus,
) ([]models.Groupable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	return getGroups(ctx, func(ctx context.Context, g models.Groupable) bool { return true }, errs, service)
}

// GetTeams retrieves all Teams.
func (c Groups) GetTeams(
	ctx context.Context,
	errs *fault.Bus,
) ([]models.Groupable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	return getGroups(ctx, FetchOnlyTeams, errs, service)
}

// GetAll retrieves all groups.
func getGroups(
	ctx context.Context,
	filterGroupsData func(ctx context.Context, g models.Groupable) bool,
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
		models.CreateGroupCollectionResponseFromDiscriminatorValue)
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
			if filterGroupsData(ctx, item) {
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

func FetchOnlyTeams(ctx context.Context, g models.Groupable) bool {
	log := logger.Ctx(ctx)

	if g.GetAdditionalData()[ResourceProvisioningOptions] != nil {
		val, _ := tform.AnyValueToT[[]any](ResourceProvisioningOptions, g.GetAdditionalData())
		for _, v := range val {
			s, err := str.AnyToString(v)
			if err != nil {
				log.Debug("could not be converted to string value: ", ResourceProvisioningOptions)
				return false
			}

			if s == teamsAdditionalDataLabel {
				return true
			}
		}
	}

	return false
}

// GetID retrieves group by groupID.
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
		err := graph.Wrap(ctx, err, "getting group by id")

		return nil, err
	}

	return resp, graph.Stack(ctx, err).OrNil()
}

// GetTeamByID retrieves group by groupID.
func (c Groups) GetTeamByID(
	ctx context.Context,
	identifier string,
) (models.Groupable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	resp, err := service.Client().Groups().ByGroupId(identifier).Get(ctx, nil)
	if err != nil {
		err := graph.Wrap(ctx, err, "getting group by id")

		return nil, err
	}

	if !FetchOnlyTeams(ctx, resp) {
		err := clues.New("given teamID is not related to any team")

		return nil, err
	}

	return resp, graph.Stack(ctx, err).OrNil()
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
		return clues.New("missing display name")
	}

	return nil
}
