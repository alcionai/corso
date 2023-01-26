package path

import (
	"strings"
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
			_, _, err := validateServiceAndCategoryStrings(test.service, test.category)
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
		{
			name:             "ExchangeContacts",
			service:          ExchangeService.String(),
			category:         ContactsCategory.String(),
			expectedService:  ExchangeService,
			expectedCategory: ContactsCategory,
			check:            assert.NoError,
		},
		{
			name:             "ExchangeEvents",
			service:          ExchangeService.String(),
			category:         EventsCategory.String(),
			expectedService:  ExchangeService,
			expectedCategory: EventsCategory,
			check:            assert.NoError,
		},
		{
			name:             "OneDriveFiles",
			service:          OneDriveService.String(),
			category:         FilesCategory.String(),
			expectedService:  OneDriveService,
			expectedCategory: FilesCategory,
			check:            assert.NoError,
		},
		{
			name:             "SharePointLibraries",
			service:          SharePointService.String(),
			category:         LibrariesCategory.String(),
			expectedService:  SharePointService,
			expectedCategory: LibrariesCategory,
			check:            assert.NoError,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			s, c, err := validateServiceAndCategoryStrings(test.service, test.category)
			test.check(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, test.expectedService, s)
			assert.Equal(t, test.expectedCategory, c)
		})
	}
}

func (suite *ServiceCategoryUnitSuite) TestToServiceType() {
	table := []struct {
		name     string
		service  string
		expected ServiceType
	}{
		{
			name:     "SameCase",
			service:  ExchangeMetadataService.String(),
			expected: ExchangeMetadataService,
		},
		{
			name:     "DifferentCase",
			service:  strings.ToUpper(ExchangeMetadataService.String()),
			expected: ExchangeMetadataService,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, toServiceType(test.service))
		})
	}
}

func (suite *ServiceCategoryUnitSuite) TestToCategoryType() {
	table := []struct {
		name     string
		category string
		expected CategoryType
	}{
		{
			name:     "SameCase",
			category: EmailCategory.String(),
			expected: EmailCategory,
		},
		{
			name:     "DifferentCase",
			category: strings.ToUpper(EmailCategory.String()),
			expected: EmailCategory,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ToCategoryType(test.category))
		})
	}
}
