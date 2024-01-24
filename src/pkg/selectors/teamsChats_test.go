package selectors

import (
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

type TeamsChatsSelectorSuite struct {
	tester.Suite
}

func TestTeamsChatsSelectorSuite(t *testing.T) {
	suite.Run(t, &TeamsChatsSelectorSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *TeamsChatsSelectorSuite) TestNewTeamsChatsBackup() {
	t := suite.T()
	eb := NewTeamsChatsBackup(nil)
	assert.Equal(t, eb.Service, ServiceTeamsChats)
	assert.NotZero(t, eb.Scopes())
}

func (suite *TeamsChatsSelectorSuite) TestToTeamsChatsBackup() {
	t := suite.T()
	eb := NewTeamsChatsBackup(nil)
	s := eb.Selector
	eb, err := s.ToTeamsChatsBackup()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, eb.Service, ServiceTeamsChats)
	assert.NotZero(t, eb.Scopes())
}

func (suite *TeamsChatsSelectorSuite) TestNewTeamsChatsRestore() {
	t := suite.T()
	er := NewTeamsChatsRestore(nil)
	assert.Equal(t, er.Service, ServiceTeamsChats)
	assert.NotZero(t, er.Scopes())
}

func (suite *TeamsChatsSelectorSuite) TestToTeamsChatsRestore() {
	t := suite.T()
	eb := NewTeamsChatsRestore(nil)
	s := eb.Selector
	eb, err := s.ToTeamsChatsRestore()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, eb.Service, ServiceTeamsChats)
	assert.NotZero(t, eb.Scopes())
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsSelector_Exclude_TeamsChats() {
	t := suite.T()

	const (
		user   = "user"
		folder = AnyTgt
		c1     = "c1"
		c2     = "c2"
	)

	sel := NewTeamsChatsBackup([]string{user})
	sel.Exclude(sel.Chats([]string{c1, c2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		TeamsChatsScope(scopes[0]),
		map[categorizer][]string{
			TeamsChatsChat: {c1, c2},
		})
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsSelector_Include_TeamsChats() {
	t := suite.T()

	const (
		user   = "user"
		folder = AnyTgt
		c1     = "c1"
		c2     = "c2"
	)

	sel := NewTeamsChatsBackup([]string{user})
	sel.Include(sel.Chats([]string{c1, c2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	scopeMustHave(
		t,
		TeamsChatsScope(scopes[0]),
		map[categorizer][]string{
			TeamsChatsChat: {c1, c2},
		})

	assert.Equal(t, sel.Scopes()[0].Category(), TeamsChatsChat)
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsSelector_Exclude_AllData() {
	t := suite.T()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel := NewTeamsChatsBackup([]string{u1, u2})
	sel.Exclude(sel.AllData())
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		if sc[scopeKeyCategory].Compare(TeamsChatsChat.String()) {
			scopeMustHave(
				t,
				TeamsChatsScope(sc),
				map[categorizer][]string{
					TeamsChatsChat: Any(),
				})
		}
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsSelector_Include_AllData() {
	t := suite.T()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel := NewTeamsChatsBackup([]string{u1, u2})
	sel.Include(sel.AllData())
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		if sc[scopeKeyCategory].Compare(TeamsChatsChat.String()) {
			scopeMustHave(
				t,
				TeamsChatsScope(sc),
				map[categorizer][]string{
					TeamsChatsChat: Any(),
				})
		}
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsBackup_Scopes() {
	eb := NewTeamsChatsBackup(Any())
	eb.Include(eb.AllData())

	scopes := eb.Scopes()
	assert.Len(suite.T(), scopes, 1)

	for _, sc := range scopes {
		cat := sc.Category()
		suite.Run(cat.String(), func() {
			t := suite.T()

			switch sc.Category() {
			case TeamsChatsChat:
				assert.True(t, sc.IsAny(TeamsChatsChat))
			}
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsScope_Category() {
	table := []struct {
		is     teamsChatsCategory
		expect teamsChatsCategory
		check  assert.ComparisonAssertionFunc
	}{
		{TeamsChatsCategoryUnknown, TeamsChatsCategoryUnknown, assert.Equal},
		{TeamsChatsCategoryUnknown, TeamsChatsUser, assert.NotEqual},
		{TeamsChatsChat, TeamsChatsChat, assert.Equal},
		{TeamsChatsUser, TeamsChatsUser, assert.Equal},
		{TeamsChatsUser, TeamsChatsCategoryUnknown, assert.NotEqual},
	}
	for _, test := range table {
		suite.Run(test.is.String()+test.expect.String(), func() {
			eb := NewTeamsChatsBackup(Any())
			eb.Includes = []scope{
				{scopeKeyCategory: filters.Identity(test.is.String())},
			}
			scope := eb.Scopes()[0]
			test.check(suite.T(), test.expect, scope.Category())
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsScope_IncludesCategory() {
	table := []struct {
		is     teamsChatsCategory
		expect teamsChatsCategory
		check  assert.BoolAssertionFunc
	}{
		{TeamsChatsCategoryUnknown, TeamsChatsCategoryUnknown, assert.False},
		{TeamsChatsCategoryUnknown, TeamsChatsUser, assert.True},
		{TeamsChatsUser, TeamsChatsUser, assert.True},
		{TeamsChatsUser, TeamsChatsCategoryUnknown, assert.True},
	}
	for _, test := range table {
		suite.Run(test.is.String()+test.expect.String(), func() {
			eb := NewTeamsChatsBackup(Any())
			eb.Includes = []scope{
				{scopeKeyCategory: filters.Identity(test.is.String())},
			}
			scope := eb.Scopes()[0]
			test.check(suite.T(), scope.IncludesCategory(test.expect))
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsScope_Get() {
	eb := NewTeamsChatsBackup(Any())
	eb.Include(eb.AllData())

	scopes := eb.Scopes()

	table := []teamsChatsCategory{
		TeamsChatsChat,
	}
	for _, test := range table {
		suite.Run(test.String(), func() {
			t := suite.T()

			for _, sc := range scopes {
				switch sc.Category() {
				case TeamsChatsChat:
					assert.Equal(t, Any(), sc.Get(TeamsChatsChat))
				}
				assert.Equal(t, None(), sc.Get(TeamsChatsCategoryUnknown))
			}
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsScope_MatchesInfo() {
	cs := NewTeamsChatsRestore(Any())

	const (
		name    = "smarf mcfnords"
		member  = "cooks@2many.smarf"
		subject = "I have seen the fnords!"
	)

	var (
		now    = time.Now()
		future = now.Add(1 * time.Minute)
	)

	infoWith := func(itype details.ItemType) details.ItemInfo {
		return details.ItemInfo{
			TeamsChats: &details.TeamsChatsInfo{
				ItemType: itype,
				Chat: details.ChatInfo{
					CreatedAt:          now,
					HasExternalMembers: false,
					LastMessageAt:      future,
					LastMessagePreview: "preview",
					Members:            []string{member},
					MessageCount:       1,
					Topic:              name,
				},
			},
		}
	}

	table := []struct {
		name   string
		itype  details.ItemType
		scope  []TeamsChatsScope
		expect assert.BoolAssertionFunc
	}{
		{"chat with a different member", details.TeamsChat, cs.ChatMember("blarps"), assert.False},
		{"chat with the same member", details.TeamsChat, cs.ChatMember(member), assert.True},
		{"chat with a member submatch search", details.TeamsChat, cs.ChatMember(member[2:5]), assert.True},
		{"chat with a different name", details.TeamsChat, cs.ChatName("blarps"), assert.False},
		{"chat with the same name", details.TeamsChat, cs.ChatName(name), assert.True},
		{"chat with a subname search", details.TeamsChat, cs.ChatName(name[2:5]), assert.True},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			scopes := setScopesToDefault(test.scope)
			for _, scope := range scopes {
				test.expect(t, scope.matchesInfo(infoWith(test.itype)))
			}
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsScope_MatchesPath() {
	const (
		user = "userID"
		chat = "chatID"
	)

	repoRef, err := path.Build("tid", user, path.TeamsChatsService, path.ChatsCategory, true, chat)
	require.NoError(suite.T(), err, clues.ToCore(err))

	var (
		loc   = strings.Join([]string{chat}, "/")
		short = "thisisahashofsomekind"
		cs    = NewTeamsChatsRestore(Any())
		ent   = details.Entry{
			RepoRef:     repoRef.String(),
			ShortRef:    short,
			ItemRef:     chat,
			LocationRef: loc,
		}
	)

	table := []struct {
		name     string
		scope    []TeamsChatsScope
		shortRef string
		expect   assert.BoolAssertionFunc
	}{
		{"all items", cs.AllData(), "", assert.True},
		{"all chats", cs.Chats(Any()), "", assert.True},
		{"no chats", cs.Chats(None()), "", assert.False},
		{"matching chats", cs.Chats([]string{chat}), "", assert.True},
		{"non-matching chats", cs.Chats([]string{"smarf"}), "", assert.False},
		{"one of multiple chats", cs.Chats([]string{"smarf", chat}), "", assert.True},
		{"chats short ref", cs.Chats([]string{short}), short, assert.True},
		{"non-leaf short ref", cs.Chats([]string{"foo"}), short, assert.False},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			scopes := setScopesToDefault(test.scope)
			var aMatch bool
			for _, scope := range scopes {
				pvs, err := TeamsChatsChat.pathValues(repoRef, ent, Config{})
				require.NoError(t, err)

				if matchesPathValues(scope, TeamsChatsChat, pvs) {
					aMatch = true
					break
				}
			}
			test.expect(t, aMatch)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsRestore_Reduce() {
	chat, err := path.Build("tid", "uid", path.TeamsChatsService, path.ChatsCategory, true, "cid")
	require.NoError(suite.T(), err, clues.ToCore(err))

	toRR := func(p path.Path) string {
		newElems := []string{}

		for _, e := range p.Folders() {
			newElems = append(newElems, e+".d")
		}

		joinedFldrs := strings.Join(newElems, "/")

		return stubRepoRef(p.Service(), p.Category(), p.ProtectedResource(), joinedFldrs, p.Item())
	}

	makeDeets := func(refs ...path.Path) *details.Details {
		deets := &details.Details{
			DetailsModel: details.DetailsModel{
				Entries: []details.Entry{},
			},
		}

		for _, r := range refs {
			itype := details.UnknownType

			switch r {
			case chat:
				itype = details.TeamsChat
			}

			deets.Entries = append(deets.Entries, details.Entry{
				RepoRef: toRR(r),
				// Don't escape because we assume nice paths.
				LocationRef: r.Folder(false),
				ItemInfo: details.ItemInfo{
					TeamsChats: &details.TeamsChatsInfo{
						ItemType: itype,
					},
				},
			})
		}

		return deets
	}

	table := []struct {
		name         string
		deets        *details.Details
		makeSelector func() *TeamsChatsRestore
		expect       []string
	}{
		{
			"no refs",
			makeDeets(),
			func() *TeamsChatsRestore {
				er := NewTeamsChatsRestore(Any())
				er.Include(er.AllData())
				return er
			},
			[]string{},
		},
		{
			"chat only",
			makeDeets(chat),
			func() *TeamsChatsRestore {
				er := NewTeamsChatsRestore(Any())
				er.Include(er.AllData())
				return er
			},
			[]string{toRR(chat)},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel := test.makeSelector()
			results := sel.Reduce(ctx, test.deets, fault.New(true))
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsRestore_Reduce_locationRef() {
	var (
		chat         = stubRepoRef(path.TeamsChatsService, path.ChatsCategory, "uid", "", "cid")
		chatLocation = "chatname"
	)

	makeDeets := func(refs ...string) *details.Details {
		deets := &details.Details{
			DetailsModel: details.DetailsModel{
				Entries: []details.Entry{},
			},
		}

		for _, r := range refs {
			var (
				location string
				itype    = details.UnknownType
			)

			switch r {
			case chat:
				itype = details.TeamsChat
				location = chatLocation
			}

			deets.Entries = append(deets.Entries, details.Entry{
				RepoRef:     r,
				LocationRef: location,
				ItemInfo: details.ItemInfo{
					TeamsChats: &details.TeamsChatsInfo{
						ItemType: itype,
					},
				},
			})
		}

		return deets
	}

	arr := func(s ...string) []string {
		return s
	}

	table := []struct {
		name         string
		deets        *details.Details
		makeSelector func() *TeamsChatsRestore
		expect       []string
	}{
		{
			"no refs",
			makeDeets(),
			func() *TeamsChatsRestore {
				er := NewTeamsChatsRestore(Any())
				er.Include(er.AllData())
				return er
			},
			[]string{},
		},
		{
			"chat only",
			makeDeets(chat),
			func() *TeamsChatsRestore {
				er := NewTeamsChatsRestore(Any())
				er.Include(er.AllData())
				return er
			},
			arr(chat),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel := test.makeSelector()
			results := sel.Reduce(ctx, test.deets, fault.New(true))
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestScopesByCategory() {
	var (
		cs         = NewTeamsChatsRestore(Any())
		teamsChats = cs.Chats(Any())
	)

	type expect struct {
		chat int
	}

	type input []scope

	makeInput := func(cs ...[]TeamsChatsScope) []scope {
		mss := []scope{}

		for _, sl := range cs {
			for _, s := range sl {
				mss = append(mss, scope(s))
			}
		}

		return mss
	}
	cats := map[path.CategoryType]teamsChatsCategory{
		path.ChatsCategory: TeamsChatsChat,
	}

	table := []struct {
		name   string
		scopes input
		expect expect
	}{
		{"teamsChats only", makeInput(teamsChats), expect{1}},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := scopesByCategory[TeamsChatsScope](test.scopes, cats, false)
			assert.Len(t, result[TeamsChatsChat], test.expect.chat)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestPasses() {
	const (
		chatID = "chatID"
		cat    = TeamsChatsChat
	)

	short := "thisisahashofsomekind"
	entry := details.Entry{
		ShortRef: short,
		ItemRef:  chatID,
	}

	repoRef, err := path.Build("tid", "user", path.TeamsChatsService, path.ChatsCategory, true, chatID)
	require.NoError(suite.T(), err, clues.ToCore(err))

	var (
		cs        = NewTeamsChatsRestore(Any())
		otherChat = setScopesToDefault(cs.Chats([]string{"smarf"}))
		chat      = setScopesToDefault(cs.Chats([]string{chatID}))
		noChat    = setScopesToDefault(cs.Chats(None()))
		allChats  = setScopesToDefault(cs.Chats(Any()))
		ent       = details.Entry{
			RepoRef: repoRef.String(),
		}
	)

	table := []struct {
		name                        string
		excludes, filters, includes []TeamsChatsScope
		expect                      assert.BoolAssertionFunc
	}{
		{"empty", nil, nil, nil, assert.False},
		{"in Chat", nil, nil, chat, assert.True},
		{"in Other", nil, nil, otherChat, assert.False},
		{"in no Chat", nil, nil, noChat, assert.False},
		{"ex None filter chat", allChats, chat, nil, assert.False},
		{"ex Chat", chat, nil, allChats, assert.False},
		{"ex Other", otherChat, nil, allChats, assert.True},
		{"in and ex Chat", chat, nil, chat, assert.False},
		{"filter Chat", nil, chat, allChats, assert.True},
		{"filter Other", nil, otherChat, allChats, assert.False},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			pvs, err := cat.pathValues(repoRef, ent, Config{})
			require.NoError(t, err)

			result := passes(
				cat,
				pvs,
				entry,
				test.excludes,
				test.filters,
				test.includes)
			test.expect(t, result)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestContains() {
	target := "fnords"

	var (
		cs      = NewTeamsChatsRestore(Any())
		noChat  = setScopesToDefault(cs.Chats(None()))
		does    = setScopesToDefault(cs.Chats([]string{target}))
		doesNot = setScopesToDefault(cs.Chats([]string{"smarf"}))
	)

	table := []struct {
		name   string
		scopes []TeamsChatsScope
		expect assert.BoolAssertionFunc
	}{
		{"no chat", noChat, assert.False},
		{"does contain", does, assert.True},
		{"does not contain", doesNot, assert.False},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			var result bool
			for _, scope := range test.scopes {
				if scope.Matches(TeamsChatsChat, target) {
					result = true
					break
				}
			}
			test.expect(t, result)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestIsAny() {
	var (
		cs           = NewTeamsChatsRestore(Any())
		specificChat = setScopesToDefault(cs.Chats([]string{"chat"}))
		anyChat      = setScopesToDefault(cs.Chats(Any()))
	)

	table := []struct {
		name   string
		scopes []TeamsChatsScope
		cat    teamsChatsCategory
		expect assert.BoolAssertionFunc
	}{
		{"specific chat", specificChat, TeamsChatsChat, assert.False},
		{"any chat", anyChat, TeamsChatsChat, assert.True},
		{"wrong category", anyChat, TeamsChatsUser, assert.False},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			var result bool
			for _, scope := range test.scopes {
				if scope.IsAny(test.cat) {
					result = true
					break
				}
			}
			test.expect(t, result)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsCategory_leafCat() {
	table := []struct {
		cat    teamsChatsCategory
		expect teamsChatsCategory
	}{
		{teamsChatsCategory("foo"), teamsChatsCategory("foo")},
		{TeamsChatsCategoryUnknown, TeamsChatsCategoryUnknown},
		{TeamsChatsUser, TeamsChatsUser},
		{TeamsChatsChat, TeamsChatsChat},
	}
	for _, test := range table {
		suite.Run(test.cat.String(), func() {
			assert.Equal(suite.T(), test.expect, test.cat.leafCat(), test.cat.String())
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsCategory_PathValues() {
	t := suite.T()

	chatPath, err := path.Build("tid", "u", path.TeamsChatsService, path.ChatsCategory, true, "chatitem.d")
	require.NoError(t, err, clues.ToCore(err))

	chatLoc, err := path.Build("tid", "u", path.TeamsChatsService, path.ChatsCategory, true, "chatitem")
	require.NoError(t, err, clues.ToCore(err))

	var (
		chatMap = map[categorizer][]string{
			TeamsChatsChat: {chatPath.Item(), "chat-short"},
		}
		chatOnlyNameMap = map[categorizer][]string{
			TeamsChatsChat: {"chat-short"},
		}
	)

	table := []struct {
		cat            teamsChatsCategory
		path           path.Path
		loc            path.Path
		short          string
		expect         map[categorizer][]string
		expectOnlyName map[categorizer][]string
	}{
		{TeamsChatsChat, chatPath, chatLoc, "chat-short", chatMap, chatOnlyNameMap},
	}
	for _, test := range table {
		suite.Run(string(test.cat), func() {
			t := suite.T()
			ent := details.Entry{
				RepoRef:     test.path.String(),
				ShortRef:    test.short,
				LocationRef: test.loc.Folder(true),
				ItemRef:     test.path.Item(),
			}

			pvs, err := test.cat.pathValues(test.path, ent, Config{})
			require.NoError(t, err)

			for k := range test.expect {
				assert.ElementsMatch(t, test.expect[k], pvs[k])
			}

			pvs, err = test.cat.pathValues(test.path, ent, Config{OnlyMatchItemNames: true})
			require.NoError(t, err)

			for k := range test.expectOnlyName {
				assert.ElementsMatch(t, test.expectOnlyName[k], pvs[k], k)
			}
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestTeamsChatsCategory_PathKeys() {
	chat := []categorizer{TeamsChatsChat}
	user := []categorizer{TeamsChatsUser}

	var empty []categorizer

	table := []struct {
		cat    teamsChatsCategory
		expect []categorizer
	}{
		{TeamsChatsCategoryUnknown, empty},
		{TeamsChatsChat, chat},
		{TeamsChatsUser, user},
	}
	for _, test := range table {
		suite.Run(string(test.cat), func() {
			assert.Equal(suite.T(), test.cat.pathKeys(), test.expect)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestCategoryFromItemType() {
	table := []struct {
		name   string
		input  details.ItemType
		expect teamsChatsCategory
	}{
		{
			name:   "chat",
			input:  details.TeamsChat,
			expect: TeamsChatsChat,
		},
		{
			name:   "unknown",
			input:  details.UnknownType,
			expect: TeamsChatsCategoryUnknown,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := teamsChatsCategoryFromItemType(test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}

func (suite *TeamsChatsSelectorSuite) TestCategory_PathType() {
	table := []struct {
		cat      teamsChatsCategory
		pathType path.CategoryType
	}{
		{TeamsChatsCategoryUnknown, path.UnknownCategory},
		{TeamsChatsChat, path.ChatsCategory},
		{TeamsChatsUser, path.UnknownCategory},
	}
	for _, test := range table {
		suite.Run(test.cat.String(), func() {
			assert.Equal(suite.T(), test.pathType, test.cat.PathType())
		})
	}
}
