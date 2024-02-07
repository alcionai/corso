package utils

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/common/idname"
	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

// UsersMap retrieves all users in the tenant and returns them in an idname.Cacher
func UsersMap(
	ctx context.Context,
	acct account.Account,
	co control.Options,
	counter *count.Bus,
	errs *fault.Bus,
) (idname.Cacher, error) {
	au, err := makeUserAPI(acct, co, counter)
	if err != nil {
		return nil, clues.Wrap(err, "constructing a graph client")
	}

	return au.GetAllIDsAndNames(ctx, errs)
}

func makeUserAPI(
	acct account.Account,
	co control.Options,
	counter *count.Bus,
) (api.Users, error) {
	creds, err := acct.M365Config()
	if err != nil {
		return api.Users{}, clues.Wrap(err, "getting m365 account creds")
	}

	cli, err := api.NewClient(creds, co, counter)
	if err != nil {
		return api.Users{}, clues.Wrap(err, "constructing api client")
	}

	return cli.Users(), nil
}
