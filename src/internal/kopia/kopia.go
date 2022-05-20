package kopia

import (
	"context"

	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/pkg/errors"
)

const (
	defaultKopiaConfigFilePath = "/tmp/repository.config"
	defaultKopiaConfigPasswd   = "todo:passwd"
)

var (
	errInit    = errors.New("initializing repo")
	errConnect = errors.New("connecting repo")
)

type StorageMaker interface {
	BlobStorage(ctx context.Context, create bool) (blob.Storage, error)
}

type KopiaWrapper struct {
	storage              blob.Storage
	newRepositoryOptions *repo.NewRepositoryOptions
	connectOptions       *repo.ConnectOptions
}

func NewInitializer(
	ctx context.Context,
	sm StorageMaker,
	newOpts *repo.NewRepositoryOptions,
	connOpts *repo.ConnectOptions,
) (KopiaWrapper, error) {
	return newInitConnector(ctx, sm, true, newOpts, connOpts)
}

func NewConnector(ctx context.Context, sm StorageMaker, opts *repo.ConnectOptions) (KopiaWrapper, error) {
	return newInitConnector(ctx, sm, false, nil, opts)
}

func newInitConnector(
	ctx context.Context,
	sm StorageMaker,
	createStorage bool,
	newOpts *repo.NewRepositoryOptions,
	connOpts *repo.ConnectOptions,
) (KopiaWrapper, error) {
	bst, err := sm.BlobStorage(ctx, createStorage)
	if err != nil {
		return KopiaWrapper{}, err
	}
	return KopiaWrapper{bst, newOpts, connOpts}, nil
}

func (kw KopiaWrapper) Initialize(ctx context.Context) error {
	if err := repo.Initialize(
		ctx,
		kw.storage,
		kw.newRepositoryOptions,
		defaultKopiaConfigPasswd,
	); err != nil {
		return errors.Wrap(err, errInit.Error())
	}

	if err := repo.Connect(
		ctx,
		defaultKopiaConfigFilePath,
		kw.storage,
		defaultKopiaConfigPasswd,
		kw.connectOptions,
	); err != nil {
		return errors.Wrap(err, errConnect.Error())
	}

	return nil
}

func (kw KopiaWrapper) Connect(ctx context.Context) error {
	if err := repo.Connect(
		ctx,
		defaultKopiaConfigFilePath,
		kw.storage,
		defaultKopiaConfigPasswd,
		kw.connectOptions,
	); err != nil {
		return errors.Wrap(err, errConnect.Error())
	}
	return nil
}
