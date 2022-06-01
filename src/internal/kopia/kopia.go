package kopia

import (
	"context"

	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/pkg/storage"
)

const (
	defaultKopiaConfigFilePath = "/tmp/repository.config"
)

var (
	errInit             = errors.New("initializing repo")
	errConnect          = errors.New("connecting repo")
	errRequriesPassword = errors.New("corso password required")
)

type kopiaWrapper struct {
	storage storage.Storage
}

func New(s storage.Storage) kopiaWrapper {
	return kopiaWrapper{s}
}

func (kw kopiaWrapper) Initialize(ctx context.Context) error {
	bst, err := blobStoreByProvider(ctx, kw.storage)
	if err != nil {
		return errors.Wrap(err, errInit.Error())
	}
	defer bst.Close(ctx)

	cfg := kw.storage.CommonConfig()
	if len(cfg.CorsoPassword) == 0 {
		return errRequriesPassword
	}

	// todo - issue #75: nil here should be a storage.NewRepoOptions()
	if err = repo.Initialize(ctx, bst, nil, cfg.CorsoPassword); err != nil {
		return errors.Wrap(err, errInit.Error())
	}

	// todo - issue #75: nil here should be a storage.ConnectOptions()
	if err := repo.Connect(
		ctx,
		defaultKopiaConfigFilePath,
		bst,
		cfg.CorsoPassword,
		nil,
	); err != nil {
		return errors.Wrap(err, errConnect.Error())
	}

	return nil
}

func (kw kopiaWrapper) Connect(ctx context.Context) error {
	bst, err := blobStoreByProvider(ctx, kw.storage)
	if err != nil {
		return errors.Wrap(err, errInit.Error())
	}
	defer bst.Close(ctx)

	cfg := kw.storage.CommonConfig()
	if len(cfg.CorsoPassword) == 0 {
		return errRequriesPassword
	}

	// todo - issue #75: nil here should be storage.ConnectOptions()
	if err := repo.Connect(
		ctx,
		defaultKopiaConfigFilePath,
		bst,
		cfg.CorsoPassword,
		nil,
	); err != nil {
		return errors.Wrap(err, errConnect.Error())
	}
	return nil
}

func blobStoreByProvider(ctx context.Context, s storage.Storage) (blob.Storage, error) {
	switch s.Provider {
	case storage.ProviderS3:
		return s3BlobStorage(ctx, s.S3Config())
	default:
		return nil, errors.New("storage provider details are required")
	}
}
