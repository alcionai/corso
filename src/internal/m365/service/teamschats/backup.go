package teamschats

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/teamschats"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type teamsChatsBackup struct{}

// NewBackup provides a struct that matches standard apis
// across m365/service handlers.
func NewBackup() *teamsChatsBackup {
	return &teamsChatsBackup{}
}

func (teamsChatsBackup) ProduceBackupCollections(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	creds account.M365Config,
	su support.StatusUpdater,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, *prefixmatcher.StringSetMatcher, bool, error) {
	b, err := bpc.Selector.ToTeamsChatsBackup()
	if err != nil {
		return nil, nil, true, clues.WrapWC(ctx, err, "parsing selector")
	}

	var (
		el          = errs.Local()
		collections = []data.BackupCollection{}
		categories  = map[path.CategoryType]struct{}{}
	)

	ctx = clues.Add(
		ctx,
		"user_id", clues.Hide(bpc.ProtectedResource.ID()),
		"user_name", clues.Hide(bpc.ProtectedResource.Name()))

	bc := backupCommon{
		apiCli:         ac,
		producerConfig: bpc,
		creds:          creds,
		user:           bpc.ProtectedResource,
		statusUpdater:  su,
	}

	for _, scope := range b.Scopes() {
		if el.Failure() != nil {
			break
		}

		cl := counter.Local()
		ictx := clues.AddLabelCounter(ctx, cl.PlainAdder())
		ictx = clues.Add(ictx, "category", scope.Category().PathType())

		var colls []data.BackupCollection

		switch scope.Category().PathType() {
		case path.ChatsCategory:
			colls, err = backupChats(
				ictx,
				bc,
				scope,
				cl,
				el)
		}

		if err != nil {
			el.AddRecoverable(ctx, clues.Stack(err))
			continue
		}

		collections = append(collections, colls...)

		categories[scope.Category().PathType()] = struct{}{}
	}

	if len(collections) > 0 {
		baseCols, err := graph.BaseCollections(
			ctx,
			collections,
			creds.AzureTenantID,
			bpc.ProtectedResource.ID(),
			path.TeamsChatsService,
			categories,
			su,
			counter,
			errs)
		if err != nil {
			return nil, nil, true, err
		}

		collections = append(collections, baseCols...)
	}

	counter.Add(count.Collections, int64(len(collections)))

	logger.Ctx(ctx).Infow("produced collections", "stats", counter.Values())

	return collections, nil, true, clues.Stack(el.Failure()).OrNil()
}

type backupCommon struct {
	apiCli         api.Client
	producerConfig inject.BackupProducerConfig
	creds          account.M365Config
	user           idname.Provider
	statusUpdater  support.StatusUpdater
}

func backupChats(
	ctx context.Context,
	bc backupCommon,
	scope selectors.TeamsChatsScope,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	var colls []data.BackupCollection

	progressMessage := observe.MessageWithCompletion(
		ctx,
		observe.ProgressCfg{
			Indent:            1,
			CompletionMessage: func() string { return "(done)" },
		},
		scope.Category().PathType().HumanString())
	defer close(progressMessage)

	bh := teamschats.NewUsersChatsBackupHandler(
		bc.creds.AzureTenantID,
		bc.producerConfig.ProtectedResource.ID(),
		bc.apiCli.Chats())

	// Always disable lazy reader for channels until #4321 support is added
	useLazyReader := false

	colls, _, err := teamschats.CreateCollections(
		ctx,
		bc.producerConfig,
		bh,
		bc.creds.AzureTenantID,
		scope,
		bc.statusUpdater,
		useLazyReader,
		counter,
		errs)

	return colls, clues.Stack(err).OrNil()
}
