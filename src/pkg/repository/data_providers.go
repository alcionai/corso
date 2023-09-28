package repository

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
)

type DataProvider interface {
	inject.BackupProducer
	inject.ExportConsumer
	inject.RestoreConsumer
}

type DataProviderConnector interface {
	// ConnectDataProvider initializes configurations
	// and establishes the client connection with the
	// data provider for this operation.
	ConnectDataProvider(
		ctx context.Context,
		pst path.ServiceType,
	) error
}

func (r *repository) ConnectDataProvider(
	ctx context.Context,
	pst path.ServiceType,
) error {
	var (
		provider DataProvider
		err      error
	)

	switch r.Account.Provider {
	case account.ProviderM365:
		provider, err = connectToM365(ctx, *r, pst)
	default:
		err = clues.New("unrecognized provider")
	}

	r.Provider = provider

	return clues.Wrap(err, "connecting data provider").
		WithClues(ctx).
		OrNil()
}

func connectToM365(
	ctx context.Context,
	r repository,
	pst path.ServiceType,
) (*m365.Controller, error) {
	if r.Provider != nil {
		ctrl, ok := r.Provider.(*m365.Controller)
		if !ok {
			// if the provider is initialized to a non-m365 controller, we should not
			// attempt to connnect to m365 afterward.
			return nil, clues.New("Attempted to connect to multiple data providers")
		}

		return ctrl, nil
	}

	progressBar := observe.MessageWithCompletion(ctx, "Connecting to M365")
	defer close(progressBar)

	ctrl, err := m365.NewController(ctx, r.Account, pst, r.Opts)
	if err != nil {
		return nil, err
	}

	return ctrl, nil
}
