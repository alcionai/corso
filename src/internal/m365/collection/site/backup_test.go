package site

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type SharePointSuite struct {
	tester.Suite
}

func TestSharePointSuite(t *testing.T) {
	suite.Run(t, &SharePointSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SharePointSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, false, 4)
}

func (suite *SharePointSuite) TestCollectPages() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID  = tconfig.M365SiteID(t)
		a       = tconfig.NewM365Account(t)
		counter = count.New()
	)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := api.NewClient(
		creds,
		control.DefaultOptions(),
		counter)
	require.NoError(t, err, clues.ToCore(err))

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: mock.NewProvider(siteID, siteID),
	}

	sel := selectors.NewSharePointBackup([]string{siteID})

	col, err := CollectPages(
		ctx,
		bpc,
		creds,
		ac,
		sel.Lists(selectors.Any())[0],
		(&MockGraphService{}).UpdateStatus,
		counter,
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, col)
}

func (suite *SharePointSuite) TestCollectLists() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID  = tconfig.M365SiteID(t)
		a       = tconfig.NewM365Account(t)
		counter = count.New()
	)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := api.NewClient(
		creds,
		control.DefaultOptions(),
		counter)
	require.NoError(t, err, clues.ToCore(err))

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: mock.NewProvider(siteID, siteID),
	}

	sel := selectors.NewSharePointBackup([]string{siteID})

	bh := NewListsBackupHandler(bpc.ProtectedResource.ID(), ac.Lists())

	col, _, err := CollectLists(
		ctx,
		bh,
		bpc,
		ac,
		creds.AzureTenantID,
		sel.Lists(selectors.Any())[0],
		(&MockGraphService{}).UpdateStatus,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	metadataFound := false

	for _, c := range col {
		if c.FullPath().Service() == path.SharePointMetadataService {
			metadataFound = true
			break
		}
	}

	assert.Less(t, 0, len(col))
	assert.True(t, metadataFound)
}

func (suite *SharePointSuite) TestParseListsMetadataCollections() {
	type fileValues struct {
		fileName string
		value    string
	}

	table := []struct {
		name                 string
		cat                  path.CategoryType
		wantedCategorycat    path.CategoryType
		data                 []fileValues
		expect               map[string]metadata.DeltaPath
		canUsePreviousBackup bool
		expectError          assert.ErrorAssertionFunc
	}{
		{
			name:              "previous path only",
			cat:               path.ListsCategory,
			wantedCategorycat: path.ListsCategory,
			data: []fileValues{
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]metadata.DeltaPath{
				"key": {
					Path: "prev-path",
				},
			},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name:              "multiple previous paths",
			cat:               path.ListsCategory,
			wantedCategorycat: path.ListsCategory,
			data: []fileValues{
				{metadata.PreviousPathFileName, "prev-path"},
				{metadata.PreviousPathFileName, "prev-path-2"},
			},
			canUsePreviousBackup: false,
			expectError:          assert.Error,
		},
		{
			name:              "unwanted category",
			cat:               path.LibrariesCategory,
			wantedCategorycat: path.ListsCategory,
			data: []fileValues{
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expectError: assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			entries := []graph.MetadataCollectionEntry{}

			for _, d := range test.data {
				entries = append(
					entries,
					graph.NewMetadataEntry(d.fileName, map[string]string{"key": d.value}))
			}

			pathPrefix, err := path.BuildMetadata(
				"t", "u",
				path.SharePointService,
				test.cat,
				false)
			require.NoError(t, err, "path prefix")

			coll, err := graph.MakeMetadataCollection(
				pathPrefix,
				entries,
				func(cos *support.ControllerOperationStatus) {},
				count.New())
			require.NoError(t, err, clues.ToCore(err))

			dps, canUsePreviousBackup, err := parseListsMetadataCollections(
				ctx,
				test.wantedCategorycat,
				[]data.RestoreCollection{
					dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: coll}),
				})
			test.expectError(t, err, clues.ToCore(err))

			if test.cat != test.wantedCategorycat {
				assert.Len(t, dps, 0)
			} else {
				assert.Equal(t, test.canUsePreviousBackup, canUsePreviousBackup, "can use previous backup")

				assert.Len(t, dps, len(test.expect))

				for k, v := range dps {
					assert.Equal(t, v.Path, test.expect[k].Path, "path")
				}
			}
		})
	}
}

type failingColl struct {
	t *testing.T
}

func (f failingColl) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ic := make(chan data.Item)
	defer close(ic)

	errs.AddRecoverable(ctx, assert.AnError)

	return ic
}

func (f failingColl) FullPath() path.Path {
	tmp, err := path.Build(
		"tenant",
		"siteid",
		path.SharePointService,
		path.ListsCategory,
		false,
		"list1")
	require.NoError(f.t, err, clues.ToCore(err))

	return tmp
}

func (f failingColl) FetchItemByName(context.Context, string) (data.Item, error) {
	// no fetch calls will be made
	return nil, nil
}

func (suite *SharePointSuite) TestParseListsMetadataCollections_ReadFailure() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	fc := failingColl{t}

	_, canUsePreviousBackup, err := parseListsMetadataCollections(ctx, path.ListsCategory, []data.RestoreCollection{fc})
	require.NoError(t, err)
	require.False(t, canUsePreviousBackup)
}
