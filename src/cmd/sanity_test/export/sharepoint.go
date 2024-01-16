package export

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/cmd/sanity_test/driveish"
	"github.com/alcionai/corso/src/cmd/sanity_test/restore"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func CheckSharePointExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	CheckSharepointListsExport(ctx, ac, envs)

	drive, err := ac.Sites().GetDefaultDrive(ctx, envs.SiteID)
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	envs.RestoreContainer = filepath.Join(envs.RestoreContainer, "Libraries/Documents") // check in default loc
	driveish.CheckExport(
		ctx,
		ac,
		drive,
		envs)
}

func CheckSharepointListsExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	exportFolderName := "Lists"

	sourceTree := restore.BuildListsSanitree(ctx, ac, envs.SiteID, envs.SourceContainer, exportFolderName)

	listsExportDir := filepath.Join(envs.RestoreContainer, exportFolderName)
	exportedTree := common.BuildFilepathSanitreeForSharepointLists(ctx, listsExportDir)

	ctx = clues.Add(
		ctx,
		"export_container_id", exportedTree.ID,
		"export_container_name", exportedTree.Name,
		"source_container_id", sourceTree.ID,
		"source_container_name", sourceTree.Name)

	comparator := func(
		ctx context.Context,
		expect *common.Sanitree[models.Siteable, models.Listable],
		result *common.Sanitree[fs.FileInfo, fs.FileInfo],
	) {
		modifiedResultLeaves := map[string]*common.Sanileaf[fs.FileInfo, fs.FileInfo]{}

		for key, val := range result.Leaves {
			fixedName := strings.TrimSuffix(key, ".json")

			modifiedResultLeaves[fixedName] = val
		}

		common.CompareLeaves(ctx, expect.Leaves, modifiedResultLeaves, nil)
	}

	common.CompareDiffTrees(
		ctx,
		sourceTree,
		exportedTree,
		comparator)

	common.Infof(ctx, "Success")
}
