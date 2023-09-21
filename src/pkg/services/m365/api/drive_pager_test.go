package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
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
			[][]string{tconfig.M365AcctCredEnvs}),
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
			driveID:      suite.its.user.driveID,
			rootFolderID: suite.its.user.driveRootFolderID,
		},
		{
			name:         "site drive",
			driveID:      suite.its.site.driveID,
			rootFolderID: suite.its.site.driveRootFolderID,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			t.Log("drive", test.driveID)
			t.Log("rootFolder", test.rootFolderID)

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
			expect := make([]api.DriveItemIDType, 0, len(ims))

			assert.NotEmptyf(
				t,
				ims,
				"need at least one item to compare in user %s drive %s folder %s",
				suite.its.user.id, test.driveID, test.rootFolderID)

			results, err := suite.its.ac.
				Drives().
				GetItemsInContainerByCollisionKey(ctx, test.driveID, test.rootFolderID)
			require.NoError(t, err, clues.ToCore(err))
			require.NotEmpty(t, results)

			for _, k := range expect {
				t.Log("expects key", k)
			}

			for k := range results {
				t.Log("results key", k)
			}

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

func (suite *DrivePagerIntgSuite) TestDrives_GetItemIDsInContainer() {
	table := []struct {
		name         string
		driveID      string
		rootFolderID string
	}{
		{
			name:         "user drive",
			driveID:      suite.its.user.driveID,
			rootFolderID: suite.its.user.driveRootFolderID,
		},
		{
			name:         "site drive",
			driveID:      suite.its.site.driveID,
			rootFolderID: suite.its.site.driveRootFolderID,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			t.Log("drive", test.driveID)
			t.Log("rootFolder", test.rootFolderID)

			items, err := suite.its.ac.Stable.
				Client().
				Drives().
				ByDriveId(test.driveID).
				Items().
				ByDriveItemId(test.rootFolderID).
				Children().
				Get(ctx, nil)
			require.NoError(t, err, clues.ToCore(err))

			igv := items.GetValue()
			expect := map[string]api.DriveItemIDType{}

			assert.NotEmptyf(
				t,
				igv,
				"need at least one item to compare in user %s drive %s folder %s",
				suite.its.user.id, test.driveID, test.rootFolderID)

			for _, itm := range igv {
				expect[ptr.Val(itm.GetId())] = api.DriveItemIDType{
					ItemID:   ptr.Val(itm.GetId()),
					IsFolder: itm.GetFolder() != nil,
				}
			}

			results, err := suite.its.ac.
				Drives().
				GetItemIDsInContainer(ctx, test.driveID, test.rootFolderID)
			require.NoError(t, err, clues.ToCore(err))
			require.NotEmpty(t, results)
			require.Equal(t, len(expect), len(results), "must have same count of items")

			for k := range expect {
				t.Log("expects key", k)
			}

			for k, v := range results {
				t.Log("results key", k)
				assert.NotEmpty(t, v, "all values should be populated")
			}

			assert.Equal(t, expect, results)
		})
	}
}
