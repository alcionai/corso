package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	msDrive "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
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
			expect := make([]DriveItemIDType, 0, len(ims))

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
			expect := map[string]DriveItemIDType{}

			assert.NotEmptyf(
				t,
				igv,
				"need at least one item to compare in user %s drive %s folder %s",
				suite.its.user.id, test.driveID, test.rootFolderID)

			for _, itm := range igv {
				expect[ptr.Val(itm.GetId())] = DriveItemIDType{
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

func (suite *DrivePagerIntgSuite) TestEnumerateDriveItems() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	items := []models.DriveItemable{}

	pager := suite.its.
		ac.
		Drives().
		EnumerateDriveItemsDelta(
			ctx,
			suite.its.user.driveID,
			"",
			CallConfig{
				Select: DefaultDriveItemProps(),
			})

	for page, reset, done := pager.NextPage(); !done; page, reset, done = pager.NextPage() {
		items = append(items, page...)

		assert.False(t, reset, "should not reset")
	}

	du, err := pager.Results()

	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, items, "should find items in user's drive")
	assert.NotEmpty(t, du.URL, "should have a delta link")
}

func (suite *DrivePagerIntgSuite) TestDriveDeltaPagerQueryParams() {
	tests := []struct {
		name   string
		setupf func()
		expect assert.ErrorAssertionFunc
	}{
		{
			name: "validate select and top params",
			setupf: func() {
				delta := msDrive.NewItemItemsItemDeltaResponse()
				delta.SetValue([]models.DriveItemable{
					models.NewDriveItem(),
				})

				str := "deltaLink"
				delta.SetOdataDeltaLink(&str)

				var selectParams string
				for _, v := range DefaultDriveItemProps() {
					selectParams += v + ","
				}

				// remove last comma
				selectParams = selectParams[:len(selectParams)-1]

				queryParams := map[string]string{
					"$top":    string(maxDeltaPageSize),
					"$select": selectParams,
				}

				interceptV1Path("drives", "drive", "items", "root", "delta()").
					MatchParams(queryParams).
					Reply(200).
					JSON(requireParseableToMap(suite.T(), delta))
			},
			expect: assert.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			defer gock.Off()
			test.setupf()

			pager := suite.
				its.
				gockAC.
				Drives().
				EnumerateDriveItemsDelta(
					ctx,
					"drive",
					"",
					CallConfig{
						Select: DefaultDriveItemProps(),
					})
			for _, reset, done := pager.NextPage(); !done; _, reset, done = pager.NextPage() {
				assert.False(t, reset, "should not reset")
			}

			_, err := pager.Results()
			require.Error(t, err, clues.ToCore(err))
		})
	}
}
