package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
)

type ServiceTypeUnitSuite struct {
	tester.Suite
}

func TestServiceTypeUnitSuite(t *testing.T) {
	suite.Run(t, &ServiceTypeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

var knownServices = []ServiceType{
	UnknownService,
	ExchangeService,
	OneDriveService,
	SharePointService,
	ExchangeMetadataService,
	OneDriveMetadataService,
	SharePointMetadataService,
	GroupsService,
	GroupsMetadataService,
}

func (suite *ServiceTypeUnitSuite) TestValildateServiceAndSubService() {
	table := map[ServiceType]map[ServiceType]assert.ErrorAssertionFunc{}

	for _, si := range knownServices {
		table[si] = map[ServiceType]assert.ErrorAssertionFunc{}
		for _, sj := range knownServices {
			table[si][sj] = assert.Error
		}
	}

	// expected successful
	table[GroupsService][SharePointService] = assert.NoError

	for srv, ti := range table {
		for sub, expect := range ti {
			suite.Run(srv.String()+"-"+sub.String(), func() {
				err := ValidateServiceAndSubService(srv, sub)
				expect(suite.T(), err, clues.ToCore(err))
			})
		}
	}
}
