package export

import (
	"context"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/cmd/sanity_test/driveish"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func CheckSharePointExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	drive, err := ac.Sites().GetDefaultDrive(ctx, envs.SiteID)
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	driveish.CheckExport(
		ctx,
		ac,
		drive,
		envs)
}
