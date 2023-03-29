package tester_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
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

type TesterLoadSuite struct {
	tester.Suite
	called bool
}

func TestTesterLoadSuite(t *testing.T) {
	suite.Run(t, &TesterLoadSuite{Suite: tester.NewLoadSuite(t, nil)})
}

func (suite *TesterLoadSuite) SetupSuite() {
	suite.called = true
}

func (suite *TesterLoadSuite) TestLoadSuite() {
	require.True(suite.T(), suite.called)
}

type TesterNightlySuite struct {
	tester.Suite
	called bool
}

func TestTesterNightlySuite(t *testing.T) {
	suite.Run(t, &TesterNightlySuite{Suite: tester.NewNightlySuite(t, nil)})
}

func (suite *TesterNightlySuite) SetupSuite() {
	suite.called = true
}

func (suite *TesterNightlySuite) TestNightlySuite() {
	require.True(suite.T(), suite.called)
}
