package m365

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/pkg/fault"
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
func (c client) UsersCompatNoInfo(ctx context.Context) ([]*UserNoInfo, error) {
	errs := fault.New(true)

	us, err := usersNoInfo(ctx, c.ac, errs)
	if err != nil {
		return nil, err
	}

	return us, errs.Failure()
}

// UserHasMailbox returns true if the user has an exchange mailbox enabled
// false otherwise, and a nil pointer and an error in case of error
func (c client) UserHasMailbox(ctx context.Context, userID string) (bool, error) {
	return exchange.IsServiceEnabled(ctx, c.ac.Users(), userID)
}

func (c client) UserGetMailboxInfo(
	ctx context.Context,
	userID string,
) (api.MailboxInfo, error) {
	return exchange.GetMailboxInfo(ctx, c.ac.Users(), userID)
}

// UserHasDrives returns true if the user has any drives
// false otherwise, and a nil pointer and an error in case of error
func (c client) UserHasDrives(ctx context.Context, userID string) (bool, error) {
	return onedrive.IsServiceEnabled(ctx, c.ac.Users(), userID)
}

// usersNoInfo returns a list of users in the specified M365 tenant - with no info
func usersNoInfo(
	ctx context.Context,
	ac api.Client,
	errs *fault.Bus,
) ([]*UserNoInfo, error) {
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

func (c client) UserAssignedLicenses(ctx context.Context, userID string) (int, error) {
	us, err := c.ac.Users().GetByID(
		ctx,
		userID,
		api.CallConfig{Select: api.SelectProps("assignedLicenses")})
	if err != nil {
		return 0, err
	}

	if us.GetAssignedLicenses() != nil {
		for _, license := range us.GetAssignedLicenses() {
			fmt.Println(license.GetSkuId())
		}

		return len(us.GetAssignedLicenses()), nil
	}

	return 0, clues.New("user missing assigned licenses")
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
