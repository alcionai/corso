package exchange

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/backup/details"
)

type ContactSuite struct {
	suite.Suite
}

func TestContactSuite(t *testing.T) {
	suite.Run(t, &ContactSuite{})
}

func (suite *ContactSuite) TestContactInfo() {
	tests := []struct {
		name         string
		contactAndRP func() (models.Contactable, *details.ExchangeInfo)
	}{
		{
			name: "Empty Contact",
			contactAndRP: func() (models.Contactable, *details.ExchangeInfo) {
				return models.NewContact(), &details.ExchangeInfo{}
			},
		}, {
			name: "Only Name",
			contactAndRP: func() (models.Contactable, *details.ExchangeInfo) {
				aPerson := "Whole Person"
				contact := models.NewContact()
				contact.SetDisplayName(&aPerson)
				return contact, &details.ExchangeInfo{ContactName: aPerson}
			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			contact, expected := test.contactAndRP()
			suite.Equal(expected, ContactInfo(contact))
		})
	}
}
