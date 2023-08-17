package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

// GroupsCompat returns a list of groups in the specified M365 tenant.
func GroupsCompat(ctx context.Context, acct account.Account) ([]*Group, error) {
	errs := fault.New(true)

	us, err := Groups(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	return us, errs.Failure()
}

// Groups returns a list of groups in the specified M365 tenant
func Groups(
	ctx context.Context,
	acct account.Account,
	errs *fault.Bus,
) ([]*Group, error) {
	ac, err := makeAC(ctx, acct, path.GroupsService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return getAllGroups(ctx, ac.Groups())
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

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// parseUser extracts information from `models.Userable` we care about
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
