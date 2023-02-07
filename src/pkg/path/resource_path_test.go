package path_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/path"
)

const (
	testTenant = "aTenant"
	testUser   = "aUser"
)

var (
	// Purposely doesn't have characters that need escaping so it can be easily
	// computed using strings.Join().
	rest = []string{"some", "folder", "path", "with", "possible", "item"}

	missingInfo = []struct {
		name   string
		tenant string
		user   string
		rest   []string
	}{
		{
			name:   "NoTenant",
			tenant: "",
			user:   testUser,
			rest:   rest,
		},
		{
			name:   "NoResourceOwner",
			tenant: testTenant,
			user:   "",
			rest:   rest,
		},
		{
			name:   "NoFolderOrItem",
			tenant: testTenant,
			user:   testUser,
			rest:   nil,
		},
	}

	modes = []struct {
		name            string
		isItem          bool
		expectedFolders []string
		expectedItem    string
	}{
		{
			name:            "Folder",
			isItem:          false,
			expectedFolders: rest,
			expectedItem:    "",
		},
		{
			name:            "Item",
			isItem:          true,
			expectedFolders: rest[:len(rest)-1],
			expectedItem:    rest[len(rest)-1],
		},
	}

	// Set of acceptable service/category mixtures.
	serviceCategories = []struct {
		service  path.ServiceType
		category path.CategoryType
		pathFunc func(pb *path.Builder, tenant, user string, isItem bool) (path.Path, error)
	}{
		{
			service:  path.ExchangeService,
			category: path.EmailCategory,
			pathFunc: func(pb *path.Builder, tenant, user string, isItem bool) (path.Path, error) {
				return pb.ToDataLayerExchangePathForCategory(tenant, user, path.EmailCategory, isItem)
			},
		},
		{
			service:  path.ExchangeService,
			category: path.ContactsCategory,
			pathFunc: func(pb *path.Builder, tenant, user string, isItem bool) (path.Path, error) {
				return pb.ToDataLayerExchangePathForCategory(tenant, user, path.ContactsCategory, isItem)
			},
		},
		{
			service:  path.ExchangeService,
			category: path.EventsCategory,
			pathFunc: func(pb *path.Builder, tenant, user string, isItem bool) (path.Path, error) {
				return pb.ToDataLayerExchangePathForCategory(tenant, user, path.EventsCategory, isItem)
			},
		},
		{
			service:  path.OneDriveService,
			category: path.FilesCategory,
			pathFunc: func(pb *path.Builder, tenant, user string, isItem bool) (path.Path, error) {
				return pb.ToDataLayerOneDrivePath(tenant, user, isItem)
			},
		},
		{
			service:  path.SharePointService,
			category: path.LibrariesCategory,
			pathFunc: func(pb *path.Builder, tenant, site string, isItem bool) (path.Path, error) {
				return pb.ToDataLayerSharePointPath(tenant, site, path.LibrariesCategory, isItem)
			},
		},
		{
			service:  path.SharePointService,
			category: path.ListsCategory,
			pathFunc: func(pb *path.Builder, tenant, site string, isItem bool) (path.Path, error) {
				return pb.ToDataLayerSharePointPath(tenant, site, path.ListsCategory, isItem)
			},
		},
		{
			service:  path.SharePointService,
			category: path.PagesCategory,
			pathFunc: func(pb *path.Builder, tenant, site string, isItem bool) (path.Path, error) {
				return pb.ToDataLayerSharePointPath(tenant, site, path.PagesCategory, isItem)
			},
		},
	}
)

type DataLayerResourcePath struct {
	suite.Suite
}

func TestDataLayerResourcePath(t *testing.T) {
	suite.Run(t, new(DataLayerResourcePath))
}

func (suite *DataLayerResourcePath) TestMissingInfoErrors() {
	for _, types := range serviceCategories {
		suite.T().Run(types.service.String()+types.category.String(), func(t1 *testing.T) {
			for _, m := range modes {
				t1.Run(m.name, func(t2 *testing.T) {
					for _, test := range missingInfo {
						t2.Run(test.name, func(t *testing.T) {
							b := path.Builder{}.Append(test.rest...)

							_, err := types.pathFunc(
								b,
								test.tenant,
								test.user,
								m.isItem,
							)
							assert.Error(t, err)
						})
					}
				})
			}
		})
	}
}

func (suite *DataLayerResourcePath) TestMailItemNoFolder() {
	item := "item"
	b := path.Builder{}.Append(item)

	for _, types := range serviceCategories {
		suite.T().Run(types.service.String()+types.category.String(), func(t *testing.T) {
			p, err := types.pathFunc(
				b,
				testTenant,
				testUser,
				true,
			)
			require.NoError(t, err)

			assert.Empty(t, p.Folder(false))
			assert.Empty(t, p.Folders())
			assert.Equal(t, item, p.Item())
		})
	}
}

func (suite *DataLayerResourcePath) TestPopFront() {
	expected := path.Builder{}.Append(append(
		[]string{path.ExchangeService.String(), testUser, path.EmailCategory.String()},
		rest...,
	)...)

	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			pb := path.Builder{}.Append(rest...)
			p, err := pb.ToDataLayerExchangePathForCategory(
				testTenant,
				testUser,
				path.EmailCategory,
				m.isItem,
			)
			require.NoError(t, err)

			b := p.PopFront()
			assert.Equal(t, expected.String(), b.String())
		})
	}
}

func (suite *DataLayerResourcePath) TestDir() {
	elements := []string{
		testTenant,
		path.ExchangeService.String(),
		testUser,
		path.EmailCategory.String(),
	}

	for _, m := range modes {
		suite.T().Run(m.name, func(t1 *testing.T) {
			pb := path.Builder{}.Append(rest...)
			p, err := pb.ToDataLayerExchangePathForCategory(
				testTenant,
				testUser,
				path.EmailCategory,
				m.isItem,
			)
			require.NoError(t1, err)

			for i := 1; i <= len(rest); i++ {
				t1.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
					p, err = p.Dir()
					require.NoError(t, err)

					expected := path.Builder{}.Append(elements...).Append(rest[:len(rest)-i]...)
					assert.Equal(t, expected.String(), p.String())
					assert.Empty(t, p.Item())
				})
			}

			t1.Run("All", func(t *testing.T) {
				p, err = p.Dir()
				assert.Error(t, err)
			})
		})
	}
}

func (suite *DataLayerResourcePath) TestToServiceCategoryMetadataPath() {
	tenant := "a-tenant"
	user := "a-user"
	table := []struct {
		name            string
		service         path.ServiceType
		category        path.CategoryType
		postfix         []string
		expectedService path.ServiceType
		check           assert.ErrorAssertionFunc
	}{
		{
			name:            "NoPostfixPasses",
			service:         path.ExchangeService,
			category:        path.EmailCategory,
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "PostfixPasses",
			service:         path.ExchangeService,
			category:        path.EmailCategory,
			postfix:         []string{"a", "b"},
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:     "Fails",
			service:  path.ExchangeService,
			category: path.FilesCategory,
			check:    assert.Error,
		},
		{
			name:            "Passes",
			service:         path.ExchangeService,
			category:        path.ContactsCategory,
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			service:         path.ExchangeService,
			category:        path.EventsCategory,
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			service:         path.OneDriveService,
			category:        path.FilesCategory,
			expectedService: path.OneDriveMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			service:         path.SharePointService,
			category:        path.LibrariesCategory,
			expectedService: path.SharePointMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			service:         path.SharePointService,
			category:        path.ListsCategory,
			expectedService: path.SharePointMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			service:         path.SharePointService,
			category:        path.PagesCategory,
			expectedService: path.SharePointMetadataService,
			check:           assert.NoError,
		},
	}

	for _, test := range table {
		suite.T().Run(strings.Join([]string{
			test.name,
			test.service.String(),
			test.category.String(),
		}, "_"), func(t *testing.T) {
			pb := path.Builder{}.Append(test.postfix...)
			p, err := pb.ToServiceCategoryMetadataPath(
				tenant,
				user,
				test.service,
				test.category,
				false,
			)

			test.check(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, test.expectedService, p.Service())
		})
	}
}

func (suite *DataLayerResourcePath) TestToExchangePathForCategory() {
	b := path.Builder{}.Append(rest...)
	table := []struct {
		category path.CategoryType
		check    assert.ErrorAssertionFunc
	}{
		{
			category: path.UnknownCategory,
			check:    assert.Error,
		},
		{
			category: path.CategoryType(-1),
			check:    assert.Error,
		},
		{
			category: path.EmailCategory,
			check:    assert.NoError,
		},
		{
			category: path.ContactsCategory,
			check:    assert.NoError,
		},
		{
			category: path.EventsCategory,
			check:    assert.NoError,
		},
	}

	for _, m := range modes {
		suite.T().Run(m.name, func(t1 *testing.T) {
			for _, test := range table {
				t1.Run(test.category.String(), func(t *testing.T) {
					p, err := b.ToDataLayerExchangePathForCategory(
						testTenant,
						testUser,
						test.category,
						m.isItem,
					)

					test.check(t, err)

					if err != nil {
						return
					}

					assert.Equal(t, testTenant, p.Tenant())
					assert.Equal(t, path.ExchangeService, p.Service())
					assert.Equal(t, test.category, p.Category())
					assert.Equal(t, testUser, p.ResourceOwner())
					assert.Equal(t, strings.Join(m.expectedFolders, "/"), p.Folder(false))
					assert.Equal(t, m.expectedFolders, p.Folders())
					assert.Equal(t, m.expectedItem, p.Item())
				})
			}
		})
	}
}

type PopulatedDataLayerResourcePath struct {
	suite.Suite
	// Bool value is whether the path is an item path or a folder path.
	paths map[bool]path.Path
}

func TestPopulatedDataLayerResourcePath(t *testing.T) {
	suite.Run(t, new(PopulatedDataLayerResourcePath))
}

func (suite *PopulatedDataLayerResourcePath) SetupSuite() {
	suite.paths = make(map[bool]path.Path, 2)
	base := path.Builder{}.Append(rest...)

	for _, t := range []bool{true, false} {
		p, err := base.ToDataLayerExchangePathForCategory(
			testTenant,
			testUser,
			path.EmailCategory,
			t,
		)
		require.NoError(suite.T(), err)

		suite.paths[t] = p
	}
}

func (suite *PopulatedDataLayerResourcePath) TestTenant() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			assert.Equal(t, testTenant, suite.paths[m.isItem].Tenant())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestService() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			assert.Equal(t, path.ExchangeService, suite.paths[m.isItem].Service())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestCategory() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			assert.Equal(t, path.EmailCategory, suite.paths[m.isItem].Category())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestResourceOwner() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			assert.Equal(t, testUser, suite.paths[m.isItem].ResourceOwner())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestFolder() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			assert.Equal(
				t,
				strings.Join(m.expectedFolders, "/"),
				suite.paths[m.isItem].Folder(false),
			)
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestFolders() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			assert.Equal(t, m.expectedFolders, suite.paths[m.isItem].Folders())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestItem() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			assert.Equal(t, m.expectedItem, suite.paths[m.isItem].Item())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestAppend() {
	newElement := "someElement"
	isItem := []struct {
		name    string
		hasItem bool
		// Used if the starting path is a folder.
		expectedFolder string
		expectedItem   string
	}{
		{
			name:           "Item",
			hasItem:        true,
			expectedFolder: strings.Join(rest, "/"),
			expectedItem:   newElement,
		},
		{
			name:    "Directory",
			hasItem: false,
			expectedFolder: strings.Join(
				append(append([]string{}, rest...), newElement),
				"/",
			),
			expectedItem: "",
		},
	}

	for _, m := range modes {
		suite.T().Run(m.name, func(t1 *testing.T) {
			for _, test := range isItem {
				t1.Run(test.name, func(t *testing.T) {
					newPath, err := suite.paths[m.isItem].Append(newElement, test.hasItem)

					// Items don't allow appending.
					if m.isItem {
						assert.Error(t, err)
						return
					}

					assert.Equal(t, test.expectedFolder, newPath.Folder(false))
					assert.Equal(t, test.expectedItem, newPath.Item())
				})
			}
		})
	}
}
