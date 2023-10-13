package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// UserNoInfo is the minimal information required to identify and display a user.
// TODO(pandeyabs): Rename this to User now that `Info` support has been removed.
type UserNoInfo struct {
	PrincipalName string
	ID            string
	Name          string
}

// UsersCompatNoInfo returns a list of users in the specified M365 tenant.
// TODO(pandeyabs): Rename this to Users now that `Info` support has been removed. Would
// need corresponding changes in SDK consumers.
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
	ac, err := makeAC(ctx, acct, path.ExchangeService, count.New())
	if err != nil {
		return false, clues.Stack(err)
	}

	return exchange.IsServiceEnabled(ctx, ac.Users(), userID)
}

func UserGetMailboxInfo(
	ctx context.Context,
	acct account.Account,
	userID string,
) (api.MailboxInfo, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService, count.New())
	if err != nil {
		return api.MailboxInfo{}, clues.Stack(err)
	}

	return exchange.GetMailboxInfo(ctx, ac.Users(), userID)
}

// UserHasDrives returns true if the user has any drives
// false otherwise, and a nil pointer and an error in case of error
func UserHasDrives(ctx context.Context, acct account.Account, userID string) (bool, error) {
	ac, err := makeAC(ctx, acct, path.OneDriveService, count.New())
	if err != nil {
		return false, clues.Stack(err)
	}

	return onedrive.IsServiceEnabled(ctx, ac.Users(), userID)
}

// usersNoInfo returns a list of users in the specified M365 tenant - with no info
func usersNoInfo(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*UserNoInfo, error) {
	ac, err := makeAC(ctx, acct, path.UnknownService, count.New())
	if err != nil {
		return nil, clues.Stack(err)
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

// parseUser extracts information from `models.Userable` we care about
func parseUser(item models.Userable) (*UserNoInfo, error) {
	if item.GetUserPrincipalName() == nil {
		return nil, clues.New("user missing principal name").
			With("user_id", ptr.Val(item.GetId()))
	}

	u := &UserNoInfo{
		PrincipalName: ptr.Val(item.GetUserPrincipalName()),
		ID:            ptr.Val(item.GetId()),
		Name:          ptr.Val(item.GetDisplayName()),
	}

	return u, nil
}
