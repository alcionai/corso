package print

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PrintUnitSuite struct {
	suite.Suite
}

func TestPrintUnitSuite(t *testing.T) {
	suite.Run(t, new(PrintUnitSuite))
}

func (suite *PrintUnitSuite) TestOnly() {
	t := suite.T()
	c := &cobra.Command{}
	// cannot use tester.NewContext() here: circular imports
	ctx := SetRootCmd(context.Background(), c)
	assert.NoError(t, Only(ctx, nil))
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
