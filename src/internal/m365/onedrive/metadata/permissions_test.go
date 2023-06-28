package metadata

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type PermissionsUnitTestSuite struct {
	tester.Suite
}

func TestPermissionsUnitTestSuite(t *testing.T) {
	suite.Run(t, &PermissionsUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PermissionsUnitTestSuite) TestDiffPermissions() {
	perm1 := Permission{
		ID:       "id1",
		Roles:    []string{"read"},
		EntityID: "user-id1",
	}

	perm2 := Permission{
		ID:       "id2",
		Roles:    []string{"write"},
		EntityID: "user-id2",
	}

	perm3 := Permission{
		ID:       "id3",
		Roles:    []string{"write"},
		EntityID: "user-id3",
	}

	// The following two permissions have same id and user but
	// different roles, this is a valid scenario for permissions.
	sameidperm1 := Permission{
		ID:       "id0",
		Roles:    []string{"write"},
		EntityID: "user-id0",
	}
	sameidperm2 := Permission{
		ID:       "id0",
		Roles:    []string{"read"},
		EntityID: "user-id0",
	}

	emailperm1 := Permission{
		ID:    "id1",
		Roles: []string{"read"},
		Email: "email1@provider.com",
	}

	emailperm2 := Permission{
		ID:    "id1",
		Roles: []string{"read"},
		Email: "email2@provider.com",
	}

	table := []struct {
		name    string
		before  []Permission
		after   []Permission
		added   []Permission
		removed []Permission
	}{
		{
			name:    "single permission added",
			before:  []Permission{},
			after:   []Permission{perm1},
			added:   []Permission{perm1},
			removed: []Permission{},
		},
		{
			name:    "single permission removed",
			before:  []Permission{perm1},
			after:   []Permission{},
			added:   []Permission{},
			removed: []Permission{perm1},
		},
		{
			name:    "multiple permission added",
			before:  []Permission{},
			after:   []Permission{perm1, perm2},
			added:   []Permission{perm1, perm2},
			removed: []Permission{},
		},
		{
			name:    "single permission removed",
			before:  []Permission{perm1, perm2},
			after:   []Permission{},
			added:   []Permission{},
			removed: []Permission{perm1, perm2},
		},
		{
			name:    "extra permissions",
			before:  []Permission{perm1, perm2},
			after:   []Permission{perm1, perm2, perm3},
			added:   []Permission{perm3},
			removed: []Permission{},
		},
		{
			name:    "less permissions",
			before:  []Permission{perm1, perm2, perm3},
			after:   []Permission{perm1, perm2},
			added:   []Permission{},
			removed: []Permission{perm3},
		},
		{
			name:    "same id different role",
			before:  []Permission{sameidperm1},
			after:   []Permission{sameidperm2},
			added:   []Permission{sameidperm2},
			removed: []Permission{sameidperm1},
		},
		{
			name:    "email based extra permissions",
			before:  []Permission{emailperm1},
			after:   []Permission{emailperm1, emailperm2},
			added:   []Permission{emailperm2},
			removed: []Permission{},
		},
		{
			name:    "email based less permissions",
			before:  []Permission{emailperm1, emailperm2},
			after:   []Permission{emailperm1},
			added:   []Permission{},
			removed: []Permission{emailperm2},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			added, removed := DiffPermissions(test.before, test.after)

			assert.Equal(t, added, test.added, "added permissions")
			assert.Equal(t, removed, test.removed, "removed permissions")
		})
	}
}

func (suite *PermissionsUnitTestSuite) TestDiffLinkShares() {
	entities1 := []Entity{{ID: "e1"}}
	ls1 := LinkShare{
		ID:       "id1",
		Entities: entities1,
		Link:     LinkShareLink{WebURL: "id1"},
	}

	lsempty := LinkShare{
		ID:   "id1",
		Link: LinkShareLink{WebURL: "id1"},
	}

	table := []struct {
		name    string
		before  []LinkShare
		after   []LinkShare
		added   []LinkShare
		removed []LinkShare
	}{
		{
			name:    "single link share added",
			before:  []LinkShare{},
			after:   []LinkShare{ls1},
			added:   []LinkShare{ls1},
			removed: []LinkShare{},
		},
		{
			name:    "empty filtered from before",
			before:  []LinkShare{lsempty},
			after:   []LinkShare{},
			added:   []LinkShare{},
			removed: []LinkShare{},
		},
		{
			name:    "empty filtered from after",
			before:  []LinkShare{},
			after:   []LinkShare{lsempty},
			added:   []LinkShare{},
			removed: []LinkShare{},
		},
		{
			name:    "empty filtered from both",
			before:  []LinkShare{lsempty, ls1},
			after:   []LinkShare{lsempty},
			added:   []LinkShare{},
			removed: []LinkShare{ls1},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			added, removed := DiffLinkShares(test.before, test.after)

			assert.Equal(t, added, test.added, "added link shares")
			assert.Equal(t, removed, test.removed, "removed link shares")
		})
	}
}

func getPermsAndResourceOwnerPerms(
	permID, resourceOwner string,
	gv2t GV2Type,
	scopes []string,
) (models.Permissionable, Permission) {
	sharepointIdentitySet := models.NewSharePointIdentitySet()

	switch gv2t {
	case GV2App, GV2Device, GV2Group, GV2User:
		identity := models.NewIdentity()
		identity.SetId(&resourceOwner)
		identity.SetAdditionalData(map[string]any{"email": &resourceOwner})

		switch gv2t {
		case GV2User:
			sharepointIdentitySet.SetUser(identity)
		case GV2Group:
			sharepointIdentitySet.SetGroup(identity)
		case GV2App:
			sharepointIdentitySet.SetApplication(identity)
		case GV2Device:
			sharepointIdentitySet.SetDevice(identity)
		}

	case GV2SiteUser, GV2SiteGroup:
		spIdentity := models.NewSharePointIdentity()
		spIdentity.SetId(&resourceOwner)
		spIdentity.SetAdditionalData(map[string]any{"email": &resourceOwner})

		switch gv2t {
		case GV2SiteUser:
			sharepointIdentitySet.SetSiteUser(spIdentity)
		case GV2SiteGroup:
			sharepointIdentitySet.SetSiteGroup(spIdentity)
		}
	}

	perm := models.NewPermission()
	perm.SetId(&permID)
	perm.SetRoles([]string{"read"})
	perm.SetGrantedToV2(sharepointIdentitySet)

	ownersPerm := Permission{
		ID:         permID,
		Roles:      []string{"read"},
		EntityID:   resourceOwner,
		EntityType: gv2t,
	}

	return perm, ownersPerm
}

func (suite *PermissionsUnitTestSuite) TestDrivePermissionsFilter() {
	var (
		pID  = "fakePermId"
		uID  = "fakeuser@provider.com"
		uID2 = "fakeuser2@provider.com"
		own  = []string{"owner"}
		r    = []string{"read"}
		rw   = []string{"read", "write"}
	)

	userOwnerPerm, userOwnerROperm := getPermsAndResourceOwnerPerms(pID, uID, GV2User, own)
	userReadPerm, userReadROperm := getPermsAndResourceOwnerPerms(pID, uID, GV2User, r)
	userReadWritePerm, userReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, GV2User, rw)
	siteUserOwnerPerm, siteUserOwnerROperm := getPermsAndResourceOwnerPerms(pID, uID, GV2SiteUser, own)
	siteUserReadPerm, siteUserReadROperm := getPermsAndResourceOwnerPerms(pID, uID, GV2SiteUser, r)
	siteUserReadWritePerm, siteUserReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, GV2SiteUser, rw)

	groupReadPerm, groupReadROperm := getPermsAndResourceOwnerPerms(pID, uID, GV2Group, r)
	groupReadWritePerm, groupReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, GV2Group, rw)
	siteGroupReadPerm, siteGroupReadROperm := getPermsAndResourceOwnerPerms(pID, uID, GV2SiteGroup, r)
	siteGroupReadWritePerm, siteGroupReadWriteROperm := getPermsAndResourceOwnerPerms(pID, uID2, GV2SiteGroup, rw)

	noPerm, _ := getPermsAndResourceOwnerPerms(pID, uID, "user", []string{"read"})
	noPerm.SetGrantedToV2(nil) // eg: link shares

	cases := []struct {
		name              string
		graphPermissions  []models.Permissionable
		parsedPermissions []Permission
	}{
		{
			name:              "no perms",
			graphPermissions:  []models.Permissionable{},
			parsedPermissions: []Permission{},
		},
		{
			name:              "no user bound to perms",
			graphPermissions:  []models.Permissionable{noPerm},
			parsedPermissions: []Permission{},
		},

		// user
		{
			name:              "user with read permissions",
			graphPermissions:  []models.Permissionable{userReadPerm},
			parsedPermissions: []Permission{userReadROperm},
		},
		{
			name:              "user with owner permissions",
			graphPermissions:  []models.Permissionable{userOwnerPerm},
			parsedPermissions: []Permission{userOwnerROperm},
		},
		{
			name:              "user with read and write permissions",
			graphPermissions:  []models.Permissionable{userReadWritePerm},
			parsedPermissions: []Permission{userReadWriteROperm},
		},
		{
			name:              "multiple users with separate permissions",
			graphPermissions:  []models.Permissionable{userReadPerm, userReadWritePerm},
			parsedPermissions: []Permission{userReadROperm, userReadWriteROperm},
		},

		// site-user
		{
			name:              "site user with read permissions",
			graphPermissions:  []models.Permissionable{siteUserReadPerm},
			parsedPermissions: []Permission{siteUserReadROperm},
		},
		{
			name:              "site user with owner permissions",
			graphPermissions:  []models.Permissionable{siteUserOwnerPerm},
			parsedPermissions: []Permission{siteUserOwnerROperm},
		},
		{
			name:              "site user with read and write permissions",
			graphPermissions:  []models.Permissionable{siteUserReadWritePerm},
			parsedPermissions: []Permission{siteUserReadWriteROperm},
		},
		{
			name:              "multiple site users with separate permissions",
			graphPermissions:  []models.Permissionable{siteUserReadPerm, siteUserReadWritePerm},
			parsedPermissions: []Permission{siteUserReadROperm, siteUserReadWriteROperm},
		},

		// group
		{
			name:              "group with read permissions",
			graphPermissions:  []models.Permissionable{groupReadPerm},
			parsedPermissions: []Permission{groupReadROperm},
		},
		{
			name:              "group with read and write permissions",
			graphPermissions:  []models.Permissionable{groupReadWritePerm},
			parsedPermissions: []Permission{groupReadWriteROperm},
		},
		{
			name:              "multiple groups with separate permissions",
			graphPermissions:  []models.Permissionable{groupReadPerm, groupReadWritePerm},
			parsedPermissions: []Permission{groupReadROperm, groupReadWriteROperm},
		},

		// site-group
		{
			name:              "site group with read permissions",
			graphPermissions:  []models.Permissionable{siteGroupReadPerm},
			parsedPermissions: []Permission{siteGroupReadROperm},
		},
		{
			name:              "site group with read and write permissions",
			graphPermissions:  []models.Permissionable{siteGroupReadWritePerm},
			parsedPermissions: []Permission{siteGroupReadWriteROperm},
		},
		{
			name:              "multiple site groups with separate permissions",
			graphPermissions:  []models.Permissionable{siteGroupReadPerm, siteGroupReadWritePerm},
			parsedPermissions: []Permission{siteGroupReadROperm, siteGroupReadWriteROperm},
		},
	}
	for _, tc := range cases {
		suite.Run(tc.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			actual := FilterPermissions(ctx, tc.graphPermissions)
			assert.ElementsMatch(t, tc.parsedPermissions, actual)
		})
	}
}
