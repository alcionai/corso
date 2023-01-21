package api

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

type ContactsAPIUnitSuite struct {
	suite.Suite
}

func TestContactsAPIUnitSuite(t *testing.T) {
	suite.Run(t, new(ContactsAPIUnitSuite))
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
		suite.T().Run(test.name, func(t *testing.T) {
			contact, expected := test.contactAndRP()
			assert.Equal(t, expected, ContactInfo(contact))
		})
	}
}
