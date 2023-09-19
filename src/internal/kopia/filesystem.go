package kopia

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/filesystem"

	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

func filesystemStorage(
	ctx context.Context,
	repoOpts repository.Options,
	s storage.Storage,
) (blob.Storage, error) {
	cfg, err := s.StorageConfig()
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	fsCfg := cfg.(*storage.FilesystemConfig)
	opts := filesystem.Options{
		Path: fsCfg.Path,
	}

	store, err := filesystem.New(ctx, &opts, true)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return store, nil
}
