package restore

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/cmd/sanity_test/common"
	"github.com/alcionai/canario/src/cmd/sanity_test/driveish"
	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

func CheckSharePointRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	if envs.Category == "lists" {
		CheckSharePointListsRestoration(ctx, ac, envs)
	}

	if envs.Category == "libraries" {
		drive, err := ac.Sites().GetDefaultDrive(ctx, envs.SiteID)
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
}

func CheckSharePointListsRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	restoredTree := BuildListsSanitree(ctx, ac, envs.SiteID, envs.RestoreContainerPrefix, "")
	sourceTree := BuildListsSanitree(ctx, ac, envs.SiteID, envs.SourceContainer, "")

	ctx = clues.Add(
		ctx,
		"restore_container_id", restoredTree.ID,
		"restore_container_name", restoredTree.Name,
		"source_container_id", sourceTree.ID,
		"source_container_name", sourceTree.Name)

	common.CompareDiffTrees[models.Siteable, models.Listable](
		ctx,
		sourceTree,
		restoredTree,
		nil)

	common.Infof(ctx, "Success")
}

func BuildListsSanitree(
	ctx context.Context,
	ac api.Client,
	siteID string,
	restoreContainerPrefix, exportFolderName string,
) *common.Sanitree[models.Siteable, models.Listable] {
	common.Infof(ctx, "building sanitree for lists of site: %s", siteID)

	site, err := ac.Sites().GetByID(ctx, siteID, api.CallConfig{})
	if err != nil {
		common.Fatal(
			ctx,
			fmt.Sprintf("finding site by id %q", siteID),
			err)
	}

	cfg := api.CallConfig{
		Select: idAnd("displayName", "list", "lastModifiedDateTime"),
	}

	lists, err := ac.Lists().GetLists(ctx, siteID, cfg)
	if err != nil {
		common.Fatal(
			ctx,
			fmt.Sprintf("finding lists of site with id %q", siteID),
			err)
	}

	lists = filterToSupportedLists(lists)

	filteredLists := filterListsByPrefix(lists, restoreContainerPrefix)

	rootTreeName := ptr.Val(site.GetDisplayName())
	// lists get stored into the local dir at destination/Lists/
	if len(exportFolderName) > 0 {
		rootTreeName = exportFolderName
	}

	root := &common.Sanitree[models.Siteable, models.Listable]{
		Self:        site,
		ID:          ptr.Val(site.GetId()),
		Name:        rootTreeName,
		CountLeaves: len(filteredLists),
		Leaves:      map[string]*common.Sanileaf[models.Siteable, models.Listable]{},
	}

	for _, list := range filteredLists {
		listID := ptr.Val(list.GetId())

		listItems, err := ac.Lists().GetListItems(ctx, siteID, listID, api.CallConfig{})
		if err != nil {
			common.Fatal(
				ctx,
				fmt.Sprintf("finding listItems of list with id %q", listID),
				err)
		}

		m := &common.Sanileaf[models.Siteable, models.Listable]{
			Parent: root,
			Self:   list,
			ID:     listID,
			Name:   ptr.Val(list.GetDisplayName()),
			// using list item count as size for lists
			Size: int64(len(listItems)),
		}

		root.Leaves[m.ID] = m
	}

	return root
}

func filterToSupportedLists(lists []models.Listable) []models.Listable {
	filteredLists := make([]models.Listable, 0)

	for _, list := range lists {
		if !api.SkipListTemplates.HasKey(ptr.Val(list.GetList().GetTemplate())) {
			filteredLists = append(filteredLists, list)
		}
	}

	return filteredLists
}

func filterListsByPrefix(lists []models.Listable, prefix string) []models.Listable {
	result := []models.Listable{}

	for _, list := range lists {
		for _, pfx := range strings.Split(prefix, ",") {
			if strings.HasPrefix(ptr.Val(list.GetDisplayName()), pfx) {
				result = append(result, list)
				break
			}
		}
	}

	return result
}

func idAnd(ss ...string) []string {
	id := []string{"id"}

	if len(ss) == 0 {
		return id
	}

	return append(id, ss...)
}
