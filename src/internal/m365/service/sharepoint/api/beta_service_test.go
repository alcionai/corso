package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph/betasdk/models"
)

type BetaUnitSuite struct {
	tester.Suite
}

func TestBetaUnitSuite(t *testing.T) {
	suite.Run(t, &BetaUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BetaUnitSuite) TestBetaService_Adapter() {
	t := suite.T()
	a := tconfig.NewFakeM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	adpt, err := graph.CreateAdapter(
		m365.AzureTenantID,
		m365.AzureClientID,
		m365.AzureClientSecret,
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	service := NewBetaService(adpt)
	require.NotNil(t, service)

	testPage := models.NewSitePage()
	name := "testFile"
	desc := "working with parsing"

	testPage.SetName(&name)
	testPage.SetDescription(&desc)

	byteArray, err := service.Serialize(testPage)
	assert.NotEmpty(t, byteArray)
	assert.NoError(t, err, clues.ToCore(err))
}
