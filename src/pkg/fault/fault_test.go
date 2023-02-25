package fault_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type FaultErrorsUnitSuite struct {
	tester.Suite
}

func TestFaultErrorsUnitSuite(t *testing.T) {
	suite.Run(t, &FaultErrorsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *FaultErrorsUnitSuite) TestNew() {
	t := suite.T()

	n := fault.New(false)
	assert.NotNil(t, n)

	n = fault.New(true)
	assert.NotNil(t, n)
}

func (suite *FaultErrorsUnitSuite) TestErr() {
	table := []struct {
		name     string
		failFast bool
		fail     error
		add      error
		expect   assert.ErrorAssertionFunc
	}{
		{
			name:   "nil",
			expect: assert.NoError,
		},
		{
			name:     "nil, failFast",
			failFast: true,
			expect:   assert.NoError,
		},
		{
			name:   "failed",
			fail:   assert.AnError,
			expect: assert.Error,
		},
		{
			name:     "failed, failFast",
			fail:     assert.AnError,
			failFast: true,
			expect:   assert.Error,
		},
		{
			name:   "added",
			add:    assert.AnError,
			expect: assert.NoError,
		},
		{
			name:     "added, failFast",
			add:      assert.AnError,
			failFast: true,
			expect:   assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			n := fault.New(test.failFast)
			require.NotNil(t, n)
			require.NoError(t, n.Failure())
			require.Empty(t, n.Recovered())

			e := n.Fail(test.fail)
			require.NotNil(t, e)

			e = n.AddRecoverable(test.add)
			require.NotNil(t, e)

			test.expect(t, n.Failure())
		})
	}
}

func (suite *FaultErrorsUnitSuite) TestFail() {
	t := suite.T()

	n := fault.New(false)
	require.NotNil(t, n)
	require.NoError(t, n.Failure())
	require.Empty(t, n.Recovered())

	n.Fail(assert.AnError)
	assert.Error(t, n.Failure())
	assert.Empty(t, n.Recovered())

	n.Fail(assert.AnError)
	assert.Error(t, n.Failure())
	assert.NotEmpty(t, n.Recovered())
}

func (suite *FaultErrorsUnitSuite) TestErrs() {
	table := []struct {
		name     string
		failFast bool
		fail     error
		add      error
		expect   assert.ValueAssertionFunc
	}{
		{
			name:   "nil",
			expect: assert.Empty,
		},
		{
			name:     "nil, failFast",
			failFast: true,
			expect:   assert.Empty,
		},
		{
			name:   "failed",
			fail:   assert.AnError,
			expect: assert.Empty,
		},
		{
			name:     "failed, failFast",
			fail:     assert.AnError,
			failFast: true,
			expect:   assert.Empty,
		},
		{
			name:   "added",
			add:    assert.AnError,
			expect: assert.NotEmpty,
		},
		{
			name:     "added, failFast",
			add:      assert.AnError,
			failFast: true,
			expect:   assert.NotEmpty,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			n := fault.New(test.failFast)
			require.NotNil(t, n)

			e := n.Fail(test.fail)
			require.NotNil(t, e)

			e = n.AddRecoverable(test.add)
			require.NotNil(t, e)

			test.expect(t, n.Recovered())
		})
	}
}

func (suite *FaultErrorsUnitSuite) TestAdd() {
	t := suite.T()

	n := fault.New(true)
	require.NotNil(t, n)

	n.AddRecoverable(assert.AnError)
	assert.Error(t, n.Failure())
	assert.Len(t, n.Recovered(), 1)

	n.AddRecoverable(assert.AnError)
	assert.Error(t, n.Failure())
	assert.Len(t, n.Recovered(), 2)
}

func (suite *FaultErrorsUnitSuite) TestErrors() {
	t := suite.T()

	// not fail-fast
	n := fault.New(false)
	require.NotNil(t, n)

	n.Fail(errors.New("fail"))
	n.AddRecoverable(errors.New("1"))
	n.AddRecoverable(errors.New("2"))

	d := n.Errors()
	assert.Equal(t, n.Failure(), d.Failure)
	assert.ElementsMatch(t, n.Recovered(), d.Recovered)
	assert.False(t, d.FailFast)

	// fail-fast
	n = fault.New(true)
	require.NotNil(t, n)

	n.Fail(errors.New("fail"))
	n.AddRecoverable(errors.New("1"))
	n.AddRecoverable(errors.New("2"))

	d = n.Errors()
	assert.Equal(t, n.Failure(), d.Failure)
	assert.ElementsMatch(t, n.Recovered(), d.Recovered)
	assert.True(t, d.FailFast)
}

func (suite *FaultErrorsUnitSuite) TestMarshalUnmarshal() {
	t := suite.T()

	// not fail-fast
	n := fault.New(false)
	require.NotNil(t, n)

	n.AddRecoverable(errors.New("1"))
	n.AddRecoverable(errors.New("2"))

	bs, err := json.Marshal(n.Errors())
	require.NoError(t, err)

	err = json.Unmarshal(bs, &fault.Errors{})
	require.NoError(t, err)
}

type legacyErrorsData struct {
	Err      error   `json:"err"`
	Errs     []error `json:"errs"`
	FailFast bool    `json:"failFast"`
}

func (suite *FaultErrorsUnitSuite) TestUnmarshalLegacy() {
	t := suite.T()

	oldData := &legacyErrorsData{
		Errs: []error{fmt.Errorf("foo error"), fmt.Errorf("foo error"), fmt.Errorf("foo error")},
	}

	jsonStr, err := json.Marshal(oldData)
	require.NoError(t, err)

	t.Logf("jsonStr is %s\n", jsonStr)

	um := fault.Errors{}

	err = json.Unmarshal(jsonStr, &um)
	require.NoError(t, err)
}

func (suite *FaultErrorsUnitSuite) TestTracker() {
	t := suite.T()

	eb := fault.New(false)

	lb := eb.Local()
	assert.NoError(t, lb.Failure())
	assert.Empty(t, eb.Recovered())

	lb.AddRecoverable(assert.AnError)
	assert.NoError(t, lb.Failure())
	assert.NoError(t, eb.Failure())
	assert.NotEmpty(t, eb.Recovered())

	ebt := fault.New(true)

	lbt := ebt.Local()
	assert.NoError(t, lbt.Failure())
	assert.Empty(t, ebt.Recovered())

	lbt.AddRecoverable(assert.AnError)
	assert.Error(t, lbt.Failure())
	assert.Error(t, ebt.Failure())
	assert.NotEmpty(t, ebt.Recovered())
}
