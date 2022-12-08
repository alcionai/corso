package graph_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

type MetadataCollectionUnitSuite struct {
	suite.Suite
}

func TestMetadataCollectionUnitSuite(t *testing.T) {
	suite.Run(t, new(MetadataCollectionUnitSuite))
}

func (suite *MetadataCollectionUnitSuite) TestFullPath() {
	t := suite.T()

	p, err := path.Builder{}.
		Append("foo").
		ToDataLayerExchangePathForCategory(
			"a-tenant",
			"a-user",
			path.EmailCategory,
			false,
		)
	require.NoError(t, err)

	c := graph.NewMetadataCollection(p, nil, nil)

	assert.Equal(t, p.String(), c.FullPath().String())
}

func (suite *MetadataCollectionUnitSuite) TestItems() {
	t := suite.T()

	itemNames := []string{
		"a",
		"aa",
	}
	itemData := [][]byte{
		[]byte("a"),
		[]byte("aa"),
	}

	require.Equal(
		t,
		len(itemNames),
		len(itemData),
		"Requires same number of items and data",
	)

	items := []graph.MetadataItem{}

	for i := 0; i < len(itemNames); i++ {
		items = append(items, graph.NewMetadataItem(itemNames[i], itemData[i]))
	}

	p, err := path.Builder{}.
		Append("foo").
		ToDataLayerExchangePathForCategory(
			"a-tenant",
			"a-user",
			path.EmailCategory,
			false,
		)
	require.NoError(t, err)

	c := graph.NewMetadataCollection(
		p,
		items,
		func(c *support.ConnectorOperationStatus) {
			assert.Equal(t, len(itemNames), c.ObjectCount)
			assert.Equal(t, len(itemNames), c.Successful)
		},
	)

	gotData := [][]byte{}
	gotNames := []string{}

	for s := range c.Items() {
		gotNames = append(gotNames, s.UUID())

		buf, err := io.ReadAll(s.ToReader())
		if !assert.NoError(t, err) {
			continue
		}

		gotData = append(gotData, buf)
	}

	assert.ElementsMatch(t, itemNames, gotNames)
	assert.ElementsMatch(t, itemData, gotData)
}
