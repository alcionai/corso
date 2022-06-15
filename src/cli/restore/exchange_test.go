package restore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RestoreSuite struct {
	suite.Suite
}

func TestRestoreSuite(t *testing.T) {
	suite.Run(t, new(RestoreSuite))
}

func (suite *RestoreSuite) TestValidateRestoreFlags() {
	table := []struct {
		name     string
		u, f, m  string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", "u", "f", "m", assert.NoError},
		{"folder missing user", "", "f", "m", assert.Error},
		{"mail missing user", "", "", "m", assert.Error},
		{"mail missing folder", "u", "", "m", assert.Error},
		{"all missing", "", "", "", assert.NoError},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.errCheck(
				t,
				validateRestoreFlags(test.u, test.f, test.m),
			)
		})
	}
}
