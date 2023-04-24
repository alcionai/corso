package onedrive

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type DataCollectionsUnitSuite struct {
	tester.Suite
}

func TestDataCollectionsUnitSuite(t *testing.T) {
	suite.Run(t, &DataCollectionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DataCollectionsUnitSuite) TestMigrationCollections() {
	u := selectors.Selector{}
	u = u.SetDiscreteOwnerIDName("i", "n")

	od := path.OneDriveService.String()
	fc := path.FilesCategory.String()

	type migr struct {
		full string
		prev string
	}

	table := []struct {
		name            string
		version         int
		forceSkip       bool
		expectLen       int
		expectMigration []migr
	}{
		{
			name:            "no backup version",
			version:         version.NoBackup,
			forceSkip:       false,
			expectLen:       0,
			expectMigration: []migr{},
		},
		{
			name:            "above current version",
			version:         version.Backup + 5,
			forceSkip:       false,
			expectLen:       0,
			expectMigration: []migr{},
		},
		{
			name:      "user pn to id",
			version:   version.AllXMigrateUserPNToID - 1,
			forceSkip: false,
			expectLen: 1,
			expectMigration: []migr{
				{
					full: strings.Join([]string{"t", od, "i", fc}, "/"),
					prev: strings.Join([]string{"t", od, "n", fc}, "/"),
				},
			},
		},
		{
			name:            "skipped",
			version:         version.Backup + 5,
			forceSkip:       true,
			expectLen:       0,
			expectMigration: []migr{},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			opts := control.Options{
				ToggleFeatures: control.Toggles{
					RunMigrations: !test.forceSkip,
				},
			}

			mc, err := migrationCollections(nil, test.version, "t", u, nil, opts)
			require.NoError(t, err, clues.ToCore(err))

			if test.expectLen == 0 {
				assert.Nil(t, mc)
				return
			}

			assert.Len(t, mc, test.expectLen)

			migrs := []migr{}

			for _, col := range mc {
				var fp, pp string

				if col.FullPath() != nil {
					fp = col.FullPath().String()
				}

				if col.PreviousPath() != nil {
					pp = col.PreviousPath().String()
				}

				t.Logf("Found migration collection:\n* full: %s\n* prev: %s\n", fp, pp)

				migrs = append(migrs, test.expectMigration...)
			}

			for i, m := range migrs {
				assert.Contains(t, migrs, m, "expected to find migration: %+v", test.expectMigration[i])
			}
		})
	}
}
