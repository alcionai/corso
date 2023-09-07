package groups

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type GroupsBackupUnitSuite struct {
	tester.Suite
}

func TestGroupsBackupUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsBackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type mockRestoreProducer struct {
	rc  []data.RestoreCollection
	err error
}

func (mr mockRestoreProducer) ProduceRestoreCollections(
	ctx context.Context,
	snapshotID string,
	paths []path.RestorePaths,
	bc kopia.ByteCounter,
	errs *fault.Bus,
) ([]data.RestoreCollection, error) {
	return mr.rc, mr.err
}

type mockCollection struct {
	items []mockItem
}

type mockItem struct {
	name string
	data string
}

func (mi mockItem) ToReader() io.ReadCloser { return io.NopCloser(strings.NewReader(mi.data)) }
func (mi mockItem) ID() string              { return mi.name }
func (mi mockItem) Deleted() bool           { return false }

func (mc mockCollection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ch := make(chan data.Item)

	go func() {
		defer close(ch)

		for _, item := range mc.items {
			ch <- item
		}
	}()

	return ch
}
func (mc mockCollection) FullPath() path.Path { panic("unimplemented") }
func (mc mockCollection) FetchItemByName(ctx context.Context, name string) (data.Item, error) {
	panic("unimplemented")
}

func (suite *GroupsBackupUnitSuite) TestMetadataFiles() {
	tests := []struct {
		name      string
		reason    identity.Reasoner
		r         inject.RestoreProducer
		manID     manifest.ID
		result    [][]string
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "error",
			reason:    kopia.NewReason("tenant", "user", path.GroupsService, path.LibrariesCategory),
			manID:     "manifestID",
			r:         mockRestoreProducer{err: assert.AnError},
			expectErr: require.Error,
		},
		{
			name:   "single site",
			reason: kopia.NewReason("tenant", "user", path.GroupsService, path.LibrariesCategory),
			manID:  "manifestID",
			r: mockRestoreProducer{
				rc: []data.RestoreCollection{
					mockCollection{
						items: []mockItem{
							{name: "previouspath", data: `{"id1": "path/to/id1"}`},
						},
					},
				},
			},
			result:    [][]string{{"sites", "id1", "delta"}, {"sites", "id1", "previouspath"}},
			expectErr: require.NoError,
		},
		{
			name:   "multiple sites",
			reason: kopia.NewReason("tenant", "user", path.GroupsService, path.LibrariesCategory),
			manID:  "manifestID",
			r: mockRestoreProducer{
				rc: []data.RestoreCollection{
					mockCollection{
						items: []mockItem{
							{name: "previouspath", data: `{"id1": "path/to/id1", "id2": "path/to/id2"}`},
						},
					},
				},
			},
			result: [][]string{
				{"sites", "id1", "delta"},
				{"sites", "id1", "previouspath"},
				{"sites", "id2", "delta"},
				{"sites", "id2", "previouspath"},
			},
			expectErr: require.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			res, err := MetadataFiles(ctx, test.reason, test.r, test.manID, fault.New(true))

			test.expectErr(t, err)
			assert.Equal(t, test.result, res)
		})
	}
}
