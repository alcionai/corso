package selectors

import (
	"testing"

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

func (suite *SelectorSuite) TestBadCastErr() {
	err := badCastErr(ServiceUnknown, ServiceExchange)
	assert.Error(suite.T(), err)
}

func (suite *SelectorSuite) TestResourceOwnersIn() {
	rootCat := rootCatStub.String()

	table := []struct {
		name   string
		input  []scope
		expect []string
	}{
		{
			name:   "nil",
			input:  nil,
			expect: []string{},
		},
		{
			name:   "empty",
			input:  []scope{},
			expect: []string{},
		},
		{
			name:   "single",
			input:  []scope{{rootCat: filters.Identity("foo")}},
			expect: []string{"foo"},
		},
		{
			name:   "multiple values",
			input:  []scope{{rootCat: filters.Identity(join("foo", "bar"))}},
			expect: []string{"foo", "bar"},
		},
		{
			name:   "with any",
			input:  []scope{{rootCat: filters.Identity(join("foo", "bar", AnyTgt))}},
			expect: []string{"foo", "bar"},
		},
		{
			name:   "with none",
			input:  []scope{{rootCat: filters.Identity(join("foo", "bar", NoneTgt))}},
			expect: []string{"foo", "bar"},
		},
		{
			name: "multiple scopes",
			input: []scope{
				{rootCat: filters.Identity(join("foo", "bar"))},
				{rootCat: filters.Identity(join("baz"))},
			},
			expect: []string{"foo", "bar", "baz"},
		},
		{
			name: "multiple scopes with duplicates",
			input: []scope{
				{rootCat: filters.Identity(join("foo", "bar"))},
				{rootCat: filters.Identity(join("baz", "foo"))},
			},
			expect: []string{"foo", "bar", "baz"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := resourceOwnersIn(test.input, rootCat)
			assert.ElementsMatch(t, test.expect, result)
		})
	}
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
	does[key.String()] = filterize(scopeConfig{}, target)
	doesNot := stubScope("")
	doesNot[key.String()] = filterize(scopeConfig{}, "smarf")

	assert.True(t, matches(does, key, target), "does contain")
	assert.False(t, matches(doesNot, key, target), "does not contain")
}

func (suite *SelectorSuite) TestIsAnyResourceOwner() {
	t := suite.T()
	assert.False(t, isAnyResourceOwner(newSelector(ServiceUnknown, []string{"foo"})))
	assert.False(t, isAnyResourceOwner(newSelector(ServiceUnknown, []string{})))
	assert.False(t, isAnyResourceOwner(newSelector(ServiceUnknown, nil)))
	assert.True(t, isAnyResourceOwner(newSelector(ServiceUnknown, []string{AnyTgt})))
	assert.True(t, isAnyResourceOwner(newSelector(ServiceUnknown, Any())))
}

func (suite *SelectorSuite) TestIsNoneResourceOwner() {
	t := suite.T()
	assert.False(t, isNoneResourceOwner(newSelector(ServiceUnknown, []string{"foo"})))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, []string{})))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, nil)))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, []string{NoneTgt})))
	assert.True(t, isNoneResourceOwner(newSelector(ServiceUnknown, None())))
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
			result := splitByResourceOwner[mockScope](s, allOwners, rootCatStub)

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
				sel.Include(sel.Libraries([]string{"A directory"}, SuffixMatch()))

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
			test.isErr(t, err)
		})
	}
}
