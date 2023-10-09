package fault

import (
	"testing"

	"github.com/alcionai/clues"
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
	i := ContainerErr(clues.New("foo"), "ns", "id", "name", addtl)

	expect := Item{
		Namespace:  "ns",
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
	i := FileErr(clues.New("foo"), "ns", "id", "name", addtl)

	expect := Item{
		Namespace:  "ns",
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
	i := OwnerErr(clues.New("foo"), "ns", "id", "name", addtl)

	expect := Item{
		Namespace:  "ns",
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
			item:   FileErr(assert.AnError, "ns", "id", "name", addtl),
			expect: []string{"Error", FileType.Printable(), "name", "cname", cause},
		},
		{
			name:   "container",
			item:   ContainerErr(assert.AnError, "ns", "id", "name", addtl),
			expect: []string{"Error", ContainerType.Printable(), "name", "cname", cause},
		},
		{
			name:   "owner",
			item:   OwnerErr(assert.AnError, "ns", "id", "name", nil),
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
	i := ContainerSkip(SkipMalware, "ns", "id", "name", addtl)

	expect := Item{
		Namespace:  "ns",
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
	i := FileSkip(SkipMalware, "ns", "id", "name", addtl)

	expect := Item{
		Namespace:  "ns",
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
	i := OwnerSkip(SkipMalware, "ns", "id", "name", addtl)

	expect := Item{
		Namespace:  "ns",
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
			skip:   FileSkip(SkipMalware, "ns", "id", "name", addtl),
			expect: []string{"Skip", FileType.Printable(), "name", "cname", string(SkipMalware)},
		},
		{
			name:   "container",
			skip:   ContainerSkip(SkipMalware, "ns", "id", "name", addtl),
			expect: []string{"Skip", ContainerType.Printable(), "name", "cname", string(SkipMalware)},
		},
		{
			name:   "owner",
			skip:   OwnerSkip(SkipMalware, "ns", "id", "name", nil),
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

func (suite *ItemUnitSuite) TestAlert_String() {
	var (
		t = suite.T()
		a Alert
	)

	assert.Contains(t, a.String(), "Alert: <nil>")

	a = Alert{
		Item:    Item{},
		Message: "",
	}
	assert.Contains(t, a.String(), "Alert: <nil>")

	a = Alert{
		Item: Item{
			ID: "item_id",
		},
		Message: "msg",
	}
	assert.NotContains(t, a.String(), "item_id")
	assert.Contains(t, a.String(), "Alert: msg")
}

func (suite *ItemUnitSuite) TestNewAlert() {
	t := suite.T()
	addtl := map[string]any{"foo": "bar"}
	a := NewAlert("message-to-show", "ns", "item_id", "item_name", addtl)

	expect := Alert{
		Item: Item{
			Namespace:  "ns",
			ID:         "item_id",
			Name:       "item_name",
			Additional: addtl,
		},
		Message: "message-to-show",
	}

	assert.Equal(t, expect, *a)
}

func (suite *ItemUnitSuite) TestAlert_HeadersValues() {
	addtl := map[string]any{
		AddtlContainerID:   "cid",
		AddtlContainerName: "cname",
	}

	table := []struct {
		name   string
		alert  *Alert
		expect []string
	}{
		{
			name:   "new alert",
			alert:  NewAlert("message-to-show", "ns", "id", "name", addtl),
			expect: []string{"Alert", "message-to-show", "cname", "name", "id"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, []string{"Action", "Message", "Container", "Name", "ID"}, test.alert.Headers())
			assert.Equal(t, test.expect, test.alert.Values())
		})
	}
}
