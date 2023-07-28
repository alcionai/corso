package restore

import (
	"context"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/path"
)

func CheckSharePointRestoration(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	siteID, userID, folderName, dataFolder string,
	startTime time.Time,
) {
	drive, err := client.
		Sites().
		BySiteId(siteID).
		Drive().
		Get(ctx, nil)
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	checkDriveRestoration(
		ctx,
		client,
		path.SharePointService,
		folderName,
		ptr.Val(drive.GetId()),
		ptr.Val(drive.GetName()),
		dataFolder,
		startTime,
		true)
}
