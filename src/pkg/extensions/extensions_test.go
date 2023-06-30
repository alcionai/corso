package extensions

// Tests for extensions.go

import (
	"bytes"
	"context"
	"hash/crc32"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

var _ CorsoItemExtension = &MockExtension{}

// Temporary, testing purposes only
type MockExtension struct {
	// TODO: Add cumlulative crc32 checksum
	numBytes    int
	crc32       uint32
	info        details.ItemInfo
	extInfo     *details.ExtensionInfo
	innerRc     io.ReadCloser
	ctx         context.Context
	failOnRead  bool
	failOnClose bool
}

func (me *MockExtension) Read(p []byte) (int, error) {
	if me.failOnRead {
		return 0, clues.New("mock read error")
	}

	n, err := me.innerRc.Read(p)
	if err != nil && err != io.EOF {
		logger.CtxErr(me.ctx, err).Error("inner read error")
		return n, err
	}

	me.numBytes += n
	me.crc32 = crc32.Update(me.crc32, crc32.IEEETable, p[:n])

	if err == io.EOF {
		logger.Ctx(me.ctx).Info("mock extension reached EOF")
		me.extInfo.Data["numBytes"] = me.numBytes
		me.extInfo.Data["crc32"] = me.crc32
	}

	return n, err
}

func (me *MockExtension) Close() error {
	if me.failOnClose {
		return clues.New("mock close error")
	}

	err := me.innerRc.Close()
	if err != nil {
		return err
	}

	me.extInfo.Data["numBytes"] = me.numBytes
	me.extInfo.Data["crc32"] = me.crc32
	logger.Ctx(me.ctx).Infow(
		"mock extension closed",
		"numBytes", me.numBytes, "crc32", me.crc32)

	return nil
}

func NewMockExtension(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	extInfo *details.ExtensionInfo,
) (CorsoItemExtension, error) {
	return &MockExtension{
		ctx:     ctx,
		innerRc: rc,
		info:    info,
		extInfo: extInfo,
	}, nil
}

type ExtensionsUnitSuite struct {
	tester.Suite
}

func TestExtensionsUnitSuite(t *testing.T) {
	suite.Run(t, &ExtensionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// func readFrom(rc io.ReadCloser) error {
// 	defer rc.Close()

// 	p := make([]byte, 4)

// 	for {
// 		_, err := rc.Read(p)
// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

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
		factories       []CorsoItemExtensionFactory
		rc              io.ReadCloser
		validateOutputs outputValidationFunc
	}{
		{
			name: "happy path",
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
			factories: []CorsoItemExtensionFactory{
				NewMockExtension,
				NewMockExtension,
				NewMockExtension,
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
			factories: []CorsoItemExtensionFactory{
				NewMockExtension,
				nil,
				NewMockExtension,
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
			factories: []CorsoItemExtensionFactory{
				func(
					ctx context.Context,
					rc io.ReadCloser,
					info details.ItemInfo,
					extInfo *details.ExtensionInfo,
				) (CorsoItemExtension, error) {
					return nil, clues.New("creating extension")
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
			factories: []CorsoItemExtensionFactory{
				NewMockExtension,
				func(
					ctx context.Context,
					rc io.ReadCloser,
					info details.ItemInfo,
					extInfo *details.ExtensionInfo,
				) (CorsoItemExtension, error) {
					return nil, clues.New("creating extension")
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
			factories: []CorsoItemExtensionFactory{
				NewMockExtension,
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
			ith := &ItemExtensionHandler{}

			extRc, extInfo, err := ith.AddItemExtensions(
				ctx,
				test.rc,
				testItemInfo,
				test.factories)
			require.True(t, test.validateOutputs(extRc, extInfo, err))
		})
	}
}

// TODO: tests for loggerExtension
// TODO: Tests to verify RC wrapper ordering by AddItemExtensioner
