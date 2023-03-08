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

	i := ContainerErr(errors.New("foo"), "id", "name", "containerID", "containerName")

	expect := Item{
		ID:            "id",
		Name:          "name",
		ContainerID:   "containerID",
		ContainerName: "containerName",
		Type:          ContainerType,
		Cause:         "foo",
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestFileErr() {
	t := suite.T()

	i := FileErr(errors.New("foo"), "id", "name", "containerID", "containerName")

	expect := Item{
		ID:            "id",
		Name:          "name",
		ContainerID:   "containerID",
		ContainerName: "containerName",
		Type:          FileType,
		Cause:         "foo",
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestOwnerErr() {
	t := suite.T()

	i := OwnerErr(errors.New("foo"), "id", "name")

	expect := Item{
		ID:    "id",
		Name:  "name",
		Type:  ResourceOwnerType,
		Cause: "foo",
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

	i := ContainerSkip(SkipMalware, "id", "name", "containerID", "containerName")

	expect := Item{
		ID:            "id",
		Name:          "name",
		ContainerID:   "containerID",
		ContainerName: "containerName",
		Type:          ContainerType,
		Cause:         string(SkipMalware),
	}

	assert.Equal(t, Skipped{expect}, *i)
}

func (suite *ItemUnitSuite) TestFileSkip() {
	t := suite.T()

	i := FileSkip(SkipMalware, "id", "name", "containerID", "containerName")

	expect := Item{
		ID:            "id",
		Name:          "name",
		ContainerID:   "containerID",
		ContainerName: "containerName",
		Type:          FileType,
		Cause:         string(SkipMalware),
	}

	assert.Equal(t, Skipped{expect}, *i)
}

func (suite *ItemUnitSuite) TestOwnerSkip() {
	t := suite.T()

	i := OwnerSkip(SkipMalware, "id", "name")

	expect := Item{
		ID:    "id",
		Name:  "name",
		Type:  ResourceOwnerType,
		Cause: string(SkipMalware),
	}

	assert.Equal(t, Skipped{expect}, *i)
}
