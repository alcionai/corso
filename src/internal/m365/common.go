package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
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

type getMailboxer interface {
	GetMailInbox(ctx context.Context, userID string) (models.MailFolderable, error)
	GetMailboxSettings(ctx context.Context, userID string) (models.Userable, error)
	GetFirstInboxMessage(ctx context.Context, userID, inboxID string) error
}

func GetMailboxInfo(
	ctx context.Context,
	gmb getMailboxer,
	userID string,
) (api.MailboxInfo, error) {
	mi := api.MailboxInfo{
		ErrGetMailBoxSetting: []error{},
	}

	// First check whether the user is able to access their inbox.
	inbox, err := gmb.GetMailInbox(ctx, userID)
	if err != nil {
		if err := api.EvaluateMailboxError(graph.Stack(ctx, err)); err != nil {
			logger.CtxErr(ctx, err).Error("getting user's mail folder")

			return mi, err
		}

		logger.Ctx(ctx).Info("resource owner does not have a mailbox enabled")

		mi.ErrGetMailBoxSetting = append(
			mi.ErrGetMailBoxSetting,
			api.ErrMailBoxSettingsNotFound)

		return mi, nil
	}

	mboxSettings, err := gmb.GetMailboxSettings(ctx, userID)
	if err != nil {
		logger.CtxErr(ctx, err).Info("err getting user's mailbox settings")

		if !graph.IsErrAccessDenied(err) {
			return mi, graph.Wrap(ctx, err, "getting user's mailbox settings")
		}

		logger.CtxErr(ctx, err).Info("mailbox settings access denied")

		mi.ErrGetMailBoxSetting = append(
			mi.ErrGetMailBoxSetting,
			api.ErrMailBoxSettingsAccessDenied,
		)
	} else {
		mi = api.ParseMailboxSettings(mboxSettings, mi)
	}

	err = gmb.GetFirstInboxMessage(ctx, userID, ptr.Val(inbox.GetId()))
	if err != nil {
		if !graph.IsErrQuotaExceeded(err) {
			return mi, clues.Stack(err)
		}

		mi.QuotaExceeded = graph.IsErrQuotaExceeded(err)
	}

	return mi, nil
}
