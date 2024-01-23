package metadata

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
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
		ID:   "id2",
		Link: LinkShareLink{WebURL: "id2"},
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

func (suite *PermissionsUnitTestSuite) TestEqual() {
	table := []struct {
		name     string
		perm1    Permission
		perm2    Permission
		expected bool
	}{
		{
			name: "same id no email",
			perm1: Permission{
				Roles:      []string{"read"},
				EntityID:   "user-id1",
				EntityType: GV2User,
			},
			perm2: Permission{
				Roles:      []string{"read"},
				EntityID:   "user-id1",
				EntityType: GV2User,
			},
			expected: true,
		},
		{
			name: "no id same email",
			perm1: Permission{
				Roles:      []string{"read"},
				Email:      "id1@provider.com",
				EntityType: GV2User,
			},
			perm2: Permission{
				Roles:      []string{"read"},
				Email:      "id1@provider.com",
				EntityType: GV2User,
			},
			expected: true,
		},
		{
			// Can happen if user changes email
			name: "same id different email",
			perm1: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				Email:      "id1@provider.com",
				EntityType: GV2User,
			},
			perm2: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				Email:      "id1-new@provider.com",
				EntityType: GV2User,
			},
			expected: true,
		},
		{
			name: "different id different email",
			perm1: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				Email:      "id1@provider.com",
				EntityType: GV2User,
			},
			perm2: Permission{
				EntityID:   "user-id2",
				Roles:      []string{"read"},
				Email:      "id2@provider.com",
				EntityType: GV2User,
			},
			expected: false,
		},
		{
			name: "different id same email",
			perm1: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				Email:      "id1@provider.com",
				EntityType: GV2User,
			},
			perm2: Permission{
				EntityID:   "user-id2",
				Roles:      []string{"read"},
				Email:      "id1@provider.com",
				EntityType: GV2User,
			},
			expected: false,
		},
		{
			name: "one with id one with email",
			perm1: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			perm2: Permission{
				Email:      "id2@provider.com",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			expected: false,
		},
		{
			name: "same email one with no id",
			perm1: Permission{
				EntityID:   "user-id1",
				Email:      "id1@provider.com",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			perm2: Permission{
				Email:      "id1@provider.com",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			expected: false,
		},
		{
			// should not ideally happen, not entirely sure if it
			// should be false as we could just be missing the id
			name: "same id one with no email",
			perm1: Permission{
				EntityID:   "user-id1",
				Email:      "id1@provider.com",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			perm2: Permission{
				Email:      "id1@provider.com",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			expected: false,
		},
		{
			name: "same id different role",
			perm1: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			perm2: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"write"},
				EntityType: GV2User,
			},
			expected: false,
		},
		{
			name: "same id same role",
			perm1: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			perm2: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			expected: true,
		},
		{
			name: "same email different role",
			perm1: Permission{
				Email:      "id1@provider.com",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			perm2: Permission{
				Email:      "id1@provider.com",
				Roles:      []string{"write"},
				EntityType: GV2User,
			},
			expected: false,
		},
		{
			name: "same id different entity type",
			perm1: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				EntityType: GV2User,
			},
			perm2: Permission{
				EntityID:   "user-id1",
				Roles:      []string{"read"},
				EntityType: GV2Group,
			},
			expected: false,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, test.expected, test.perm1.Equals(test.perm2), "permissions equality")
		})
	}
}

func (suite *PermissionsUnitTestSuite) TestFilterLinkShares() {
	table := []struct {
		name     string
		perms    func() []models.Permissionable
		expected [][]Entity
	}{
		{
			name: "with id and email",
			perms: func() []models.Permissionable {
				perm1 := models.NewPermission()
				perm1.SetId(ptr.To("id1"))
				perm1.SetRoles([]string{"read"})

				spi11 := models.NewUser()
				spi11.SetId(ptr.To("user-id1"))
				spi11.SetAdditionalData(map[string]any{"email": ptr.To("id1@provider")})

				spi12 := models.NewUser()
				spi12.SetId(ptr.To("user-id2"))
				spi12.SetAdditionalData(map[string]any{"email": ptr.To("id2@provider")})

				gv21 := models.NewSharePointIdentitySet()
				gv21.SetUser(spi11)

				gv22 := models.NewSharePointIdentitySet()
				gv22.SetUser(spi12)

				perm1.SetGrantedToIdentitiesV2([]models.SharePointIdentitySetable{gv21, gv22})

				li1 := models.NewSharingLink()
				li1.SetWebUrl(ptr.To("https://link1"))
				perm1.SetLink(li1)

				return []models.Permissionable{perm1}
			},
			expected: [][]Entity{
				{
					{ID: "user-id1", Email: "id1@provider", EntityType: GV2User},
					{ID: "user-id2", Email: "id2@provider", EntityType: GV2User},
				},
			},
		},
		{
			name: "only email",
			perms: func() []models.Permissionable {
				perm1 := models.NewPermission()
				perm1.SetId(ptr.To("id1"))
				perm1.SetRoles([]string{"read"})

				spi11 := models.NewUser()
				spi11.SetAdditionalData(map[string]any{"email": ptr.To("id1@provider")})

				spi12 := models.NewUser()
				spi12.SetAdditionalData(map[string]any{"email": ptr.To("id2@provider")})

				gv21 := models.NewSharePointIdentitySet()
				gv21.SetUser(spi11)

				gv22 := models.NewSharePointIdentitySet()
				gv22.SetUser(spi12)

				perm1.SetGrantedToIdentitiesV2([]models.SharePointIdentitySetable{gv21, gv22})

				li1 := models.NewSharingLink()
				li1.SetWebUrl(ptr.To("https://link1"))
				perm1.SetLink(li1)

				return []models.Permissionable{perm1}
			},
			expected: [][]Entity{
				{
					{Email: "id1@provider", EntityType: GV2User},
					{Email: "id2@provider", EntityType: GV2User},
				},
			},
		},
		{
			name: "one with id one with email",
			perms: func() []models.Permissionable {
				perm1 := models.NewPermission()
				perm1.SetId(ptr.To("id1"))
				perm1.SetRoles([]string{"read"})

				spi11 := models.NewUser()
				spi11.SetId(ptr.To("user-id1"))
				spi11.SetAdditionalData(map[string]any{"email": ptr.To("id1@provider")})

				spi12 := models.NewUser()
				spi12.SetAdditionalData(map[string]any{"email": ptr.To("id2@provider")})

				gv21 := models.NewSharePointIdentitySet()
				gv21.SetUser(spi11)

				gv22 := models.NewSharePointIdentitySet()
				gv22.SetUser(spi12)

				perm1.SetGrantedToIdentitiesV2([]models.SharePointIdentitySetable{gv21, gv22})

				li1 := models.NewSharingLink()
				li1.SetWebUrl(ptr.To("https://link1"))
				perm1.SetLink(li1)

				return []models.Permissionable{perm1}
			},
			expected: [][]Entity{
				{
					{ID: "user-id1", Email: "id1@provider", EntityType: GV2User},
					{Email: "id2@provider", EntityType: GV2User},
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			actual := FilterLinkShares(ctx, test.perms())
			assert.Equal(t, len(test.expected), len(actual), "number of link shares")

			for i, expected := range test.expected {
				assert.ElementsMatch(t, expected, actual[i].Entities, "link share entities")
			}
		})
	}
}
