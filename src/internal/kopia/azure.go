package kopia

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/blob/azure"

	"github.com/alcionai/corso/src/pkg/storage"
)

func azBlobStorage(ctx context.Context, s storage.Storage) (blob.Storage, error) {
	cfg, err := s.AzureConfig()
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	opts := azure.Options{
		Container: cfg.Container,
		Prefix:    cfg.Prefix,
	}

	store, err := azure.New(ctx, &opts, false)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return store, nil
}
