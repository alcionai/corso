package export

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/cmd/sanity_test/restore"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func CheckEmailExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	sourceTree := restore.BuildEmailSanitree(ctx, ac, envs.UserID, envs.SourceContainer)

	emailsExportDir := filepath.Join(envs.RestoreContainer, "Emails")
	exportedTree := common.BuildFilepathSanitree(ctx, emailsExportDir)

	ctx = clues.Add(
		ctx,
		"export_container_id", exportedTree.ID,
		"export_container_name", exportedTree.Name,
		"source_container_id", sourceTree.ID,
		"source_container_name", sourceTree.Name)

	comparator := func(
		ctx context.Context,
		expect *common.Sanitree[models.MailFolderable, any],
		result *common.Sanitree[fs.FileInfo, fs.FileInfo],
	) {
		modifiedExpectedLeaves := map[string]*common.Sanileaf[models.MailFolderable, any]{}
		modifiedResultLeaves := map[string]*common.Sanileaf[fs.FileInfo, fs.FileInfo]{}

		for key, val := range expect.Leaves {
			val.Size = 0 // we cannot match up sizes
			modifiedExpectedLeaves[key] = val
		}

		for key, val := range result.Leaves {
			fixedName := strings.TrimSuffix(key, ".eml")
			val.Size = 0

			modifiedResultLeaves[fixedName] = val
		}

		common.CompareLeaves(ctx, expect.Leaves, modifiedResultLeaves, nil)
	}

	common.CompareDiffTrees(
		ctx,
		sourceTree,
		exportedTree.Children[envs.SourceContainer],
		comparator)

	common.Infof(ctx, "Success")
}
