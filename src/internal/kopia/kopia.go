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

	if err := kw.open(ctx, cfg.CorsoPassword); err != nil {
		return err
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

	if err := kw.open(ctx, cfg.CorsoPassword); err != nil {
		return err
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

	err := kw.rep.Close(ctx)
	kw.rep = nil

	if err != nil {
		return errors.Wrap(err, "closing repository connection")
	}

	return nil
}

func (kw kopiaWrapper) open(ctx context.Context, password string) error {
	// TODO(ashmrtnz): issue #75: nil here should be storage.ConnectionOptions().
	rep, err := repo.Open(ctx, defaultKopiaConfigFilePath, password, nil)
	if err != nil {
		return errors.Wrap(err, "opening repository connection")
	}

	kw.rep = rep
	return nil
}
