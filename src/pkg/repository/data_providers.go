package repository

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

type DataProvider interface {
	// ConnectToM365 establishes graph api connections
	// and initializes api client configurations.
	ConnectToM365(
		ctx context.Context,
		pst path.ServiceType,
	) (*m365.Controller, error)
}

func (r repository) ConnectToM365(
	ctx context.Context,
	pst path.ServiceType,
) (*m365.Controller, error) {
	ctrl, err := connectToM365(ctx, pst, r.Account, r.Opts)
	if err != nil {
		return nil, clues.Wrap(err, "connecting to m365")
	}

	return ctrl, nil
}

var m365nonce bool

func connectToM365(
	ctx context.Context,
	pst path.ServiceType,
	acct account.Account,
	co control.Options,
) (*m365.Controller, error) {
	if !m365nonce {
		m365nonce = true

		progressBar := observe.MessageWithCompletion(ctx, "Connecting to M365")
		defer close(progressBar)
	}

	ctrl, err := m365.NewController(ctx, acct, pst, co)
	if err != nil {
		return nil, err
	}

	return ctrl, nil
}
