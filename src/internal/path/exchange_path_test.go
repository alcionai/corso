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
	tenant = "aTenant"
	user   = "aUser"
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
			user:   user,
			rest:   rest,
		},
		{
			name:   "NoUser",
			tenant: tenant,
			user:   "",
			rest:   rest,
		},
		{
			name:   "NoRest",
			tenant: tenant,
			user:   user,
			rest:   nil,
		},
	}

	modes = []struct {
		name           string
		builderFunc    func(b path.Builder, tenant, user string) (path.Path, error)
		expectedFolder string
		expectedItem   string
	}{
		{
			name:           "Folder",
			builderFunc:    path.Builder.ToDataLayerExchangeMailFolder,
			expectedFolder: strings.Join(rest, "/"),
			expectedItem:   "",
		},
		{
			name:           "Item",
			builderFunc:    path.Builder.ToDataLayerExchangeMailItem,
			expectedFolder: strings.Join(rest[0:len(rest)-1], "/"),
			expectedItem:   rest[len(rest)-1],
		},
	}
)

type ExchangeMailUnitSuite struct {
	suite.Suite
}

func TestExchangeMailUnitSuite(t *testing.T) {
	suite.Run(t, new(ExchangeMailUnitSuite))
}

func (suite *ExchangeMailUnitSuite) TestMissingInfoErrors() {
	for _, m := range modes {
		suite.T().Run(m.name, func(tOuter *testing.T) {
			for _, test := range missingInfo {
				tOuter.Run(test.name, func(t *testing.T) {
					b := path.Builder{}.Append(test.rest...)

					_, err := m.builderFunc(*b, test.tenant, test.user)
					assert.Error(t, err)
				})
			}
		})
	}
}

type PopulatedExchangeMailUnitSuite struct {
	suite.Suite
	b *path.Builder
}

func TestPopulatedExchangeMailUnitSuite(t *testing.T) {
	suite.Run(t, new(PopulatedExchangeMailUnitSuite))
}

func (suite *PopulatedExchangeMailUnitSuite) SetupSuite() {
	suite.b = path.Builder{}.Append(rest...)
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetTenant() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := m.builderFunc(*suite.b, tenant, user)
			require.NoError(t, err)

			assert.Equal(suite.T(), tenant, p.Tenant())
		})
	}
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetUser() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := m.builderFunc(*suite.b, tenant, user)
			require.NoError(t, err)

			assert.Equal(suite.T(), user, p.User())
		})
	}
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetFolder() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := m.builderFunc(*suite.b, tenant, user)
			require.NoError(t, err)

			assert.Equal(suite.T(), m.expectedFolder, p.Folder())
		})
	}
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetItem() {
	for _, m := range modes {
		suite.T().Run(m.name, func(t *testing.T) {
			p, err := m.builderFunc(*suite.b, tenant, user)
			require.NoError(t, err)

			assert.Equal(suite.T(), m.expectedItem, p.Item())
		})
	}
}
