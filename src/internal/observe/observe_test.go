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

	from := make([]byte, 100)
	prog := observe.ItemProgress(
		io.NopCloser(bytes.NewReader(from)),
		"test",
		100)
	require.NotNil(t, prog)

	for {
		to := make([]byte, 25)
		n, err := prog.Read(to)

		if errors.Is(err, io.EOF) {
			break
		}

		assert.NoError(t, err)
		assert.Less(t, 0, n)
	}

	recorded := recorder.String()
	assert.Contains(t, recorded, "25%")
	assert.Contains(t, recorded, "50%")
	assert.Contains(t, recorded, "75%")
}
