package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OneDriveSelectorSuite struct {
	suite.Suite
}

func TestOneDriveSelectorSuite(t *testing.T) {
	suite.Run(t, new(OneDriveSelectorSuite))
}

func (suite *OneDriveSelectorSuite) TestNewOneDriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup()
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OneDriveSelectorSuite) TestToOneDriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup()
	s := ob.Selector
	ob, err := s.ToOneDriveBackup()
	require.NoError(t, err)
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OneDriveSelectorSuite) TestOneDriveBackup_DiscreteScopes() {
	usrs := []string{"u1", "u2"}
	table := []struct {
		name     string
		include  []string
		discrete []string
		expect   []string
	}{
		{
			name:     "any user",
			include:  Any(),
			discrete: usrs,
			expect:   usrs,
		},
		{
			name:     "discrete user",
			include:  []string{"u3"},
			discrete: usrs,
			expect:   []string{"u3"},
		},
		{
			name:     "nil discrete slice",
			include:  Any(),
			discrete: nil,
			expect:   Any(),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			eb := NewOneDriveBackup()
			eb.Include(eb.Users(test.include))

			scopes := eb.DiscreteScopes(test.discrete)
			for _, sc := range scopes {
				users := sc.Get(OneDriveUser)
				assert.Equal(t, test.expect, users)
			}
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	userScopes := sel.Users([]string{u1, u2})
	for _, scope := range userScopes {
		// Scope value is either u1 or u2
		assert.Contains(t, join(u1, u2), scope[OneDriveUser.String()].Target)
	}

	// Initialize the selector Include, Exclude, Filter
	sel.Exclude(userScopes)
	sel.Include(userScopes)
	sel.Filter(userScopes)

	table := []struct {
		name          string
		scopesToCheck []scope
	}{
		{"Include Scopes", sel.Includes},
		{"Exclude Scopes", sel.Excludes},
		{"Filter Scopes", sel.Filters},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			require.Len(t, test.scopesToCheck, 1)
			for _, scope := range test.scopesToCheck {
				// Scope value is u1,u2
				assert.Contains(t, join(u1, u2), scope[OneDriveUser.String()].Target)
			}
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Include_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Include(sel.Users([]string{u1, u2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			OneDriveScope(sc),
			map[categorizer]string{OneDriveUser: join(u1, u2)},
		)
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Exclude_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Exclude(sel.Users([]string{u1, u2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			OneDriveScope(sc),
			map[categorizer]string{OneDriveUser: join(u1, u2)},
		)
	}
}

func (suite *OneDriveSelectorSuite) TestNewOneDriveRestore() {
	t := suite.T()
	or := NewOneDriveRestore()
	assert.Equal(t, or.Service, ServiceOneDrive)
	assert.NotZero(t, or.Scopes())
}

func (suite *OneDriveSelectorSuite) TestToOneDriveRestore() {
	t := suite.T()
	eb := NewOneDriveRestore()
	s := eb.Selector
	or, err := s.ToOneDriveRestore()
	require.NoError(t, err)
	assert.Equal(t, or.Service, ServiceOneDrive)
	assert.NotZero(t, or.Scopes())
}
