package extensions

// Tests for extensions.go

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type ExtensionsUnitSuite struct {
	tester.Suite
}

func TestExtensionsUnitSuite(t *testing.T) {
	suite.Run(t, &ExtensionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type mockExtension struct {
	numBytes int
	data     map[string]any
}

func (me *mockExtension) WrapItem(
	_ details.ItemInfo,
	rc io.ReadCloser,
) (io.ReadCloser, error) {
	p := make([]byte, 4)

	n, err := rc.Read(p)
	if err != nil && err != io.EOF {
		return nil, err
	}

	me.numBytes += n

	return rc, nil
}

func (me *mockExtension) OutputData() map[string]any {
	me.data["numBytes"] = me.numBytes
	return me.data
}

func readFrom(rc io.ReadCloser) error {
	defer rc.Close()

	p := make([]byte, 4)

	for {
		_, err := rc.Read(p)
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (suite *ExtensionsUnitSuite) TestExtensionsBasic() {
	table := []struct {
		name        string
		factories   []CorsoItemExtensionFactory
		payload     []byte
		expectedErr require.ErrorAssertionFunc
		rc          io.ReadCloser
	}{
		{
			name: "basic",
			factories: []CorsoItemExtensionFactory{
				func() CorsoItemExtension {
					return &mockExtension{data: map[string]any{}}
				},
			},
			payload:     []byte("hello world"),
			expectedErr: require.NoError,
			rc:          io.NopCloser(bytes.NewReader([]byte("hello world"))),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)
			defer flush()

			ehFactory := ExtensionHandlerFactory(func(
				info details.ItemInfo,
				rc io.ReadCloser,
				factory []CorsoItemExtensionFactory,
			) (ExtensionHandler, error) {
				return NewExtensionHandler(info, rc, factory)
			})

			ext, err := ehFactory(details.ItemInfo{}, test.rc, test.factories)
			require.NoError(suite.T(), err)

			err = readFrom(ext)
			require.NoError(suite.T(), err)

			kv, err := ext.GetExtensionData(ctx)
			require.NoError(suite.T(), err)

			require.Equal(suite.T(), len(test.payload), kv["numBytes"])
		})
	}
}
