package selectors

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/filters"
	"github.com/alcionai/corso/src/pkg/path"
)

type SelectorSuite struct {
	tester.Suite
}

func TestSelectorSuite(t *testing.T) {
	suite.Run(t, &SelectorSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SelectorSuite) TestNewSelector() {
	t := suite.T()
	s := newSelector(ServiceUnknown, Any())
	assert.NotNil(t, s)
	assert.Equal(t, s.Service, ServiceUnknown)
	assert.NotNil(t, s.Includes)
}

// set the clues hashing to mask for the span of this suite
func (suite *SelectorSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *SelectorSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
}

func (suite *SelectorSuite) TestBadCastErr() {
	err := badCastErr(ServiceUnknown, ServiceExchange)
	assert.Error(suite.T(), err, clues.ToCore(err))
}

func (suite *SelectorSuite) TestPathCategoriesIn() {
	leafCat := leafCatStub.String()
	f := filters.Identity(leafCat)

	table := []struct {
		name   string
		input  []scope
		expect []path.CategoryType
	}{
		{
			name:   "nil",
			input:  nil,
			expect: []path.CategoryType{},
		},
		{
			name:   "empty",
			input:  []scope{},
			expect: []path.CategoryType{},
		},
		{
			name:   "single",
			input:  []scope{{leafCat: f, scopeKeyCategory: f}},
			expect: []path.CategoryType{leafCatStub.PathType()},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := pathCategoriesIn[mockScope, mockCategorizer](test.input)
			assert.ElementsMatch(t, test.expect, result)
		})
	}
}

func (suite *SelectorSuite) TestContains() {
	t := suite.T()
	key := rootCatStub
	target := "fnords"
	does := stubScope("")
	does[key.String()] = filterFor(scopeConfig{}, target)
	doesNot := stubScope("")
	doesNot[key.String()] = filterFor(scopeConfig{}, "smarf")

	assert.True(t, matches(does, key, target), "does contain")
	assert.False(t, matches(doesNot, key, target), "does not contain")
}

func (suite *SelectorSuite) TestIsAnyResourceOwner() {
	t := suite.T()
	assert.False(t, isAnyProtectedResource(newSelector(ServiceUnknown, []string{"foo"})))
	assert.False(t, isAnyProtectedResource(newSelector(ServiceUnknown, []string{})))
	assert.False(t, isAnyProtectedResource(newSelector(ServiceUnknown, nil)))
	assert.True(t, isAnyProtectedResource(newSelector(ServiceUnknown, []string{AnyTgt})))
	assert.True(t, isAnyProtectedResource(newSelector(ServiceUnknown, Any())))
}

func (suite *SelectorSuite) TestIsNoneResourceOwner() {
	t := suite.T()
	assert.False(t, isNoneProtectedResource(newSelector(ServiceUnknown, []string{"foo"})))
	assert.True(t, isNoneProtectedResource(newSelector(ServiceUnknown, []string{})))
	assert.True(t, isNoneProtectedResource(newSelector(ServiceUnknown, nil)))
	assert.True(t, isNoneProtectedResource(newSelector(ServiceUnknown, []string{NoneTgt})))
	assert.True(t, isNoneProtectedResource(newSelector(ServiceUnknown, None())))
}

func (suite *SelectorSuite) TestSplitByResourceOnwer() {
	allOwners := []string{"foo", "bar", "baz", "qux"}

	table := []struct {
		name           string
		input          []string
		expectLen      int
		expectDiscrete []string
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []string{},
		},
		{
			name:  "noneTgt",
			input: []string{NoneTgt},
		},
		{
			name:  "none",
			input: None(),
		},
		{
			name:           "AnyTgt",
			input:          []string{AnyTgt},
			expectLen:      len(allOwners),
			expectDiscrete: allOwners,
		},
		{
			name:           "Any",
			input:          Any(),
			expectLen:      len(allOwners),
			expectDiscrete: allOwners,
		},
		{
			name:           "one owner",
			input:          []string{"fnord"},
			expectLen:      1,
			expectDiscrete: []string{"fnord"},
		},
		{
			name:           "two owners",
			input:          []string{"fnord", "smarf"},
			expectLen:      2,
			expectDiscrete: []string{"fnord", "smarf"},
		},
		{
			name:  "two owners and NoneTgt",
			input: []string{"fnord", "smarf", NoneTgt},
		},
		{
			name:           "two owners and AnyTgt",
			input:          []string{"fnord", "smarf", AnyTgt},
			expectLen:      len(allOwners),
			expectDiscrete: allOwners,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			s := newSelector(ServiceUnknown, test.input)
			result := splitByProtectedResource[mockScope](s, allOwners, rootCatStub)

			assert.Len(t, result, test.expectLen)

			for _, expect := range test.expectDiscrete {
				var found bool

				for _, sel := range result {
					if sel.DiscreteOwner == expect {
						found = true
						break
					}
				}

				assert.Truef(t, found, "%s in list of discrete owners", expect)
			}
		})
	}
}

func (suite *SelectorSuite) TestIDName() {
	table := []struct {
		title                string
		id, name             string
		expectID, expectName string
	}{
		{"empty", "", "", "", ""},
		{"only id", "id", "", "id", "id"},
		{"only name", "", "name", "", "name"},
		{"both", "id", "name", "id", "name"},
	}
	for _, test := range table {
		suite.Run(test.title, func() {
			sel := Selector{DiscreteOwner: test.id, DiscreteOwnerName: test.name}
			assert.Equal(suite.T(), test.expectID, sel.ID())
			assert.Equal(suite.T(), test.expectName, sel.Name())
		})
	}
}

func (suite *SelectorSuite) TestSetDiscreteOwnerIDName() {
	table := []struct {
		title                string
		initID, initName     string
		id, name             string
		expectID, expectName string
	}{
		{"empty", "", "", "", "", "", ""},
		{"only id", "", "", "id", "", "id", "id"},
		{"only name", "", "", "", "", "", ""},
		{"both", "", "", "id", "name", "id", "name"},
		{"both", "init-id", "", "", "name", "init-id", "name"},
	}
	for _, test := range table {
		suite.Run(test.title, func() {
			sel := Selector{DiscreteOwner: test.initID, DiscreteOwnerName: test.initName}
			sel = sel.SetDiscreteOwnerIDName(test.id, test.name)
			assert.Equal(suite.T(), test.expectID, sel.ID())
			assert.Equal(suite.T(), test.expectName, sel.Name())
		})
	}
}

// TestPathCategories verifies that no scope produces a `path.UnknownCategory`
func (suite *SelectorSuite) TestPathCategories_includes() {
	users := []string{"someuser@onmicrosoft.com"}

	table := []struct {
		name        string
		getSelector func(t *testing.T) *Selector
		isErr       assert.ErrorAssertionFunc
	}{
		{
			name:  "empty",
			isErr: assert.Error,
			getSelector: func(t *testing.T) *Selector {
				return &Selector{}
			},
		},
		{
			name:  "Mail_B",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewExchangeBackup(users)
				sel.Include(sel.MailFolders([]string{"MailFolder"}, PrefixMatch()))
				sel.Mails([]string{"MailFolder2"}, []string{"Mail"})
				return &sel.Selector
			},
		},
		{
			name:  "Mail_R",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewExchangeRestore(users)
				sel.Include(sel.MailFolders([]string{"MailFolder"}, PrefixMatch()))

				return &sel.Selector
			},
		},
		{
			name:  "Contacts",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewExchangeBackup(users)
				sel.Include(sel.ContactFolders([]string{"Contact Folder"}, PrefixMatch()))
				return &sel.Selector
			},
		},
		{
			name:  "Contacts_R",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewExchangeRestore(users)
				sel.Include(sel.ContactFolders([]string{"Contact Folder"}, PrefixMatch()))
				return &sel.Selector
			},
		},
		{
			name:  "Events",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewExchangeBackup(users)
				sel.Include(sel.EventCalendars([]string{"July"}, PrefixMatch()))
				return &sel.Selector
			},
		},
		{
			name:  "Events_R",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewExchangeRestore(users)
				sel.Include(sel.EventCalendars([]string{"July"}, PrefixMatch()))
				sel.EventCalendars([]string{"Independence Day EventID"})
				return &sel.Selector
			},
		},
		{
			name:  "SharePoint Pages",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewSharePointBackup(users)
				sel.Include(sel.Pages([]string{"Something"}, SuffixMatch()))
				sel.PageItems([]string{"Home Directory"}, []string{"Event Page"})

				return &sel.Selector
			},
		},
		{
			name:  "SharePoint Lists",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewSharePointBackup(users)
				sel.Include(sel.Lists([]string{"Lists from website"}, SuffixMatch()))

				return &sel.Selector
			},
		},
		{
			name:  "SharePoint Libraries",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewSharePointBackup(users)
				sel.Include(sel.LibraryFolders([]string{"A directory"}, SuffixMatch()))

				return &sel.Selector
			},
		},
		{
			name:  "OneDrive",
			isErr: assert.NoError,
			getSelector: func(t *testing.T) *Selector {
				sel := NewOneDriveBackup(users)
				sel.Include(sel.Folders([]string{"Single Folder"}, PrefixMatch()))

				return &sel.Selector
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			obj := test.getSelector(t)

			cats, err := obj.PathCategories()
			for _, entry := range cats.Includes {
				assert.NotEqual(t, entry, path.UnknownCategory)
			}

			test.isErr(t, err, clues.ToCore(err))
		})
	}
}

func (suite *SelectorSuite) TestSelector_pii() {
	table := []struct {
		name        string
		sel         func() Selector
		expect      string
		expectPlain string
	}{
		{
			name:        "empty selector",
			sel:         func() Selector { return Selector{} },
			expect:      `{"resourceOwners":"UnknownComparison:"}`,
			expectPlain: `{"resourceOwners":"UnknownComparison:"}`,
		},
		{
			name: "no scopes",
			sel: func() Selector {
				return Selector{
					Service:        ServiceUnknown,
					DiscreteOwner:  "owner",
					ResourceOwners: filterFor(scopeConfig{}, "owner_1", "owner_2"),
				}
			},
			expect:      `{"resourceOwners":"EQ:***,***","discreteOwner":"***"}`,
			expectPlain: `{"resourceOwners":"EQ:owner_1,owner_2","discreteOwner":"owner"}`,
		},
		{
			name: "one scope each type",
			sel: func() Selector {
				s := NewExchangeBackup([]string{"owner_1", "owner_2"})
				s.DiscreteOwner = "owner"

				s.Exclude(s.MailFolders([]string{"e"}))
				s.Filter(s.MailFolders([]string{"f"}, SuffixMatch()))
				s.Include(s.MailFolders([]string{"i"}, PrefixMatch()))

				return s.Selector
			},
			//nolint:lll
			expect: `{"service":1,` +
				`"resourceOwners":"EQ:***,***",` +
				`"discreteOwner":"***",` +
				`"exclusions":[{"ExchangeMail":"Pass","ExchangeMailFolder":"PathCont:***","category":"Identity:***","type":"Identity:***"}],` +
				`"filters":[{"ExchangeMail":"Pass","ExchangeMailFolder":"PathSfx:***","category":"Identity:***","type":"Identity:***"}],` +
				`"includes":[{"ExchangeMail":"Pass","ExchangeMailFolder":"PathPfx:***","category":"Identity:***","type":"Identity:***"}]}`,
			//nolint:lll
			expectPlain: `{"service":1,` +
				`"resourceOwners":"EQ:owner_1,owner_2",` +
				`"discreteOwner":"owner",` +
				`"exclusions":[{"ExchangeMail":"Pass","ExchangeMailFolder":"PathCont:e","category":"Identity:ExchangeMailFolder","type":"Identity:ExchangeMail"}],` +
				`"filters":[{"ExchangeMail":"Pass","ExchangeMailFolder":"PathSfx:f","category":"Identity:ExchangeMailFolder","type":"Identity:ExchangeMail"}],` +
				`"includes":[{"ExchangeMail":"Pass","ExchangeMailFolder":"PathPfx:i","category":"Identity:ExchangeMailFolder","type":"Identity:ExchangeMail"}]}`,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := test.sel().Conceal()
			assert.Equal(t, test.expect, result, "conceal")

			result = test.sel().String()
			assert.Equal(t, test.expect, result, "string")

			result = test.sel().PlainString()
			assert.Equal(t, test.expectPlain, result, "plainString")

			result = fmt.Sprintf("%s", test.sel())
			assert.Equal(t, test.expect, result, "fmt %%s")

			result = fmt.Sprintf("%v", test.sel())
			assert.Equal(t, test.expect, result, "fmt %%v")

			result = fmt.Sprintf("%+v", test.sel())
			assert.Equal(t, test.expect, result, "fmt %%+v")
		})
	}
}
