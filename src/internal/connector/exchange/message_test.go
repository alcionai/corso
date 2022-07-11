package exchange

import (
	"testing"
	"time"

	"github.com/alcionai/corso/pkg/backup"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/suite"
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
		msgAndRP func() (models.Messageable, *backup.ExchangeInfo)
	}{
		{
			name: "Empty message",
			msgAndRP: func() (models.Messageable, *backup.ExchangeInfo) {
				return models.NewMessage(), &backup.ExchangeInfo{}
			},
		},
		{
			name: "Just sender",
			msgAndRP: func() (models.Messageable, *backup.ExchangeInfo) {
				sender := "foo@bar.com"
				sr := models.NewRecipient()
				sea := models.NewEmailAddress()
				msg := models.NewMessage()
				sea.SetAddress(&sender)
				sr.SetEmailAddress(sea)
				msg.SetSender(sr)
				return msg, &backup.ExchangeInfo{Sender: sender}
			},
		},
		{
			name: "Just subject",
			msgAndRP: func() (models.Messageable, *backup.ExchangeInfo) {
				subject := "Hello world"
				msg := models.NewMessage()
				msg.SetSubject(&subject)
				return msg, &backup.ExchangeInfo{Subject: subject}
			},
		},
		{
			name: "Just receivedtime",
			msgAndRP: func() (models.Messageable, *backup.ExchangeInfo) {
				now := time.Now()
				msg := models.NewMessage()
				msg.SetReceivedDateTime(&now)
				return msg, &backup.ExchangeInfo{Received: now}
			},
		},
		{
			name: "All fields",
			msgAndRP: func() (models.Messageable, *backup.ExchangeInfo) {
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
				return msg, &backup.ExchangeInfo{Sender: sender, Subject: subject, Received: now}
			},
		}}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			msg, expected := tt.msgAndRP()
			suite.Equal(expected, MessageInfo(msg))
		})
	}
}
