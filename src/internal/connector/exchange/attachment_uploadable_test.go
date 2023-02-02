package exchange

import (
	"testing"
	"time"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ExchangeRestoreAttachmentSuite struct {
	suite.Suite
	gs          graph.Servicer
	credentials account.M365Config
	ac          api.Client
}

func TestExchangeRestoreAttachmentSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoConnectorRestoreExchangeCollectionTests,
	)

	suite.Run(t, new(ExchangeRestoreAttachmentSuite))
}

func (suite *ExchangeRestoreAttachmentSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365
	suite.ac, err = api.NewClient(m365)
	require.NoError(t, err)

	adpt, err := graph.CreateAdapter(m365.AzureTenantID, m365.AzureClientID, m365.AzureClientSecret)
	require.NoError(t, err)

	suite.gs = graph.NewService(adpt)

	require.NoError(suite.T(), err)
}

// TestAttachmentUploads verifies that various attachnments can be uploaded
// to certain base objects.
func (suite *ExchangeRestoreAttachmentSuite) TestAttachmentUploadsMail() {
	ctx, flush := tester.NewContext()
	defer flush()

	// base message --> outlookable ??
	// attachment -->
	t := suite.T()
	//userID := tester.M365UserID(suite.T())

	userID := "dustina@8qzvrj.onmicrosoft.com"
	service, err := createService(suite.ac.Credentials)
	require.NoError(t, err)

	now := time.Now()
	folderName := "TestRestoreMailwithLargeAttachment: " + common.FormatSimpleDateTime(now)
	folder, err := suite.ac.Mail().CreateMailFolder(ctx, userID, folderName)
	require.NoError(t, err)

	folderID := *folder.GetId()
	byteArray := mockconnector.GetMockMessageBytes("Base Message")

	deets, err := RestoreExchangeObject(ctx, byteArray, path.EmailCategory, control.Copy, service, folderID, userID)
	require.NoError(t, err)
	assert.NotNil(t, deets)

	tests := []struct {
		name          string
		getAttachment func(t *testing.T) models.Attachmentable
	}{
		{
			name: "Reference Attachment",
			getAttachment: func(t *testing.T) models.Attachmentable {
				byteArray := mockconnector.GetMockAttachmentReference()

				attach, err := support.CreateAttachmentFromBytes(byteArray)
				require.NoError(t, err)

				return attach
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("Test compiled fine")
		})
	}
	// attach to upload

}
