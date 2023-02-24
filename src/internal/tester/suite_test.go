package tester_test

import (
	"testing"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TesterUnitSuite struct {
	tester.Suite
	called bool
}

func TestTesterUnitSuite(t *testing.T) {
	suite.Run(t, &TesterUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *TesterUnitSuite) SetupSuite() {
	suite.called = true
}

func (suite *TesterUnitSuite) TestUnitSuite() {
	require.True(suite.T(), suite.called)
}

type TesterIntegrationSuite struct {
	tester.Suite
	called bool
}

func TestTesterIntegrationSuite(t *testing.T) {
	suite.Run(t, &TesterIntegrationSuite{Suite: tester.NewIntegrationSuite(t, nil)})
}

func (suite *TesterIntegrationSuite) SetupSuite() {
	suite.called = true
}

func (suite *TesterIntegrationSuite) TestIntegrationSuite() {
	require.True(suite.T(), suite.called)
}

type TesterE2ESuite struct {
	tester.Suite
	called bool
}

func TestTesterE2ESuite(t *testing.T) {
	suite.Run(t, &TesterE2ESuite{Suite: tester.NewE2ESuite(t, nil)})
}

func (suite *TesterE2ESuite) SetupSuite() {
	suite.called = true
}

func (suite *TesterE2ESuite) TestE2ESuite() {
	require.True(suite.T(), suite.called)
}
