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
	pst path.ServiceType,
	counter *count.Bus,
) (api.Client, error) {
	api.InitConcurrencyLimit(ctx, pst)

	creds, err := acct.M365Config()
	if err != nil {
		return api.Client{}, clues.Wrap(err, "getting m365 account creds").WithClues(ctx)
	}

	cli, err := api.NewClient(
		creds,
		control.DefaultOptions(),
		counter)
	if err != nil {
		return api.Client{}, clues.Wrap(err, "constructing api client").WithClues(ctx)
	}

	return cli, nil
}
