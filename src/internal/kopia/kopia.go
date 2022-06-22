package kopia

import (
	"context"
	"io"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/blob"
	"github.com/kopia/kopia/repo/manifest"
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
	errInit           = errors.New("initializing repo")
	errConnect        = errors.New("connecting repo")
	errNotConnected   = errors.New("not connected to repo")
	errUnsupportedDir = errors.New("unsupported static children in streaming directory")
)

type BackupStats struct {
	SnapshotID          string
	TotalFileCount      int
	TotalDirectoryCount int
	IgnoredErrorCount   int
	ErrorCount          int
	Incomplete          bool
	IncompleteReason    string
}

func manifestToStats(man *snapshot.Manifest) BackupStats {
	return BackupStats{
		SnapshotID:          string(man.ID),
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

// getStreamItemFunc returns a function that can be used by kopia's
// virtualfs.StreamingDirectory to iterate through directory entries and call
// kopia callbacks on directory entries. It binds the directory to the given
// DataCollection.
func getStreamItemFunc(
	collection connector.DataCollection,
) func(context.Context, func(context.Context, fs.Entry) error) error {
	return func(ctx context.Context, cb func(context.Context, fs.Entry) error) error {
		for {
			e, err := collection.NextItem()
			if err != nil {
				if err == io.EOF {
					return nil
				}

				return errors.Wrap(err, "materializing directory entry")
			}

			entry := virtualfs.StreamingFileFromReader(e.UUID(), e.ToReader())
			if err = cb(ctx, entry); err != nil {
				return errors.Wrap(err, "executing callback")
			}
		}
	}
}

// buildKopiaDirs recursively builds a directory hierarchy from the roots up.
// Returned directories are either virtualfs.StreamingDirectory or
// virtualfs.staticDirectory.
func buildKopiaDirs(dirName string, dir *treeMap) (fs.Directory, error) {
	// Don't support directories that have both a DataCollection and a set of
	// static child directories.
	if dir.collection != nil && len(dir.childDirs) > 0 {
		return nil, errors.New(errUnsupportedDir.Error())
	}

	if dir.collection != nil {
		return virtualfs.NewStreamingDirectory(dirName, getStreamItemFunc(dir.collection)), nil
	}

	// Need to build the directory tree from the leaves up because intermediate
	// directories need to have all their entries at creation time.
	childDirs := []fs.Entry{}

	for childName, childDir := range dir.childDirs {
		child, err := buildKopiaDirs(childName, childDir)
		if err != nil {
			return nil, err
		}

		childDirs = append(childDirs, child)
	}

	return virtualfs.NewStaticDirectory(dirName, childDirs), nil
}

type treeMap struct {
	childDirs  map[string]*treeMap
	collection connector.DataCollection
}

func newTreeMap() *treeMap {
	return &treeMap{
		childDirs: map[string]*treeMap{},
	}
}

// inflateDirTree returns an fs.Directory tree rooted at the oldest common
// ancestor of the streams and uses virtualfs.StaticDirectory for internal nodes
// in the hierarchy. Leaf nodes are virtualfs.StreamingDirectory with the given
// DataCollections.
func inflateDirTree(ctx context.Context, collections []connector.DataCollection) (fs.Directory, error) {
	roots := make(map[string]*treeMap)

	for _, s := range collections {
		path := s.FullPath()

		if len(path) == 0 {
			return nil, errors.New("no identifier for collection")
		}

		dir, ok := roots[path[0]]
		if !ok {
			dir = newTreeMap()
			roots[path[0]] = dir
		}

		// Single DataCollection with no ancestors.
		if len(path) == 1 {
			dir.collection = s
			continue
		}

		for _, p := range path[1 : len(path)-1] {
			newDir, ok := dir.childDirs[p]
			if !ok {
				newDir = newTreeMap()

				if dir.childDirs == nil {
					dir.childDirs = map[string]*treeMap{}
				}

				dir.childDirs[p] = newDir
			}

			dir = newDir
		}

		// At this point we have all the ancestor directories of this DataCollection
		// as treeMap objects and `dir` is the parent directory of this
		// DataCollection.

		end := len(path) - 1

		// Make sure this entry doesn't already exist.
		if _, ok := dir.childDirs[path[end]]; ok {
			return nil, errors.New(errUnsupportedDir.Error())
		}

		sd := newTreeMap()
		sd.collection = s
		dir.childDirs[path[end]] = sd
	}

	if len(roots) > 1 {
		return nil, errors.New("multiple root directories")
	}

	var res fs.Directory
	for dirName, dir := range roots {
		tmp, err := buildKopiaDirs(dirName, dir)
		if err != nil {
			return nil, err
		}

		res = tmp
	}

	return res, nil
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

// RestoreSingleItem looks up the item at the given path in the snapshot with id
// snapshotID. The path should be the full path of the item from the root.
// If the item is a file in kopia then it returns a DataCollection with the item
// as its sole element and DataCollection.FullPath() set to
// split(dirname(itemPath), "/"). If the item does not exist in kopia or is not
// a file an error is returned. The UUID of the returned DataStreams will be the
// name of the kopia file the data is sourced from.
func (kw KopiaWrapper) RestoreSingleItem(
	ctx context.Context,
	snapshotID string,
	itemPath []string,
) (connector.DataCollection, error) {
	manifest, err := snapshot.LoadSnapshot(ctx, kw.rep, manifest.ID(snapshotID))
	if err != nil {
		return nil, errors.Wrap(err, "getting snapshot handle")
	}

	rootDirEntry, err := snapshotfs.SnapshotRoot(kw.rep, manifest)
	if err != nil {
		return nil, errors.Wrap(err, "getting root directory")
	}

	// Fine if rootDirEntry is nil, will be checked in called function.
	return kw.restoreSingleItem(ctx, rootDirEntry, itemPath[1:])
}

// restoreSingleItem looks up the item at the given path starting from rootDir
// where rootDir is the root of a snapshot. If the item is a file in kopia then
// it returns a DataCollection with the item as its sole element and
// DataCollection.FullPath() set to split(dirname(itemPath), "/"). If the item
// does not exist in kopia or is not a file an error is returned. The UUID of
// the returned DataStreams will be the name of the kopia file the data is
// sourced from.
func (kw KopiaWrapper) restoreSingleItem(
	ctx context.Context,
	rootDir fs.Entry,
	itemPath []string,
) (connector.DataCollection, error) {
	e, err := snapshotfs.GetNestedEntry(ctx, rootDir, itemPath)
	if err != nil {
		return nil, errors.Wrap(err, "getting object handle")
	}

	f, ok := e.(fs.File)
	if !ok {
		return nil, errors.New("not a file")
	}

	r, err := f.Open(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}

	pathWithRoot := []string{rootDir.Name()}
	pathWithRoot = append(pathWithRoot, itemPath[:len(itemPath)-1]...)

	return &singleItemCollection{
		stream: kopiaDataStream{
			uuid:   itemPath[len(itemPath)-1],
			reader: r,
		},
		path: pathWithRoot,
	}, nil
}
