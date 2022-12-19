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
		name        string
		data        []fileValues
		expect      map[string]DeltaPath
		expectError assert.ErrorAssertionFunc
	}{
		{
			name: "delta urls only",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
			},
			expect:      map[string]DeltaPath{},
			expectError: assert.NoError,
		},
		{
			name: "multiple delta urls",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
				{graph.DeltaURLsFileName, "delta-link-2"},
			},
			expectError: assert.Error,
		},
		{
			name: "previous path only",
			data: []fileValues{
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect:      map[string]DeltaPath{},
			expectError: assert.NoError,
		},
		{
			name: "multiple previous paths",
			data: []fileValues{
				{graph.PreviousPathFileName, "prev-path"},
				{graph.PreviousPathFileName, "prev-path-2"},
			},
			expectError: assert.Error,
		},
		{
			name: "delta urls and previous paths",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					delta: "delta-link",
					path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "delta urls and empttyprevious paths",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
				{graph.PreviousPathFileName, ""},
			},
			expect:      map[string]DeltaPath{},
			expectError: assert.NoError,
		},
		{
			name: "empty delta urls and previous paths",
			data: []fileValues{
				{graph.DeltaURLsFileName, ""},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect:      map[string]DeltaPath{},
			expectError: assert.NoError,
		},
		{
			name: "delta urls with special chars",
			data: []fileValues{
				{graph.DeltaURLsFileName, "`!@#$%^&*()_[]{}/\"\\"},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					delta: "`!@#$%^&*()_[]{}/\"\\",
					path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "delta urls with escaped chars",
			data: []fileValues{
				{graph.DeltaURLsFileName, `\n\r\t\b\f\v\0\\`},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					delta: "\\n\\r\\t\\b\\f\\v\\0\\\\",
					path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "delta urls with newline char runes",
			data: []fileValues{
				// rune(92) = \, rune(110) = n.  Ensuring it's not possible to
				// error in serializing/deserializing and produce a single newline
				// character from those two runes.
				{graph.DeltaURLsFileName, string([]rune{rune(92), rune(110)})},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					delta: "\\n",
					path:  "prev-path",
				},
			},
			expectError: assert.NoError,
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

			cdps, err := ParseMetadataCollections(ctx, colls)
			test.expectError(t, err)

			emails := cdps[path.EmailCategory]

			assert.Len(t, emails, len(test.expect))

			for k, v := range emails {
				assert.Equal(t, v.delta, emails[k].delta, "delta")
				assert.Equal(t, v.path, emails[k].path, "path")
			}
		})
	}
}
