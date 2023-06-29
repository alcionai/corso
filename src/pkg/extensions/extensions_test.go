package extensions

// Tests for extensions.go

import (
	"bytes"
	"context"
	"hash/crc32"
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
				func(
					ctx context.Context,
					rc io.ReadCloser,
					info details.ItemInfo,
					extInfo *details.ExtensionInfo,
				) (CorsoItemExtension, error) {
					return NewMockExtension(ctx, rc, info, extInfo)
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

			extRc, extInfo, err := AddItemExtensions(
				ctx,
				test.rc,
				details.ItemInfo{},
				test.factories)
			require.NoError(suite.T(), err)

			err = readFrom(extRc)
			require.NoError(suite.T(), err)

			require.Equal(suite.T(), len(test.payload), extInfo.Data["numBytes"])

			// verify crc32
			c := extInfo.Data["crc32"].(uint32)
			require.Equal(suite.T(), c, crc32.ChecksumIEEE(test.payload))
		})
	}
}
