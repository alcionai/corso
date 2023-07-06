package api_test

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/config"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ContactsAPIUnitSuite struct {
	tester.Suite
}

func TestContactsAPIUnitSuite(t *testing.T) {
	suite.Run(t, &ContactsAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ContactsAPIUnitSuite) TestContactInfo() {
	initial := time.Now()

	tests := []struct {
		name         string
		contactAndRP func() (models.Contactable, *details.ExchangeInfo)
	}{
		{
			name: "Empty Contact",
			contactAndRP: func() (models.Contactable, *details.ExchangeInfo) {
				contact := models.NewContact()
				contact.SetCreatedDateTime(&initial)
				contact.SetLastModifiedDateTime(&initial)

				i := &details.ExchangeInfo{
					ItemType: details.ExchangeContact,
					Created:  initial,
					Modified: initial,
				}
				return contact, i
			},
		}, {
			name: "Only Name",
			contactAndRP: func() (models.Contactable, *details.ExchangeInfo) {
				aPerson := "Whole Person"
				contact := models.NewContact()
				contact.SetCreatedDateTime(&initial)
				contact.SetLastModifiedDateTime(&initial)
				contact.SetDisplayName(&aPerson)
				i := &details.ExchangeInfo{
					ItemType:    details.ExchangeContact,
					ContactName: aPerson,
					Created:     initial,
					Modified:    initial,
				}
				return contact, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			contact, expected := test.contactAndRP()
			assert.Equal(suite.T(), expected, api.ContactInfo(contact))
		})
	}
}

func (suite *ContactsAPIUnitSuite) TestBytesToContactable() {
	table := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       "empty bytes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "invalid bytes",
			byteArray:  []byte("A random sentence doesn't make an object"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid Contact",
			byteArray:  exchMock.ContactBytes("Support Test"),
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := api.BytesToContactable(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}

type ContactsAPIIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestContactsAPIntgSuite(t *testing.T) {
	suite.Run(t, &ContactsAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{config.M365AcctCredEnvs}),
	})
}

func (suite *ContactsAPIIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ContactsAPIIntgSuite) TestContacts_GetContainerByName() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// contacts cannot filter for the parent "contacts" folder, so we
	// have to hack this by creating a folder to match beforehand.

	rc := testdata.DefaultRestoreConfig("contacts_api")

	cc, err := suite.its.ac.Contacts().CreateContainer(
		ctx,
		suite.its.userID,
		"",
		rc.Location)
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name      string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      ptr.Val(cc.GetDisplayName()),
			expectErr: assert.NoError,
		},
		{
			name:      "smarfs",
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := suite.its.ac.
				Contacts().
				GetContainerByName(ctx, suite.its.userID, "", test.name)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
