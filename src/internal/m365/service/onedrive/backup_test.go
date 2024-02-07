package onedrive

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/operations/inject"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/version"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/selectors"
)

type BackupUnitSuite struct {
	tester.Suite
}

func TestBackupUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupUnitSuite) TestMigrationCollections() {
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
			version:   version.All8MigrateUserPNToID - 1,
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
				ToggleFeatures: control.Toggles{},
			}

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: test.version,
				Options:           opts,
				ProtectedResource: u,
			}

			mc, err := migrationCollections(bpc, "t", nil, count.New())
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
