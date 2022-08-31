package exchange

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/backup/details"
)

type MessageSuite struct {
	suite.Suite
}

func TestMessageSuite(t *testing.T) {
	suite.Run(t, &MessageSuite{})
}

func (suite *MessageSuite) TestMessageInfo() {
	tests := []struct {
		name     string
		msgAndRP func() (models.Messageable, *details.ExchangeInfo)
	}{
		{
			name: "Empty message",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				i := &details.ExchangeInfo{ItemType: details.ExchangeMail}
				return models.NewMessage(), i
			},
		},
		{
			name: "Just sender",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				sender := "foo@bar.com"
				sr := models.NewRecipient()
				sea := models.NewEmailAddress()
				msg := models.NewMessage()
				sea.SetAddress(&sender)
				sr.SetEmailAddress(sea)
				msg.SetSender(sr)
				i := &details.ExchangeInfo{
					ItemType: details.ExchangeMail,
					Sender:   sender,
				}
				return msg, i
			},
		},
		{
			name: "Just subject",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				subject := "Hello world"
				msg := models.NewMessage()
				msg.SetSubject(&subject)
				i := &details.ExchangeInfo{
					ItemType: details.ExchangeMail,
					Subject:  subject,
				}
				return msg, i
			},
		},
		{
			name: "Just receivedtime",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				now := time.Now()
				msg := models.NewMessage()
				msg.SetReceivedDateTime(&now)
				i := &details.ExchangeInfo{
					ItemType: details.ExchangeMail,
					Received: now,
				}
				return msg, i
			},
		},
		{
			name: "All fields",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				sender := "foo@bar.com"
				subject := "Hello world"
				now := time.Now()
				sr := models.NewRecipient()
				sea := models.NewEmailAddress()
				msg := models.NewMessage()
				sea.SetAddress(&sender)
				sr.SetEmailAddress(sea)
				msg.SetSender(sr)
				msg.SetSubject(&subject)
				msg.SetReceivedDateTime(&now)
				i := &details.ExchangeInfo{
					ItemType: details.ExchangeMail,
					Sender:   sender,
					Subject:  subject,
					Received: now,
				}
				return msg, i
			},
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			msg, expected := tt.msgAndRP()
			suite.Equal(expected, MessageInfo(msg))
		})
	}
}
