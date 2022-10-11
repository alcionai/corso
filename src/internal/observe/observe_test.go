package observe_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/observe"
)

type ObserveProgressUnitSuite struct {
	suite.Suite
}

func TestObserveProgressUnitSuite(t *testing.T) {
	suite.Run(t, new(ObserveProgressUnitSuite))
}

func (suite *ObserveProgressUnitSuite) TestDoesThings() {
	t := suite.T()

	recorder := strings.Builder{}
	observe.SeedWriter(&recorder)

	defer func() {
		// don't cross-contaminate other tests.
		observe.Complete()
		observe.SeedWriter(nil)
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
