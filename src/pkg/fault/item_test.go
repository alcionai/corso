package fault_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
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
		i *fault.Item
	)

	assert.Contains(t, i.Error(), "nil")

	i = &fault.Item{}
	assert.Contains(t, i.Error(), "unknown kind")

	i = &fault.Item{Kind: fault.ItemKindFile}
	assert.Contains(t, i.Error(), fault.ItemKindFile)
}

func (suite *ItemUnitSuite) TestContainerItem() {
	t := suite.T()

	i := fault.ContainerItem(errors.New("foo"), "id", "name", "containerID", "containerName")

	expect := fault.Item{
		ID:            "id",
		Name:          "name",
		ContainerID:   "containerID",
		ContainerName: "containerName",
		Kind:          fault.ItemKindContainer,
		Cause:         "foo",
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestFileItem() {
	t := suite.T()

	i := fault.FileItem(errors.New("foo"), "id", "name", "containerID", "containerName")

	expect := fault.Item{
		ID:            "id",
		Name:          "name",
		ContainerID:   "containerID",
		ContainerName: "containerName",
		Kind:          fault.ItemKindFile,
		Cause:         "foo",
	}

	assert.Equal(t, expect, *i)
}

func (suite *ItemUnitSuite) TestOwnerItem() {
	t := suite.T()

	i := fault.OwnerItem(errors.New("foo"), "id", "name")

	expect := fault.Item{
		ID:    "id",
		Name:  "name",
		Kind:  fault.ItemKindResourceOwner,
		Cause: "foo",
	}

	assert.Equal(t, expect, *i)
}
