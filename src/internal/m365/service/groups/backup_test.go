package groups

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// mockRestoreProducer copied from operations/backup_tet.go
type mockRestoreProducer struct {
	gotPaths  []path.Path
	colls     []data.RestoreCollection
	collsByID map[string][]data.RestoreCollection // snapshotID: []RestoreCollection
	err       error
	onRestore restoreFunc
}

type restoreFunc func(
	id string,
	ps []path.RestorePaths,
) ([]data.RestoreCollection, error)

func (mr *mockRestoreProducer) buildRestoreFunc(
	t *testing.T,
	oid string,
	ops []path.Path,
) {
	mr.onRestore = func(
		id string,
		ps []path.RestorePaths,
	) ([]data.RestoreCollection, error) {
		gotPaths := make([]path.Path, 0, len(ps))

		for _, rp := range ps {
			gotPaths = append(gotPaths, rp.StoragePath)
		}

		assert.Equal(t, oid, id, "manifest id")
		checkPaths(t, ops, gotPaths)

		return mr.colls, mr.err
	}
}

func (mr *mockRestoreProducer) ProduceRestoreCollections(
	ctx context.Context,
	snapshotID string,
	paths []path.RestorePaths,
	bc kopia.ByteCounter,
	errs *fault.Bus,
) ([]data.RestoreCollection, error) {
	for _, ps := range paths {
		mr.gotPaths = append(mr.gotPaths, ps.StoragePath)
	}

	if mr.onRestore != nil {
		return mr.onRestore(snapshotID, paths)
	}

	if len(mr.collsByID) > 0 {
		return mr.collsByID[snapshotID], mr.err
	}

	return mr.colls, mr.err
}

func makeMetadataBasePath(
	t *testing.T,
	tenant string,
	service path.ServiceType,
	resourceOwner string,
	category path.CategoryType,
	elems ...string,
) path.Path {
	t.Helper()

	p, err := path.Builder{}.ToServiceCategoryMetadataPath(
		tenant,
		resourceOwner,
		service,
		category,
		false)
	require.NoError(t, err, clues.ToCore(err))

	p, err = p.Append(false, elems...)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func checkPaths(t *testing.T, expected, got []path.Path) {
	assert.ElementsMatch(t, expected, got)
}

type BackupUnitSuite struct {
	tester.Suite
}

func TestOperationsManifestsUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupUnitSuite) TestCollectSiteMetadata() {
	const (
		tid = "tenantid"
		gid = "groupid"
		sid = "siteid"
	)

	itemPath := makeMetadataBasePath(
		suite.T(),
		tid,
		path.GroupsService,
		gid,
		path.LibrariesCategory,
		odConsts.SitesPathDir,
		sid)

	table := []struct {
		name        string
		manID       string
		expectPaths func(*testing.T, []string) []path.Path
		expectErr   error
	}{
		{
			name:  "simple",
			manID: "simple",
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				for _, f := range files {
					p, err := itemPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			paths := test.expectPaths(t, graph.AllMetadataFileNames())

			mr := mockRestoreProducer{err: test.expectErr}
			mr.buildRestoreFunc(t, test.manID, paths)

			man := kopia.ManifestEntry{
				Manifest: &snapshot.Manifest{ID: manifest.ID(test.manID)},
			}

			_, err := collectSiteMetadata(ctx, &mr, man, tid, gid, sid, fault.New(true))
			assert.ErrorIs(t, err, test.expectErr, clues.ToCore(err))
		})
	}
}
