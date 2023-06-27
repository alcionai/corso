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
) io.ReadCloser {
	p := make([]byte, 11)

	n, err := rc.Read(p)
	if err != nil && err != io.EOF {
		return nil
	}

	me.numBytes += n

	return rc
}

func (me *mockExtension) OutputData() map[string]any {
	me.data["numBytes"] = me.numBytes
	return me.data
}

func readFrom(rc io.ReadCloser) error {
	defer rc.Close()

	p := make([]byte, 11)

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
		factory     BackupItemExtensionFactory
		payload     []byte
		expectedErr require.ErrorAssertionFunc
		rc          io.ReadCloser
	}{
		{
			name: "basic",
			factory: func() BackupItemExtension {
				return &mockExtension{
					numBytes: 0,
					data:     map[string]any{},
				}
			},
			payload:     []byte("hello world"),
			expectedErr: require.NoError,
			rc:          io.NopCloser(bytes.NewReader([]byte("hello world"))),
		},
	}

	for _, test := range table {
		ext, err := newExtension(
			details.ItemInfo{},
			test.rc,
			[]BackupItemExtensionFactory{test.factory})
		require.NoError(suite.T(), err)

		err = readFrom(ext)
		require.NoError(suite.T(), err)

		kv, err := ext.GetExtensionData()
		require.NoError(suite.T(), err)

		require.Equal(suite.T(), 1, len(kv))
		require.Equal(suite.T(), len(test.payload), kv["numBytes"])
	}
}
