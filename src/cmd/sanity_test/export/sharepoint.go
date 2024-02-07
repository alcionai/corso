package export

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/tidwall/gjson"

	"github.com/alcionai/canario/src/cmd/sanity_test/common"
	"github.com/alcionai/canario/src/cmd/sanity_test/driveish"
	"github.com/alcionai/canario/src/cmd/sanity_test/restore"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

func CheckSharePointExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	if envs.Category == path.ListsCategory.String() {
		CheckSharepointListsExport(ctx, ac, envs)
	}

	if envs.Category == path.LibrariesCategory.String() {
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
}

func CheckSharepointListsExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	exportFolderName := path.ListsCategory.HumanString()

	sourceTree := restore.BuildListsSanitree(ctx, ac, envs.SiteID, envs.SourceContainer, exportFolderName)

	listsExportDir := filepath.Join(envs.RestoreContainer, exportFolderName)
	exportedTree := BuildFilepathSanitreeForSharepointLists(ctx, listsExportDir)

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

func BuildFilepathSanitreeForSharepointLists(
	ctx context.Context,
	rootDir string,
) *common.Sanitree[fs.FileInfo, fs.FileInfo] {
	var root *common.Sanitree[fs.FileInfo, fs.FileInfo]

	walker := func(
		p string,
		info os.FileInfo,
		err error,
	) error {
		if root == nil {
			root = common.CreateNewRoot(info, false)
			return nil
		}

		relPath := common.GetRelativePath(
			ctx,
			rootDir,
			p,
			info,
			err)

		if !info.IsDir() {
			file, err := os.Open(p)
			if err != nil {
				common.Fatal(ctx, "opening file to read", err)
			}
			defer file.Close()

			content, err := io.ReadAll(file)
			if err != nil {
				common.Fatal(ctx, "reading file", err)
			}

			res := gjson.Get(string(content), "items.#")
			itemsCount := res.Num

			elems := path.Split(relPath)

			node := root.NodeAt(ctx, elems[:len(elems)-2])
			node.CountLeaves++
			node.Leaves[info.Name()] = &common.Sanileaf[fs.FileInfo, fs.FileInfo]{
				Parent: node,
				Self:   info,
				ID:     info.Name(),
				Name:   info.Name(),
				// using list item count as size for lists
				Size: int64(itemsCount),
			}
		}

		return nil
	}

	err := filepath.Walk(rootDir, walker)
	if err != nil {
		common.Fatal(ctx, "walking filepath", err)
	}

	return root
}
