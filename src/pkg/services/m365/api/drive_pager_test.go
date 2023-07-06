package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/config"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type DrivePagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestDrivePagerIntgSuite(t *testing.T) {
	suite.Run(t, &DrivePagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{config.M365AcctCredEnvs}),
	})
}

func (suite *DrivePagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *DrivePagerIntgSuite) TestDrives_GetItemsInContainerByCollisionKey() {
	table := []struct {
		name         string
		driveID      string
		rootFolderID string
	}{
		{
			name:         "user drive",
			driveID:      suite.its.userDriveID,
			rootFolderID: suite.its.userDriveRootFolderID,
		},
		{
			name:         "site drive",
			driveID:      suite.its.siteDriveID,
			rootFolderID: suite.its.siteDriveRootFolderID,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			items, err := suite.its.ac.Stable.
				Client().
				Drives().
				ByDriveId(test.driveID).
				Items().
				ByDriveItemId(test.rootFolderID).
				Children().
				Get(ctx, nil)
			require.NoError(t, err, clues.ToCore(err))

			ims := items.GetValue()
			expect := make([]api.DriveCollisionItem, 0, len(ims))

			assert.NotEmptyf(
				t,
				ims,
				"need at least one item to compare in user %s drive %s folder %s",
				suite.its.userID, test.driveID, test.rootFolderID)

			results, err := suite.its.ac.Drives().GetItemsInContainerByCollisionKey(ctx, test.driveID, test.rootFolderID)
			require.NoError(t, err, clues.ToCore(err))
			require.NotEmpty(t, results)

			for k, v := range results {
				assert.NotEmpty(t, k, "all keys should be populated")
				assert.NotEmpty(t, v, "all values should be populated")
			}

			for _, e := range expect {
				r, ok := results[e.ItemID]
				assert.Truef(t, ok, "expected results to contain collision key: %s", e)
				assert.Equal(t, e, r)
			}
		})
	}
}
