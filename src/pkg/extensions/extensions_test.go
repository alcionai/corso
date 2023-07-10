package extensions

// Tests for extensions.go

import (
	"bytes"
	"errors"
	"hash/crc32"
	"io"
	"testing"

	"github.com/alcionai/clues"
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

func (suite *ExtensionsUnitSuite) TestAddItemExtensions() {
	type outputValidationFunc func(
		extRc io.ReadCloser,
		extData *details.ExtensionData,
		err error,
	) bool

	var (
		testRc       = io.NopCloser(bytes.NewReader([]byte("some data")))
		testItemInfo = details.ItemInfo{
			OneDrive: &details.OneDriveInfo{
				DriveID: "driveID",
			},
		}
	)

	table := []struct {
		name            string
		factories       []CreateItemExtensioner
		rc              io.ReadCloser
		validateOutputs outputValidationFunc
	}{
		{
			name: "happy path",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{},
			},
			rc: testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extData *details.ExtensionData,
				err error,
			) bool {
				return err == nil && extRc != nil && extData != nil
			},
		},
		{
			name: "multiple valid factories",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{},
				&MockItemExtensionFactory{},
			},
			rc: testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extData *details.ExtensionData,
				err error,
			) bool {
				return err == nil && extRc != nil && extData != nil
			},
		},
		{
			name:      "no factories supplied",
			factories: nil,
			rc:        testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extData *details.ExtensionData,
				err error,
			) bool {
				return err != nil && extRc == nil && extData == nil
			},
		},
		{
			name: "factory slice contains nil",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{},
				nil,
				&MockItemExtensionFactory{},
			},
			rc: testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extData *details.ExtensionData,
				err error,
			) bool {
				return err != nil && extRc == nil && extData == nil
			},
		},
		{
			name: "factory call returns error",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{
					FailOnFactoryCreation: true,
				},
			},
			rc: testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extData *details.ExtensionData,
				err error,
			) bool {
				return err != nil && extRc == nil && extData == nil
			},
		},
		{
			name: "one or more factory calls return error",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{},
				&MockItemExtensionFactory{
					FailOnFactoryCreation: true,
				},
			},
			rc: testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extData *details.ExtensionData,
				err error,
			) bool {
				return err != nil && extRc == nil && extData == nil
			},
		},
		{
			name: "nil inner rc",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{},
			},
			rc: nil,
			validateOutputs: func(
				extRc io.ReadCloser,
				extData *details.ExtensionData,
				err error,
			) bool {
				return err != nil && extRc == nil && extData == nil
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)
			defer flush()

			extRc, extData, err := AddItemExtensions(
				ctx,
				test.rc,
				testItemInfo,
				test.factories)
			require.True(t, test.validateOutputs(extRc, extData, err))
		})
	}
}

func readFrom(rc io.ReadCloser) error {
	defer rc.Close()

	var err error

	p := make([]byte, 4)

	for err == nil {
		_, err := rc.Read(p)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (suite *ExtensionsUnitSuite) TestReadCloserWrappers() {
	data := []byte("hello world!")

	table := []struct {
		name      string
		factories []CreateItemExtensioner
		payload   []byte
		check     require.ErrorAssertionFunc
		rc        io.ReadCloser
	}{
		{
			name: "happy path",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{},
			},
			payload: data,
			check:   require.NoError,
			rc:      io.NopCloser(bytes.NewReader(data)),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)
			defer flush()

			extRc, extData, err := AddItemExtensions(
				ctx,
				test.rc,
				details.ItemInfo{},
				test.factories)
			require.NoError(suite.T(), err)

			err = readFrom(extRc)
			test.check(t, err, clues.ToCore(err))

			if err == nil {
				require.Equal(suite.T(), len(test.payload), int(extData.Data[KNumBytes].(int64)))
				c := extData.Data[KCrc32].(uint32)
				require.Equal(suite.T(), c, crc32.ChecksumIEEE(test.payload))
			}
		})
	}
}

// TODO(pandeyabs): Tests to verify RC wrapper ordering by AddItemExtensioner
