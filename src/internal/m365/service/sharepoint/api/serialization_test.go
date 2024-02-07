package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	spMock "github.com/alcionai/canario/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/canario/src/internal/tester"
)

type SerializationUnitSuite struct {
	tester.Suite
}

func TestSerializationUnitSuite(t *testing.T) {
	suite.Run(t, &SerializationUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SerializationUnitSuite) TestCreateFromBytes() {
	listBytes, err := spMock.ListBytes("DataSupportSuite")
	require.NoError(suite.T(), err)

	tests := []struct {
		name          string
		byteArray     []byte
		parseableFunc serialization.ParsableFactory
		checkError    assert.ErrorAssertionFunc
		isNil         assert.ValueAssertionFunc
	}{
		{
			name:          "Valid List",
			byteArray:     listBytes,
			parseableFunc: models.CreateListFromDiscriminatorValue,
			checkError:    assert.NoError,
			isNil:         assert.NotNil,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := createFromBytes(test.byteArray, test.parseableFunc)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}
