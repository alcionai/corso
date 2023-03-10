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

func (suite *ItemUnitSuite) TestSkipped_String() {
	var (
		t = suite.T()
		i *Skipped
	)

	assert.Contains(t, i.String(), "nil")

	i = &Skipped{Item{}}
	assert.Contains(t, i.String(), "unknown type")

	i = &Skipped{Item{Type: FileType}}
	assert.Contains(t, i.item.Error(), FileType)
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
