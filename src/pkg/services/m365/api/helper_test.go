package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

type intgTesterSetup struct {
	ac                    api.Client
	gockAC                api.Client
	userID                string
	userDriveID           string
	userDriveRootFolderID string
	userDriveTestFolderID string // root:/test
	siteID                string
	siteDriveID           string
	siteDriveRootFolderID string
	siteDriveTestFolderID string // root:/test
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	a := tester.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	its.gockAC, err = mock.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	// user drive

	its.userID = tester.M365UserID(t)

	userDrive, err := its.ac.Users().GetDefaultDrive(ctx, its.userID)
	require.NoError(t, err, clues.ToCore(err))

	its.userDriveID = ptr.Val(userDrive.GetId())

	userDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.userDriveID)
	require.NoError(t, err, clues.ToCore(err))

	its.userDriveRootFolderID = ptr.Val(userDriveRootFolder.GetId())

	its.siteID = tester.M365SiteID(t)

	// site

	siteDrive, err := its.ac.Sites().GetDefaultDrive(ctx, its.siteID)
	require.NoError(t, err, clues.ToCore(err))

	its.siteDriveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.siteDriveID)
	require.NoError(t, err, clues.ToCore(err))

	its.siteDriveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	return its
}
