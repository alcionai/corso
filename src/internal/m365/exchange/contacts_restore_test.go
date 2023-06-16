package exchange

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ContactsRestoreIntgSuite struct {
	tester.Suite
	creds  account.M365Config
	ac     api.Client
	userID string
}

func TestContactsRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &ContactsRestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *ContactsRestoreIntgSuite) SetupSuite() {
	t := suite.T()

	a := tester.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = creds

	suite.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	suite.userID = tester.M365UserID(t)
}

// Testing to ensure that cache system works for in multiple different environments
func (suite *ContactsRestoreIntgSuite) TestCreateContainerDestination() {
	runCreateDestinationTest(
		suite.T(),
		newMailRestoreHandler(suite.ac),
		path.EmailCategory,
		suite.creds.AzureTenantID,
		suite.userID,
		testdata.DefaultRestoreConfig("").Location,
		[]string{"Hufflepuff"},
		[]string{"Ravenclaw"})
}
