package restore

import (
	"context"

	"github.com/alcionai/canario/src/cmd/sanity_test/common"
	"github.com/alcionai/canario/src/cmd/sanity_test/driveish"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

func CheckOneDriveRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	drive, err := ac.Users().GetDefaultDrive(ctx, envs.UserID)
	if err != nil {
		common.Fatal(ctx, "getting user's default drive:", err)
	}

	driveish.CheckRestoration(
		ctx,
		ac,
		drive,
		envs,
		driveish.ComparatorEqualPerms)
}
