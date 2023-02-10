package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

type PointerSuite struct {
	suite.Suite
}

func TestPointerSuite(t *testing.T) {
	suite.Run(t, new(PointerSuite))
}

func (suite *PointerSuite) TestVal() {
	word := "Hello World"
	tables := []struct {
		name    string
		pointer *string
		isEmpty assert.ValueAssertionFunc
	}{
		{
			name:    "Nil",
			pointer: nil,
			isEmpty: assert.Empty,
		},
		{
			name:    "Generic",
			pointer: &word,
			isEmpty: assert.NotEmpty,
		},
	}

	for _, test := range tables {
		suite.T().Run(test.name, func(t *testing.T) {
			value := ptr.Val(test.pointer)
			test.isEmpty(t, value)
		})
	}
}
