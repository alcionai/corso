package mock

import (
	"testing"

	"github.com/alcionai/clues"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/connector/support"
)

type MockSuite struct {
	tester.Suite
}

func TestMockSuite(t *testing.T) {
	suite.Run(t, &MockSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MockSuite) TestMockByteHydration() {
	subject := "Mock Hydration"
	tests := []struct {
		name           string
		transformation func(t *testing.T) error
	}{
		{
			name: "SharePoint: List Empty",
			transformation: func(t *testing.T) error {
				emptyMap := make(map[string]string)
				temp := List(subject, "Artist", emptyMap)
				writer := kioser.NewJsonSerializationWriter()
				err := writer.WriteObjectValue("", temp)
				require.NoError(t, err, clues.ToCore(err))

				bytes, err := writer.GetSerializedContent()
				require.NoError(t, err, clues.ToCore(err))

				_, err = support.CreateListFromBytes(bytes)

				return err
			},
		},
		{
			name: "SharePoint: List 6 Items",
			transformation: func(t *testing.T) error {
				bytes, err := ListBytes(subject)
				require.NoError(t, err, clues.ToCore(err))
				_, err = support.CreateListFromBytes(bytes)
				return err
			},
		},
		{
			name: "SharePoint: Page",
			transformation: func(t *testing.T) error {
				bytes := Page(subject)
				_, err := support.CreatePageFromBytes(bytes)

				return err
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := test.transformation(t)
			assert.NoError(t, err, clues.ToCore(err))
		})
	}
}
