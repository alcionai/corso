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

type KopiaWrapper struct {
	storage              blob.Storage
	newRepositoryOptions *repo.NewRepositoryOptions
	connectOptions       *repo.ConnectOptions
}

func NewInitializer(
	st blob.Storage,
	newOpts *repo.NewRepositoryOptions,
	connOpts *repo.ConnectOptions,
) KopiaWrapper {
	return KopiaWrapper{
		storage:              st,
		newRepositoryOptions: newOpts,
		connectOptions:       connOpts,
	}
}

func NewConnector(st blob.Storage, opts *repo.ConnectOptions) KopiaWrapper {
	return KopiaWrapper{
		storage:        st,
		connectOptions: opts,
	}
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

	if err := kw.Connect(ctx); err != nil {
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
