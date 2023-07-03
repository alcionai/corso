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

func (suite *ExtensionsUnitSuite) TestAddItemExtensions() {
	type outputValidationFunc func(
		extRc io.ReadCloser,
		extInfo *details.ExtensionInfo,
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
				extInfo *details.ExtensionInfo,
				err error,
			) bool {
				return err == nil && extRc != nil && extInfo != nil
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
				extInfo *details.ExtensionInfo,
				err error,
			) bool {
				return err == nil && extRc != nil && extInfo != nil
			},
		},
		{
			name:      "no factories supplied",
			factories: nil,
			rc:        testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extInfo *details.ExtensionInfo,
				err error,
			) bool {
				return err != nil && extRc == nil && extInfo == nil
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
				extInfo *details.ExtensionInfo,
				err error,
			) bool {
				return err != nil && extRc == nil && extInfo == nil
			},
		},
		{
			name: "factory call returns error",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{
					shouldReturnError: true,
				},
			},
			rc: testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extInfo *details.ExtensionInfo,
				err error,
			) bool {
				return err != nil && extRc == nil && extInfo == nil
			},
		},
		{
			name: "one or more factory calls return error",
			factories: []CreateItemExtensioner{
				&MockItemExtensionFactory{},
				&MockItemExtensionFactory{
					shouldReturnError: true,
				},
			},
			rc: testRc,
			validateOutputs: func(
				extRc io.ReadCloser,
				extInfo *details.ExtensionInfo,
				err error,
			) bool {
				return err != nil && extRc == nil && extInfo == nil
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
				extInfo *details.ExtensionInfo,
				err error,
			) bool {
				return err != nil && extRc == nil && extInfo == nil
			},
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
				testItemInfo,
				test.factories)
			require.True(t, test.validateOutputs(extRc, extInfo, err))
		})
	}
}

// TODO(pandeyabs): Tests to verify RC wrapper ordering by AddItemExtensioner
