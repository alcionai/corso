package restore

import (
	"context"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/cmd/sanity_test/driveish"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// only checks drive restoration.  channel messages are not
// supported for restore actions.
func CheckGroupsRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	drive, err := ac.Sites().GetDefaultDrive(ctx, envs.TeamSiteID)
	if err != nil {
		common.Fatal(ctx, "getting site's default drive:", err)
	}

	driveish.CheckRestoration(
		ctx,
		ac,
		drive,
		envs,
		// skip permissions tests
		nil)
}
