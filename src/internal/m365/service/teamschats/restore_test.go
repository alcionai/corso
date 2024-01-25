package teamschats

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type RestoreUnitSuite struct {
	tester.Suite
}

func TestRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreUnitSuite) TestConsumeRestoreCollections_noErrorOnChats() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	rcc := inject.RestoreConsumerConfig{}
	pth, err := path.BuildPrefix(
		"t",
		"pr",
		path.TeamsChatsService,
		path.ChatsCategory)
	require.NoError(t, err, clues.ToCore(err))

	dcs := []data.RestoreCollection{
		mock.Collection{Path: pth},
	}

	_, _, err = NewTeamsChatsHandler(api.Client{}, nil).
		ConsumeRestoreCollections(
			ctx,
			rcc,
			dcs,
			fault.New(false),
			nil)
	assert.NoError(t, err, "Chats restore")
}
