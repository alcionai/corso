package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type OnedriveSourceSuite struct {
	suite.Suite
}

func TestOnedriveSourceSuite(t *testing.T) {
	suite.Run(t, new(OnedriveSourceSuite))
}

func (suite *OnedriveSourceSuite) TestNewOnedriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup()
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OnedriveSourceSuite) TestToOnedriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup()
	s := ob.Selector
	ob, err := s.ToOneDriveBackup()
	require.NoError(t, err)
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OnedriveSourceSuite) TestOnedriveBackup_DiscreteScopes() {
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

func (suite *OnedriveSourceSuite) TestOnedriveSelector_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)
	userScopes := sel.Users([]string{u1, u2})
	for _, scope := range userScopes {
		// Scope value is either u1 or u2
		assert.Contains(t, []string{u1, u2}, scope[OneDriveUser.String()])
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
			require.Equal(t, 2, len(test.scopesToCheck))
			for _, scope := range test.scopesToCheck {
				// Scope value is u1,u2
				assert.Contains(t, join(u1, u2), scope[OneDriveUser.String()])
			}
		})
	}
}

func (suite *OnedriveSourceSuite) TestOneDriveSelector_Include_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Include(sel.Users([]string{u1, u2}))
	scopes := sel.Includes
	require.Len(t, scopes, 2)

	for _, scope := range scopes {
		assert.Contains(t, join(u1, u2), scope[OneDriveUser.String()])
	}
}

func (suite *OnedriveSourceSuite) TestOneDriveSelector_Exclude_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Exclude(sel.Users([]string{u1, u2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 2)

	for _, scope := range scopes {
		assert.Contains(t, join(u1, u2), scope[OneDriveUser.String()])
	}
}
