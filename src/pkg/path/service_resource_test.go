package path

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ServiceResourceUnitSuite struct {
	tester.Suite
}

func TestServiceResourceUnitSuite(t *testing.T) {
	suite.Run(t, &ServiceResourceUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ServiceResourceUnitSuite) TestNewServiceResource() {
	table := []struct {
		name         string
		input        []any
		expectErr    assert.ErrorAssertionFunc
		expectResult []ServiceResource
	}{
		{
			name:         "empty",
			input:        []any{},
			expectErr:    assert.Error,
			expectResult: nil,
		},
		{
			name:         "odd elems: 1",
			input:        []any{ExchangeService},
			expectErr:    assert.Error,
			expectResult: nil,
		},
		{
			name:         "odd elems: 3",
			input:        []any{ExchangeService, "mailbox", OneDriveService},
			expectErr:    assert.Error,
			expectResult: nil,
		},
		{
			name:         "non-service even index",
			input:        []any{"foo", "bar"},
			expectErr:    assert.Error,
			expectResult: nil,
		},
		{
			name:         "non-string odd index",
			input:        []any{ExchangeService, OneDriveService},
			expectErr:    assert.Error,
			expectResult: nil,
		},
		{
			name:         "valid single",
			input:        []any{ExchangeService, "mailbox"},
			expectErr:    assert.NoError,
			expectResult: []ServiceResource{{ExchangeService, "mailbox"}},
		},
		{
			name:      "valid multiple",
			input:     []any{ExchangeService, "mailbox", OneDriveService, "user"},
			expectErr: assert.NoError,
			expectResult: []ServiceResource{
				{ExchangeService, "mailbox"},
				{OneDriveService, "user"},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := NewServiceResources(test.input...)
			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectResult, result)
		})
	}
}

func (suite *ServiceResourceUnitSuite) TestValidateServiceResources() {
	table := []struct {
		name   string
		srs    []ServiceResource
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "empty",
			srs:    []ServiceResource{},
			expect: assert.Error,
		},
		{
			name:   "invalid resource",
			srs:    []ServiceResource{{ExchangeService, ""}},
			expect: assert.Error,
		},
		{
			name: "invalid subservice",
			srs: []ServiceResource{
				{ExchangeService, "mailbox"},
				{OneDriveService, "user"},
			},
			expect: assert.Error,
		},
		{
			name: "valid",
			srs: []ServiceResource{
				{GroupsService, "group"},
				{SharePointService, "site"},
			},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			err := validateServiceResources(test.srs)
			test.expect(t, err, clues.ToCore(err))
		})
	}
}

func (suite *ServiceResourceUnitSuite) TestServiceResourceToElements() {
	table := []struct {
		name   string
		srs    []ServiceResource
		expect Elements
	}{
		{
			name:   "empty",
			srs:    []ServiceResource{},
			expect: Elements{},
		},
		{
			name:   "single",
			srs:    []ServiceResource{{ExchangeService, "user"}},
			expect: Elements{ExchangeService.String(), "user"},
		},
		{
			name: "multiple",
			srs: []ServiceResource{
				{ExchangeService, "mailbox"},
				{OneDriveService, "user"},
			},
			expect: Elements{
				ExchangeService.String(), "mailbox",
				OneDriveService.String(), "user",
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := serviceResourcesToElements(test.srs)

			// not ElementsMatch, order matters
			assert.Equal(t, test.expect, result)
		})
	}
}

func (suite *ServiceResourceUnitSuite) TestElementsToServiceResource() {
	table := []struct {
		name      string
		elems     Elements
		expectErr assert.ErrorAssertionFunc
		expectIdx int
		expectSRS []ServiceResource
	}{
		{
			name:      "empty",
			elems:     Elements{},
			expectErr: assert.Error,
			expectIdx: -1,
			expectSRS: nil,
		},
		{
			name:      "nil",
			elems:     nil,
			expectErr: assert.Error,
			expectIdx: -1,
			expectSRS: nil,
		},
		{
			name:      "non-service 0th elem",
			elems:     Elements{"fnords"},
			expectErr: assert.Error,
			expectIdx: -1,
			expectSRS: nil,
		},
		{
			name:      "non-service 2nd elem",
			elems:     Elements{ExchangeService.String(), "fnords", "smarf"},
			expectErr: assert.Error,
			expectIdx: -1,
			expectSRS: nil,
		},
		{
			name:      "single serviceResource",
			elems:     Elements{ExchangeService.String(), "fnords"},
			expectErr: assert.NoError,
			expectIdx: 2,
			expectSRS: []ServiceResource{{ExchangeService, "fnords"}},
		},
		{
			name:      "single serviceResource and extra value",
			elems:     Elements{ExchangeService.String(), "fnords", "smarf"},
			expectErr: assert.NoError,
			expectIdx: 2,
			expectSRS: []ServiceResource{{ExchangeService, "fnords"}},
		},
		{
			name:      "multiple serviceResource",
			elems:     Elements{ExchangeService.String(), "fnords", OneDriveService.String(), "smarf"},
			expectErr: assert.NoError,
			expectIdx: 4,
			expectSRS: []ServiceResource{{ExchangeService, "fnords"}, {OneDriveService, "smarf"}},
		},
		{
			name:      "multiple serviceResource and extra value",
			elems:     Elements{ExchangeService.String(), "fnords", OneDriveService.String(), "smarf", "flaboigans"},
			expectErr: assert.NoError,
			expectIdx: 4,
			expectSRS: []ServiceResource{{ExchangeService, "fnords"}, {OneDriveService, "smarf"}},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			srs, idx, err := elementsToServiceResources(test.elems)
			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectIdx, idx)
			assert.Equal(t, test.expectSRS, srs)
		})
	}
}
