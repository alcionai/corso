package api

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type ExchangeServiceSuite struct {
	suite.Suite
	gs          graph.Servicer
	credentials account.M365Config
}

func TestExchangeServiceSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

	suite.Run(t, new(ExchangeServiceSuite))
}

func (suite *ExchangeServiceSuite) SetupSuite() {
	t := suite.T()
	tester.MustGetEnvSets(t, tester.M365AcctCredEnvs)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365

	adpt, err := graph.CreateAdapter(
		m365.AzureTenantID,
		m365.AzureClientID,
		m365.AzureClientSecret)
	require.NoError(t, err)

	suite.gs = graph.NewService(adpt)
}

func (suite *ExchangeServiceSuite) TestOptionsForCalendars() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Empty Literal",
			params:     []string{},
			checkError: assert.NoError,
		},
		{
			name:       "Invalid Parameter",
			params:     []string{"status"},
			checkError: assert.Error,
		},
		{
			name:       "Invalid Parameters",
			params:     []string{"status", "height", "month"},
			checkError: assert.Error,
		},
		{
			name:       "Valid Parameters",
			params:     []string{"changeKey", "events", "owner"},
			checkError: assert.NoError,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := optionsForCalendars(test.params)
			test.checkError(t, err)
		})
	}
}

// TestOptionsForFolders ensures that approved query options
// are added to the RequestBuildConfiguration. Expected will always be +1
// on than the input as "id" are always included within the select parameters
func (suite *ExchangeServiceSuite) TestOptionsForFolders() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
		expected   int
	}{
		{
			name:       "Valid Folder Option",
			params:     []string{"parentFolderId"},
			checkError: assert.NoError,
			expected:   2,
		},
		{
			name:       "Multiple Folder Options: Valid",
			params:     []string{"displayName", "isHidden"},
			checkError: assert.NoError,
			expected:   3,
		},
		{
			name:       "Invalid Folder option param",
			params:     []string{"status"},
			checkError: assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			config, err := optionsForMailFolders(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(test.expected, len(config.QueryParameters.Select))
			}
		})
	}
}

// TestOptionsForContacts similar to TestExchangeService_optionsForFolders
func (suite *ExchangeServiceSuite) TestOptionsForContacts() {
	tests := []struct {
		name       string
		params     []string
		checkError assert.ErrorAssertionFunc
		expected   int
	}{
		{
			name:       "Valid Contact Option",
			params:     []string{"displayName"},
			checkError: assert.NoError,
			expected:   2,
		},
		{
			name:       "Multiple Contact Options: Valid",
			params:     []string{"displayName", "parentFolderId"},
			checkError: assert.NoError,
			expected:   3,
		},
		{
			name:       "Invalid Contact Option param",
			params:     []string{"status"},
			checkError: assert.Error,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			options, err := optionsForContacts(test.params)
			test.checkError(t, err)
			if err == nil {
				suite.Equal(test.expected, len(options.QueryParameters.Select))
			}
		})
	}
}

// TestGraphQueryFunctions verifies if Query functions APIs
// through Microsoft Graph are functional
func (suite *ExchangeServiceSuite) TestGraphQueryFunctions() {
	ctx, flush := tester.NewContext()
	defer flush()

	c, err := NewClient(suite.credentials)
	require.NoError(suite.T(), err)

	userID := tester.M365UserID(suite.T())
	tests := []struct {
		name     string
		function GraphQuery
	}{
		{
			name:     "GraphQuery: Get All ContactFolders",
			function: c.Contacts().GetAllContactFolderNamesForUser,
		},
		{
			name:     "GraphQuery: Get All Calendars for User",
			function: c.Events().GetAllCalendarNamesForUser,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			response, err := test.function(ctx, userID)
			assert.NoError(t, err)
			assert.NotNil(t, response)
		})
	}
}

//nolint:lll
var stubHTMLContent = "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style type=\"text/css\" style=\"display:none\">\r\n<!--\r\np\r\n\t{margin-top:0;\r\n\tmargin-bottom:0}\r\n-->\r\n</style></head><body dir=\"ltr\"><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Happy New Year,</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">In accordance with TPS report guidelines, there have been questions about how to address our activities SharePoint Cover page. Do you believe this is the best picture?&nbsp;</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><img class=\"FluidPluginCopy ContentPasted0 w-2070 h-1380\" size=\"5854817\" data-outlook-trace=\"F:1|T:1\" src=\"cid:85f4faa3-9851-40c7-ba0a-e63dce1185f9\" style=\"max-width:100%\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Let me know if this meets our culture requirements.</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Warm Regards,</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Dustin</div></body></html>"

func (suite *ExchangeServiceSuite) TestHasAttachments() {
	tests := []struct {
		name          string
		hasAttachment assert.BoolAssertionFunc
		getBodyable   func(t *testing.T) models.ItemBodyable
	}{
		{
			name:          "Mock w/out attachment",
			hasAttachment: assert.False,
			getBodyable: func(t *testing.T) models.ItemBodyable {
				byteArray := mockconnector.GetMockMessageWithBodyBytes(
					"Test",
					"This is testing",
					"This is testing",
				)
				message, err := support.CreateMessageFromBytes(byteArray)
				require.NoError(t, err)
				return message.GetBody()
			},
		},
		{
			name:          "Mock w/ inline attachment",
			hasAttachment: assert.True,
			getBodyable: func(t *testing.T) models.ItemBodyable {
				byteArray := mockconnector.GetMessageWithOneDriveAttachment("Test legacy")
				message, err := support.CreateMessageFromBytes(byteArray)
				require.NoError(t, err)
				return message.GetBody()
			},
		},
		{
			name:          "Edge Case",
			hasAttachment: assert.True,
			getBodyable: func(t *testing.T) models.ItemBodyable {
				body := models.NewItemBody()
				body.SetContent(&stubHTMLContent)
				cat := models.HTML_BODYTYPE
				body.SetContentType(&cat)
				return body
			},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			found := HasAttachments(test.getBodyable(t))
			test.hasAttachment(t, found)
		})
	}
}
