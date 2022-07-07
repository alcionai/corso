package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/kopia"
)

type OperationSuite struct {
	suite.Suite
}

func TestOperationSuite(t *testing.T) {
	suite.Run(t, new(OperationSuite))
}

func (suite *OperationSuite) TestNewOperation() {
	t := suite.T()
	op := newOperation(Options{}, nil, nil)
	assert.NotNil(t, op.ID)
}

func (suite *OperationSuite) TestOperation_Validate() {
	kwStub := &kopia.Wrapper{}
	msStub := &kopia.ModelStore{}
	table := []struct {
		name     string
		kw       *kopia.Wrapper
		ms       *kopia.ModelStore
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", kwStub, msStub, assert.NoError},
		{"missing kopia", nil, msStub, assert.Error},
		{"missing kopia", kwStub, nil, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			op := newOperation(Options{}, test.kw, test.ms)
			test.errCheck(t, op.validate())
		})
	}
}
