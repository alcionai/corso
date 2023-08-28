package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// Contains m365 code shared between CLI & SDK

type getMailInboxer interface {
	GetMailInbox(ctx context.Context, userID string) (models.MailFolderable, error)
}

func IsExchangeServiceEnabled(
	ctx context.Context,
	gmi getMailInboxer,
	resource string,
) (bool, error) {
	_, err := gmi.GetMailInbox(ctx, resource)
	if err != nil {
		if err := api.EvaluateMailboxError(err); err != nil {
			logger.CtxErr(ctx, err).Error("getting user's mail folder")
			return false, clues.Stack(err)
		}

		logger.Ctx(ctx).Info("resource owner does not have a mailbox enabled")

		return false, nil
	}

	return true, nil
}

type getDefaultDriver interface {
	GetDefaultDrive(ctx context.Context, userID string) (models.Driveable, error)
}

func IsOneDriveServiceEnabled(
	ctx context.Context,
	gdd getDefaultDriver,
	resource string,
) (bool, error) {
	_, err := gdd.GetDefaultDrive(ctx, resource)
	if err != nil {
		// we consider this a non-error case, since it
		// answers the question the caller is asking.
		if clues.HasLabel(err, graph.LabelsMysiteNotFound) || clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return false, nil
		}

		if graph.IsErrUserNotFound(err) {
			logger.CtxErr(ctx, err).Info("resource owner not found")

			return false, clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		logger.CtxErr(ctx, err).Error("getting user's drive")

		return false, clues.Stack(err)
	}

	return true, nil
}

type getSiteRooter interface {
	GetRoot(ctx context.Context) (models.Siteable, error)
}

func IsSharePointServiceEnabled(
	ctx context.Context,
	gsr getSiteRooter,
	resource string,
) (bool, error) {
	_, err := gsr.GetRoot(ctx)
	if err != nil {
		if clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return false, nil
		}

		return false, clues.Stack(err)
	}

	return true, nil
}
