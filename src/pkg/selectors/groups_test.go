package selectors

import (
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"

	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type GroupsSelectorSuite struct {
	tester.Suite
}

func TestGroupsSelectorSuite(t *testing.T) {
	suite.Run(t, &GroupsSelectorSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsSelectorSuite) TestNewGroupsBackup() {
	t := suite.T()
	ob := NewGroupsBackup(nil)
	assert.Equal(t, ob.Service, ServiceGroups)
	assert.NotZero(t, ob.Scopes())
}

func (suite *GroupsSelectorSuite) TestToGroupsBackup() {
	t := suite.T()
	ob := NewGroupsBackup(nil)
	s := ob.Selector
	ob, err := s.ToGroupsBackup()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ob.Service, ServiceGroups)
	assert.NotZero(t, ob.Scopes())
}

func (suite *GroupsSelectorSuite) TestNewGroupsRestore() {
	t := suite.T()
	or := NewGroupsRestore(nil)
	assert.Equal(t, or.Service, ServiceGroups)
	assert.NotZero(t, or.Scopes())
}

func (suite *GroupsSelectorSuite) TestToGroupsRestore() {
	t := suite.T()
	eb := NewGroupsRestore(nil)
	s := eb.Selector
	or, err := s.ToGroupsRestore()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, or.Service, ServiceGroups)
	assert.NotZero(t, or.Scopes())
}

func (suite *GroupsSelectorSuite) TestGroupsRestore_Reduce() {
	toRR := func(cat path.CategoryType, midID string, folders []string, item string) string {
		var (
			folderElems = make([]string, 0, len(folders))
			isDrive     = cat == path.LibrariesCategory
		)

		for _, f := range folders {
			if isDrive {
				f = f + ".d"
			}

			folderElems = append(folderElems, f)
		}

		return stubRepoRef(
			path.GroupsService,
			cat,
			midID,
			strings.Join(folderElems, "/"),
			item)
	}

	var (
		drivePrefixElems = []string{
			odConsts.DrivesPathDir,
			"drive!id",
			odConsts.RootPathDir,
		}
		itemElems1 = []string{"folderA", "folderB"}
		itemElems2 = []string{"folderA", "folderC"}
		itemElems3 = []string{"folderD", "folderE"}
		pairAC     = "folderA/folderC"
		libItem    = toRR(
			path.LibrariesCategory,
			"sid",
			append(slices.Clone(drivePrefixElems), itemElems1...),
			"item")
		libItem2 = toRR(
			path.LibrariesCategory,
			"sid",
			append(slices.Clone(drivePrefixElems), itemElems2...),
			"item2")
		libItem3 = toRR(
			path.LibrariesCategory,
			"sid",
			append(slices.Clone(drivePrefixElems), itemElems3...),
			"item3")
		chanItem  = toRR(path.ChannelMessagesCategory, "gid", slices.Clone(itemElems1), "chitem")
		chanItem2 = toRR(path.ChannelMessagesCategory, "gid", slices.Clone(itemElems2), "chitem2")
		chanItem3 = toRR(path.ChannelMessagesCategory, "gid", slices.Clone(itemElems3), "chitem3")
		convItem  = toRR(path.ConversationPostsCategory, "gid", slices.Clone(itemElems1), "convitem")
		convItem2 = toRR(path.ConversationPostsCategory, "gid", slices.Clone(itemElems2), "convitem2")
		convItem3 = toRR(path.ConversationPostsCategory, "gid", slices.Clone(itemElems3), "convitem3")
	)

	deets := &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: []details.Entry{
				{
					RepoRef:     libItem,
					ItemRef:     "item",
					LocationRef: strings.Join(append([]string{odConsts.RootPathDir}, itemElems1...), "/"),
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.SharePointLibrary,
							ItemName:   "itemName",
							ParentPath: strings.Join(itemElems1, "/"),
						},
					},
				},
				{
					RepoRef:     libItem2,
					LocationRef: strings.Join(append([]string{odConsts.RootPathDir}, itemElems2...), "/"),
					// ItemRef intentionally blank to test fallback case
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.SharePointLibrary,
							ItemName:   "itemName2",
							ParentPath: strings.Join(itemElems2, "/"),
						},
					},
				},
				{
					RepoRef:     libItem3,
					ItemRef:     "item3",
					LocationRef: strings.Join(append([]string{odConsts.RootPathDir}, itemElems3...), "/"),
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.SharePointLibrary,
							ItemName:   "itemName3",
							ParentPath: strings.Join(itemElems3, "/"),
						},
					},
				},
				{
					RepoRef:     chanItem,
					ItemRef:     "citem",
					LocationRef: strings.Join(itemElems1, "/"),
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.GroupsChannelMessage,
							ParentPath: strings.Join(itemElems1, "/"),
						},
					},
				},
				{
					RepoRef:     chanItem2,
					LocationRef: strings.Join(itemElems2, "/"),
					// ItemRef intentionally blank to test fallback case
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.GroupsChannelMessage,
							ParentPath: strings.Join(itemElems2, "/"),
						},
					},
				},
				{
					RepoRef:     chanItem3,
					ItemRef:     "citem3",
					LocationRef: strings.Join(itemElems3, "/"),
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.GroupsChannelMessage,
							ParentPath: strings.Join(itemElems3, "/"),
						},
					},
				},
				{
					RepoRef:     convItem,
					ItemRef:     "convitem",
					LocationRef: strings.Join(itemElems1, "/"),
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.GroupsConversationPost,
							ParentPath: strings.Join(itemElems1, "/"),
						},
					},
				},
				{
					RepoRef:     convItem2,
					LocationRef: strings.Join(itemElems2, "/"),
					// ItemRef intentionally blank to test fallback case
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.GroupsConversationPost,
							ParentPath: strings.Join(itemElems2, "/"),
						},
					},
				},
				{
					RepoRef:     convItem3,
					ItemRef:     "convitem3",
					LocationRef: strings.Join(itemElems3, "/"),
					ItemInfo: details.ItemInfo{
						Groups: &details.GroupsInfo{
							ItemType:   details.GroupsConversationPost,
							ParentPath: strings.Join(itemElems3, "/"),
						},
					},
				},
			},
		},
	}

	arr := func(s ...string) []string {
		return s
	}

	table := []struct {
		name         string
		makeSelector func() *GroupsRestore
		expect       []string
		cfg          Config
	}{
		{
			name: "all",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.AllData())
				return sel
			},
			expect: arr(
				libItem, libItem2, libItem3,
				chanItem, chanItem2, chanItem3),
			// TODO: re-add when we release conversations
			// convItem, convItem2, convItem3),
		},
		{
			name: "only match library item",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.LibraryItems(Any(), []string{"item2"}))
				return sel
			},
			expect: arr(libItem2),
		},
		{
			name: "library id doesn't match name",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.LibraryItems(Any(), []string{"item2"}))
				return sel
			},
			expect: []string{},
			cfg:    Config{OnlyMatchItemNames: true},
		},
		{
			name: "library only match item name",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.LibraryItems(Any(), []string{"itemName2"}))
				return sel
			},
			expect: arr(libItem2),
			cfg:    Config{OnlyMatchItemNames: true},
		},
		{
			name: "name doesn't match",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.LibraryItems(Any(), []string{"itemName2"}))
				return sel
			},
			expect: []string{},
		},
		{
			name: "only match folder",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore([]string{"sid"})
				sel.Include(sel.LibraryFolders([]string{"folderA/folderB", pairAC}))
				return sel
			},
			expect: arr(libItem, libItem2),
		},
		{
			name: "only match channel item",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.ChannelMessages(Any(), []string{"chitem2"}))
				return sel
			},
			expect: arr(chanItem2),
		},
		{
			name: "channel id doesn't match name",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.ChannelMessages(Any(), []string{"item2"}))
				return sel
			},
			expect: []string{},
			cfg:    Config{OnlyMatchItemNames: true},
		},
		{
			name: "only match conversation item",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.ConversationPosts(Any(), []string{"convitem2"}))
				return sel
			},
			expect: arr(convItem2),
		},
		{
			name: "conversation id doesn't match name",
			makeSelector: func() *GroupsRestore {
				sel := NewGroupsRestore(Any())
				sel.Include(sel.ConversationPosts(Any(), []string{"item2"}))
				return sel
			},
			expect: []string{},
			cfg:    Config{OnlyMatchItemNames: true},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel := test.makeSelector()
			sel.Configure(test.cfg)
			results := sel.Reduce(ctx, deets, fault.New(true))
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *GroupsSelectorSuite) TestGroupsCategory_PathValues() {
	var (
		itemName = "item"
		itemID   = "item-id"
		shortRef = "short"
		elems    = []string{itemID}
	)

	table := []struct {
		name       string
		sc         groupsCategory
		pathElems  []string
		locRef     string
		parentPath string
		expected   map[categorizer][]string
		cfg        Config
	}{
		{
			name:      "Groups Channel Messages",
			sc:        GroupsChannelMessage,
			pathElems: elems,
			locRef:    "",
			expected: map[categorizer][]string{
				GroupsChannel:        {""},
				GroupsChannelMessage: {itemID, shortRef},
			},
			cfg: Config{},
		},
		{
			name:      "Groups Conversation Posts",
			sc:        GroupsConversationPost,
			pathElems: elems,
			locRef:    "",
			expected: map[categorizer][]string{
				GroupsConversation:     {""},
				GroupsConversationPost: {itemID, shortRef},
			},
			cfg: Config{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			itemPath, err := path.Build(
				"tenant",
				"site",
				path.GroupsService,
				test.sc.PathType(),
				true,
				test.pathElems...)
			require.NoError(t, err, clues.ToCore(err))

			ent := details.Entry{
				RepoRef:     itemPath.String(),
				ShortRef:    shortRef,
				ItemRef:     itemPath.Item(),
				LocationRef: test.locRef,
				ItemInfo: details.ItemInfo{
					Groups: &details.GroupsInfo{
						ItemName:   itemName,
						ParentPath: test.parentPath,
					},
				},
			}

			pv, err := test.sc.pathValues(itemPath, ent, test.cfg)
			require.NoError(t, err)
			assert.Equal(t, test.expected, pv)
		})
	}
}

func (suite *GroupsSelectorSuite) TestGroupsScope_MatchesInfo() {
	var (
		sel  = NewGroupsRestore(Any())
		user = "user@mail.com"
		host = "www.website.com"
		// pth  = "/foo"
		// url          = host + pth
		epoch  = time.Time{}
		now    = time.Now()
		mod    = now.Add(15 * time.Minute)
		future = now.Add(45 * time.Minute)
		dgcm   = details.GroupsChannelMessage
		dspl   = details.SharePointLibrary
	)

	type expectation func(t assert.TestingT, value bool, msg string, args ...any) bool

	table := []struct {
		name     string
		itemType details.ItemType
		creator  string
		scope    []GroupsScope
		expect   expectation
	}{
		{"file create after the epoch", dspl, user, sel.CreatedAfter(dttm.Format(epoch)), assert.Truef},
		{"file create after the epoch wrong type", dgcm, user, sel.CreatedAfter(dttm.Format(epoch)), assert.Falsef},
		{"file create after now", dspl, user, sel.CreatedAfter(dttm.Format(now)), assert.Falsef},
		{"file create after later", dspl, user, sel.CreatedAfter(dttm.Format(future)), assert.Falsef},
		{"file create before future", dspl, user, sel.CreatedBefore(dttm.Format(future)), assert.Truef},
		{"file create before future wrong type", dgcm, user, sel.CreatedBefore(dttm.Format(future)), assert.Falsef},
		{"file create before now", dspl, user, sel.CreatedBefore(dttm.Format(now)), assert.Falsef},
		{"file create before modification", dspl, user, sel.CreatedBefore(dttm.Format(mod)), assert.Truef},
		{"file create before epoch", dspl, user, sel.CreatedBefore(dttm.Format(now)), assert.Falsef},
		{"file modified after the epoch", dspl, user, sel.ModifiedAfter(dttm.Format(epoch)), assert.Truef},
		{"file modified after now", dspl, user, sel.ModifiedAfter(dttm.Format(now)), assert.Truef},
		{"file modified after later", dspl, user, sel.ModifiedAfter(dttm.Format(future)), assert.Falsef},
		{"file modified before future", dspl, user, sel.ModifiedBefore(dttm.Format(future)), assert.Truef},
		{"file modified before now", dspl, user, sel.ModifiedBefore(dttm.Format(now)), assert.Falsef},
		{"file modified before epoch", dspl, user, sel.ModifiedBefore(dttm.Format(now)), assert.Falsef},
		{"in library", dspl, user, sel.Library("included-library"), assert.Truef},
		{"not in library", dspl, user, sel.Library("not-included-library"), assert.Falsef},
		{"site id", dspl, user, sel.Site("site1"), assert.Truef},
		{"web url", dspl, user, sel.Site(user), assert.Truef},
		{"library id", dspl, user, sel.Library("1234"), assert.Truef},
		{"not library id", dspl, user, sel.Library("abcd"), assert.Falsef},

		{"channel message created by", dgcm, user, sel.MessageCreator(user), assert.Truef},
		{"channel message not created by", dgcm, user, sel.MessageCreator(host), assert.Falsef},
		{"chan msg create after the epoch", dgcm, user, sel.MessageCreatedAfter(dttm.Format(epoch)), assert.Truef},
		{
			"chan msg create after the epoch wrong type",
			dspl,
			user,
			sel.MessageCreatedAfter(dttm.Format(epoch)),
			assert.Falsef,
		},
		{"chan msg create after now", dgcm, user, sel.MessageCreatedAfter(dttm.Format(now)), assert.Falsef},
		{"chan msg create after later", dgcm, user, sel.MessageCreatedAfter(dttm.Format(future)), assert.Falsef},
		{"chan msg create before future", dgcm, user, sel.MessageCreatedBefore(dttm.Format(future)), assert.Truef},
		{
			"chan msg create before future wrong type",
			dspl,
			user,
			sel.MessageCreatedBefore(dttm.Format(future)),
			assert.Falsef,
		},
		{"chan msg create before now", dgcm, user, sel.MessageCreatedBefore(dttm.Format(now)), assert.Falsef},
		{"chan msg create before reply", dgcm, user, sel.MessageCreatedBefore(dttm.Format(mod)), assert.Truef},
		{"chan msg create before reply wrong type", dspl, user, sel.MessageCreatedBefore(dttm.Format(mod)), assert.Falsef},
		{"chan msg create before epoch", dgcm, user, sel.MessageCreatedBefore(dttm.Format(now)), assert.Falsef},
		{"chan msg last reply after the epoch", dgcm, user, sel.MessageLastReplyAfter(dttm.Format(epoch)), assert.Truef},
		{"chan msg last reply after now", dgcm, user, sel.MessageLastReplyAfter(dttm.Format(now)), assert.Truef},
		{"chan msg last reply after later", dgcm, user, sel.MessageLastReplyAfter(dttm.Format(future)), assert.Falsef},
		{"chan msg last reply before future", dgcm, user, sel.MessageLastReplyBefore(dttm.Format(future)), assert.Truef},
		{"chan msg last reply before now", dgcm, user, sel.MessageLastReplyBefore(dttm.Format(now)), assert.Falsef},
		{"chan msg last reply before epoch", dgcm, user, sel.MessageLastReplyBefore(dttm.Format(now)), assert.Falsef},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			itemInfo := details.ItemInfo{
				Groups: &details.GroupsInfo{
					ItemType:  test.itemType,
					Created:   now,
					WebURL:    test.creator,
					Modified:  mod,
					DriveName: "included-library",
					DriveID:   "1234",
					SiteID:    "site1",
					Message: details.ChannelMessageInfo{
						Creator:   test.creator,
						CreatedAt: now,
					},
					LastReply: details.ChannelMessageInfo{
						CreatedAt: mod,
					},
				},
			}

			scopes := setScopesToDefault(test.scope)
			for _, scope := range scopes {
				test.expect(
					t,
					scope.matchesInfo(itemInfo),
					"not matching:\nscope:\n\t%+v\ninfo:\n\t%+v",
					scope,
					itemInfo.Groups)
			}
		})
	}
}

func (suite *GroupsSelectorSuite) TestCategory_PathType() {
	table := []struct {
		cat      groupsCategory
		pathType path.CategoryType
	}{
		{GroupsCategoryUnknown, path.UnknownCategory},
		{GroupsChannel, path.ChannelMessagesCategory},
		{GroupsChannelMessage, path.ChannelMessagesCategory},
		{GroupsConversation, path.ConversationPostsCategory},
		{GroupsConversationPost, path.ConversationPostsCategory},
		{GroupsInfoChannelMessageCreator, path.ChannelMessagesCategory},
		{GroupsInfoChannelMessageCreatedAfter, path.ChannelMessagesCategory},
		{GroupsInfoChannelMessageCreatedBefore, path.ChannelMessagesCategory},
		{GroupsInfoChannelMessageLastReplyAfter, path.ChannelMessagesCategory},
		{GroupsInfoChannelMessageLastReplyBefore, path.ChannelMessagesCategory},
		{GroupsLibraryFolder, path.LibrariesCategory},
		{GroupsLibraryItem, path.LibrariesCategory},
		{GroupsInfoSiteLibraryDrive, path.LibrariesCategory},
		{GroupsInfoSite, path.LibrariesCategory},
	}
	for _, test := range table {
		suite.Run(test.cat.String(), func() {
			assert.Equal(
				suite.T(),
				test.pathType.String(),
				test.cat.PathType().String())
		})
	}
}
