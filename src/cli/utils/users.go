package utils

import (
	"context"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// usersMap retrieves all users in the tenant and returns them in an idname.Cacher
func UsersMap(
	ctx context.Context,
	acct account.Account,
	errs *fault.Bus,
) (idname.Cacher, error) {
	au, err := makeUserAPI(acct)
	if err != nil {
		return nil, clues.Wrap(err, "constructing a graph client")
	}

	users, err := au.GetAll(ctx, errs)
	if err != nil {
		return nil, clues.Wrap(err, "getting all users")
	}

	idToName := make(map[string]string, len(users))

	for _, u := range users {
		id := strings.ToLower(ptr.Val(u.GetId()))
		name := strings.ToLower(ptr.Val(u.GetUserPrincipalName()))

		idToName[id] = name
	}

	return idname.NewCache(idToName), nil
}

func makeUserAPI(acct account.Account) (api.Users, error) {
	creds, err := acct.M365Config()
	if err != nil {
		return api.Users{}, clues.Wrap(err, "getting m365 account creds")
	}

	cli, err := api.NewClient(creds)
	if err != nil {
		return api.Users{}, clues.Wrap(err, "constructing api client")
	}

	return cli.Users(), nil
}
