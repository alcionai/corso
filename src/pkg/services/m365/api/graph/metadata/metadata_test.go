package metadata_test

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph/metadata"
)

type boolfAssertionFunc func(assert.TestingT, bool, string, ...any) bool

type testCase struct {
	service  path.ServiceType
	category path.CategoryType
	expected boolfAssertionFunc
}

var (
	tenant = "a-tenant"
	user   = "a-user"

	notMetaSuffixes = []string{
		"",
		metadata.DataFileSuffix,
	}

	metaSuffixes = []string{
		metadata.MetaFileSuffix,
		metadata.DirMetaFileSuffix,
	}

	cases = []testCase{
		{
			service:  path.ExchangeService,
			category: path.EmailCategory,
			expected: assert.Falsef,
		},
		{
			service:  path.ExchangeService,
			category: path.ContactsCategory,
			expected: assert.Falsef,
		},
		{
			service:  path.ExchangeService,
			category: path.EventsCategory,
			expected: assert.Falsef,
		},
		{
			service:  path.OneDriveService,
			category: path.FilesCategory,
			expected: assert.Truef,
		},
		{
			service:  path.SharePointService,
			category: path.LibrariesCategory,
			expected: assert.Truef,
		},
		{
			service:  path.SharePointService,
			category: path.ListsCategory,
			expected: assert.Falsef,
		},
		{
			service:  path.SharePointService,
			category: path.PagesCategory,
			expected: assert.Falsef,
		},
		{
			service:  path.GroupsService,
			category: path.LibrariesCategory,
			expected: assert.Truef,
		},
	}
)

type MetadataUnitSuite struct {
	tester.Suite
}

func TestMetadataUnitSuite(t *testing.T) {
	suite.Run(t, &MetadataUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MetadataUnitSuite) TestIsMetadataFile_Files_MetaSuffixes() {
	for _, test := range cases {
		for _, ext := range metaSuffixes {
			suite.Run(fmt.Sprintf("%s %s %s", test.service, test.category, ext), func() {
				t := suite.T()

				p, err := path.Build(
					tenant,
					user,
					test.service,
					test.category,
					true,
					"file"+ext)
				require.NoError(t, err, clues.ToCore(err))

				test.expected(t, metadata.IsMetadataFile(p), "extension %s", ext)
			})
		}
	}
}

func (suite *MetadataUnitSuite) TestIsMetadataFile_Files_NotMetaSuffixes() {
	for _, test := range cases {
		for _, ext := range notMetaSuffixes {
			suite.Run(fmt.Sprintf("%s %s %s", test.service, test.category, ext), func() {
				t := suite.T()

				p, err := path.Build(
					tenant,
					user,
					test.service,
					test.category,
					true,
					"file"+ext)
				require.NoError(t, err, clues.ToCore(err))

				assert.Falsef(t, metadata.IsMetadataFile(p), "extension %s", ext)
			})
		}
	}
}

func (suite *MetadataUnitSuite) TestIsMetadataFile_Directories() {
	suffixes := append(append([]string{}, notMetaSuffixes...), metaSuffixes...)

	for _, test := range cases {
		for _, ext := range suffixes {
			suite.Run(fmt.Sprintf("%s %s %s", test.service, test.category, ext), func() {
				t := suite.T()

				p, err := path.Build(
					tenant,
					user,
					test.service,
					test.category,
					false,
					"file"+ext)
				require.NoError(t, err, clues.ToCore(err))

				assert.Falsef(t, metadata.IsMetadataFile(p), "extension %s", ext)
			})
		}
	}
}

func (suite *MetadataUnitSuite) TestIsMetadataFile() {
	table := []struct {
		name       string
		service    path.ServiceType
		category   path.CategoryType
		isMetaFile bool
		expected   bool
	}{
		{
			name:     "onedrive .data file",
			service:  path.OneDriveService,
			category: path.FilesCategory,
		},
		{
			name:     "sharepoint library .data file",
			service:  path.SharePointService,
			category: path.LibrariesCategory,
		},
		{
			name:     "group library .data file",
			service:  path.GroupsService,
			category: path.LibrariesCategory,
		},
		{
			name:     "group conversations .data file",
			service:  path.GroupsService,
			category: path.ConversationPostsCategory,
		},
		{
			name:       "onedrive .meta file",
			service:    path.OneDriveService,
			category:   path.FilesCategory,
			isMetaFile: true,
			expected:   true,
		},
		{
			name:       "sharepoint library .meta file",
			service:    path.SharePointService,
			category:   path.LibrariesCategory,
			isMetaFile: true,
			expected:   true,
		},
		{
			name:       "group library .meta file",
			service:    path.GroupsService,
			category:   path.LibrariesCategory,
			isMetaFile: true,
			expected:   true,
		},
		{
			name:       "group conversations .meta file",
			service:    path.GroupsService,
			category:   path.ConversationPostsCategory,
			isMetaFile: true,
			expected:   true,
		},
		// For services which don't have metadata files, make sure the function
		// returns false. We don't want .meta suffix (assuming it exists) in
		// these cases to be interpreted as metadata files.
		{
			name:       "exchange service",
			service:    path.ExchangeService,
			category:   path.EmailCategory,
			isMetaFile: true,
		},
		{
			name:       "group channels",
			service:    path.GroupsService,
			category:   path.ChannelMessagesCategory,
			isMetaFile: true,
		},
		{
			name:       "lists",
			service:    path.SharePointService,
			category:   path.ListsCategory,
			isMetaFile: true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			fileName := "file"
			if test.isMetaFile {
				fileName += metadata.MetaFileSuffix
			} else {
				fileName += metadata.DataFileSuffix
			}

			p, err := path.Build(
				"t",
				"u",
				test.service,
				test.category,
				true,
				"some", "path", "for", fileName)
			require.NoError(t, err, clues.ToCore(err))

			actual := metadata.IsMetadataFile(p)
			assert.Equal(t, test.expected, actual)
		})
	}
}
