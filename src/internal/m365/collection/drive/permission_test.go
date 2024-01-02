package drive

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/syncd"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
)

type PermissionsUnitTestSuite struct {
	tester.Suite
}

func TestPermissionsUnitTestSuite(t *testing.T) {
	suite.Run(t, &PermissionsUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PermissionsUnitTestSuite) TestComputeParentPermissions_oneDrive() {
	runComputeParentPermissionsTest(suite, path.OneDriveService, path.FilesCategory, "user")
}

func (suite *PermissionsUnitTestSuite) TestComputeParentPermissions_sharePoint() {
	runComputeParentPermissionsTest(suite, path.SharePointService, path.LibrariesCategory, "site")
}

func runComputeParentPermissionsTest(
	suite *PermissionsUnitTestSuite,
	service path.ServiceType,
	category path.CategoryType,
	resourceOwner string,
) {
	entryPath := odConsts.DriveFolderPrefixBuilder("drive-id").String() + "/level0/level1/level2/entry"
	rootEntryPath := odConsts.DriveFolderPrefixBuilder("drive-id").String() + "/entry"

	entry, err := path.Build(
		"tenant",
		resourceOwner,
		service,
		category,
		false,
		strings.Split(entryPath, "/")...)
	require.NoError(suite.T(), err, "creating path")

	rootEntry, err := path.Build(
		"tenant",
		resourceOwner,
		service,
		category,
		false,
		strings.Split(rootEntryPath, "/")...)
	require.NoError(suite.T(), err, "creating path")

	level2, err := entry.Dir()
	require.NoError(suite.T(), err, "level2 path")

	level1, err := level2.Dir()
	require.NoError(suite.T(), err, "level1 path")

	level0, err := level1.Dir()
	require.NoError(suite.T(), err, "level0 path")

	md := metadata.Metadata{
		SharingMode: metadata.SharingModeCustom,
		Permissions: []metadata.Permission{
			{
				Roles:    []string{"write"},
				EntityID: "user-id",
			},
		},
	}

	metadata2 := metadata.Metadata{
		SharingMode: metadata.SharingModeCustom,
		Permissions: []metadata.Permission{
			{
				Roles:    []string{"read"},
				EntityID: "user-id",
			},
		},
	}

	inherited := metadata.Metadata{
		SharingMode: metadata.SharingModeInherited,
		Permissions: []metadata.Permission{},
	}

	table := []struct {
		name        string
		item        path.Path
		meta        metadata.Metadata
		parentPerms map[string]metadata.Metadata
	}{
		{
			name:        "root level entry",
			item:        rootEntry,
			meta:        metadata.Metadata{},
			parentPerms: map[string]metadata.Metadata{},
		},
		{
			name:        "root level directory",
			item:        level0,
			meta:        metadata.Metadata{},
			parentPerms: map[string]metadata.Metadata{},
		},
		{
			name: "direct parent perms",
			item: entry,
			meta: md,
			parentPerms: map[string]metadata.Metadata{
				level2.String(): md,
			},
		},
		{
			name: "top level parent perms",
			item: entry,
			meta: md,
			parentPerms: map[string]metadata.Metadata{
				level2.String(): inherited,
				level1.String(): inherited,
				level0.String(): md,
			},
		},
		{
			name: "all inherited",
			item: entry,
			meta: metadata.Metadata{},
			parentPerms: map[string]metadata.Metadata{
				level2.String(): inherited,
				level1.String(): inherited,
				level0.String(): inherited,
			},
		},
		{
			name: "multiple custom permission",
			item: entry,
			meta: md,
			parentPerms: map[string]metadata.Metadata{
				level2.String(): inherited,
				level1.String(): md,
				level0.String(): metadata2,
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			input := syncd.NewMapTo[metadata.Metadata]()
			for k, v := range test.parentPerms {
				input.Store(k, v)
			}

			m, err := computePreviousMetadata(ctx, test.item, input)
			require.NoError(t, err, "compute permissions")

			assert.Equal(t, m, test.meta)
		})
	}
}

type mockUdip struct {
	// cannot use []bool as this is not passed by ref
	success chan bool
}

func (mockUdip) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	return nil
}

func (m mockUdip) PostItemPermissionUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	if ptr.Val(body.GetRecipients()[0].GetObjectId()) == "failure" {
		m.success <- false

		err := graphTD.ODataErrWithMsg("InvalidRequest", string("One or more users could not be resolved"))

		return nil, err
	}

	m.success <- true

	resp := drives.NewItemItemsItemInviteResponse()
	perm := models.NewPermission()
	perm.SetId(ptr.To(itemID))
	resp.SetValue([]models.Permissionable{perm})

	return resp, nil
}

func (suite *PermissionsUnitTestSuite) TestPermissionRestoreNonExistentUser() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	successChan := make(chan bool, 3)
	m := mockUdip{
		success: successChan,
	}

	err := UpdatePermissions(
		ctx,
		&m,
		"drive-id",
		"item-id",
		[]metadata.Permission{
			{Roles: []string{"write"}, EntityID: "user-id1"},
			{Roles: []string{"write"}, EntityID: "failure"},
			{Roles: []string{"write"}, EntityID: "user-id2"},
		},
		[]metadata.Permission{},
		syncd.NewMapTo[string](),
		fault.New(true))

	assert.NoError(t, err, "update permissions")
	close(successChan)

	var successValues []bool
	for success := range successChan {
		successValues = append(successValues, success)
	}

	expectedSuccessValues := []bool{true, false, true}
	assert.Equal(t, expectedSuccessValues, successValues)
}

type mockUpils struct {
	success chan bool
}

func (mockUpils) DeleteItemPermission(
	ctx context.Context,
	driveID, itemID, permissionID string,
) error {
	return nil
}

func (m mockUpils) PostItemLinkShareUpdate(
	ctx context.Context,
	driveID, itemID string,
	body *drives.ItemItemsItemCreateLinkPostRequestBody,
) (models.Permissionable, error) {
	shouldFail := false

	recip := body.GetAdditionalData()["recipients"].([]map[string]string)
	for _, r := range recip {
		if r["objectId"] == "failure" {
			shouldFail = true
		}
	}

	if shouldFail {
		m.success <- false

		err := graphTD.ODataErrWithMsg("InvalidRequest", string("One or more users could not be resolved"))

		return nil, err
	}

	m.success <- true

	perm := models.NewPermission()
	perm.SetId(ptr.To(itemID))

	return perm, nil
}

func (suite *PermissionsUnitTestSuite) TestLinkShareRestoreNonExistentUser() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	successChan := make(chan bool, 10) // 7 required
	m := mockUpils{
		success: successChan,
	}

	_, err := UpdateLinkShares(
		ctx,
		&m,
		"drive-id",
		"item-id",
		[]metadata.LinkShare{
			{Roles: []string{"write"}, Entities: []metadata.Entity{{ID: "user-id"}}},
			{Roles: []string{"write"}, Entities: []metadata.Entity{{ID: "failure"}}},
			{Roles: []string{"write"}, Entities: []metadata.Entity{{ID: "user-id"}, {ID: "failure"}}},
			{Roles: []string{"write"}, Entities: []metadata.Entity{{ID: "user-id"}, {ID: "failure"}, {ID: "user-id"}}},
		},
		[]metadata.LinkShare{},
		syncd.NewMapTo[string](),
		fault.New(true))

	assert.NoError(t, err, "update permissions")
	close(successChan)

	var successValues []bool
	for success := range successChan {
		successValues = append(successValues, success)
	}

	expectedSuccessValues := []bool{true, false, false, false}
	assert.Equal(t, expectedSuccessValues, successValues)
}

func (suite *PermissionsUnitTestSuite) TestFilterUnavailableEntitiesInPermissions() {
	table := []struct {
		name               string
		permissions        []metadata.Permission
		expected           []metadata.Permission
		availableEntities  AvailableEntities
		skippedPermissions []string
	}{
		{
			name: "single item",
			permissions: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
			},
			expected: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
			},
			availableEntities: AvailableEntities{
				Users:  idname.NewCache(map[string]string{"e1": "e1"}),
				Groups: idname.NewCache(map[string]string{}),
			},
		},
		{
			name: "single item with missing entity",
			permissions: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
			},
			expected: []metadata.Permission{},
			availableEntities: AvailableEntities{
				Users:  idname.NewCache(map[string]string{}),
				Groups: idname.NewCache(map[string]string{}),
			},
			skippedPermissions: []string{"p1"},
		},
		{
			name: "multiple items",
			permissions: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
				{ID: "p2", EntityID: "e2", EntityType: metadata.GV2User},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2User},
			},
			expected: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
				{ID: "p2", EntityID: "e2", EntityType: metadata.GV2User},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2User},
			},
			availableEntities: AvailableEntities{
				Users:  idname.NewCache(map[string]string{"e1": "e1", "e2": "e2", "e3": "e3"}),
				Groups: idname.NewCache(map[string]string{}),
			},
		},
		{
			name: "multiple items with missing entity",
			permissions: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
				{ID: "p2", EntityID: "e2", EntityType: metadata.GV2User},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2Group},
			},
			expected: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2Group},
			},
			availableEntities: AvailableEntities{
				Users:  idname.NewCache(map[string]string{"e1": "e1"}),
				Groups: idname.NewCache(map[string]string{"e3": "e3"}),
			},
			skippedPermissions: []string{"p2"},
		},
		{
			name: "multiple items with missing entity and multiple types",
			permissions: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
				{ID: "p2", EntityID: "e2", EntityType: metadata.GV2User},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2Group},
				{ID: "p4", EntityID: "e4", EntityType: metadata.GV2Group},
				{ID: "p5", EntityID: "e5", EntityType: metadata.GV2User},
			},
			expected: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2User},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2Group},
				{ID: "p5", EntityID: "e5", EntityType: metadata.GV2User},
			},
			availableEntities: AvailableEntities{
				Users:  idname.NewCache(map[string]string{"e1": "e1", "e5": "e5"}),
				Groups: idname.NewCache(map[string]string{"e3": "e3"}),
			},
			skippedPermissions: []string{"p2", "p4"},
		},
		{
			name: "single item different type",
			permissions: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2Group},
			},
			expected: []metadata.Permission{},
			availableEntities: AvailableEntities{
				Users:  idname.NewCache(map[string]string{"e1": "e1"}),
				Groups: idname.NewCache(map[string]string{}),
			},
			skippedPermissions: []string{"p1"},
		},
		{
			name: "unhandled types",
			permissions: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2Device},
				{ID: "p2", EntityID: "e2", EntityType: metadata.GV2App},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2SiteUser},
				{ID: "p4", EntityID: "e4", EntityType: metadata.GV2SiteGroup},
			},
			expected: []metadata.Permission{
				{ID: "p1", EntityID: "e1", EntityType: metadata.GV2Device},
				{ID: "p2", EntityID: "e2", EntityType: metadata.GV2App},
				{ID: "p3", EntityID: "e3", EntityType: metadata.GV2SiteUser},
				{ID: "p4", EntityID: "e4", EntityType: metadata.GV2SiteGroup},
			},
			availableEntities: AvailableEntities{
				// these are users and not what we have
				Users:  idname.NewCache(map[string]string{"e1": "e1", "e2": "e2", "e3": "e3", "e4": "e4"}),
				Groups: idname.NewCache(map[string]string{}),
			},
			skippedPermissions: []string{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			oldToNew := syncd.NewMapTo[string]()
			filtered := filterUnavailableEntitiesInPermissions(ctx, test.permissions, test.availableEntities, oldToNew)

			assert.Equal(t, test.expected, filtered, "filtered permissions")

			for _, id := range test.skippedPermissions {
				_, ok := oldToNew.Load(id)
				assert.True(t, ok, fmt.Sprintf("skipped id %s", id))
			}
		})
	}
}

type eidtype struct {
	id    string
	etype metadata.GV2Type
}

func (suite *PermissionsUnitTestSuite) TestFilterUnavailableEntitiesInLinkShare() {
	ls := func(lsid string, entities []eidtype) metadata.LinkShare {
		ents := []metadata.Entity{}
		for _, e := range entities {
			ents = append(ents, metadata.Entity{ID: e.id, EntityType: e.etype})
		}

		return metadata.LinkShare{
			ID:       lsid,
			Entities: ents,
		}
	}

	ae := func(uids, gids []string) AvailableEntities {
		users := map[string]string{}
		for _, id := range uids {
			users[id] = id
		}

		groups := map[string]string{}
		for _, id := range gids {
			groups[id] = id
		}

		return AvailableEntities{
			Users:  idname.NewCache(users),
			Groups: idname.NewCache(groups),
		}
	}

	table := []struct {
		name              string
		linkShares        []metadata.LinkShare
		expected          []metadata.LinkShare
		availableEntities AvailableEntities
		skippedLinkShares []string
	}{
		{
			name:              "single item, single available entity",
			linkShares:        []metadata.LinkShare{ls("ls1", []eidtype{{"e1", metadata.GV2User}})},
			expected:          []metadata.LinkShare{ls("ls1", []eidtype{{"e1", metadata.GV2User}})},
			availableEntities: ae([]string{"e1"}, []string{}),
		},
		{
			name:              "single item, single missing entity",
			linkShares:        []metadata.LinkShare{ls("ls1", []eidtype{{"e1", metadata.GV2User}})},
			expected:          []metadata.LinkShare{},
			availableEntities: ae([]string{}, []string{}),
			skippedLinkShares: []string{"ls1"},
		},
		{
			name: "multiple items, multiple available entities",
			linkShares: []metadata.LinkShare{
				ls("ls1", []eidtype{{"e1", metadata.GV2User}, {"e2", metadata.GV2User}}),
				ls("ls2", []eidtype{{"e3", metadata.GV2User}, {"e4", metadata.GV2User}}),
			},
			expected: []metadata.LinkShare{
				ls("ls1", []eidtype{{"e1", metadata.GV2User}, {"e2", metadata.GV2User}}),
				ls("ls2", []eidtype{{"e3", metadata.GV2User}, {"e4", metadata.GV2User}}),
			},
			availableEntities: ae([]string{"e1", "e2", "e3", "e4"}, []string{}),
		},
		{
			name: "multiple items, missing entities",
			linkShares: []metadata.LinkShare{
				ls("ls1", []eidtype{{"e1", metadata.GV2User}, {"e2", metadata.GV2Group}}),
				ls("ls2", []eidtype{{"e3", metadata.GV2User}, {"e4", metadata.GV2Group}}),
			},
			expected: []metadata.LinkShare{
				ls("ls1", []eidtype{{"e1", metadata.GV2User}}),
				ls("ls2", []eidtype{{"e4", metadata.GV2Group}}),
			},
			availableEntities: ae([]string{"e1"}, []string{"e4"}),
			skippedLinkShares: []string{},
		},
		{
			name: "unhandled items",
			linkShares: []metadata.LinkShare{
				ls("ls1", []eidtype{{"e1", metadata.GV2Device}, {"e2", metadata.GV2App}}),
				ls("ls2", []eidtype{{"e3", metadata.GV2SiteUser}, {"e4", metadata.GV2SiteGroup}}),
			},
			expected: []metadata.LinkShare{
				ls("ls1", []eidtype{{"e1", metadata.GV2Device}, {"e2", metadata.GV2App}}),
				ls("ls2", []eidtype{{"e3", metadata.GV2SiteUser}, {"e4", metadata.GV2SiteGroup}}),
			},
			availableEntities: ae([]string{}, []string{}),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			oldToNew := syncd.NewMapTo[string]()
			filtered := filterUnavailableEntitiesInLinkShare(ctx, test.linkShares, test.availableEntities, oldToNew)

			assert.Equal(t, test.expected, filtered, "filtered link shares")

			for _, id := range test.skippedLinkShares {
				_, ok := oldToNew.Load(id)
				assert.True(t, ok, fmt.Sprintf("skipped id %s", id))
			}
		})
	}
}
