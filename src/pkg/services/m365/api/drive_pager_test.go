package api

import (
	"strconv"
	"strings"
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
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
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

// TestDriveDeltaPagerQueryParams is a regression test to check if
// the delta pager is setting query params and prefer headers as
// expected.
func (suite *DrivePagerIntgSuite) TestDriveDeltaPagerQueryParams() {
	tests := []struct {
		name        string
		setupf      func()
		forURLCache bool
		expect      assert.ErrorAssertionFunc
	}{
		// Separate out tests for $select and $prefer because otherwise
		// it's hard to tell which gock check is failing.
		{
			name: "check $top",
			setupf: func() {
				delta := msDrive.NewItemItemsItemDeltaResponse()
				delta.SetValue([]models.DriveItemable{
					models.NewDriveItem(),
				})

				str := "deltaLink"
				delta.SetOdataDeltaLink(&str)

				interceptV1Path("drives", "drive", "items", "root", "delta()").
					MatchParam("$top", strconv.Itoa(int(maxDeltaPageSize))).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), delta))
			},
			expect: assert.NoError,
		},
		{
			name: "check $select",
			setupf: func() {
				delta := msDrive.NewItemItemsItemDeltaResponse()
				delta.SetValue([]models.DriveItemable{
					models.NewDriveItem(),
				})

				str := "deltaLink"
				delta.SetOdataDeltaLink(&str)

				// This is a copy of DefaultDriveItemProps. This list should be
				// adjusted while changing DefaultDriveItemProps. This however
				// doesn't protect against adding new params to DefaultDriveItemProps
				// since gock doesn't allow asserting for presence of unexpected
				// params.
				params := idAnd(
					"content.downloadUrl",
					"createdBy",
					"createdDateTime",
					"file",
					"folder",
					"lastModifiedDateTime",
					"name",
					"package",
					"parentReference",
					"root",
					"size",
					"deleted",
					"malware",
					"shared")

				selectParams := strings.Join(params, ",")

				interceptV1Path("drives", "drive", "items", "root", "delta()").
					MatchParam("$select", selectParams).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), delta))
			},
			expect: assert.NoError,
		},
		{
			name: "prefer headers",
			setupf: func() {
				delta := msDrive.NewItemItemsItemDeltaResponse()
				delta.SetValue([]models.DriveItemable{
					models.NewDriveItem(),
				})

				str := "deltaLink"
				delta.SetOdataDeltaLink(&str)

				preferHeaderItems := []string{
					"deltashowremovedasdeleted",
					"deltatraversepermissiongaps",
					"deltashowsharingchanges",
					"hierarchicalsharing",
				}
				preferParams := strings.Join(preferHeaderItems, ",")

				interceptV1Path("drives", "drive", "items", "root", "delta()").
					MatchHeaders(map[string]string{
						"Prefer": preferParams,
					}).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), delta))
			},
			expect: assert.NoError,
		},
		{
			name: "check url cache $select",
			setupf: func() {
				delta := msDrive.NewItemItemsItemDeltaResponse()
				delta.SetValue([]models.DriveItemable{
					models.NewDriveItem(),
				})

				str := "deltaLink"
				delta.SetOdataDeltaLink(&str)

				// This is a copy of URLCacheDriveItemProps. This list should be
				// adjusted while changing URLCacheDriveItemProps.
				params := idAnd(
					"content.downloadUrl",
					"deleted",
					"file",
					"folder")

				selectParams := strings.Join(params, ",")

				interceptV1Path("drives", "drive", "items", "root", "delta()").
					MatchParam("$select", selectParams).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), delta))
			},
			forURLCache: true,
			expect:      assert.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			defer gock.Off()
			test.setupf()

			cfg := CallConfig{
				Select: DefaultDriveItemProps(),
			}

			if test.forURLCache {
				cfg = CallConfig{
					Select: URLCacheDriveItemProps(),
				}
			}

			pager := suite.
				its.
				gockAC.
				Drives().
				EnumerateDriveItemsDelta(
					ctx,
					"drive",
					"",
					cfg)
			for _, reset, done := pager.NextPage(); !done; _, reset, done = pager.NextPage() {
				assert.False(t, reset, "should not reset")
			}

			_, err := pager.Results()
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

// Non delta pager equivalent of above test.
func (suite *DrivePagerIntgSuite) TestDriveNonDeltaPagerQueryParams() {
	tests := []struct {
		name        string
		setupf      func()
		forURLCache bool
		expect      assert.ErrorAssertionFunc
	}{
		{
			name: "check $top",
			setupf: func() {
				resp := models.NewDriveItemCollectionResponse()
				resp.SetValue([]models.DriveItemable{
					models.NewDriveItem(),
				})

				interceptV1Path("drives", "drive", "items", "root", "children").
					MatchParam("$top", strconv.Itoa(int(maxNonDeltaPageSize))).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), resp))
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

			_, err := suite.its.gockAC.
				Drives().
				GetItemIDsInContainer(ctx, "drive", "root")

			require.NoError(t, err, clues.ToCore(err))
		})
	}
}
