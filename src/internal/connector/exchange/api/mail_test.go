package api

import (
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type MailAPIUnitSuite struct {
	tester.Suite
}

func TestMailAPIUnitSuite(t *testing.T) {
	suite.Run(t, &MailAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MailAPIUnitSuite) TestMailInfo() {
	initial := time.Now()

	tests := []struct {
		name     string
		msgAndRP func() (models.Messageable, *details.ExchangeInfo)
	}{
		{
			name: "Empty message",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				msg := models.NewMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)

				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Recipient: []string{},
					Created:   initial,
					Modified:  initial,
				}
				return msg, i
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
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetSender(sr)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Recipient: []string{},
					Sender:    sender,
					Created:   initial,
					Modified:  initial,
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
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Subject:   subject,
					Created:   initial,
					Recipient: []string{},
					Modified:  initial,
				}
				return msg, i
			},
		},
		{
			name: "Just receivedtime",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				now := time.Now()
				msg := models.NewMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetReceivedDateTime(&now)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Recipient: []string{},
					Received:  now,
					Created:   initial,
					Modified:  initial,
				}
				return msg, i
			},
		},
		{
			name: "All fields",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				sender := "foo@bar.com"
				receiver := "foofoo@bar.com"
				subject := "Hello world"
				now := time.Now()
				sr := models.NewRecipient()
				sea := models.NewEmailAddress()
				recv := models.NewRecipient()
				req := models.NewEmailAddress()
				recvs := make([]models.Recipientable, 0)
				msg := models.NewMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				sea.SetAddress(&sender)
				sr.SetEmailAddress(sea)
				msg.SetSender(sr)
				req.SetAddress(&receiver)
				recv.SetEmailAddress(req)
				msg.SetSubject(&subject)
				msg.SetReceivedDateTime(&now)
				recvs = append(recvs, recv, sr)
				msg.SetToRecipients(recvs)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Sender:    sender,
					Subject:   subject,
					Recipient: []string{receiver, sender},
					Received:  now,
					Created:   initial,
					Modified:  initial,
				}
				return msg, i
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			msg, expected := tt.msgAndRP()
			assert.Equal(suite.T(), expected, MailInfo(msg))
		})
	}
}
