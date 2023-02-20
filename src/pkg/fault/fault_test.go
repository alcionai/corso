package fault_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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
		suite.T().Run(test.name, func(t *testing.T) {
			n := fault.New(test.failFast)
			require.NotNil(t, n)
			require.NoError(t, n.Err())
			require.Empty(t, n.Errs())

			e := n.Fail(test.fail)
			require.NotNil(t, e)

			e = n.Add(test.add)
			require.NotNil(t, e)

			test.expect(t, n.Err())
		})
	}
}

func (suite *FaultErrorsUnitSuite) TestFail() {
	t := suite.T()

	n := fault.New(false)
	require.NotNil(t, n)
	require.NoError(t, n.Err())
	require.Empty(t, n.Errs())

	n.Fail(assert.AnError)
	assert.Error(t, n.Err())
	assert.Empty(t, n.Errs())

	n.Fail(assert.AnError)
	assert.Error(t, n.Err())
	assert.NotEmpty(t, n.Errs())
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

			e = n.Add(test.add)
			require.NotNil(t, e)

			test.expect(t, n.Errs())
		})
	}
}

func (suite *FaultErrorsUnitSuite) TestAdd() {
	t := suite.T()

	n := fault.New(true)
	require.NotNil(t, n)

	n.Add(assert.AnError)
	assert.Error(t, n.Err())
	assert.Len(t, n.Errs(), 1)

	n.Add(assert.AnError)
	assert.Error(t, n.Err())
	assert.Len(t, n.Errs(), 2)
}

func (suite *FaultErrorsUnitSuite) TestData() {
	t := suite.T()

	// not fail-fast
	n := fault.New(false)
	require.NotNil(t, n)

	n.Fail(errors.New("fail"))
	n.Add(errors.New("1"))
	n.Add(errors.New("2"))

	d := n.Data()
	assert.Equal(t, n.Err(), d.Err)
	assert.ElementsMatch(t, n.Errs(), d.Errs)
	assert.False(t, d.FailFast)

	// fail-fast
	n = fault.New(true)
	require.NotNil(t, n)

	n.Fail(errors.New("fail"))
	n.Add(errors.New("1"))
	n.Add(errors.New("2"))

	d = n.Data()
	assert.Equal(t, n.Err(), d.Err)
	assert.ElementsMatch(t, n.Errs(), d.Errs)
	assert.True(t, d.FailFast)
}

func (suite *FaultErrorsUnitSuite) TestMarshalUnmarshal() {
	t := suite.T()

	// not fail-fast
	n := fault.New(false)
	require.NotNil(t, n)

	n.Add(errors.New("1"))
	n.Add(errors.New("2"))

	data := n.Data()

	jsonStr, err := json.Marshal(data)
	require.NoError(t, err)

	um := fault.ErrorsData{}

	err = json.Unmarshal(jsonStr, &um)
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

	um := fault.ErrorsData{}

	err = json.Unmarshal(jsonStr, &um)
	require.NoError(t, err)
}
