package kopia

import (
	"context"
	"os"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/filesystem"

	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

func localFSBlobStorage(
	ctx context.Context,
	repoOpts repository.Options,
	s storage.Storage,
) (blob.Storage, error) {
	opts := filesystem.Options{
		Path: os.Getenv("filesystem_path"),
	}

	store, err := filesystem.New(ctx, &opts, false)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return store, nil
}
