package data_test

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

type errReader struct {
	io.ReadCloser
	readCount int
	errAfter  int
	err       error
}

func (er *errReader) Read(p []byte) (int, error) {
	if er.err != nil && er.readCount == er.errAfter {
		return 0, er.err
	}

	toRead := len(p)
	if er.readCount+toRead > er.errAfter {
		toRead = er.errAfter - er.readCount
	}

	n, err := er.ReadCloser.Read(p[:toRead])
	er.readCount += n

	return n, err
}

type ItemUnitSuite struct {
	tester.Suite
}

func TestItemUnitSuite(t *testing.T) {
	suite.Run(t, &ItemUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemUnitSuite) TestUnindexedPrefetchedItem() {
	prefetch, err := data.NewPrefetchedItem(
		io.NopCloser(bytes.NewReader([]byte{})),
		"foo",
		time.Time{})
	require.NoError(suite.T(), err, clues.ToCore(err))

	var item data.Item = prefetch

	_, ok := item.(data.ItemInfo)
	assert.False(suite.T(), ok, "unindexedPrefetchedItem implements Info()")
}

func (suite *ItemUnitSuite) TestUnindexedLazyItem() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	lazy := data.NewLazyItem(
		ctx,
		nil,
		"foo",
		time.Time{},
		fault.New(true))

	var item data.Item = lazy

	_, ok := item.(data.ItemInfo)
	assert.False(t, ok, "unindexedLazyItem implements Info()")
}

func (suite *ItemUnitSuite) TestDeletedItem() {
	var (
		t = suite.T()

		id   = "foo"
		item = data.NewDeletedItem(id)
	)

	assert.Equal(t, id, item.ID(), "ID")
	assert.True(t, item.Deleted(), "deleted")
}

func (suite *ItemUnitSuite) TestPrefetchedItem() {
	var (
		id  = "foo"
		now = time.Now()

		baseData = []byte("hello world")
	)

	table := []struct {
		name   string
		reader io.ReadCloser
		info   details.ItemInfo

		readErr    require.ErrorAssertionFunc
		expectData []byte
	}{
		{
			name:       "EmptyReader",
			reader:     io.NopCloser(bytes.NewReader([]byte{})),
			info:       details.ItemInfo{Exchange: &details.ExchangeInfo{Modified: now}},
			readErr:    require.NoError,
			expectData: []byte{},
		},
		{
			name:       "ReaderWithData",
			reader:     io.NopCloser(bytes.NewReader(baseData)),
			info:       details.ItemInfo{Exchange: &details.ExchangeInfo{Modified: now}},
			readErr:    require.NoError,
			expectData: baseData,
		},
		{
			name:       "ReaderWithData DifferentService",
			reader:     io.NopCloser(bytes.NewReader(baseData)),
			info:       details.ItemInfo{OneDrive: &details.OneDriveInfo{Modified: now}},
			readErr:    require.NoError,
			expectData: baseData,
		},
		{
			name: "ReaderWithData ReadError",
			reader: &errReader{
				ReadCloser: io.NopCloser(bytes.NewReader(baseData)),
				errAfter:   5,
				err:        assert.AnError,
			},
			info:       details.ItemInfo{Exchange: &details.ExchangeInfo{Modified: now}},
			readErr:    require.Error,
			expectData: baseData[:5],
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			item, err := data.NewPrefetchedItemWithInfo(test.reader, id, test.info)
			require.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, id, item.ID(), "ID")
			assert.False(t, item.Deleted(), "deleted")
			assert.Equal(
				t,
				test.info.Modified(),
				item.ModTime(),
				"mod time")

			r, err := readers.NewVersionedRestoreReader(item.ToReader())
			require.NoError(t, err, "version error: %v", clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, readers.DefaultSerializationVersion, r.Format().Version)
			assert.False(t, r.Format().DelInFlight)

			readData, err := io.ReadAll(r)
			test.readErr(t, err, "read error: %v", clues.ToCore(err))
			assert.Equal(t, test.expectData, readData, "read data")
		})
	}
}

type mockItemDataGetter struct {
	getCalled bool

	reader      io.ReadCloser
	info        *details.ItemInfo
	delInFlight bool
	err         error
}

func (mid *mockItemDataGetter) check(t *testing.T, expectCalled bool) {
	assert.Equal(t, expectCalled, mid.getCalled, "GetData() called")
}

func (mid *mockItemDataGetter) GetData(
	ctx context.Context,
	errs *fault.Bus,
) (io.ReadCloser, *details.ItemInfo, bool, error) {
	mid.getCalled = true

	if mid.err != nil {
		errs.AddRecoverable(ctx, mid.err)
	}

	return mid.reader, mid.info, mid.delInFlight, mid.err
}

func (suite *ItemUnitSuite) TestLazyItem() {
	var (
		id  = "foo"
		now = time.Now()

		baseData = []byte("hello world")
	)

	table := []struct {
		name         string
		mid          *mockItemDataGetter
		versionErr   assert.ErrorAssertionFunc
		readErr      assert.ErrorAssertionFunc
		infoErr      assert.ErrorAssertionFunc
		expectData   []byte
		expectBusErr bool
	}{
		{
			name: "EmptyReader",
			mid: &mockItemDataGetter{
				reader: io.NopCloser(bytes.NewReader([]byte{})),
				info:   &details.ItemInfo{Exchange: &details.ExchangeInfo{Modified: now}},
			},
			versionErr: assert.NoError,
			readErr:    assert.NoError,
			infoErr:    assert.NoError,
			expectData: []byte{},
		},
		{
			name: "ReaderWithData",
			mid: &mockItemDataGetter{
				reader: io.NopCloser(bytes.NewReader(baseData)),
				info:   &details.ItemInfo{Exchange: &details.ExchangeInfo{Modified: now}},
			},
			versionErr: assert.NoError,
			readErr:    assert.NoError,
			infoErr:    assert.NoError,
			expectData: baseData,
		},
		{
			name: "ReaderWithData",
			mid: &mockItemDataGetter{
				reader: io.NopCloser(bytes.NewReader(baseData)),
				info:   &details.ItemInfo{OneDrive: &details.OneDriveInfo{Modified: now}},
			},
			versionErr: assert.NoError,
			readErr:    assert.NoError,
			infoErr:    assert.NoError,
			expectData: baseData,
		},
		{
			name: "ReaderWithData GetDataError",
			mid: &mockItemDataGetter{
				err: assert.AnError,
			},
			versionErr:   assert.Error,
			readErr:      assert.Error,
			infoErr:      assert.Error,
			expectData:   []byte{},
			expectBusErr: true,
		},
		{
			name: "ReaderWithData ReadError",
			mid: &mockItemDataGetter{
				reader: &errReader{
					ReadCloser: io.NopCloser(bytes.NewReader(baseData)),
					errAfter:   5,
					err:        assert.AnError,
				},
				info: &details.ItemInfo{OneDrive: &details.OneDriveInfo{Modified: now}},
			},
			versionErr: assert.NoError,
			readErr:    assert.Error,
			infoErr:    assert.NoError,
			expectData: baseData[:5],
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			errs := fault.New(true)

			defer test.mid.check(t, true)

			item := data.NewLazyItemWithInfo(
				ctx,
				test.mid,
				id,
				now,
				errs)

			assert.Equal(t, id, item.ID(), "ID")
			assert.False(t, item.Deleted(), "deleted")
			assert.Equal(
				t,
				now,
				item.ModTime(),
				"mod time")

			// Read data to execute lazy reader.
			r, err := readers.NewVersionedRestoreReader(item.ToReader())
			test.versionErr(t, err, "version error: %v", clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, readers.DefaultSerializationVersion, r.Format().Version)
			assert.False(t, r.Format().DelInFlight)

			readData, err := io.ReadAll(r)
			test.readErr(t, err, clues.ToCore(err), "read error")
			assert.Equal(t, test.expectData, readData, "read data")

			_, err = item.Info()
			test.infoErr(t, err, "Info(): %v", clues.ToCore(err))

			e := errs.Errors()

			if !test.expectBusErr {
				assert.Nil(t, e.Failure, "hard failure")
				assert.Empty(t, e.Recovered, "recovered")

				return
			}

			assert.NotNil(t, e.Failure, "hard failure")
		})
	}
}

func (suite *ItemUnitSuite) TestLazyItem_DeletedInFlight() {
	var (
		id  = "foo"
		now = time.Now()
	)

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	errs := fault.New(true)

	mid := &mockItemDataGetter{delInFlight: true}
	defer mid.check(t, true)

	item := data.NewLazyItemWithInfo(ctx, mid, id, now, errs)

	assert.Equal(t, id, item.ID(), "ID")
	assert.False(t, item.Deleted(), "deleted")
	assert.Equal(
		t,
		now,
		item.ModTime(),
		"mod time")

	// Read data to execute lazy reader.
	r, err := readers.NewVersionedRestoreReader(item.ToReader())
	require.NoError(t, err, "version error: %v", clues.ToCore(err))

	assert.Equal(t, readers.DefaultSerializationVersion, r.Format().Version)
	assert.True(t, r.Format().DelInFlight)

	readData, err := io.ReadAll(r)
	require.NoError(t, err, clues.ToCore(err), "read error")
	assert.Empty(t, readData, "read data")

	_, err = item.Info()
	assert.ErrorIs(t, err, data.ErrNotFound, "Info() error")

	e := errs.Errors()

	assert.Nil(t, e.Failure, "hard failure")
	assert.Empty(t, e.Recovered, "recovered")
}

func (suite *ItemUnitSuite) TestLazyItem_InfoBeforeReadErrors() {
	var (
		id  = "foo"
		now = time.Now()
	)

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	errs := fault.New(true)

	mid := &mockItemDataGetter{}
	defer mid.check(t, false)

	item := data.NewLazyItemWithInfo(ctx, mid, id, now, errs)

	assert.Equal(t, id, item.ID(), "ID")
	assert.False(t, item.Deleted(), "deleted")
	assert.Equal(
		t,
		now,
		item.ModTime(),
		"mod time")

	_, err := item.Info()
	assert.Error(t, err, "Info() error")
}
