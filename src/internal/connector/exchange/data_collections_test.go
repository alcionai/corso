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
	type fileValues struct {
		fileName string
		value    string
	}

	table := []struct {
		name         string
		data         []fileValues
		expectDeltas map[string]string
	}{
		{
			name: "delta urls",
			data: []fileValues{
				{graph.DeltaTokenFileName, "delta-link"},
			},
			expectDeltas: map[string]string{
				"key": "delta-link",
			},
		},
		{
			name: "delta urls with special chars",
			data: []fileValues{
				{graph.DeltaTokenFileName, "`!@#$%^&*()_[]{}/\"\\"},
			},
			expectDeltas: map[string]string{
				"key": "`!@#$%^&*()_[]{}/\"\\",
			},
		},
		{
			name: "delta urls with escaped chars",
			data: []fileValues{
				{graph.DeltaTokenFileName, `\n\r\t\b\f\v\0\\`},
			},
			expectDeltas: map[string]string{
				"key": "\\n\\r\\t\\b\\f\\v\\0\\\\",
			},
		},
		{
			name: "delta urls with newline char runes",
			data: []fileValues{
				// rune(92) = \, rune(110) = n.  If a parsing error were possible
				// by serializing/deserializing those two runes and producing a
				// single newline character, this would produce it.
				{graph.DeltaTokenFileName, string([]rune{rune(92), rune(110)})},
			},
			expectDeltas: map[string]string{
				"key": "\\n",
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			colls := []data.Collection{}

			for _, d := range test.data {
				bs, err := json.Marshal(map[string]string{"key": d.value})
				require.NoError(t, err)

				p, err := path.Builder{}.ToServiceCategoryMetadataPath(
					"t", "u",
					path.ExchangeService,
					path.EmailCategory,
					false,
				)
				require.NoError(t, err)

				item := []graph.MetadataItem{graph.NewMetadataItem(d.fileName, bs)}
				coll := graph.NewMetadataCollection(p, item, func(cos *support.ConnectorOperationStatus) {})
				colls = append(colls, coll)
			}

			_, deltas, err := ParseMetadataCollections(ctx, colls)
			require.NoError(t, err)
			assert.NotEmpty(t, deltas, "deltas")
			for k, v := range test.expectDeltas {
				assert.Equal(t, v, deltas[k], "deltas elements")
			}
		})
	}
}
