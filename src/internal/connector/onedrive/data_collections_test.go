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
		expectLen       int
		expectMigration []migr
	}{
		{
			name:            "no backup version",
			version:         version.NoBackup,
			expectLen:       0,
			expectMigration: []migr{},
		},
		{
			name:            "above current version",
			version:         version.Backup + 5,
			expectLen:       0,
			expectMigration: []migr{},
		},
		{
			name:      "user pn to id",
			version:   version.All7MigrateUserPNToID - 1,
			expectLen: 1,
			expectMigration: []migr{
				{
					full: strings.Join([]string{"t", od, "i", fc}, "/"),
					prev: strings.Join([]string{"t", od, "n", fc}, "/"),
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			mc, err := migrationCollections(nil, test.version, "t", u, nil, control.Options{})
			require.NoError(t, err, clues.ToCore(err))

			if test.expectLen == 0 {
				assert.Nil(t, mc)
				return
			}

			assert.LessOrEqual(t, test.expectLen, len(mc))

			migrs := make([]bool, len(test.expectMigration))

			for _, col := range mc {
				var fp, pp string

				if col.FullPath() != nil {
					fp = col.FullPath().String()
				}

				if col.PreviousPath() != nil {
					pp = col.PreviousPath().String()
				}

				t.Logf("Found migration collection:\n* full: %s\n* prev: %s\n", fp, pp)

				for i, cm := range test.expectMigration {
					if cm.full == fp && cm.prev == pp {
						migrs[i] = true
					}
				}
			}

			for i, m := range migrs {
				assert.Truef(t, m, "expected to find migration: %+v", test.expectMigration[i])
			}
		})
	}
}
