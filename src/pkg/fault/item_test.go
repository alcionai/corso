package fault

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ItemUnitSuite struct {
	tester.Suite
}

func TestItemUnitSuite(t *testing.T) {
	suite.Run(t, &ItemUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemUnitSuite) TestItem_Error() {
	var (
		t = suite.T()
		i *Item
	)

	assert.Contains(t, i.Error(), "nil")

	i = &Item{}
	assert.Contains(t, i.Error(), "unknown type")

	i = &Item{Type: FileType}
	assert.Contains(t, i.Error(), FileType)
}

func (suite *ItemUnitSuite) TestContainerErr() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := ContainerErr(errors.New("foo"), "id", "name", addtl)

	expect := Item{
		ID:         "id",
		Name:       "name",
		Type:       ContainerType,
		Cause:      "foo",
		Additional: addtl,
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestFileErr() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := FileErr(errors.New("foo"), "id", "name", addtl)

	expect := Item{
		ID:         "id",
		Name:       "name",
		Type:       FileType,
		Cause:      "foo",
		Additional: addtl,
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestOwnerErr() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := OwnerErr(errors.New("foo"), "id", "name", addtl)

	expect := Item{
		ID:         "id",
		Name:       "name",
		Type:       ResourceOwnerType,
		Cause:      "foo",
		Additional: addtl,
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestItemType_Printable() {
	table := []struct {
		t      itemType
		expect string
	}{
		{
			t:      FileType,
			expect: "File",
		},
		{
			t:      ContainerType,
			expect: "Container",
		},
		{
			t:      ResourceOwnerType,
			expect: "Resource Owner",
		},
		{
			t:      itemType("foo"),
			expect: "Unknown",
		},
	}
	for _, test := range table {
		suite.Run(string(test.t), func() {
			assert.Equal(suite.T(), test.expect, test.t.Printable())
		})
	}
}

func (suite *ItemUnitSuite) TestItem_HeadersValues() {
	var (
		err   = assert.AnError
		cause = err.Error()
		addtl = map[string]any{
			AddtlContainerID:   "cid",
			AddtlContainerName: "cname",
		}
	)

	table := []struct {
		name   string
		item   *Item
		expect []string
	}{
		{
			name:   "file",
			item:   FileErr(assert.AnError, "id", "name", addtl),
			expect: []string{"Error", FileType.Printable(), "name", "cname", cause},
		},
		{
			name:   "container",
			item:   ContainerErr(assert.AnError, "id", "name", addtl),
			expect: []string{"Error", ContainerType.Printable(), "name", "cname", cause},
		},
		{
			name:   "owner",
			item:   OwnerErr(assert.AnError, "id", "name", nil),
			expect: []string{"Error", ResourceOwnerType.Printable(), "name", "", cause},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, []string{"Action", "Type", "Name", "Container", "Cause"}, test.item.Headers())
			assert.Equal(t, test.expect, test.item.Values())
		})
	}
}

func (suite *ItemUnitSuite) TestSkipped_String() {
	var (
		t = suite.T()
		i *Skipped
	)

	assert.Contains(t, i.String(), "nil")

	i = &Skipped{Item{}}
	assert.Contains(t, i.String(), "unknown type")

	i = &Skipped{Item{Type: FileType}}
	assert.Contains(t, i.Item.Error(), FileType)
}

func (suite *ItemUnitSuite) TestContainerSkip() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := ContainerSkip(SkipMalware, "id", "name", addtl)

	expect := Item{
		ID:         "id",
		Name:       "name",
		Type:       ContainerType,
		Cause:      string(SkipMalware),
		Additional: addtl,
	}

	assert.Equal(t, Skipped{expect}, *i)
}

func (suite *ItemUnitSuite) TestFileSkip() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := FileSkip(SkipMalware, "id", "name", addtl)

	expect := Item{
		ID:         "id",
		Name:       "name",
		Type:       FileType,
		Cause:      string(SkipMalware),
		Additional: addtl,
	}

	assert.Equal(t, Skipped{expect}, *i)
}

func (suite *ItemUnitSuite) TestOwnerSkip() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	i := OwnerSkip(SkipMalware, "id", "name", addtl)

	expect := Item{
		ID:         "id",
		Name:       "name",
		Type:       ResourceOwnerType,
		Cause:      string(SkipMalware),
		Additional: addtl,
	}

	assert.Equal(t, Skipped{expect}, *i)
}

func (suite *ItemUnitSuite) TestSkipped_HeadersValues() {
	addtl := map[string]any{
		AddtlContainerID:   "cid",
		AddtlContainerName: "cname",
	}

	table := []struct {
		name   string
		skip   *Skipped
		expect []string
	}{
		{
			name:   "file",
			skip:   FileSkip(SkipMalware, "id", "name", addtl),
			expect: []string{"Skip", FileType.Printable(), "name", "cname", string(SkipMalware)},
		},
		{
			name:   "container",
			skip:   ContainerSkip(SkipMalware, "id", "name", addtl),
			expect: []string{"Skip", ContainerType.Printable(), "name", "cname", string(SkipMalware)},
		},
		{
			name:   "owner",
			skip:   OwnerSkip(SkipMalware, "id", "name", nil),
			expect: []string{"Skip", ResourceOwnerType.Printable(), "name", "", string(SkipMalware)},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, []string{"Action", "Type", "Name", "Container", "Cause"}, test.skip.Headers())
			assert.Equal(t, test.expect, test.skip.Values())
		})
	}
}
