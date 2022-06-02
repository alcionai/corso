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
	errInit    = errors.New("initializing repo")
	errConnect = errors.New("connecting repo")
	errOpen    = errors.New("open repo")
	errClose   = errors.New("close repo")
)

type kopiaWrapper struct {
	storage storage.Storage
	rep     repo.Repository
}

func New(s storage.Storage) kopiaWrapper {
	return kopiaWrapper{storage: s}
}

func (kw kopiaWrapper) Initialize(ctx context.Context) error {
	bst, err := blobStoreByProvider(ctx, kw.storage)
	if err != nil {
		return errors.Wrap(err, errInit.Error())
	}
	defer bst.Close(ctx)

	cfg, err := kw.storage.CommonConfig()
	if err != nil {
		return err
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

	cfg, err := kw.storage.CommonConfig()
	if err != nil {
		return err
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
		return s3BlobStorage(ctx, s)
	default:
		return nil, errors.New("storage provider details are required")
	}
}

func (kw kopiaWrapper) Close(ctx context.Context) error {
	if kw.rep == nil {
		return nil
	}

	if err := kw.rep.Close(ctx); err != nil {
		return errors.Wrap(err, errClose.Error())
	}

	kw.rep = nil
	return nil
}

func (kw kopiaWrapper) open(ctx context.Context) error {
	cfg := kw.storage.CommonConfig()
	if len(cfg.CorsoPassword) == 0 {
		return errRequriesPassword
	}

	// TODO(ashmrtnz): issue #75: nil here should be storage.ConnectionOptions().
	rep, err := repo.Open(ctx, defaultKopiaConfigFilePath, cfg.CorsoPassword, nil)
	if err != nil {
		return errors.Wrap(err, errOpen.Error())
	}

	kw.rep = rep
	return nil
}
