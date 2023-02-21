package graph

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
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

	c := NewMetadataCollection(p, nil, nil)

	assert.Equal(t, p.String(), c.FullPath().String())
}

func (suite *MetadataCollectionUnitSuite) TestItems() {
	ctx, flush := tester.NewContext()
	defer flush()

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

	items := []MetadataItem{}

	for i := 0; i < len(itemNames); i++ {
		items = append(items, NewMetadataItem(itemNames[i], itemData[i]))
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

	c := NewMetadataCollection(
		p,
		items,
		func(c *support.ConnectorOperationStatus) {
			assert.Equal(t, len(itemNames), c.Metrics.Objects)
			assert.Equal(t, len(itemNames), c.Metrics.Successes)
		},
	)

	gotData := [][]byte{}
	gotNames := []string{}

	for s := range c.Items(ctx, fault.New(true)) {
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

func (suite *MetadataCollectionUnitSuite) TestMakeMetadataCollection() {
	tenant := "a-tenant"
	user := "a-user"

	table := []struct {
		name            string
		service         path.ServiceType
		cat             path.CategoryType
		metadata        MetadataCollectionEntry
		collectionCheck assert.ValueAssertionFunc
		errCheck        assert.ErrorAssertionFunc
	}{
		{
			name:            "EmptyTokens",
			service:         path.ExchangeService,
			cat:             path.EmailCategory,
			metadata:        NewMetadataEntry("", nil),
			collectionCheck: assert.Nil,
			errCheck:        assert.Error,
		},
		{
			name:    "Tokens",
			service: path.ExchangeService,
			cat:     path.EmailCategory,
			metadata: NewMetadataEntry(
				uuid.NewString(),
				map[string]string{
					"hello": "world",
					"hola":  "mundo",
				}),
			collectionCheck: assert.NotNil,
			errCheck:        assert.NoError,
		},
		{
			name:    "BadCategory",
			service: path.ExchangeService,
			cat:     path.FilesCategory,
			metadata: NewMetadataEntry(
				uuid.NewString(),
				map[string]string{
					"hello": "world",
					"hola":  "mundo",
				}),
			collectionCheck: assert.Nil,
			errCheck:        assert.Error,
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			col, err := MakeMetadataCollection(
				tenant,
				user,
				test.service,
				test.cat,
				[]MetadataCollectionEntry{test.metadata},
				func(*support.ConnectorOperationStatus) {})

			test.errCheck(t, err)
			if err != nil {
				return
			}

			test.collectionCheck(t, col)
			if col == nil {
				return
			}

			itemCount := 0
			for item := range col.Items(ctx, fault.New(true)) {
				assert.Equal(t, test.metadata.fileName, item.UUID())

				gotMap := map[string]string{}
				decoder := json.NewDecoder(item.ToReader())
				itemCount++

				err := decoder.Decode(&gotMap)
				if !assert.NoError(t, err) {
					continue
				}

				assert.Equal(t, test.metadata.data, gotMap)
			}

			assert.Equal(t, 1, itemCount)
		})
	}
}
