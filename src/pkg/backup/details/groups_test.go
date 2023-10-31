package details_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type GroupsUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsUnitSuite) TestGroupsPrintable() {
	now := time.Now()
	then := now.Add(time.Minute)

	table := []struct {
		name     string
		info     details.GroupsInfo
		expectHs []string
		expectVs []string
	}{
		{
			name: "channel message",
			info: details.GroupsInfo{
				ItemType:   details.GroupsChannelMessage,
				ParentPath: "parentpath",
				Message: details.ChannelMessageInfo{
					Preview:    "preview",
					ReplyCount: 1,
					Creator:    "creator",
					CreatedAt:  now,
					Subject:    "subject",
				},
				LastReply: details.ChannelMessageInfo{
					CreatedAt: then,
				},
			},
			expectHs: []string{"Message", "Channel", "Subject", "Replies", "Creator", "Created", "Last Reply"},
			expectVs: []string{
				"preview",
				"parentpath",
				"subject",
				"1",
				"creator",
				dttm.FormatToTabularDisplay(now),
				dttm.FormatToTabularDisplay(then),
			},
		},
		{
			name: "conversation post",
			info: details.GroupsInfo{
				ItemType: details.GroupsConversationPost,
				Post: details.ConversationPostInfo{
					Preview:   "preview",
					Creator:   "creator",
					CreatedAt: now,
					Topic:     "topic",
				},
			},
			expectHs: []string{"Post", "Conversation", "Sender", "Created"},
			expectVs: []string{
				"preview",
				"topic",
				"creator",
				dttm.FormatToTabularDisplay(now),
			},
		},
		{
			name: "sharepoint library",
			info: details.GroupsInfo{
				ItemType:   details.SharePointLibrary,
				ParentPath: "parentPath",
				Created:    now,
				Modified:   then,
				DriveName:  "librarydrive",
				ItemName:   "item",
				Size:       42,
				Owner:      "user",
			},
			expectHs: []string{"ItemName", "Library", "ParentPath", "Size", "Owner", "Created", "Modified"},
			expectVs: []string{
				"item",
				"librarydrive",
				"parentPath",
				"42 B",
				"user",
				dttm.FormatToTabularDisplay(now),
				dttm.FormatToTabularDisplay(then),
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
