package onedrive

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
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
	entryPath := fmt.Sprintf(rootDrivePattern, "drive-id") + "/level0/level1/level2/entry"
	rootEntryPath := fmt.Sprintf(rootDrivePattern, "drive-id") + "/entry"

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
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			m, err := computeParentPermissions(ctx, test.item, test.parentPerms)
			require.NoError(t, err, "compute permissions")

			assert.Equal(t, m, test.meta)
		})
	}
}
