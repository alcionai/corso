package restore

import (
	"context"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func CheckSharePointRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	drive, err := ac.Sites().GetDefaultDrive(ctx, envs.SiteID)
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	checkDriveRestoration(
		ctx,
		ac,
		path.SharePointService,
		envs.FolderName,
		ptr.Val(drive.GetId()),
		ptr.Val(drive.GetName()),
		envs.DataFolder,
		envs.StartTime,
		true)
}
