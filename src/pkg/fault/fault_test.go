package fault_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester/aw"
	"github.com/alcionai/corso/src/pkg/fault"
)

type FaultErrorsUnitSuite struct {
	suite.Suite
}

func TestFaultErrorsUnitSuite(t *testing.T) {
	suite.Run(t, new(FaultErrorsUnitSuite))
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
			expect: aw.NoErr,
		},
		{
			name:     "nil, failFast",
			failFast: true,
			expect:   aw.NoErr,
		},
		{
			name:   "failed",
			fail:   assert.AnError,
			expect: aw.Err,
		},
		{
			name:     "failed, failFast",
			fail:     assert.AnError,
			failFast: true,
			expect:   aw.Err,
		},
		{
			name:   "added",
			add:    assert.AnError,
			expect: aw.NoErr,
		},
		{
			name:     "added, failFast",
			add:      assert.AnError,
			failFast: true,
			expect:   aw.Err,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			n := fault.New(test.failFast)
			require.NotNil(t, n)
			aw.MustNoErr(t, n.Failure())
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
	aw.MustNoErr(t, n.Failure())
	require.Empty(t, n.Recovered())

	n.Fail(assert.AnError)
	aw.Err(t, n.Failure())
	assert.Empty(t, n.Recovered())

	n.Fail(assert.AnError)
	aw.Err(t, n.Failure())
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
		suite.T().Run(test.name, func(t *testing.T) {
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
	aw.Err(t, n.Failure())
	assert.Len(t, n.Recovered(), 1)

	n.AddRecoverable(assert.AnError)
	aw.Err(t, n.Failure())
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
	aw.MustNoErr(t, err)

	err = json.Unmarshal(bs, &fault.Errors{})
	aw.MustNoErr(t, err)
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
	aw.MustNoErr(t, err)

	t.Logf("jsonStr is %s\n", jsonStr)

	um := fault.Errors{}

	err = json.Unmarshal(jsonStr, &um)
	aw.MustNoErr(t, err)
}

func (suite *FaultErrorsUnitSuite) TestTracker() {
	t := suite.T()

	eb := fault.New(false)

	lb := eb.Local()
	aw.NoErr(t, lb.Failure())
	assert.Empty(t, eb.Recovered())

	lb.AddRecoverable(assert.AnError)
	aw.NoErr(t, lb.Failure())
	aw.NoErr(t, eb.Failure())
	assert.NotEmpty(t, eb.Recovered())

	ebt := fault.New(true)

	lbt := ebt.Local()
	aw.NoErr(t, lbt.Failure())
	assert.Empty(t, ebt.Recovered())

	lbt.AddRecoverable(assert.AnError)
	aw.Err(t, lbt.Failure())
	aw.Err(t, ebt.Failure())
	assert.NotEmpty(t, ebt.Recovered())
}
