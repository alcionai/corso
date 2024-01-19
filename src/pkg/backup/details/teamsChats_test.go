package details_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/dttm"
)

type ChatsUnitSuite struct {
	tester.Suite
}

func TestChatsUnitSuite(t *testing.T) {
	suite.Run(t, &ChatsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ChatsUnitSuite) TestChatsPrintable() {
	now := time.Now()
	then := now.Add(time.Minute)

	table := []struct {
		name     string
		info     details.TeamsChatsInfo
		expectHs []string
		expectVs []string
	}{
		{
			name: "channel message",
			info: details.TeamsChatsInfo{
				ItemType:   details.TeamsChat,
				ParentPath: "parentpath",
				Chat: details.ChatInfo{
					CreatedAt:          now,
					HasExternalMembers: true,
					LastMessageAt:      then,
					LastMessagePreview: "last message preview",
					Members:            []string{"foo@bar.baz", "fnords@smarf.zoomba"},
					MessageCount:       42,
					Name:               "chat name",
				},
			},
			expectHs: []string{"Name", "Last message", "Last message at", "Message count", "Created", "Members"},
			expectVs: []string{
				"chat name",
				"last message preview",
				dttm.FormatToTabularDisplay(then),
				"42",
				dttm.FormatToTabularDisplay(now),
				"foo@bar.baz, and 1 more",
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			hs := test.info.Headers()
			vs := test.info.Values()

			assert.Equal(t, len(hs), len(vs))
			assert.Equal(t, test.expectHs, hs)
			assert.Equal(t, test.expectVs, vs)
		})
	}
}
