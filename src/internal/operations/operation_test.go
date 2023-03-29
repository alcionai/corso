package operations

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/store"
)

type OperationSuite struct {
	tester.Suite
}

func TestOperationSuite(t *testing.T) {
	suite.Run(t, &OperationSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OperationSuite) TestNewOperation() {
	t := suite.T()
	op := newOperation(control.Options{}, events.Bus{}, nil, nil, nil)
	assert.Greater(t, op.CreatedAt, time.Time{})
}

func (suite *OperationSuite) TestOperation_Validate() {
	kwStub := &kopia.Wrapper{}
	swStub := &store.Wrapper{}
	gcStub := &connector.GraphConnector{}

	table := []struct {
		name     string
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		gc       *connector.GraphConnector
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", kwStub, swStub, gcStub, assert.NoError},
		{"missing kopia wrapper", nil, swStub, gcStub, assert.Error},
		{"missing store wrapper", kwStub, nil, gcStub, assert.Error},
		{"missing graph connector", kwStub, swStub, nil, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := newOperation(control.Options{}, events.Bus{}, test.kw, test.sw, test.gc).validate()
			test.errCheck(suite.T(), err, clues.ToCore(err))
		})
	}
}
