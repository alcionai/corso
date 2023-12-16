package sharepoint

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type SharepointRestoreUnitSuite struct {
	tester.Suite
}

func TestSharepointRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &SharepointRestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SharepointRestoreUnitSuite) TestSharePointHandler_ConsumeRestoreCollections_noErrorOnLists() {
	t := suite.T()
	siteID := "site-id"

	ctx, flush := tester.NewContext(t)
	defer flush()

	pr := idname.NewProvider(siteID, siteID)
	rcc := inject.RestoreConsumerConfig{
		ProtectedResource: pr,
	}
	pth, err := path.Builder{}.
		Append("lists").
		ToDataLayerPath(
			"tenant",
			siteID,
			path.SharePointService,
			path.ListsCategory,
			false)
	require.NoError(t, err, clues.ToCore(err))

	dcs := []data.RestoreCollection{
		mock.Collection{Path: pth},
	}

	sh := NewSharePointHandler(control.DefaultOptions(), api.Client{}, nil)

	_, _, err = sh.ConsumeRestoreCollections(
		ctx,
		rcc,
		dcs,
		fault.New(false),
		nil)
	require.NoError(t, err, "Sharepoint lists restore")
}
