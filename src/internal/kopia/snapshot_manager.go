package kopia

import (
	"context"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
)

type snapshotManager interface {
	FindManifests(
		ctx context.Context,
		tags map[string]string,
	) ([]*manifest.EntryMetadata, error)
	LoadSnapshots(ctx context.Context, ids []manifest.ID) ([]*snapshot.Manifest, error)
}
