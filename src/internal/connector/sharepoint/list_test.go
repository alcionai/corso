package sharepoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type SharePointSuite struct {
	suite.Suite
	userID string
	creds  account.M365Config
}

func (suite *SharePointSuite) SetupSuite() {
	t := suite.T()
	suite.userID = tester.SecondaryM365UserID(suite.T())

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
}

func TestSharePointSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(SharePointSuite))
}

func (suite *SharePointSuite) TestGetSite() {
	ctx, flush := tester.NewContext()
	defer flush()

	service, err := createTestService(suite.creds)
	require.NoError(suite.T(), err)

	err = GetSite(ctx, service, suite.userID)
	assert.NoError(suite.T(), err)

}

func (suite *SharePointSuite) TestLoadList() {
	ctx, flush := tester.NewContext()
	defer flush()
	service, err := createTestService(suite.creds)
	require.NoError(suite.T(), err)

	id := "8qzvrj.sharepoint.com,1c9ef309-f47c-4e69-832b-a83edd69fa7f,c57f6e0e-3e4b-472c-b528-b56a2ccd0507"
	//resp := service.Client().SitesById(id).Lists()
	require.NoError(suite.T(), err)
	_, err = loadLists(ctx, service, id)
	assert.NoError(suite.T(), err)
}

// requestBody := models.NewListItem()
// fells := models.NewFieldValueSet()
// additionalData := map[string]interface{}{
// 	"title":  "Widget",
// 	"color":  "Lavendar",
// 	"weight": int32(32),
// } Post did not work
// fells.SetAdditionalData(additionalData)
// requestBody.SetFields(fells)
// Creating a list item without knowing what list it is for did not work or produce an error
