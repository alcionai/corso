package export

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/cmd/sanity_test/restore"
	"github.com/alcionai/corso/src/cmd/sanity_test/utils"
	"github.com/alcionai/corso/src/internal/common/ptr"
)

func CheckOneDriveExport(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	userID, folderName, dataFolder string,
) {
	drive, err := client.
		Users().
		ByUserId(userID).
		Drive().
		Get(ctx, nil)
	if err != nil {
		utils.Fatal(ctx, "getting the drive:", err)
	}

	// map itemID -> item size
	fileSizes := make(map[string]int64)       // exportFileSizes = make(map[string]int64)
	exportFileSizes := make(map[string]int64) // exportFileSizes = make(map[string]int64)
	startTime := time.Now()

	err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(folderName, path)
		if err != nil {
			return err
		}

		exportFileSizes[relPath] = info.Size()
		if startTime.After(info.ModTime()) {
			startTime = info.ModTime()
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
	}

	_ = restore.PopulateDriveDetails(
		ctx,
		client,
		ptr.Val(drive.GetId()),
		folderName,
		dataFolder,
		fileSizes,
		map[string][]restore.PermissionInfo{},
		startTime,
	)

	for fileName, expected := range fileSizes {
		utils.LogAndPrint(ctx, "checking for file: %s", fileName)

		got := exportFileSizes[fileName]

		utils.Assert(
			ctx,
			func() bool { return expected == got },
			fmt.Sprintf("different file size: %s", fileName),
			expected,
			got)
	}

	fmt.Println("Success")
}
