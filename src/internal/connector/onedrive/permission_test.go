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
