package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type getMailInboxer interface {
	GetMailInbox(ctx context.Context, userID string) (models.MailFolderable, error)
}

func IsServiceEnabled(
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
		if err := api.EvaluateMailboxError(clues.Stack(err)); err != nil {
			logger.CtxErr(ctx, err).Error("getting user's mail folder")

			return mi, err
		}

		logger.Ctx(ctx).Info("resource owner does not have a mailbox enabled")

		mi.ErrGetMailBoxSetting = append(
			mi.ErrGetMailBoxSetting,
			api.ErrMailBoxNotFound)

		return mi, nil
	}

	mboxSettings, err := gmb.GetMailboxSettings(ctx, userID)
	if err != nil {
		logger.CtxErr(ctx, err).Info("err getting user's mailbox settings")

		if !graph.IsErrAccessDenied(err) {
			return mi, clues.Wrap(err, "getting user's mailbox settings")
		}

		logger.CtxErr(ctx, err).Info("mailbox settings access denied")

		mi.ErrGetMailBoxSetting = append(
			mi.ErrGetMailBoxSetting,
			api.ErrMailBoxSettingsAccessDenied)
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
