package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OpProgressSuite struct {
	suite.Suite
}

func TestOpProgressSuite(t *testing.T) {
	suite.Run(t, new(OpProgressSuite))
}

func (suite *OpProgressSuite) TestNewOpProgress() {
	t := suite.T()

	op := newOpProgress()
	assert.NotNil(t, op.progressChan)
	assert.NotNil(t, op.errorChan)

	op.Close()
	assert.Nil(t, op.progressChan)
	assert.Nil(t, op.errorChan)
}

func (suite *OpProgressSuite) TestOpProgress_Report() {
	t := suite.T()

	op := newOpProgress()
	assert.NotPanics(t,
		func() {
			op.Report("test")
		})

	ch := op.progressChan
	op.progressChan = nil
	assert.NotPanics(t,
		func() {
			op.Report("test")
		})

	op.progressChan = ch
	op.Close()

	assert.Panics(t,
		func() {
			op.Report("test")
		})
}

func (suite *OpProgressSuite) TestOpProgress_Error() {
	t := suite.T()

	op := newOpProgress()
	assert.NotPanics(t,
		func() {
			op.Error(assert.AnError)
		})

	ch := op.progressChan
	op.progressChan = nil
	assert.NotPanics(t,
		func() {
			op.Error(assert.AnError)
		})

	op.progressChan = ch
	op.Close()

	assert.Panics(t,
		func() {
			op.Error(assert.AnError)
		})
}
