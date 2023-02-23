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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/aw"
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
	tst        = Safe("test")
	testcat    = Safe("testcat")
	testertons = Safe("testertons")
)

func (suite *ObserveProgressUnitSuite) TestItemProgress() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		Complete()
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	from := make([]byte, 100)
	prog, closer := ItemProgress(
		ctx,
		io.NopCloser(bytes.NewReader(from)),
		"folder",
		tst,
		100)
	require.NotNil(t, prog)
	require.NotNil(t, closer)

	defer closer()

	var i int

	for {
		to := make([]byte, 25)
		n, err := prog.Read(to)

		if errors.Is(err, io.EOF) {
			break
		}

		aw.NoErr(t, err)
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
	ctx, flush := tester.NewContext()
	defer flush()

	ctx, cancel := context.WithCancel(ctx)

	t := suite.T()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		Complete()
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	progCh, closer := CollectionProgress(ctx, "test", testcat, testertons)
	require.NotNil(t, progCh)
	require.NotNil(t, closer)

	defer close(progCh)

	for i := 0; i < 50; i++ {
		progCh <- struct{}{}
	}

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	// blocks, but should resolve due to the ctx cancel
	closer()
}

func (suite *ObserveProgressUnitSuite) TestCollectionProgress_unblockOnChannelClose() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		Complete()
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	progCh, closer := CollectionProgress(ctx, "test", testcat, testertons)
	require.NotNil(t, progCh)
	require.NotNil(t, closer)

	for i := 0; i < 50; i++ {
		progCh <- struct{}{}
	}

	go func() {
		time.Sleep(1 * time.Second)
		close(progCh)
	}()

	// blocks, but should resolve due to the cancel
	closer()
}

func (suite *ObserveProgressUnitSuite) TestObserveProgress() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	Message(ctx, Safe(message))
	Complete()
	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithCompletion() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	ch, closer := MessageWithCompletion(ctx, Safe(message))

	// Trigger completion
	ch <- struct{}{}

	// Run the closer - this should complete because the bar was compelted above
	closer()

	Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), "done")
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithChannelClosed() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	ch, closer := MessageWithCompletion(ctx, Safe(message))

	// Close channel without completing
	close(ch)

	// Run the closer - this should complete because the channel was closed above
	closer()

	Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), "done")
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithContextCancelled() {
	ctx, flush := tester.NewContext()
	defer flush()

	ctx, cancel := context.WithCancel(ctx)

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	_, closer := MessageWithCompletion(ctx, Safe(message))

	// cancel context
	cancel()

	// Run the closer - this should complete because the context was closed above
	closer()

	Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithCount() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	header := "Header"
	message := "Test Message"
	count := 3

	ch, closer := ProgressWithCount(ctx, header, Safe(message), int64(count))

	for i := 0; i < count; i++ {
		ch <- struct{}{}
	}

	// Run the closer - this should complete because the context was closed above
	closer()

	Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), fmt.Sprintf("%d/%d", count, count))
}

func (suite *ObserveProgressUnitSuite) TestrogressWithCountChannelClosed() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		SeedWriter(context.Background(), nil, nil)
	}()

	header := "Header"
	message := "Test Message"
	count := 3

	ch, closer := ProgressWithCount(ctx, header, Safe(message), int64(count))

	close(ch)

	// Run the closer - this should complete because the context was closed above
	closer()

	Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), fmt.Sprintf("%d/%d", 0, count))
}

func (suite *ObserveProgressUnitSuite) TestListen() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t     = suite.T()
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
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t     = suite.T()
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
	ctx, flush := tester.NewContext()
	defer flush()

	ctx, cancelFn := context.WithCancel(ctx)

	var (
		t     = suite.T()
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
