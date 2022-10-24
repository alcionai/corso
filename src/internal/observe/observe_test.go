package observe_test

import (
	"bytes"
	"context"
	"errors"
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

func (suite *ObserveProgressUnitSuite) TestItemProgress() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	recorder := strings.Builder{}
	observe.SeedWriter(ctx, &recorder, false)

	defer func() {
		// don't cross-contaminate other tests.
		observe.Complete()
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, false)
	}()

	from := make([]byte, 100)
	prog, closer := observe.ItemProgress(
		io.NopCloser(bytes.NewReader(from)),
		"test",
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
	observe.SeedWriter(ctx, &recorder, false)

	defer func() {
		// don't cross-contaminate other tests.
		observe.Complete()
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, false)
	}()

	progCh, closer := observe.CollectionProgress("test", "testcat", "testertons")
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
	observe.SeedWriter(ctx, &recorder, false)

	defer func() {
		// don't cross-contaminate other tests.
		observe.Complete()
		//nolint:forbidigo
		observe.SeedWriter(context.Background(), nil, false)
	}()

	progCh, closer := observe.CollectionProgress("test", "testcat", "testertons")
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
