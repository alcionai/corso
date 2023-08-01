package observe

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ObserveProgressUnitSuite struct {
	tester.Suite
}

func TestObserveProgressUnitSuite(t *testing.T) {
	suite.Run(t, &ObserveProgressUnitSuite{
		Suite: tester.NewUnitSuite(t),
	})
}

var (
	tst        = "test"
	testcat    = "testcat"
	testertons = "testertons"
)

func (suite *ObserveProgressUnitSuite) TestObserve_ItemProgress() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	from := make([]byte, 100)
	prog, abort := ItemProgress(
		ctx,
		io.NopCloser(bytes.NewReader(from)),
		"folder",
		tst,
		100)
	require.NotNil(t, prog)
	require.NotNil(t, abort)

	var i int

	for {
		to := make([]byte, 25)
		n, err := prog.Read(to)

		if errors.Is(err, io.EOF) {
			break
		}

		assert.NoError(t, err, clues.ToCore(err))
		assert.Equal(t, 25, n)
		i++
	}

	// mpb doesn't transmit any written values to the output writer until
	// bar completion.  Since we clean up after the bars, the recorder
	// traces nothing.
	// recorded := recorder.String()
	// assert.Contains(t, recorded, "25%")
	// assert.Contains(t, recorded, "50%")
	// assert.Contains(t, recorded, "75%")
	assert.Equal(t, 4, i)
}

func (suite *ObserveProgressUnitSuite) TestCollectionProgress_unblockOnCtxCancel() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	ctx, cancel := context.WithCancel(ctx)

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	progCh := CollectionProgress(ctx, testcat, testertons)
	require.NotNil(t, progCh)

	defer close(progCh)

	for i := 0; i < 50; i++ {
		progCh <- struct{}{}
	}

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
}

func (suite *ObserveProgressUnitSuite) TestCollectionProgress_unblockOnChannelClose() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	progCh := CollectionProgress(ctx, testcat, testertons)
	require.NotNil(t, progCh)

	for i := 0; i < 50; i++ {
		progCh <- struct{}{}
	}

	go func() {
		time.Sleep(1 * time.Second)
		close(progCh)
	}()
}

func (suite *ObserveProgressUnitSuite) TestObserve_message() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	message := uuid.NewString()[:8]

	Message(ctx, message)
	Flush(ctx)
	assert.NotEmpty(t, recorder)
	assert.Contains(t, recorder.String(), message)
}

func (suite *ObserveProgressUnitSuite) TestObserve_progressWithChannelClosed() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	message := uuid.NewString()[:8]

	ch := MessageWithCompletion(ctx, message)

	// Close channel without completing
	close(ch)

	Flush(ctx)

	assert.NotEmpty(t, recorder.String())
	assert.Contains(t, recorder.String(), message)
	assert.Contains(t, recorder.String(), "done")
}

func (suite *ObserveProgressUnitSuite) TestObserve_progressWithContextCancelled() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	ctx, cancel := context.WithCancel(ctx)

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	message := uuid.NewString()[:8]

	_ = MessageWithCompletion(ctx, message)

	// cancel context
	cancel()

	Flush(ctx)

	require.NotEmpty(t, recorder.String())
	require.Contains(t, recorder.String(), message)
}

func (suite *ObserveProgressUnitSuite) TestObserve_progressWithCount() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	header := "Header"
	message := uuid.NewString()[:8]
	count := 3

	ch := ProgressWithCount(ctx, header, message, int64(count))

	for i := 0; i < count; i++ {
		ch <- struct{}{}
	}

	close(ch)

	Flush(ctx)

	assert.NotEmpty(t, recorder.String())
	assert.Contains(t, recorder.String(), message)
	assert.Contains(t, recorder.String(), fmt.Sprintf("%d/%d", count, count))
}

func (suite *ObserveProgressUnitSuite) TestObserve_progressWithCountChannelClosed() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	recorder := strings.Builder{}
	ctx = SeedObserver(ctx, &recorder, config{})

	header := "Header"
	message := uuid.NewString()[:8]
	count := 3

	ch := ProgressWithCount(ctx, header, message, int64(count))

	close(ch)

	Flush(ctx)

	assert.NotEmpty(t, recorder.String())
	assert.Contains(t, recorder.String(), message)
	assert.Contains(t, recorder.String(), fmt.Sprintf("%d/%d", 0, count))
}

func (suite *ObserveProgressUnitSuite) TestListen() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		ch    = make(chan struct{})
		end   bool
		onEnd = func() { end = true }
		inc   bool
		onInc = func() { inc = true }
	)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- struct{}{}

		time.Sleep(500 * time.Millisecond)
		close(ch)
	}()

	// regular channel close
	listen(ctx, ch, onEnd, onInc)
	assert.True(t, end)
	assert.True(t, inc)
}

func (suite *ObserveProgressUnitSuite) TestListen_close() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		ch    = make(chan struct{})
		end   bool
		onEnd = func() { end = true }
		inc   bool
		onInc = func() { inc = true }
	)

	go func() {
		time.Sleep(500 * time.Millisecond)
		close(ch)
	}()

	// regular channel close
	listen(ctx, ch, onEnd, onInc)
	assert.True(t, end)
	assert.False(t, inc)
}

func (suite *ObserveProgressUnitSuite) TestListen_cancel() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	ctx, cancelFn := context.WithCancel(ctx)

	var (
		ch    = make(chan struct{})
		end   bool
		onEnd = func() { end = true }
		inc   bool
		onInc = func() { inc = true }
	)

	go func() {
		time.Sleep(500 * time.Millisecond)
		cancelFn()
	}()

	// regular channel close
	listen(ctx, ch, onEnd, onInc)
	assert.True(t, end)
	assert.False(t, inc)
}
