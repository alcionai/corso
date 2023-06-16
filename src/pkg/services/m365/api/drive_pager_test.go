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
	suite.Run(t, &MailPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *DrivePagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *DrivePagerIntgSuite) TestGetItemsInContainerByCollisionKey() {
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

			results, err := suite.its.ac.Drives().GetItemsInContainerByCollisionKey(ctx, test.driveID, test.rootFolderID)
			require.NoError(t, err, clues.ToCore(err))
			require.Less(t, 0, len(results), "requires at least one result")

			for k, v := range results {
				assert.NotEmpty(t, k, "all keys should be populated")
				assert.NotEmpty(t, v, "all values should be populated")
			}
		})
	}
}
