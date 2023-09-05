package graph

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type MetadataCollectionUnitSuite struct {
	tester.Suite
}

func TestMetadataCollectionUnitSuite(t *testing.T) {
	suite.Run(t, &MetadataCollectionUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MetadataCollectionUnitSuite) TestFullPath() {
	t := suite.T()

	p, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"foo")
	require.NoError(t, err, clues.ToCore(err))

	c := NewMetadataCollection(p, nil, nil)

	assert.Equal(t, p.String(), c.FullPath().String())
}

func (suite *MetadataCollectionUnitSuite) TestItems() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

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

	p, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"foo")
	require.NoError(t, err, clues.ToCore(err))

	c := NewMetadataCollection(
		p,
		items,
		func(c *support.ControllerOperationStatus) {
			assert.Equal(t, len(itemNames), c.Metrics.Objects)
			assert.Equal(t, len(itemNames), c.Metrics.Successes)
		},
	)

	gotData := [][]byte{}
	gotNames := []string{}

	for s := range c.Items(ctx, fault.New(true)) {
		gotNames = append(gotNames, s.ID())

		buf, err := io.ReadAll(s.ToReader())
		if !assert.NoError(t, err, clues.ToCore(err)) {
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
		pathPrefixCheck assert.ErrorAssertionFunc
		errCheck        assert.ErrorAssertionFunc
	}{
		{
			name:            "EmptyTokens",
			service:         path.ExchangeService,
			cat:             path.EmailCategory,
			metadata:        NewMetadataEntry("", nil),
			collectionCheck: assert.Nil,
			pathPrefixCheck: assert.NoError,
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
			pathPrefixCheck: assert.NoError,
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
			pathPrefixCheck: assert.Error,
			errCheck:        assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			pathPrefix, err := path.Builder{}.ToServiceCategoryMetadataPath(
				tenant,
				user,
				test.service,
				test.cat,
				false)
			test.pathPrefixCheck(t, err, "path prefix")
			if err != nil {
				return
			}

			col, err := MakeMetadataCollection(
				pathPrefix,
				[]MetadataCollectionEntry{test.metadata},
				func(*support.ControllerOperationStatus) {})

			test.errCheck(t, err, clues.ToCore(err))
			if err != nil {
				return
			}

			test.collectionCheck(t, col)
			if col == nil {
				return
			}

			itemCount := 0
			for item := range col.Items(ctx, fault.New(true)) {
				assert.Equal(t, test.metadata.fileName, item.ID())

				gotMap := map[string]string{}
				decoder := json.NewDecoder(item.ToReader())
				itemCount++

				err := decoder.Decode(&gotMap)
				if !assert.NoError(t, err, clues.ToCore(err)) {
					continue
				}

				assert.Equal(t, test.metadata.data, gotMap)
			}

			assert.Equal(t, 1, itemCount)
		})
	}
}
