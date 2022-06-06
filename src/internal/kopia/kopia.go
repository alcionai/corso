package kopia

import (
	"context"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/pkg/storage"
)

const (
	defaultKopiaConfigFilePath = "/tmp/repository.config"

	// TODO(ashmrtnz): These should be some values from upper layer corso,
	// possibly corresponding to who is making the backup.
	kTestHost = "a-test-machine"
	kTestUser = "testUser"
)

var (
	errInit         = errors.New("initializing repo")
	errConnect      = errors.New("connecting repo")
	errNotConnected = errors.New("not connected to repo")
)

type BackupStats struct {
	TotalFileCount      int
	TotalDirectoryCount int
	IgnoredErrorCount   int
	ErrorCount          int
	Incomplete          bool
	IncompleteReason    string
}

func manifestToStats(man *snapshot.Manifest) BackupStats {
	return BackupStats{
		TotalFileCount:      int(man.Stats.TotalFileCount),
		TotalDirectoryCount: int(man.Stats.TotalDirectoryCount),
		IgnoredErrorCount:   int(man.Stats.IgnoredErrorCount),
		ErrorCount:          int(man.Stats.ErrorCount),
		Incomplete:          man.IncompleteReason != "",
		IncompleteReason:    man.IncompleteReason,
	}
}

type KopiaWrapper struct {
	storage storage.Storage
	rep     repo.Repository
}

func New(s storage.Storage) *KopiaWrapper {
	return &KopiaWrapper{storage: s}
}

func (kw *KopiaWrapper) Initialize(ctx context.Context) error {
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

func (kw *KopiaWrapper) Connect(ctx context.Context) error {
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

func (kw *KopiaWrapper) Close(ctx context.Context) error {
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

func (kw *KopiaWrapper) open(ctx context.Context, password string) error {
	// TODO(ashmrtnz): issue #75: nil here should be storage.ConnectionOptions().
	rep, err := repo.Open(ctx, defaultKopiaConfigFilePath, password, nil)
	if err != nil {
		return errors.Wrap(err, "opening repository connection")
	}

	kw.rep = rep
	return nil
}

func inflateDirTree(ctx context.Context, collections []connector.DataCollection) (fs.Directory, error) {
	// TODO(ashmrtnz): Implement when virtualfs.StreamingDirectory is available.
	return virtualfs.NewStaticDirectory("sample-dir", []fs.Entry{}), nil
}

func (kw KopiaWrapper) BackupCollections(
	ctx context.Context,
	collections []connector.DataCollection,
) (*BackupStats, error) {
	if kw.rep == nil {
		return nil, errNotConnected
	}

	dirTree, err := inflateDirTree(ctx, collections)
	if err != nil {
		return nil, errors.Wrap(err, "building kopia directories")
	}

	stats, err := kw.makeSnapshotWithRoot(ctx, dirTree)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (kw KopiaWrapper) makeSnapshotWithRoot(
	ctx context.Context,
	root fs.Directory,
) (*BackupStats, error) {
	si := snapshot.SourceInfo{
		Host:     kTestHost,
		UserName: kTestUser,
		// TODO(ashmrtnz): will this be something useful for snapshot lookups later?
		Path: root.Name(),
	}
	ctx, rw, err := kw.rep.NewWriter(ctx, repo.WriteSessionOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "get repo writer")
	}

	policyTree, err := policy.TreeForSource(ctx, kw.rep, si)
	if err != nil {
		return nil, errors.Wrap(err, "get policy tree")
	}

	u := snapshotfs.NewUploader(rw)

	man, err := u.Upload(ctx, root, policyTree, si)
	if err != nil {
		return nil, errors.Wrap(err, "uploading data")
	}

	if _, err := snapshot.SaveSnapshot(ctx, rw, man); err != nil {
		return nil, errors.Wrap(err, "saving snapshot")
	}

	if err := rw.Flush(ctx); err != nil {
		return nil, errors.Wrap(err, "flushing writer")
	}

	res := manifestToStats(man)
	return &res, nil
}
