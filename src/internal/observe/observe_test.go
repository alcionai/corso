package observe_test

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

	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/tester"
)

type ObserveProgressUnitSuite struct {
	suite.Suite
}

func TestObserveProgressUnitSuite(t *testing.T) {
	suite.Run(t, new(ObserveProgressUnitSuite))
}

var (
	tst        = observe.Safe("test")
	testcat    = observe.Safe("testcat")
	testertons = observe.Safe("testertons")
)

func (suite *ObserveProgressUnitSuite) TestItemProgress() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	recorder := strings.Builder{}
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		observe.Complete()
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	from := make([]byte, 100)
	prog, closer := observe.ItemProgress(
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

		assert.NoError(t, err)
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
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		observe.Complete()
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	progCh, closer := observe.CollectionProgress(ctx, "test", testcat, testertons)
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
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		observe.Complete()
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	progCh, closer := observe.CollectionProgress(ctx, "test", testcat, testertons)
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
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	observe.Message(ctx, observe.Safe(message))
	observe.Complete()
	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithCompletion() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	ch, closer := observe.MessageWithCompletion(ctx, observe.Safe(message))

	// Trigger completion
	ch <- struct{}{}

	// Run the closer - this should complete because the bar was compelted above
	closer()

	observe.Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), "done")
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithChannelClosed() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	ch, closer := observe.MessageWithCompletion(ctx, observe.Safe(message))

	// Close channel without completing
	close(ch)

	// Run the closer - this should complete because the channel was closed above
	closer()

	observe.Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), "done")
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithContextCancelled() {
	ctx, flush := tester.NewContext()
	defer flush()

	ctx, cancel := context.WithCancel(ctx)

	recorder := strings.Builder{}
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	message := "Test Message"

	_, closer := observe.MessageWithCompletion(ctx, observe.Safe(message))

	// cancel context
	cancel()

	// Run the closer - this should complete because the context was closed above
	closer()

	observe.Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithCount() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	header := "Header"
	message := "Test Message"
	count := 3

	ch, closer := observe.ProgressWithCount(ctx, header, observe.Safe(message), int64(count))

	for i := 0; i < count; i++ {
		ch <- struct{}{}
	}

	// Run the closer - this should complete because the context was closed above
	closer()

	observe.Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), fmt.Sprintf("%d/%d", count, count))
}

func (suite *ObserveProgressUnitSuite) TestObserveProgressWithCountChannelClosed() {
	ctx, flush := tester.NewContext()
	defer flush()

	recorder := strings.Builder{}
	observe.SeedWriter(ctx, &recorder, nil)

	defer func() {
		// don't cross-contaminate other tests.
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, nil)
	}()

	header := "Header"
	message := "Test Message"
	count := 3

	ch, closer := observe.ProgressWithCount(ctx, header, observe.Safe(message), int64(count))

	close(ch)

	// Run the closer - this should complete because the context was closed above
	closer()

	observe.Complete()

	require.NotEmpty(suite.T(), recorder.String())
	require.Contains(suite.T(), recorder.String(), message)
	require.Contains(suite.T(), recorder.String(), fmt.Sprintf("%d/%d", 0, count))
}
