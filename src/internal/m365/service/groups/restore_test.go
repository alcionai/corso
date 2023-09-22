package groups

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type GroupsUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupsUnitSuite) TestConsumeRestoreCollections_noErrorOnGroups() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rcc := inject.RestoreConsumerConfig{}
	pth, err := path.Builder{}.
		Append("General").
		ToDataLayerPath(
			"t",
			"g",
			path.GroupsService,
			path.ChannelMessagesCategory,
			false)
	require.NoError(t, err, clues.ToCore(err))

	dcs := []data.RestoreCollection{
		mock.Collection{Path: pth},
	}

	_, err = ConsumeRestoreCollections(
		ctx,
		rcc,
		api.Client{},
		idname.NewCache(map[string]string{}),
		idname.NewCache(map[string]string{}),
		dcs,
		nil,
		fault.New(false),
		nil)
	assert.NoError(t, err, "Groups Channels restore")
}
