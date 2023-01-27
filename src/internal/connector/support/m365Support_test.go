package support

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
)

type DataSupportSuite struct {
	suite.Suite
}

func TestDataSupportSuite(t *testing.T) {
	suite.Run(t, new(DataSupportSuite))
}

// TestCreateMessageFromBytes verifies approved mockdata bytes can
// be successfully transformed into M365 Message data.
func (suite *DataSupportSuite) TestCreateMessageFromBytes() {
	table := []struct {
		name        string
		byteArray   []byte
		checkError  assert.ErrorAssertionFunc
		checkObject assert.ValueAssertionFunc
	}{
		{
			name:        "Empty Bytes",
			byteArray:   make([]byte, 0),
			checkError:  assert.Error,
			checkObject: assert.Nil,
		},
		{
			name:        "aMessage bytes",
			byteArray:   mockconnector.GetMockMessageBytes("m365 mail support test"),
			checkError:  assert.NoError,
			checkObject: assert.NotNil,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateMessageFromBytes(test.byteArray)
			test.checkError(t, err)
			test.checkObject(t, result)
		})
	}
}

// TestCreateContactFromBytes verifies behavior of CreateContactFromBytes
// by ensuring correct error and object output.
func (suite *DataSupportSuite) TestCreateContactFromBytes() {
	table := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       "Empty Bytes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Invalid Bytes",
			byteArray:  []byte("A random sentence doesn't make an object"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid Contact",
			byteArray:  mockconnector.GetMockContactBytes("Support Test"),
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateContactFromBytes(test.byteArray)
			test.checkError(t, err)
			test.isNil(t, result)
		})
	}
}

func (suite *DataSupportSuite) TestCreateEventFromBytes() {
	tests := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       "Empty Byes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Invalid Bytes",
			byteArray:  []byte("Invalid byte stream \"subject:\" Not going to work"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid Event",
			byteArray:  mockconnector.GetDefaultMockEventBytes("Event Test"),
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateEventFromBytes(test.byteArray)
			test.checkError(t, err)
			test.isNil(t, result)
		})
	}
}

func (suite *DataSupportSuite) TestCreateListFromBytes() {
	listBytes, err := mockconnector.GetMockListBytes("DataSupportSuite")
	require.NoError(suite.T(), err)

	tests := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       "Empty Byes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Invalid Bytes",
			byteArray:  []byte("Invalid byte stream \"subject:\" Not going to work"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid List",
			byteArray:  listBytes,
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateListFromBytes(test.byteArray)
			test.checkError(t, err)
			test.isNil(t, result)
		})
	}
}

func (suite *DataSupportSuite) TestHasAttachments() {
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
				message, err := CreateMessageFromBytes(byteArray)
				require.NoError(t, err)
				return message.GetBody()
			},
		},
		{
			name:          "Mock w/ inline attachment",
			hasAttachment: assert.True,
			getBodyable: func(t *testing.T) models.ItemBodyable {
				byteArray := mockconnector.GetMessageWithOneDriveAttachment("Test legacy")
				message, err := CreateMessageFromBytes(byteArray)
				require.NoError(t, err)
				return message.GetBody()
			},
		},
		{
			name:          "Edge Case",
			hasAttachment: assert.True,
			getBodyable: func(t *testing.T) models.ItemBodyable {
				//nolint:lll
				content := "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style type=\"text/css\" style=\"display:none\">\r\n<!--\r\np\r\n\t{margin-top:0;\r\n\tmargin-bottom:0}\r\n-->\r\n</style></head><body dir=\"ltr\"><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Happy New Year,</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">In accordance with TPS report guidelines, there have been questions about how to address our activities SharePoint Cover page. Do you believe this is the best picture?&nbsp;</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><img class=\"FluidPluginCopy ContentPasted0 w-2070 h-1380\" size=\"5854817\" data-outlook-trace=\"F:1|T:1\" src=\"cid:85f4faa3-9851-40c7-ba0a-e63dce1185f9\" style=\"max-width:100%\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Let me know if this meets our culture requirements.</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Warm Regards,</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Dustin</div></body></html>"
				body := models.NewItemBody()
				body.SetContent(&content)
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
