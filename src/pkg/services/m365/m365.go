package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type client struct {
	ac api.Client
}

func NewM365Client(
	ctx context.Context,
	acct account.Account,
) (client, error) {
	ac, err := makeAC(ctx, acct)
	return client{ac}, clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// interfaces & structs
// ---------------------------------------------------------------------------

type getAller[T any] interface {
	GetAll(ctx context.Context, errs *fault.Bus) ([]T, error)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func makeAC(
	ctx context.Context,
	acct account.Account,
) (api.Client, error) {
	// exchange service inits a limit to concurrency.
	api.InitConcurrencyLimit(ctx, path.ExchangeService)

	creds, err := acct.M365Config()
	if err != nil {
		return api.Client{}, clues.WrapWC(ctx, err, "getting m365 account creds")
	}

	cli, err := api.NewClient(
		creds,
		control.DefaultOptions(),
		count.New())
	if err != nil {
		return api.Client{}, clues.WrapWC(ctx, err, "constructing api client")
	}

	// run a test to ensure credentials work for the client
	if err := cli.Access().GetToken(ctx); err != nil {
		return api.Client{}, clues.Wrap(err, "checking client connection")
	}

	return cli, nil
}
