package onedrive

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type PermissionsUnitTestSuite struct {
	tester.Suite
}

func TestPermissionsUnitTestSuite(t *testing.T) {
	suite.Run(t, &PermissionsUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PermissionsUnitTestSuite) TestComputeParentPermissions() {
	entryPath := fmt.Sprintf(rootDrivePattern, "drive-id") + "/level0/level1/level2/entry"
	rootEntryPath := fmt.Sprintf(rootDrivePattern, "drive-id") + "/entry"

	entry, err := path.Build(
		"tenant",
		"user",
		path.OneDriveService,
		path.FilesCategory,
		false,
		strings.Split(entryPath, "/")...,
	)
	require.NoError(suite.T(), err, "creating path")

	rootEntry, err := path.Build(
		"tenant",
		"user",
		path.OneDriveService,
		path.FilesCategory,
		false,
		strings.Split(rootEntryPath, "/")...,
	)
	require.NoError(suite.T(), err, "creating path")

	level2, err := entry.Dir()
	require.NoError(suite.T(), err, "level2 path")

	level1, err := level2.Dir()
	require.NoError(suite.T(), err, "level1 path")

	level0, err := level1.Dir()
	require.NoError(suite.T(), err, "level0 path")

	metadata := Metadata{
		SharingMode: SharingModeCustom,
		Permissions: []UserPermission{
			{
				Roles:    []string{"write"},
				EntityID: "user-id",
			},
		},
	}

	metadata2 := Metadata{
		SharingMode: SharingModeCustom,
		Permissions: []UserPermission{
			{
				Roles:    []string{"read"},
				EntityID: "user-id",
			},
		},
	}

	inherited := Metadata{
		SharingMode: SharingModeInherited,
		Permissions: []UserPermission{},
	}

	table := []struct {
		name        string
		item        path.Path
		meta        Metadata
		parentPerms map[string]Metadata
	}{
		{
			name:        "root level entry",
			item:        rootEntry,
			meta:        Metadata{},
			parentPerms: map[string]Metadata{},
		},
		{
			name:        "root level directory",
			item:        level0,
			meta:        Metadata{},
			parentPerms: map[string]Metadata{},
		},
		{
			name: "direct parent perms",
			item: entry,
			meta: metadata,
			parentPerms: map[string]Metadata{
				level2.String(): metadata,
			},
		},
		{
			name: "top level parent perms",
			item: entry,
			meta: metadata,
			parentPerms: map[string]Metadata{
				level2.String(): inherited,
				level1.String(): inherited,
				level0.String(): metadata,
			},
		},
		{
			name: "all inherited",
			item: entry,
			meta: Metadata{},
			parentPerms: map[string]Metadata{
				level2.String(): inherited,
				level1.String(): inherited,
				level0.String(): inherited,
			},
		},
		{
			name: "multiple custom permission",
			item: entry,
			meta: metadata,
			parentPerms: map[string]Metadata{
				level2.String(): inherited,
				level1.String(): metadata,
				level0.String(): metadata2,
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			_, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			m, err := computeParentPermissions(test.item, test.parentPerms)
			require.NoError(t, err, "compute permissions")

			assert.Equal(t, m, test.meta)
		})
	}
}

func (suite *PermissionsUnitTestSuite) TestDiffPermissions() {
	perm1 := UserPermission{
		ID:       "id1",
		Roles:    []string{"read"},
		EntityID: "user-id1",
	}

	perm2 := UserPermission{
		ID:       "id2",
		Roles:    []string{"write"},
		EntityID: "user-id2",
	}

	perm3 := UserPermission{
		ID:       "id3",
		Roles:    []string{"write"},
		EntityID: "user-id3",
	}

	// The following two permissions have same id and user but
	// different roles, this is a valid scenario for permissions.
	sameidperm1 := UserPermission{
		ID:       "id0",
		Roles:    []string{"write"},
		EntityID: "user-id0",
	}
	sameidperm2 := UserPermission{
		ID:       "id0",
		Roles:    []string{"read"},
		EntityID: "user-id0",
	}

	emailperm1 := UserPermission{
		ID:    "id1",
		Roles: []string{"read"},
		Email: "email1@provider.com",
	}

	emailperm2 := UserPermission{
		ID:    "id1",
		Roles: []string{"read"},
		Email: "email2@provider.com",
	}

	table := []struct {
		name    string
		before  []UserPermission
		after   []UserPermission
		added   []UserPermission
		removed []UserPermission
	}{
		{
			name:    "single permission added",
			before:  []UserPermission{},
			after:   []UserPermission{perm1},
			added:   []UserPermission{perm1},
			removed: []UserPermission{},
		},
		{
			name:    "single permission removed",
			before:  []UserPermission{perm1},
			after:   []UserPermission{},
			added:   []UserPermission{},
			removed: []UserPermission{perm1},
		},
		{
			name:    "multiple permission added",
			before:  []UserPermission{},
			after:   []UserPermission{perm1, perm2},
			added:   []UserPermission{perm1, perm2},
			removed: []UserPermission{},
		},
		{
			name:    "single permission removed",
			before:  []UserPermission{perm1, perm2},
			after:   []UserPermission{},
			added:   []UserPermission{},
			removed: []UserPermission{perm1, perm2},
		},
		{
			name:    "extra permissions",
			before:  []UserPermission{perm1, perm2},
			after:   []UserPermission{perm1, perm2, perm3},
			added:   []UserPermission{perm3},
			removed: []UserPermission{},
		},
		{
			name:    "less permissions",
			before:  []UserPermission{perm1, perm2, perm3},
			after:   []UserPermission{perm1, perm2},
			added:   []UserPermission{},
			removed: []UserPermission{perm3},
		},
		{
			name:    "same id different role",
			before:  []UserPermission{sameidperm1},
			after:   []UserPermission{sameidperm2},
			added:   []UserPermission{sameidperm2},
			removed: []UserPermission{sameidperm1},
		},
		{
			name:    "email based extra permissions",
			before:  []UserPermission{emailperm1},
			after:   []UserPermission{emailperm1, emailperm2},
			added:   []UserPermission{emailperm2},
			removed: []UserPermission{},
		},
		{
			name:    "email based less permissions",
			before:  []UserPermission{emailperm1, emailperm2},
			after:   []UserPermission{emailperm1},
			added:   []UserPermission{},
			removed: []UserPermission{emailperm2},
		},
		{
			name:    "same permissions", // add and remove one to break inheritance
			before:  []UserPermission{perm1, perm2},
			after:   []UserPermission{perm1, perm2},
			added:   []UserPermission{perm1},
			removed: []UserPermission{perm1},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			_, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			added, removed := diffPermissions(test.before, test.after)

			assert.Equal(t, added, test.added, "added permissions")
			assert.Equal(t, removed, test.removed, "removed permissions")
		})
	}
}
