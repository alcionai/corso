package source_test

import (
	"testing"

	"github.com/alcionai/corso/pkg/source"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SourceSuite struct {
	suite.Suite
}

func TestSourceSuite(t *testing.T) {
	suite.Run(t, new(SourceSuite))
}

func (suite *SourceSuite) TestSource_AddUsers() {
	table := []struct {
		name     string
		uids     []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"nil", nil, assert.NoError},
		{"zero val", []string{}, assert.NoError},
		{"single", []string{"fnord"}, assert.NoError},
		{"multi", []string{"fnord", "snarf"}, assert.NoError},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			s := source.NewSource(source.ServiceUnknown)
			err := s.AddUsers(test.uids...)
			test.errCheck(t, err)
		})
	}
}

func (suite *SourceSuite) TestSource_Users() {
	t := suite.T()

	s := source.NewSource(source.ServiceExchange)
	assert.NoError(t, s.AddUsers())
	us := s.Users()
	assert.Zero(t, len(us))
	assert.NotNil(t, us)

	assert.NoError(t, s.AddUsers("a"))
	us = s.Users()
	assert.Equal(t, len(us), 1)

	assert.NoError(t, s.AddUsers("b", "c"))
	us = s.Users()
	assert.Equal(t, len(us), 3)

	assert.Equal(t, us[0], "a")
	assert.Equal(t, us[1], "b")
	assert.Equal(t, us[2], "c")
}
