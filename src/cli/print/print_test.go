package print

import (
	"bytes"
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
	oldRoot := rootCmd
	defer SetRootCommand(oldRoot)
	SetRootCommand(c)
	assert.NoError(t, Only(nil))
	assert.True(t, c.SilenceUsage)
}

func (suite *PrintUnitSuite) TestInfo() {
	t := suite.T()
	var b bytes.Buffer
	msg := "I have seen the fnords!"
	info(&b, msg)
	assert.Contains(t, b.String(), msg)
}

func (suite *PrintUnitSuite) TestInfof() {
	t := suite.T()
	var b bytes.Buffer
	msg := "I have seen the fnords!"
	msg2 := "smarf"
	infof(&b, msg, msg2)
	bs := b.String()
	assert.Contains(t, bs, msg)
	assert.Contains(t, bs, msg2)
}
