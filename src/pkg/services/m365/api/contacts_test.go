package api

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
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
			assert.Equal(suite.T(), expected, ContactInfo(contact))
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

			result, err := BytesToContactable(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}
