package metadata

import (
	"testing"

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
