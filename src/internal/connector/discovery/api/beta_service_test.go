package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/aw"
)

type BetaUnitSuite struct {
	suite.Suite
}

func TestBetaUnitSuite(t *testing.T) {
	suite.Run(t, new(BetaUnitSuite))
}

func (suite *BetaUnitSuite) TestBetaService_Adapter() {
	t := suite.T()
	a := tester.NewMockM365Account(t)
	m365, err := a.M365Config()
	aw.MustNoErr(t, err)

	adpt, err := graph.CreateAdapter(
		m365.AzureTenantID,
		m365.AzureClientID,
		m365.AzureClientSecret,
	)
	aw.MustNoErr(t, err)

	service := NewBetaService(adpt)
	require.NotNil(t, service)

	testPage := models.NewSitePage()
	name := "testFile"
	desc := "working with parsing"

	testPage.SetName(&name)
	testPage.SetDescription(&desc)

	byteArray, err := service.Serialize(testPage)
	assert.NotEmpty(t, byteArray)
	aw.NoErr(t, err)
}
