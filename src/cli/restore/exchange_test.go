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
		name          string
		u, f, m, rpid string
		errCheck      assert.ErrorAssertionFunc
	}{
		{"all populated", "u", "f", "m", "rpid", assert.NoError},
		{"folder missing user", "", "f", "m", "rpid", assert.Error},
		{"folder with wildcard user", "*", "f", "m", "rpid", assert.Error},
		{"mail missing user", "", "", "m", "rpid", assert.Error},
		{"mail missing folder", "u", "", "m", "rpid", assert.Error},
		{"mail with wildcard folder", "u", "*", "m", "rpid", assert.Error},
		{"missing restore point id", "u", "f", "m", "", assert.Error},
		{"all missing", "", "", "", "rpid", assert.NoError},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.errCheck(
				t,
				validateRestoreFlags(test.u, test.f, test.m, test.rpid),
			)
		})
	}
}
