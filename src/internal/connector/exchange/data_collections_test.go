package exchange

import (
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
			name: "delta urls and empty previous paths",
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

			entries := []graph.MetadataCollectionEntry{}

			for _, d := range test.data {
				entries = append(
					entries,
					graph.NewMetadataEntry(d.fileName, map[string]string{"key": d.value}))
			}

			coll, err := graph.MakeMetadataCollection(
				"t", "u",
				path.ExchangeService,
				path.EmailCategory,
				entries,
				func(cos *support.ConnectorOperationStatus) {},
			)
			require.NoError(t, err)

			cdps, err := ParseMetadataCollections(ctx, []data.Collection{coll})
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
