package restore

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// CheckEmailRestoration verifies that the emails count in restored folder is equivalent to
// emails in actual m365 account
func CheckEmailRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	restoredTree := BuildEmailSanitree(ctx, ac, envs.UserID, envs.RestoreContainer)
	sourceTree := BuildEmailSanitree(ctx, ac, envs.UserID, envs.SourceContainer)

	ctx = clues.Add(
		ctx,
		"restore_container_id", restoredTree.ID,
		"restore_container_name", restoredTree.Name,
		"source_container_id", sourceTree.ID,
		"source_container_name", sourceTree.Name)

	// NOTE: We cannot compare leaves as the IDs of the restored items
	// differ from the original ones.
	common.CompareDiffTrees[models.MailFolderable, any](
		ctx,
		sourceTree,
		restoredTree.Children[envs.SourceContainer],
		nil)

	common.Infof(ctx, "Success")
}

func BuildEmailSanitree(
	ctx context.Context,
	ac api.Client,
	userID, folderName string,
) *common.Sanitree[models.MailFolderable, any] {
	gcc, err := ac.Mail().GetContainerByName(
		ctx,
		userID,
		api.MsgFolderRoot,
		folderName)
	if err != nil {
		common.Fatal(
			ctx,
			fmt.Sprintf("finding folder by name %q", folderName),
			err)
	}

	mmf, ok := gcc.(models.MailFolderable)
	if !ok {
		common.Fatal(
			ctx,
			"mail folderable required",
			clues.New("casting "+*gcc.GetDisplayName()+" to models.MailFolderable"))
	}

	root := &common.Sanitree[models.MailFolderable, any]{
		Self:        mmf,
		ID:          ptr.Val(mmf.GetId()),
		Name:        ptr.Val(mmf.GetDisplayName()),
		CountLeaves: int(ptr.Val(mmf.GetTotalItemCount())),
		Leaves:      map[string]*common.Sanileaf[models.MailFolderable, any]{},
		Children:    map[string]*common.Sanitree[models.MailFolderable, any]{},
	}

	mails, err := ac.Mail().GetItemsInContainer(
		ctx,
		userID,
		root.ID)
	if err != nil {
		common.Fatal(ctx, "getting child containers", err)
	}

	if len(mails) != root.CountLeaves {
		common.Fatal(
			ctx,
			"mails count mismatch",
			clues.New("mail message count mismatch from API"))
	}

	for _, mail := range mails {
		m := &common.Sanileaf[models.MailFolderable, any]{
			Parent: root,
			Self:   mail,
			ID:     ptr.Val(mail.GetId()),
			Name:   ptr.Val(mail.GetSubject()),
			Size:   int64(len(ptr.Val(mail.GetBody().GetContent()))),
		}

		root.Leaves[m.ID] = m
	}

	recursivelyBuildTree(ctx, ac, root, userID, root.Name+"/")

	return root
}

func recursivelyBuildTree(
	ctx context.Context,
	ac api.Client,
	stree *common.Sanitree[models.MailFolderable, any],
	userID, location string,
) {
	common.Debugf(ctx, "adding: %s", location)

	childFolders, err := ac.Mail().GetContainerChildren(
		ctx,
		userID,
		stree.ID)
	if err != nil {
		common.Fatal(ctx, "getting child containers", err)
	}

	for _, child := range childFolders {
		if int(ptr.Val(child.GetTotalItemCount()))+len(childFolders) == 0 {
			common.Infof(ctx, "skipped empty folder: %s/%s", location, ptr.Val(child.GetDisplayName()))
			continue
		}

		c := &common.Sanitree[models.MailFolderable, any]{
			Parent:      stree,
			Self:        child,
			ID:          ptr.Val(child.GetId()),
			Name:        ptr.Val(child.GetDisplayName()),
			CountLeaves: int(ptr.Val(child.GetTotalItemCount())),
			Leaves:      map[string]*common.Sanileaf[models.MailFolderable, any]{},
			Children:    map[string]*common.Sanitree[models.MailFolderable, any]{},
		}

		mails, err := ac.Mail().GetItemsInContainer(ctx, userID, c.ID)
		if err != nil {
			common.Fatal(ctx, "getting child containers", err)
		}

		for _, mail := range mails {
			m := &common.Sanileaf[models.MailFolderable, any]{
				Parent: c,
				Self:   mail,
				ID:     ptr.Val(mail.GetId()),
				Name:   ptr.Val(mail.GetSubject()),
				Size:   int64(len(ptr.Val(mail.GetBody().GetContent()))),
			}

			c.Leaves[m.ID] = m
		}

		stree.Children[c.Name] = c

		recursivelyBuildTree(ctx, ac, c, userID, location+c.Name+"/")
	}
}
