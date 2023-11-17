package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// Group is the minimal information required to identify and display a M365 Group.
type Group struct {
	ID string

	// DisplayName is the human-readable name of the group.  Normally the plaintext name that the
	// user provided when they created the group, or the updated name if it was changed.
	// Ex: displayName: "My Group"
	DisplayName string

	// IsTeam is true if the group qualifies as a Teams resource, and is able to backup and restore
	// teams data.
	IsTeam bool
}

// GroupByID retrieves a specific group.
func (c client) GroupByID(
	ctx context.Context,
	id string,
) (*Group, error) {
	cc := api.CallConfig{}

	g, err := c.ac.Groups().GetByID(ctx, id, cc)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return parseGroup(ctx, g)
}

// GroupsCompat returns a list of groups in the specified M365 tenant.
func (c client) GroupsCompat(ctx context.Context) ([]*Group, error) {
	errs := fault.New(true)

	us, err := c.Groups(ctx, errs)
	if err != nil {
		return nil, err
	}

	return us, errs.Failure()
}

// Groups returns a list of groups in the specified M365 tenant
func (c client) Groups(
	ctx context.Context,
	errs *fault.Bus,
) ([]*Group, error) {
	return getAllGroups(ctx, c.ac.Groups())
}

func getAllGroups(
	ctx context.Context,
	ga getAller[models.Groupable],
) ([]*Group, error) {
	groups, err := ga.GetAll(ctx, fault.New(true))
	if err != nil {
		return nil, clues.Wrap(err, "retrieving groups")
	}

	ret := make([]*Group, 0, len(groups))

	for _, g := range groups {
		t, err := parseGroup(ctx, g)
		if err != nil {
			return nil, clues.Wrap(err, "parsing groups")
		}

		ret = append(ret, t)
	}

	return ret, nil
}

func (c client) SitesInGroup(
	ctx context.Context,
	groupID string,
	errs *fault.Bus,
) ([]*Site, error) {
	sites, err := c.ac.Groups().GetAllSites(ctx, groupID, errs)
	if err != nil {
		return nil, clues.Stack(err)
	}

	result := make([]*Site, 0, len(sites))

	for _, site := range sites {
		result = append(result, ParseSite(ctx, site))
	}

	return result, nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// parseGroup extracts information from `models.Groupable` we care about
func parseGroup(ctx context.Context, mg models.Groupable) (*Group, error) {
	if mg.GetDisplayName() == nil {
		return nil, clues.New("group missing display name").
			With("group_id", ptr.Val(mg.GetId()))
	}

	u := &Group{
		ID:          ptr.Val(mg.GetId()),
		DisplayName: ptr.Val(mg.GetDisplayName()),
		IsTeam:      api.IsTeam(ctx, mg),
	}

	return u, nil
}

// GroupsMap retrieves an id-name cache of all groups in the tenant.
func (c client) GroupsMap(
	ctx context.Context,
	errs *fault.Bus,
) (idname.Cacher, error) {
	groups, err := c.Groups(ctx, errs)
	if err != nil {
		return idname.NewCache(nil), err
	}

	itn := make(map[string]string, len(groups))

	for _, s := range groups {
		itn[s.ID] = s.DisplayName
	}

	return idname.NewCache(itn), nil
}
