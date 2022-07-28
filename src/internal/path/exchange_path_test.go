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
	item   = "anItem"
)

var (
	// Purposely doesn't have characters that need escaping so it can be easily
	// computed using strings.Join().
	folder = []string{"some", "folder", "path"}

	missingInfo = []struct {
		name   string
		tenant string
		user   string
		folder []string
		item   string
	}{
		{
			name:   "NoTenant",
			tenant: "",
			user:   user,
			folder: folder,
			item:   item,
		},
		{
			name:   "NoUser",
			tenant: tenant,
			user:   "",
			folder: folder,
			item:   item,
		},
		{
			name:   "NoFolder",
			tenant: "",
			user:   user,
			folder: nil,
			item:   item,
		},
		{
			name:   "EmptyFolder",
			tenant: "",
			user:   user,
			folder: []string{"", ""},
			item:   item,
		},
		{
			name:   "NoItem",
			tenant: tenant,
			user:   user,
			folder: folder,
			item:   "",
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
	for _, test := range missingInfo {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := path.NewExchangeMail(
				test.tenant, test.user, test.folder, test.item)
			assert.Error(t, err)
		})
	}
}

func (suite *ExchangeMailUnitSuite) TestMissingInfoWithSegmentsErrors() {
	for _, test := range missingInfo {
		suite.T().Run(test.name, func(t *testing.T) {
			folders := strings.Join(test.folder, "")

			_, err := path.NewExchangeMailFromEscapedSegments(
				test.tenant, test.user, folders, test.item)
			assert.Error(t, err)
		})
	}
}

// Some simple escaping examples. Don't want to duplicate everything that is in
// the regular path.Base tests.
func (suite *ExchangeMailUnitSuite) TestNewExchangeMailFromRaw() {
	t := suite.T()
	localItem := `an\item`

	em, err := path.NewExchangeMail(tenant, user, folder, localItem)
	require.NoError(t, err)

	assert.Equal(t, `an\\item`, em.Item())
}

func (suite *ExchangeMailUnitSuite) TestNewExchangeMailFromEscaped() {
	t := suite.T()
	localItem := `an\\item`
	localFolder := strings.Join(folder, "/")

	em, err := path.NewExchangeMailFromEscapedSegments(tenant, user, localFolder, localItem)
	require.NoError(t, err)

	assert.Equal(t, localItem, em.Item())
}

func (suite *ExchangeMailUnitSuite) TestNewExchangeMailFromEscaped_Errors() {
	t := suite.T()
	localItem := `an\item`
	localFolder := strings.Join(folder, "/")

	_, err := path.NewExchangeMailFromEscapedSegments(tenant, user, localFolder, localItem)
	assert.Error(t, err)
}

type PopulatedExchangeMailUnitSuite struct {
	suite.Suite
	em *path.ExchangeMail
}

func TestPopulatedExchangeMailUnitSuite(t *testing.T) {
	suite.Run(t, new(PopulatedExchangeMailUnitSuite))
}

func (suite *PopulatedExchangeMailUnitSuite) SetupTest() {
	em, err := path.NewExchangeMail(tenant, user, folder, item)
	require.NoError(suite.T(), err)

	suite.em = em
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetTenant() {
	assert.Equal(suite.T(), tenant, suite.em.Tenant())
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetUser() {
	assert.Equal(suite.T(), user, suite.em.User())
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetFolder() {
	assert.Equal(suite.T(), strings.Join(folder, "/"), suite.em.Folder())
}

func (suite *PopulatedExchangeMailUnitSuite) TestGetItem() {
	assert.Equal(suite.T(), item, suite.em.Item())
}
