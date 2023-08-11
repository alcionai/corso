package metadata_test

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	odmetadata "github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/m365/graph/metadata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type boolfAssertionFunc func(assert.TestingT, bool, string, ...any) bool

type testCase struct {
	srs      []path.ServiceResource
	category path.CategoryType
	expected boolfAssertionFunc
}

var (
	tenant = "a-tenant"
	user   = "a-user"

	notMetaSuffixes = []string{
		"",
		odmetadata.DataFileSuffix,
	}

	metaSuffixes = []string{
		odmetadata.MetaFileSuffix,
		odmetadata.DirMetaFileSuffix,
	}

	cases = []testCase{
		{
			srs: []path.ServiceResource{{
				Service:           path.ExchangeService,
				ProtectedResource: user,
			}},
			category: path.EmailCategory,
			expected: assert.Falsef,
		},
		{
			srs: []path.ServiceResource{{
				Service:           path.ExchangeService,
				ProtectedResource: user,
			}},
			category: path.ContactsCategory,
			expected: assert.Falsef,
		},
		{
			srs: []path.ServiceResource{{
				Service:           path.ExchangeService,
				ProtectedResource: user,
			}},
			category: path.EventsCategory,
			expected: assert.Falsef,
		},
		{
			srs: []path.ServiceResource{{
				Service:           path.OneDriveService,
				ProtectedResource: user,
			}},
			category: path.FilesCategory,
			expected: assert.Truef,
		},
		{
			srs: []path.ServiceResource{{
				Service:           path.SharePointService,
				ProtectedResource: user,
			}},
			category: path.LibrariesCategory,
			expected: assert.Truef,
		},
		{
			srs: []path.ServiceResource{{
				Service:           path.SharePointService,
				ProtectedResource: user,
			}},
			category: path.ListsCategory,
			expected: assert.Falsef,
		},
		{
			srs: []path.ServiceResource{{
				Service:           path.SharePointService,
				ProtectedResource: user,
			}},
			category: path.PagesCategory,
			expected: assert.Falsef,
		},
		{
			srs: []path.ServiceResource{
				{
					Service:           path.OneDriveService,
					ProtectedResource: user,
				},
				{
					Service:           path.ExchangeService,
					ProtectedResource: user,
				},
			},
			category: path.EventsCategory,
			expected: assert.Falsef,
		},
		{
			srs: []path.ServiceResource{
				{
					Service:           path.ExchangeService,
					ProtectedResource: user,
				},
				{
					Service:           path.OneDriveService,
					ProtectedResource: user,
				},
			},
			category: path.FilesCategory,
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
			name := []string{}

			for _, sr := range test.srs {
				name = append(name, sr.Service.String())
			}

			name = append(name, test.category.String(), ext)

			suite.Run(strings.Join(name, " "), func() {
				t := suite.T()

				p, err := path.Build(
					tenant,
					test.srs,
					test.category,
					true,
					"file"+ext)
				require.NoError(t, err, clues.ToCore(err))

				test.expected(t, metadata.IsMetadataFilePath(p), "extension %s", ext)
			})
		}
	}
}

func (suite *MetadataUnitSuite) TestIsMetadataFile_Files_NotMetaSuffixes() {
	for _, test := range cases {
		for _, ext := range notMetaSuffixes {
			name := []string{}

			for _, sr := range test.srs {
				name = append(name, sr.Service.String())
			}

			name = append(name, test.category.String(), ext)

			suite.Run(strings.Join(name, " "), func() {
				t := suite.T()

				p, err := path.Build(
					tenant,
					test.srs,
					test.category,
					true,
					"file"+ext)
				require.NoError(t, err, clues.ToCore(err))

				assert.Falsef(t, metadata.IsMetadataFilePath(p), "extension %s", ext)
			})
		}
	}
}

func (suite *MetadataUnitSuite) TestIsMetadataFile_Directories() {
	suffixes := append(append([]string{}, notMetaSuffixes...), metaSuffixes...)

	for _, test := range cases {
		for _, ext := range suffixes {
			name := []string{}

			for _, sr := range test.srs {
				name = append(name, sr.Service.String())
			}

			name = append(name, test.category.String(), ext)

			suite.Run(strings.Join(name, " "), func() {
				t := suite.T()

				p, err := path.Build(
					tenant,
					test.srs,
					test.category,
					false,
					"file"+ext)
				require.NoError(t, err, clues.ToCore(err))

				assert.Falsef(t, metadata.IsMetadataFilePath(p), "extension %s", ext)
			})
		}
	}
}
