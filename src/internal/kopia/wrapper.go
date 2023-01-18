package kopia

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/policy"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	// TODO(ashmrtnz): These should be some values from upper layer corso,
	// possibly corresponding to who is making the backup.
	corsoHost = "corso-host"
	corsoUser = "corso"

	serializationVersion uint32 = 1
)

// common manifest tags
const (
	TagBackupID       = "backup-id"
	TagBackupCategory = "is-canon-backup"
)

var (
	errNotConnected  = errors.New("not connected to repo")
	errNoRestorePath = errors.New("no restore path given")
)

type BackupStats struct {
	SnapshotID string

	TotalHashedBytes   int64
	TotalUploadedBytes int64

	TotalFileCount      int
	CachedFileCount     int
	UncachedFileCount   int
	TotalDirectoryCount int
	IgnoredErrorCount   int
	ErrorCount          int

	Incomplete       bool
	IncompleteReason string
}

func manifestToStats(
	man *snapshot.Manifest,
	progress *corsoProgress,
	uploadCount *stats.ByteCounter,
) BackupStats {
	return BackupStats{
		SnapshotID: string(man.ID),

		TotalHashedBytes:   progress.totalBytes,
		TotalUploadedBytes: uploadCount.NumBytes,

		TotalFileCount:      int(man.Stats.TotalFileCount),
		CachedFileCount:     int(man.Stats.CachedFiles),
		UncachedFileCount:   int(man.Stats.NonCachedFiles),
		TotalDirectoryCount: int(man.Stats.TotalDirectoryCount),
		IgnoredErrorCount:   int(man.Stats.IgnoredErrorCount),
		ErrorCount:          int(man.Stats.ErrorCount),

		Incomplete:       man.IncompleteReason != "",
		IncompleteReason: man.IncompleteReason,
	}
}

func NewWrapper(c *conn) (*Wrapper, error) {
	if err := c.wrap(); err != nil {
		return nil, errors.Wrap(err, "creating Wrapper")
	}

	return &Wrapper{c}, nil
}

type Wrapper struct {
	c *conn
}

func (w *Wrapper) Close(ctx context.Context) error {
	if w.c == nil {
		return nil
	}

	err := w.c.Close(ctx)
	w.c = nil

	return errors.Wrap(err, "closing Wrapper")
}

type IncrementalBase struct {
	*snapshot.Manifest
	SubtreePaths []*path.Builder
}

// BackupCollections takes a set of collections and creates a kopia snapshot
// with the data that they contain. previousSnapshots is used for incremental
// backups and should represent the base snapshot from which metadata is sourced
// from as well as any incomplete snapshot checkpoints that may contain more
// recent data than the base snapshot. The absence of previousSnapshots causes a
// complete backup of all data.
func (w Wrapper) BackupCollections(
	ctx context.Context,
	previousSnapshots []IncrementalBase,
	collections []data.Collection,
	tags map[string]string,
	buildTreeWithBase bool,
) (*BackupStats, *details.Builder, map[string]path.Path, error) {
	if w.c == nil {
		return nil, nil, nil, errNotConnected
	}

	ctx, end := D.Span(ctx, "kopia:backupCollections")
	defer end()

	if len(collections) == 0 {
		return &BackupStats{}, &details.Builder{}, nil, nil
	}

	progress := &corsoProgress{
		pending: map[string]*itemDetails{},
		deets:   &details.Builder{},
		toMerge: map[string]path.Path{},
	}

	// When running an incremental backup, we need to pass the prior
	// snapshot bases into inflateDirTree so that the new snapshot
	// includes historical data.
	var base []IncrementalBase
	if buildTreeWithBase {
		base = previousSnapshots
	}

	dirTree, err := inflateDirTree(ctx, w.c, base, collections, progress)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "building kopia directories")
	}

	s, err := w.makeSnapshotWithRoot(
		ctx,
		previousSnapshots,
		dirTree,
		tags,
		progress,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return s, progress.deets, progress.toMerge, nil
}

func (w Wrapper) makeSnapshotWithRoot(
	ctx context.Context,
	prevSnapEntries []IncrementalBase,
	root fs.Directory,
	addlTags map[string]string,
	progress *corsoProgress,
) (*BackupStats, error) {
	var (
		man *snapshot.Manifest
		bc  = &stats.ByteCounter{}
	)

	prevSnaps := make([]*snapshot.Manifest, 0, len(prevSnapEntries))
	for _, ent := range prevSnapEntries {
		prevSnaps = append(prevSnaps, ent.Manifest)
	}

	err := repo.WriteSession(
		ctx,
		w.c,
		repo.WriteSessionOptions{
			Purpose: "KopiaWrapperBackup",
			// Always flush so we don't leak write sessions. Still uses reachability
			// for consistency.
			FlushOnFailure: true,
			OnUpload:       bc.Count,
		},
		func(innerCtx context.Context, rw repo.RepositoryWriter) error {
			si := snapshot.SourceInfo{
				Host:     corsoHost,
				UserName: corsoUser,
				// TODO(ashmrtnz): will this be something useful for snapshot lookups later?
				Path: root.Name(),
			}

			trueVal := policy.OptionalBool(true)
			errPolicy := &policy.Policy{
				ErrorHandlingPolicy: policy.ErrorHandlingPolicy{
					IgnoreFileErrors:      &trueVal,
					IgnoreDirectoryErrors: &trueVal,
				},
			}
			policyTree, err := policy.TreeForSourceWithOverride(innerCtx, w.c, si, errPolicy)
			if err != nil {
				err = errors.Wrap(err, "get policy tree")
				logger.Ctx(innerCtx).Errorw("kopia backup", err)
				return err
			}

			// By default Uploader is best-attempt.
			u := snapshotfs.NewUploader(rw)
			progress.UploadProgress = u.Progress
			u.Progress = progress

			man, err = u.Upload(innerCtx, root, policyTree, si, prevSnaps...)
			if err != nil {
				err = errors.Wrap(err, "uploading data")
				logger.Ctx(innerCtx).Errorw("kopia backup", err)
				return err
			}

			man.Tags = map[string]string{}

			for k, v := range addlTags {
				mk, mv := makeTagKV(k)

				if len(v) == 0 {
					v = mv
				}

				man.Tags[mk] = v
			}

			if _, err := snapshot.SaveSnapshot(innerCtx, rw, man); err != nil {
				err = errors.Wrap(err, "saving snapshot")
				logger.Ctx(innerCtx).Errorw("kopia backup", err)
				return err
			}

			return nil
		},
	)
	// Telling kopia to always flush may hide other errors if it fails while
	// flushing the write session (hence logging above).
	if err != nil {
		return nil, errors.Wrap(err, "kopia backup")
	}

	res := manifestToStats(man, progress, bc)

	return &res, nil
}

func (w Wrapper) getSnapshotRoot(
	ctx context.Context,
	snapshotID string,
) (fs.Entry, error) {
	man, err := snapshot.LoadSnapshot(ctx, w.c, manifest.ID(snapshotID))
	if err != nil {
		return nil, errors.Wrap(err, "getting snapshot handle")
	}

	rootDirEntry, err := snapshotfs.SnapshotRoot(w.c, man)

	return rootDirEntry, errors.Wrap(err, "getting root directory")
}

// getItemStream looks up the item at the given path starting from snapshotRoot.
// If the item is a file in kopia then it returns a data.Stream of the item. If
// the item does not exist in kopia or is not a file an error is returned. The
// UUID of the returned data.Stream will be the name of the kopia file the data
// is sourced from.
func getItemStream(
	ctx context.Context,
	itemPath path.Path,
	snapshotRoot fs.Entry,
	bcounter ByteCounter,
) (data.Stream, error) {
	if itemPath == nil {
		return nil, errors.WithStack(errNoRestorePath)
	}

	// GetNestedEntry handles nil properly.
	e, err := snapshotfs.GetNestedEntry(
		ctx,
		snapshotRoot,
		encodeElements(itemPath.PopFront().Elements()...),
	)
	if err != nil {
		if strings.Contains(err.Error(), "entry not found") {
			err = errors.Wrap(ErrNotFound, err.Error())
		}

		return nil, errors.Wrap(err, "getting nested object handle")
	}

	f, ok := e.(fs.File)
	if !ok {
		return nil, errors.New("requested object is not a file")
	}

	if bcounter != nil {
		bcounter.Count(f.Size())
	}

	r, err := f.Open(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}

	decodedName, err := decodeElement(f.Name())
	if err != nil {
		return nil, errors.Wrap(err, "decoding file name")
	}

	return &kopiaDataStream{
		uuid: decodedName,
		reader: &restoreStreamReader{
			ReadCloser:      r,
			expectedVersion: serializationVersion,
		},
		size: f.Size() - int64(versionSize),
	}, nil
}

type ByteCounter interface {
	Count(numBytes int64)
}

// RestoreMultipleItems looks up all paths- assuming each is an item declaration,
// not a directory- in the snapshot with id snapshotID. The path should be the
// full path of the item from the root.  Returns the results as a slice of single-
// item DataCollections, where the DataCollection.FullPath() matches the path.
// If the item does not exist in kopia or is not a file an error is returned.
// The UUID of the returned DataStreams will be the name of the kopia file the
// data is sourced from.
func (w Wrapper) RestoreMultipleItems(
	ctx context.Context,
	snapshotID string,
	paths []path.Path,
	bcounter ByteCounter,
) ([]data.Collection, error) {
	ctx, end := D.Span(ctx, "kopia:restoreMultipleItems")
	defer end()

	if len(paths) == 0 {
		return nil, errors.WithStack(errNoRestorePath)
	}

	snapshotRoot, err := w.getSnapshotRoot(ctx, snapshotID)
	if err != nil {
		return nil, err
	}

	var (
		errs *multierror.Error
		// Maps short ID of parent path to data collection for that folder.
		cols = map[string]*kopiaDataCollection{}
	)

	// TODO(meain): We have to update the details that is getting
	// stored to only contain the filenames, as of now it contains
	// every file that gets saved. Once we do this, we will have to
	// have the component augment the list of the files here while
	// restoring to make sure we add back in all the data and metadata
	// files that we need to form the collections.

	// NB: We have to make sure to order the files that gets fetched
	// by kopia to have it pull the metadata file for the folder
	// first, and for the metadata files for the file to be pulled
	// right after the file for it to be pulled.

	// TODO(meain) If we end converting the reads to parallel reads,
	// how would it affect files showing up at the other side or
	// should be make kopia guarantee that it will send us the file in
	// the order that we send the details list in.

	// TODO(meain): details things can avoid storing info about
	// .dirmeta and we can compute them from the list of files. The
	// only issue is that we might not be able to restore empty folder
	// as we will not be aware of them without some mention about
	// them.

	// This sort is done primarily to order .meta files after .data
	// files
	// TODO(meain): This code should be moved into the OneDrive component
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].String() < paths[j].String()
	})

	for _, itemPath := range paths {
		ds, err := getItemStream(ctx, itemPath, snapshotRoot, bcounter)
		if err != nil {
			errs = multierror.Append(errs, err)
			continue
		}

		parentPath, err := itemPath.Dir()
		if err != nil {
			errs = multierror.Append(errs, errors.Wrap(err, "making directory collection"))
			continue
		}

		c, ok := cols[parentPath.ShortRef()]
		if !ok {
			cols[parentPath.ShortRef()] = &kopiaDataCollection{path: parentPath}
			c = cols[parentPath.ShortRef()]
		}

		c.streams = append(c.streams, ds)
	}

	res := make([]data.Collection, 0, len(cols))
	for _, c := range cols {
		res = append(res, c)
	}

	return res, errs.ErrorOrNil()
}

// DeleteSnapshot removes the provided manifest from kopia.
func (w Wrapper) DeleteSnapshot(
	ctx context.Context,
	snapshotID string,
) error {
	mid := manifest.ID(snapshotID)

	if len(mid) == 0 {
		return errors.New("attempt to delete unidentified snapshot")
	}

	err := repo.WriteSession(
		ctx,
		w.c,
		repo.WriteSessionOptions{Purpose: "KopiaWrapperBackupDeletion"},
		func(innerCtx context.Context, rw repo.RepositoryWriter) error {
			if err := rw.DeleteManifest(ctx, mid); err != nil {
				return errors.Wrap(err, "deleting snapshot")
			}

			return nil
		},
	)
	// Telling kopia to always flush may hide other errors if it fails while
	// flushing the write session (hence logging above).
	if err != nil {
		return errors.Wrap(err, "kopia deleting backup manifest")
	}

	return nil
}

// FetchPrevSnapshotManifests returns a set of manifests for complete and maybe
// incomplete snapshots for the given (resource owner, service, category)
// tuples. Up to two manifests can be returned per tuple: one complete and one
// incomplete. An incomplete manifest may be returned if it is newer than the
// newest complete manifest for the tuple. Manifests are deduped such that if
// multiple tuples match the same manifest it will only be returned once.
// If tags are provided, manifests must include a superset of the k:v pairs
// specified by those tags.  Tags should pass their raw values, and will be
// normalized inside the func using MakeTagKV.
func (w Wrapper) FetchPrevSnapshotManifests(
	ctx context.Context,
	reasons []Reason,
	tags map[string]string,
) ([]*ManifestEntry, error) {
	if w.c == nil {
		return nil, errors.WithStack(errNotConnected)
	}

	return fetchPrevSnapshotManifests(ctx, w.c, reasons, tags), nil
}
