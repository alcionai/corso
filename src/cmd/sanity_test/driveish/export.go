package driveish

import (
	"context"
	"io/fs"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func CheckExport(
	ctx context.Context,
	ac api.Client,
	drive models.Driveable,
	envs common.Envs,
) {
	var (
		driveID   = ptr.Val(drive.GetId())
		driveName = ptr.Val(drive.GetName())
	)

	ctx = clues.Add(
		ctx,
		"drive_id", driveID,
		"drive_name", driveName)

	root := populateSanitree(
		ctx,
		ac,
		driveID,
		[]string{envs.SourceContainer})

	sourceTree, ok := root.Children[envs.SourceContainer]
	common.Assert(
		ctx,
		func() bool { return ok },
		"should find root-level source data folder",
		envs.SourceContainer,
		"not found")

	fpTree := common.BuildFilepathSanitree(ctx, envs.RestoreContainer)

	comparator := func(
		ctx context.Context,
		expect *common.Sanitree[models.DriveItemable, models.DriveItemable],
		result *common.Sanitree[fs.FileInfo, fs.FileInfo],
	) {
		common.CompareLeaves(ctx, expect.Leaves, result.Leaves, nil)
	}

	common.CompareDiffTrees(
		ctx,
		sourceTree,
		fpTree.Children[envs.SourceContainer],
		comparator)

	common.Infof(ctx, "Success")
}
