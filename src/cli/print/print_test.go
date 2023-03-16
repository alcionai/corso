package print

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
)

type PrintUnitSuite struct {
	tester.Suite
}

func TestPrintUnitSuite(t *testing.T) {
	suite.Run(t, &PrintUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PrintUnitSuite) TestOnly() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	c := &cobra.Command{}
	ctx = SetRootCmd(ctx, c)

	err := Only(ctx, nil)
	assert.NoError(t, err, clues.ToCore(err))
	assert.True(t, c.SilenceUsage)
}

func (suite *PrintUnitSuite) TestErr() {
	t := suite.T()
	b := bytes.Buffer{}
	msg := "I have seen the fnords!"

	err(&b, msg)
	assert.Contains(t, b.String(), "Error: ")
	assert.Contains(t, b.String(), msg)
}

func (suite *PrintUnitSuite) TestInfo() {
	t := suite.T()
	b := bytes.Buffer{}
	msg := "I have seen the fnords!"

	info(&b, msg)
	assert.Contains(t, b.String(), msg)
}

func (suite *PrintUnitSuite) TestInfof() {
	t := suite.T()
	b := bytes.Buffer{}
	msg := "I have seen the fnords!"
	msg2 := "smarf"

	infof(&b, msg, msg2)
	bs := b.String()
	assert.Contains(t, bs, msg)
	assert.Contains(t, bs, msg2)
}
