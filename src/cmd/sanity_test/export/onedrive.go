package export

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/cmd/sanity_test/restore"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func CheckOneDriveExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	drive, err := ac.Users().GetDefaultDrive(ctx, envs.UserID)
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	// map itemID -> item size
	var (
		fileSizes       = make(map[string]int64)
		exportFileSizes = make(map[string]int64)
		startTime       = time.Now()
	)

	err = filepath.Walk(
		envs.FolderName,
		common.FilepathWalker(envs.FolderName, exportFileSizes, envs.StartTime))
	if err != nil {
		fmt.Println("Error walking the path:", err)
	}

	_ = restore.PopulateDriveDetails(
		ctx,
		ac,
		ptr.Val(drive.GetId()),
		envs.FolderName,
		envs.DataFolder,
		fileSizes,
		map[string][]common.PermissionInfo{},
		startTime)

	for fileName, expected := range fileSizes {
		common.LogAndPrint(ctx, "checking for file: %s", fileName)

		got := exportFileSizes[fileName]

		common.Assert(
			ctx,
			func() bool { return expected == got },
			fmt.Sprintf("different file size: %s", fileName),
			expected,
			got)
	}

	fmt.Println("Success")
}
