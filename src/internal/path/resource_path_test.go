package path_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/path"
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
		name             string
		isItem           bool
		expectedFolder   string
		expectedItem     string
		expectedService  path.ServiceType
		expectedCategory path.CategoryType
	}{
		{
			name:             "ExchangeMailFolder",
			isItem:           false,
			expectedFolder:   strings.Join(rest, "/"),
			expectedItem:     "",
			expectedService:  path.ExchangeService,
			expectedCategory: path.EmailCategory,
		},
		{
			name:             "ExchangeMailItem",
			isItem:           true,
			expectedFolder:   strings.Join(rest[0:len(rest)-1], "/"),
			expectedItem:     rest[len(rest)-1],
			expectedService:  path.ExchangeService,
			expectedCategory: path.EmailCategory,
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
	for _, m := range modes {
		suite.T().Run(m.name, func(tOuter *testing.T) {
			for _, test := range missingInfo {
				tOuter.Run(test.name, func(t *testing.T) {
					b := path.Builder{}.Append(test.rest...)

					_, err := b.ToDataLayerExchangeMailPath(test.tenant, test.user, m.isItem)
					assert.Error(t, err)
				})
			}
		})
	}
}

func (suite *DataLayerResourcePath) TestMailItemNoFolder() {
	t := suite.T()
	item := "item"
	b := path.Builder{}.Append(item)

	p, err := b.ToDataLayerExchangeMailPath(testTenant, testUser, true)
	require.NoError(t, err)

	assert.Empty(t, p.Folder())
	assert.Equal(t, item, p.Item())
}

type PopulatedDataLayerResourcePath struct {
	suite.Suite
	b *path.Builder
}

func TestPopulatedDataLayerResourcePath(t *testing.T) {
	suite.Run(t, new(PopulatedDataLayerResourcePath))
}

func (suite *PopulatedDataLayerResourcePath) SetupSuite() {
	suite.b = path.Builder{}.Append(rest...)
}

func (suite *PopulatedDataLayerResourcePath) TestTenant() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := suite.b.ToDataLayerExchangeMailPath(testTenant, testUser, m.isItem)
			require.NoError(t, err)

			assert.Equal(t, testTenant, p.Tenant())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestService() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := suite.b.ToDataLayerExchangeMailPath(testTenant, testUser, m.isItem)
			require.NoError(t, err)

			assert.Equal(t, m.expectedService, p.Service())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestCategory() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := suite.b.ToDataLayerExchangeMailPath(testTenant, testUser, m.isItem)
			require.NoError(t, err)

			assert.Equal(t, m.expectedCategory, p.Category())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestResourceOwner() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := suite.b.ToDataLayerExchangeMailPath(testTenant, testUser, m.isItem)
			require.NoError(t, err)

			assert.Equal(t, testUser, p.ResourceOwner())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestFolder() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := suite.b.ToDataLayerExchangeMailPath(testTenant, testUser, m.isItem)
			require.NoError(t, err)

			assert.Equal(t, m.expectedFolder, p.Folder())
		})
	}
}

func (suite *PopulatedDataLayerResourcePath) TestItem() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := suite.b.ToDataLayerExchangeMailPath(testTenant, testUser, m.isItem)
			require.NoError(t, err)

			assert.Equal(t, m.expectedItem, p.Item())
		})
	}
}
