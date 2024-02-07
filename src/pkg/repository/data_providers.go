package repository

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/internal/m365"
	"github.com/alcionai/canario/src/internal/observe"
	"github.com/alcionai/canario/src/internal/operations/inject"
	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/store"
)

type DataProvider interface {
	inject.BackupProducer
	// Required for backups right now.
	inject.PopulateProtectedResourceIDAndNamer

	inject.ToServiceHandler

	VerifyAccess(ctx context.Context) error
	DeserializeMetadataFiles(
		ctx context.Context,
		colls []data.RestoreCollection,
	) ([]store.MetadataFile, error)
}

type DataProviderConnector interface {
	// ConnectDataProvider initializes configurations
	// and establishes the client connection with the
	// data provider for this operation.
	ConnectDataProvider(
		ctx context.Context,
		pst path.ServiceType,
	) error
	// DataProvider retrieves the data provider.
	DataProvider() DataProvider
}

func (r *repository) DataProvider() DataProvider {
	return r.Provider
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
		err = clues.NewWC(ctx, "unrecognized provider")
	}

	if err != nil {
		return clues.Wrap(err, "connecting data provider")
	}

	if err := provider.VerifyAccess(ctx); err != nil {
		return clues.Wrap(err, fmt.Sprintf("verifying %s account connection", r.Account.Provider))
	}

	r.Provider = provider

	return nil
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

	progressMessage := observe.MessageWithCompletion(ctx, observe.DefaultCfg(), "Connecting to M365")
	defer close(progressMessage)

	ctrl, err := m365.NewController(
		ctx,
		r.Account,
		pst,
		r.Opts,
		r.counter)
	if err != nil {
		return nil, clues.Wrap(err, "creating m365 client controller")
	}

	return ctrl, nil
}
