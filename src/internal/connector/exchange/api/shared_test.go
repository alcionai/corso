package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type SharedAPIUnitSuite struct {
	tester.Suite
}

func TestSharedAPIUnitSuite(t *testing.T) {
	suite.Run(t, &SharedAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type dummyPager struct{}

func (p *dummyPager) getNextPage(ctx context.Context) ([]getIDAndAddtler, bool, string, error) {
	return []getIDAndAddtler{}, true, "delta-url", nil
}

func (p *dummyPager) reset(nonDelta bool) {}

func (suite *SharedAPIUnitSuite) TestGetItemsAddedAndRemovedFromContainer_ensureDeltaUrl() {
	ctx, flush := tester.NewContext()
	defer flush()

	_, _, deltaURL, _ := getItemsAddedAndRemovedFromContainer(ctx, &dummyPager{})
	require.Equal(suite.T(), deltaURL, "delta-url", "get delta url")
}
