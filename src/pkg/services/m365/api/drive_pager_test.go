package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type DrivePagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestDrivePagerIntgSuite(t *testing.T) {
	suite.Run(t, &DrivePagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
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

			ac := suite.its.ac.Drives()
			items, err := ac.Stable.
				Client().
				Drives().
				ByDriveId(test.driveID).
				Items().
				ByDriveItemId(test.rootFolderID).
				Children().
				Get(ctx, nil)
			require.NoError(t, err, clues.ToCore(err))

			is := items.GetValue()
			expect := make([]string, 0, len(is))

			require.NotZero(t, len(expect), "need at least one item to compare against")

			results, err := ac.GetItemsInContainerByCollisionKey(ctx, test.driveID, test.rootFolderID)
			require.NoError(t, err, clues.ToCore(err))
			require.Less(t, 0, len(results), "requires at least one result")

			for k, v := range results {
				assert.NotEmpty(t, k, "all keys should be populated")
				assert.NotEmpty(t, v, "all values should be populated")
			}

			for _, e := range expect {
				_, ok := results[e]
				assert.Truef(t, ok, "expected results to contain collision key: %s", e)
			}
		})
	}
}
