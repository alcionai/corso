package exchange

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------------

type DataCollectionsUnitSuite struct {
	suite.Suite
}

func TestDataCollectionsUnitSuite(t *testing.T) {
	suite.Run(t, new(DataCollectionsUnitSuite))
}

func (suite *DataCollectionsUnitSuite) TestParseMetadataCollections() {
	t := suite.T()
	ctx, flush := tester.NewContext()

	defer flush()

	bs, err := json.Marshal(map[string]string{"key": "token"})
	require.NoError(t, err)

	p, err := path.Builder{}.ToServiceCategoryMetadataPath(
		"t", "u",
		path.ExchangeService,
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	item := []graph.MetadataItem{graph.NewMetadataItem(graph.DeltaTokenFileName, bs)}
	mdcoll := graph.NewMetadataCollection(p, item, func(cos *support.ConnectorOperationStatus) {})
	colls := []data.Collection{mdcoll}

	_, deltas, err := ParseMetadataCollections(ctx, colls)
	require.NoError(t, err)
	assert.NotEmpty(t, deltas, "delta urls")
	assert.Equal(t, "token", deltas["key"])
}
