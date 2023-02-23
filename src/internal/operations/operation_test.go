package operations

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester/aw"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/store"
)

type OperationSuite struct {
	suite.Suite
}

func TestOperationSuite(t *testing.T) {
	suite.Run(t, new(OperationSuite))
}

func (suite *OperationSuite) TestNewOperation() {
	t := suite.T()
	op := newOperation(control.Options{}, events.Bus{}, nil, nil)
	assert.Greater(t, op.CreatedAt, time.Time{})
}

func (suite *OperationSuite) TestOperation_Validate() {
	kwStub := &kopia.Wrapper{}
	swStub := &store.Wrapper{}

	table := []struct {
		name     string
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", kwStub, swStub, aw.NoErr},
		{"missing kopia wrapper", nil, swStub, aw.Err},
		{"missing store wrapper", kwStub, nil, aw.Err},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			op := newOperation(control.Options{}, events.Bus{}, test.kw, test.sw)
			test.errCheck(t, op.validate())
		})
	}
}
