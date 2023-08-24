package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// User is the minimal information required to identify and display a user.
type User struct {
	PrincipalName string
	ID            string
	Name          string
	Info          api.UserInfo
}

// UserNoInfo is the minimal information required to identify and display a user.
// TODO: Remove this once `UsersCompatNoInfo` is removed
type UserNoInfo struct {
	PrincipalName string
	ID            string
	Name          string
}

// UsersCompat returns a list of users in the specified M365 tenant.
// TODO(ashmrtn): Remove when upstream consumers of the SDK support the fault
// package.
func UsersCompat(ctx context.Context, acct account.Account) ([]*User, error) {
	errs := fault.New(true)

	us, err := Users(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	return us, errs.Failure()
}

// UsersCompatNoInfo returns a list of users in the specified M365 tenant.
// TODO: Remove this once `Info` is removed from the `User` struct and callers
// have switched over
func UsersCompatNoInfo(ctx context.Context, acct account.Account) ([]*UserNoInfo, error) {
	errs := fault.New(true)

	us, err := usersNoInfo(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	return us, errs.Failure()
}

// UserHasMailbox returns true if the user has an exchange mailbox enabled
// false otherwise, and a nil pointer and an error in case of error
func UserHasMailbox(ctx context.Context, acct account.Account, userID string) (bool, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService)
	if err != nil {
		return false, clues.Stack(err).WithClues(ctx)
	}

	_, err = ac.Users().GetMailInbox(ctx, userID)
	if err != nil {
		if err := api.EvaluateMailboxError(err); err != nil {
			return false, clues.Stack(err)
		}

		return false, nil
	}

	return true, nil
}

// UserHasDrives returns true if the user has any drives
// false otherwise, and a nil pointer and an error in case of error
func UserHasDrives(ctx context.Context, acct account.Account, userID string) (bool, error) {
	ac, err := makeAC(ctx, acct, path.OneDriveService)
	if err != nil {
		return false, clues.Stack(err).WithClues(ctx)
	}

	return checkUserHasDrives(ctx, ac.Users(), userID)
}

func checkUserHasDrives(ctx context.Context, dgdd getDefaultDriver, userID string) (bool, error) {
	_, err := dgdd.GetDefaultDrive(ctx, userID)
	if err != nil {
		// we consider this a non-error case, since it
		// answers the question the caller is asking.
		if clues.HasLabel(err, graph.LabelsMysiteNotFound) || clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return false, nil
		}

		if graph.IsErrUserNotFound(err) {
			return false, clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		return false, clues.Stack(err)
	}

	return true, nil
}

// usersNoInfo returns a list of users in the specified M365 tenant - with no info
// TODO: Remove this once we remove `Info` from `Users` and instead rely on the `GetUserInfo` API
// to get user information
func usersNoInfo(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*UserNoInfo, error) {
	ac, err := makeAC(ctx, acct, path.UnknownService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	us, err := ac.Users().GetAll(ctx, errs)
	if err != nil {
		return nil, err
	}

	ret := make([]*UserNoInfo, 0, len(us))

	for _, u := range us {
		pu, err := parseUser(u)
		if err != nil {
			return nil, clues.Wrap(err, "formatting user data")
		}

		puNoInfo := &UserNoInfo{
			PrincipalName: pu.PrincipalName,
			ID:            pu.ID,
			Name:          pu.Name,
		}

		ret = append(ret, puNoInfo)
	}

	return ret, nil
}

// Users returns a list of users in the specified M365 tenant
func Users(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*User, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	us, err := ac.Users().GetAll(ctx, errs)
	if err != nil {
		return nil, err
	}

	ret := make([]*User, 0, len(us))

	for _, u := range us {
		pu, err := parseUser(u)
		if err != nil {
			return nil, clues.Wrap(err, "formatting user data")
		}

		userInfo, err := ac.Users().GetInfo(ctx, pu.ID)
		if err != nil {
			return nil, clues.Wrap(err, "getting user details")
		}

		pu.Info = *userInfo

		ret = append(ret, pu)
	}

	return ret, nil
}

// parseUser extracts information from `models.Userable` we care about
func parseUser(item models.Userable) (*User, error) {
	if item.GetUserPrincipalName() == nil {
		return nil, clues.New("user missing principal name").
			With("user_id", ptr.Val(item.GetId()))
	}

	u := &User{
		PrincipalName: ptr.Val(item.GetUserPrincipalName()),
		ID:            ptr.Val(item.GetId()),
		Name:          ptr.Val(item.GetDisplayName()),
	}

	return u, nil
}

// UserInfo returns the corso-specific set of user metadata.
func GetUserInfo(
	ctx context.Context,
	acct account.Account,
	userID string,
) (*api.UserInfo, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	ui, err := ac.Users().GetInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	return ui, nil
}
