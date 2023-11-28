package export

import (
	"context"
	"path/filepath"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/cmd/sanity_test/restore"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

	common.AssertEqualTrees[models.MailFolderable, any](
		ctx,
		sourceTree,
		exportedTree.Children[envs.SourceContainer],
		nil,
		nil)

	common.Infof(ctx, "Success")
}
