package path

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ServiceCategoryUnitSuite struct {
	tester.Suite
}

func TestServiceCategoryUnitSuite(t *testing.T) {
	s := &ServiceCategoryUnitSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

func (suite *ServiceCategoryUnitSuite) TestVerifyPrefixValues() {
	table := []struct {
		name     string
		service  ServiceType
		category CategoryType
		check    assert.ErrorAssertionFunc
	}{
		{
			name:     "UnknownService",
			service:  UnknownService,
			category: EmailCategory,
			check:    assert.Error,
		},
		{
			name:     "UnknownCategory",
			service:  ExchangeService,
			category: UnknownCategory,
			check:    assert.Error,
		},
		{
			name:     "BadServiceType",
			service:  ServiceType(-1),
			category: EmailCategory,
			check:    assert.Error,
		},
		{
			name:     "BadCategoryType",
			service:  ExchangeService,
			category: CategoryType(-1),
			check:    assert.Error,
		},
		{
			name:     "ExchangeEmail",
			service:  ExchangeService,
			category: EmailCategory,
			check:    assert.NoError,
		},
		{
			name:     "ExchangeContacts",
			service:  ExchangeService,
			category: ContactsCategory,
			check:    assert.NoError,
		},
		{
			name:     "ExchangeEvents",
			service:  ExchangeService,
			category: EventsCategory,
			check:    assert.NoError,
		},
		{
			name:     "OneDriveFiles",
			service:  OneDriveService,
			category: FilesCategory,
			check:    assert.NoError,
		},
		{
			name:     "SharePointLibraries",
			service:  SharePointService,
			category: LibrariesCategory,
			check:    assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			srs := []ServiceResource{{test.service, "resource"}}

			err := verifyPrefixValues("tid", srs, test.category)
			test.check(t, err, clues.ToCore(err))
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
		suite.Run(test.name, func() {
			assert.Equal(
				suite.T(),
				test.expected,
				ToServiceType(test.service))
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
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, test.expected, ToCategoryType(test.category))
		})
	}
}
