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
	t := suite.T()
	now := time.Now()
	then := now.Add(time.Minute)

	gi := details.GroupsInfo{
		ItemType:   details.GroupsChannelMessage,
		ParentPath: "parentPath",
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
	}

	expectVs := []string{
		"preview",
		"parentPath",
		"subject",
		"1",
		"creator",
		dttm.FormatToTabularDisplay(now),
		dttm.FormatToTabularDisplay(then),
	}

	hs := gi.Headers()
	vs := gi.Values()

	assert.Equal(t, len(hs), len(vs))
	assert.Equal(t, expectVs, vs)
}
