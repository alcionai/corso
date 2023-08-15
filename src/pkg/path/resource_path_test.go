package path_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	testTenant            = "aTenant"
	testProtectedResource = "aProtectedResource"
)

func elemsWithWithoutItem(elems path.Elements) func(isItem bool) path.Elements {
	return func(isItem bool) path.Elements {
		if isItem {
			return elems[:len(elems)-1]
		}

		return elems
	}
}

func itemWithWithoutItem(elems path.Elements) func(isItem bool) string {
	return func(isItem bool) string {
		if isItem {
			return elems[len(elems)-1]
		}

		return ""
	}
}

var (
	// Purposely doesn't have characters that need escaping so it can be easily
	// computed using strings.Join().
	rest = path.Elements{"some", "folder", "path", "with", "possible", "item"}

	missingInfo = []struct {
		name   string
		tenant string
		user   string
		rest   []string
	}{
		{
			name:   "NoTenant",
			tenant: "",
			user:   testProtectedResource,
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
			user:   testProtectedResource,
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

	// Set of acceptable service[/subservice]/category mixtures.
	serviceCategories = []struct {
		name           string
		primaryService path.ServiceType
		category       path.CategoryType
		pathFunc       func(
			tenant, primaryResource string,
			isItem bool,
			suffix path.Elements,
		) (path.Path, error)
		expectFolders func(expect path.Elements) func(isItem bool) path.Elements
		expectItem    func(expect path.Elements) func(isItem bool) string
	}{
		{
			name:           path.ExchangeService.String() + path.EmailCategory.String(),
			primaryService: path.ExchangeService,
			category:       path.EmailCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.ExchangeService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.ExchangeService.String() + path.ContactsCategory.String(),
			primaryService: path.ExchangeService,
			category:       path.ContactsCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.ExchangeService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.ExchangeService.String() + path.EventsCategory.String(),
			primaryService: path.ExchangeService,
			category:       path.EventsCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.ExchangeService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.OneDriveService.String() + path.FilesCategory.String(),
			primaryService: path.OneDriveService,
			category:       path.FilesCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.OneDriveService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.SharePointService.String() + path.LibrariesCategory.String(),
			primaryService: path.SharePointService,
			category:       path.LibrariesCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.SharePointService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.SharePointService.String() + path.ListsCategory.String(),
			primaryService: path.SharePointService,
			category:       path.ListsCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.SharePointService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.SharePointService.String() + path.PagesCategory.String(),
			primaryService: path.SharePointService,
			category:       path.PagesCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.SharePointService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.GroupsService.String() + path.UnknownCategory.String(),
			primaryService: path.GroupsService,
			category:       path.UnknownCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(path.GroupsService, primaryResource)
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
		{
			name:           path.GroupsService.String() + path.SharePointService.String() + path.UnknownCategory.String(),
			primaryService: path.GroupsService,
			category:       path.LibrariesCategory,
			pathFunc: func(
				tenant, primaryResource string,
				isItem bool,
				suffix path.Elements,
			) (path.Path, error) {
				srs, err := path.NewServiceResources(
					path.GroupsService,
					primaryResource,
					path.SharePointService,
					"secondaryProtectedResource")
				if err != nil {
					return nil, err
				}

				return path.Build(tenant, srs, path.PagesCategory, isItem, suffix...)
			},
			expectFolders: elemsWithWithoutItem,
			expectItem:    itemWithWithoutItem,
		},
	}
)

type DataLayerResourcePath struct {
	tester.Suite
}

func TestDataLayerResourcePath(t *testing.T) {
	suite.Run(t, &DataLayerResourcePath{Suite: tester.NewUnitSuite(t)})
}

func (suite *DataLayerResourcePath) SetupSuite() {
	clues.SetHasher(clues.NoHash())
}

func (suite *DataLayerResourcePath) TestMissingInfoErrors() {
	for _, types := range serviceCategories {
		suite.Run(types.primaryService.String()+types.category.String(), func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					for _, test := range missingInfo {
						suite.Run(test.name, func() {
							_, err := types.pathFunc(test.tenant, test.user, m.isItem, rest)
							assert.Error(suite.T(), err, clues.ToCore(err))
						})
					}
				})
			}
		})
	}
}

func (suite *DataLayerResourcePath) TestMailItemNoFolder() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			t := suite.T()

			p, err := test.pathFunc(testTenant, testProtectedResource, true, path.Elements{"item"})
			require.NoError(t, err, clues.ToCore(err))

			assert.Empty(t, p.Folder(false))
			assert.Empty(t, p.Folders())
			assert.Equal(t, test.expectItem(path.Elements{"item"})(true), p.Item())
		})
	}
}

func (suite *DataLayerResourcePath) TestPopFront() {
	expected := path.Builder{}.Append(append(
		[]string{path.ExchangeService.String(), testProtectedResource, path.EmailCategory.String()},
		rest...,
	)...)

	for _, m := range modes {
		suite.Run(m.name, func() {
			t := suite.T()

			pb := path.Builder{}.Append(rest...)
			p, err := pb.ToDataLayerExchangePathForCategory(
				testTenant,
				testProtectedResource,
				path.EmailCategory,
				m.isItem,
			)
			require.NoError(t, err, clues.ToCore(err))

			b := p.PopFront()
			assert.Equal(t, expected.String(), b.String())
		})
	}
}

func (suite *DataLayerResourcePath) TestDir() {
	elements := []string{
		testTenant,
		path.ExchangeService.String(),
		testProtectedResource,
		path.EmailCategory.String(),
	}

	for _, m := range modes {
		suite.Run(m.name, func() {
			pb := path.Builder{}.Append(rest...)
			p, err := pb.ToDataLayerExchangePathForCategory(
				testTenant,
				testProtectedResource,
				path.EmailCategory,
				m.isItem,
			)
			require.NoError(suite.T(), err, clues.ToCore(err))

			for i := 1; i <= len(rest); i++ {
				suite.Run(fmt.Sprintf("%v", i), func() {
					t := suite.T()

					p, err = p.Dir()
					require.NoError(t, err, clues.ToCore(err))

					expected := path.Builder{}.Append(elements...).Append(rest[:len(rest)-i]...)
					assert.Equal(t, expected.String(), p.String())
					assert.Empty(t, p.Item())
				})
			}

			suite.Run("All", func() {
				p, err = p.Dir()
				assert.Error(suite.T(), err)
			})
		})
	}
}

func (suite *DataLayerResourcePath) TestToServiceCategoryMetadataPath() {
	tenant := "a-tenant"
	resource := "a-resource"
	table := []struct {
		name            string
		srs             []path.ServiceResource
		category        path.CategoryType
		postfix         []string
		expectedService path.ServiceType
		check           assert.ErrorAssertionFunc
	}{
		{
			name:            "NoPostfixPasses",
			srs:             []path.ServiceResource{{path.ExchangeService, resource}},
			category:        path.EmailCategory,
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "PostfixPasses",
			srs:             []path.ServiceResource{{path.ExchangeService, resource}},
			category:        path.EmailCategory,
			postfix:         []string{"a", "b"},
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:     "Fails",
			srs:      []path.ServiceResource{{path.ExchangeService, resource}},
			category: path.FilesCategory,
			check:    assert.Error,
		},
		{
			name:            "Passes",
			srs:             []path.ServiceResource{{path.ExchangeService, resource}},
			category:        path.ContactsCategory,
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			srs:             []path.ServiceResource{{path.ExchangeService, resource}},
			category:        path.EventsCategory,
			expectedService: path.ExchangeMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			srs:             []path.ServiceResource{{path.OneDriveService, resource}},
			category:        path.FilesCategory,
			expectedService: path.OneDriveMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			srs:             []path.ServiceResource{{path.SharePointService, resource}},
			category:        path.LibrariesCategory,
			expectedService: path.SharePointMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			srs:             []path.ServiceResource{{path.SharePointService, resource}},
			category:        path.ListsCategory,
			expectedService: path.SharePointMetadataService,
			check:           assert.NoError,
		},
		{
			name:            "Passes",
			srs:             []path.ServiceResource{{path.SharePointService, resource}},
			category:        path.PagesCategory,
			expectedService: path.SharePointMetadataService,
			check:           assert.NoError,
		},
	}

	for _, test := range table {
		name := strings.Join([]string{
			test.name,
			test.srs[0].Service.String(),
			test.category.String(),
		}, "_")

		suite.Run(name, func() {
			t := suite.T()
			pb := path.Builder{}.Append(test.postfix...)

			p, err := pb.ToServiceCategoryMetadataPath(
				tenant,
				test.srs,
				test.category,
				false)
			test.check(t, err, clues.ToCore(err))

			if err == nil {
				assert.Equal(t, test.expectedService, p.ServiceResources()[0])
			}
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
		suite.Run(m.name, func() {
			for _, test := range table {
				suite.Run(test.category.String(), func() {
					t := suite.T()

					p, err := b.ToDataLayerExchangePathForCategory(
						testTenant,
						testProtectedResource,
						test.category,
						m.isItem)
					test.check(t, err, clues.ToCore(err))

					if err != nil {
						return
					}

					assert.Equal(t, testTenant, p.Tenant())
					assert.Equal(t, path.ExchangeService, p.PrimaryService())
					assert.Equal(t, test.category, p.Category())
					assert.Equal(t, testProtectedResource, p.PrimaryProtectedResource())
					assert.Equal(t, strings.Join(m.expectedFolders, "/"), p.Folder(false))
					assert.Equal(t, path.Elements(m.expectedFolders), p.Folders())
					assert.Equal(t, m.expectedItem, p.Item())
				})
			}
		})
	}
}

type PopulatedDataLayerResourcePath struct {
	tester.Suite
	// Bool value is whether the path is an item path or a folder path.
	serviceCategoriesToIsItemToPath map[string]map[bool]path.Path
	isItemToPath                    map[bool]path.Path
}

func TestPopulatedDataLayerResourcePath(t *testing.T) {
	suite.Run(t, &PopulatedDataLayerResourcePath{Suite: tester.NewUnitSuite(t)})
}

func (suite *PopulatedDataLayerResourcePath) SetupSuite() {
	suite.serviceCategoriesToIsItemToPath = map[string]map[bool]path.Path{}

	for _, sc := range serviceCategories {
		m := make(map[bool]path.Path, 2)
		suite.serviceCategoriesToIsItemToPath[sc.name] = m

		for _, is := range []bool{true, false} {
			p, err := sc.pathFunc(testTenant, testProtectedResource, is, rest)
			require.NoError(suite.T(), err, clues.ToCore(err))

			suite.serviceCategoriesToIsItemToPath[sc.name][is] = p
			suite.isItemToPath[is] = p
		}
	}
}

func (suite *PopulatedDataLayerResourcePath) TestTenant() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					p := suite.serviceCategoriesToIsItemToPath[test.name][m.isItem]
					assert.Equal(suite.T(), testTenant, p.Tenant())
				})
			}
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestPrimaryService() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					p := suite.serviceCategoriesToIsItemToPath[test.name][m.isItem]
					assert.Equal(suite.T(), test.primaryService, p.PrimaryService())
				})
			}
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestCategory() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					p := suite.serviceCategoriesToIsItemToPath[test.name][m.isItem]
					assert.Equal(suite.T(), test.category, p.Category())
				})
			}
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestPrimaryProtectedResource() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					p := suite.serviceCategoriesToIsItemToPath[test.name][m.isItem]
					assert.Equal(suite.T(), testProtectedResource, p.PrimaryProtectedResource())
				})
			}
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestFolder() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					p := suite.serviceCategoriesToIsItemToPath[test.name][m.isItem]
					assert.Equal(suite.T(), test.expectFolders(rest)(m.isItem).String(), p.Folder(true))
				})
			}
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestFolders() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					p := suite.serviceCategoriesToIsItemToPath[test.name][m.isItem]
					assert.Equal(suite.T(), test.expectFolders(rest)(m.isItem), p.Folders())
				})
			}
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestItem() {
	for _, test := range serviceCategories {
		suite.Run(test.name, func() {
			for _, m := range modes {
				suite.Run(m.name, func() {
					p := suite.serviceCategoriesToIsItemToPath[test.name][m.isItem]
					assert.Equal(suite.T(), test.expectItem(rest)(m.isItem), p.Item())
				})
			}
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
				"/"),
			expectedItem: "",
		},
	}

	for _, m := range modes {
		suite.Run(m.name, func() {
			for _, test := range isItem {
				suite.Run(test.name, func() {
					t := suite.T()

					newPath, err := suite.isItemToPath[m.isItem].Append(test.hasItem, newElement)

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

func (suite *PopulatedDataLayerResourcePath) TestUpdateParent() {
	cases := []struct {
		name     string
		item     string
		prev     string
		cur      string
		expected string
		updated  bool
	}{
		{
			name:     "basic",
			item:     "folder/item",
			prev:     "folder",
			cur:      "new-folder",
			expected: "new-folder/item",
			updated:  true,
		},
		{
			name:     "long path",
			item:     "folder/folder1/folder2/item",
			prev:     "folder/folder1",
			cur:      "new-folder/new-folder1",
			expected: "new-folder/new-folder1/folder2/item",
			updated:  true,
		},
		{
			name:     "change to shorter path",
			item:     "folder/folder1/folder2/item",
			prev:     "folder/folder1/folder2",
			cur:      "new-folder",
			expected: "new-folder/item",
			updated:  true,
		},
		{
			name:     "change to longer path",
			item:     "folder/item",
			prev:     "folder",
			cur:      "folder/folder1/folder2/folder3",
			expected: "folder/folder1/folder2/folder3/item",
			updated:  true,
		},
		{
			name:     "not parent",
			item:     "folder/folder1/folder2/item",
			prev:     "folder1",
			cur:      "new-folder1",
			expected: "dummy",
			updated:  false,
		},
	}

	buildPath := func(t *testing.T, pth string, isItem bool) path.Path {
		pathBuilder := path.Builder{}.Append(strings.Split(pth, "/")...)
		item, err := pathBuilder.ToDataLayerOneDrivePath("tenant", "user", isItem)
		require.NoError(t, err, "err building path")

		return item
	}

	for _, tc := range cases {
		suite.Run(tc.name, func() {
			t := suite.T()

			item := buildPath(t, tc.item, true)
			prev := buildPath(t, tc.prev, false)
			cur := buildPath(t, tc.cur, false)
			expected := buildPath(t, tc.expected, true)

			updated := item.UpdateParent(prev, cur)
			assert.Equal(t, tc.updated, updated, "path updated")
			if tc.updated {
				assert.Equal(t, expected, item, "modified path")
			}
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestUpdateParent_NoopsNils() {
	oldPB := path.Builder{}.Append("hello", "world")
	newPB := path.Builder{}.Append("hola", "mundo")
	// So we can get a new copy for each test.
	testPBElems := []string{"bar", "baz"}

	table := []struct {
		name  string
		oldPB *path.Builder
		newPB *path.Builder
	}{
		{
			name:  "Nil Prev",
			newPB: newPB,
		},
		{
			name:  "Nil New",
			oldPB: oldPB,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			base := oldPB.Append(testPBElems...)
			expected := base.String()

			assert.False(t, base.UpdateParent(test.oldPB, test.newPB))
			assert.Equal(t, expected, base.String())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestHalves() {
	t := suite.T()

	onlyPrefix, err := path.BuildPrefix(
		"titd",
		[]path.ServiceResource{{path.ExchangeService, "pr"}},
		path.ContactsCategory)
	require.NoError(t, err, clues.ToCore(err))

	fullPath, err := path.Build(
		"tid",
		[]path.ServiceResource{{path.ExchangeService, "pr"}},
		path.ContactsCategory,
		true,
		"fld", "item")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name      string
		dlrp      path.Path
		expectPfx *path.Builder
		expectSfx path.Elements
	}{
		{
			name: "only prefix",
			dlrp: onlyPrefix,
			expectPfx: path.Builder{}.Append(
				"tid",
				path.ExchangeService.String(),
				"pr",
				path.ContactsCategory.String()),
			expectSfx: path.Elements{},
		},
		{
			name: "full path",
			dlrp: fullPath,
			expectPfx: path.Builder{}.Append(
				"tid",
				path.ExchangeService.String(),
				"pr",
				path.ContactsCategory.String()),
			expectSfx: path.Elements{"foo", "bar"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			pfx, sfx := test.dlrp.Halves()
			assert.Equal(t, test.expectPfx, pfx, "prefix")
			assert.Equal(t, test.expectSfx, sfx, "suffix")
		})
	}
}
