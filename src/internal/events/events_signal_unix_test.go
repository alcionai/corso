package events

import (
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/armon/go-metrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventsSignalUnitSuite struct {
	tester.Suite
}

func TestEventsSignalUnitSuite(t *testing.T) {
	suite.Run(t, &EventsSignalUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *EventsSignalUnitSuite) TestSignalDump() {
	ctx, flush := tester.NewContext()
	defer flush()
	var (
		t = suite.T()
	)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, metrics.DefaultSignal)

	go func() {
		signalDump(ctx)
	}()

	select {
	case sig := <-sigCh:
		assert.Equal(t, metrics.DefaultSignal, sig, "received wrong signal")

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout waiting for signal")
	}

	signal.Stop(sigCh)
}
