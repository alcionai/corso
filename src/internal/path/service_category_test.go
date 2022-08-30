package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceCategoryUnitSuite struct {
	suite.Suite
}

func TestServiceCategoryUnitSuite(t *testing.T) {
	suite.Run(t, new(ServiceCategoryUnitSuite))
}

func (suite *ServiceCategoryUnitSuite) TestValidateServiceAndCategoryBadStringErrors() {
	table := []struct {
		name     string
		service  string
		category string
	}{
		{
			name:     "Service",
			service:  "foo",
			category: EmailCategory.String(),
		},
		{
			name:     "Category",
			service:  ExchangeService.String(),
			category: "foo",
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, _, err := validateServiceAndCategory(test.service, test.category)
			assert.Error(suite.T(), err)
		})
	}
}

func (suite *ServiceCategoryUnitSuite) TestValidateServiceAndCategory() {
	table := []struct {
		name             string
		service          string
		category         string
		expectedService  ServiceType
		expectedCategory CategoryType
		check            assert.ErrorAssertionFunc
	}{
		{
			name:     "UnknownService",
			service:  UnknownService.String(),
			category: EmailCategory.String(),
			check:    assert.Error,
		},
		{
			name:     "UnknownCategory",
			service:  ExchangeService.String(),
			category: UnknownCategory.String(),
			check:    assert.Error,
		},
		{
			name:     "BadServiceString",
			service:  "foo",
			category: EmailCategory.String(),
			check:    assert.Error,
		},
		{
			name:     "BadCategoryString",
			service:  ExchangeService.String(),
			category: "foo",
			check:    assert.Error,
		},
		{
			name:             "ExchangeEmail",
			service:          ExchangeService.String(),
			category:         EmailCategory.String(),
			expectedService:  ExchangeService,
			expectedCategory: EmailCategory,
			check:            assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			s, c, err := validateServiceAndCategory(test.service, test.category)
			test.check(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, test.expectedService, s)
			assert.Equal(t, test.expectedCategory, c)
		})
	}
}
