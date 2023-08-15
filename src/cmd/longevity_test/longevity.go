package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/store"
)

// deleteBackups connects to the repository and deletes all backups for
// service that are at least deletionDays old. Returns the IDs of all backups
// that were deleted.
func deleteBackups(
	ctx context.Context,
	service path.ServiceType,
	deletionDays int,
) ([]string, error) {
	ctx = clues.Add(ctx, "cutoff_days", deletionDays)

	r, _, _, _, err := utils.GetAccountAndConnect(ctx, service, nil)
	if err != nil {
		return nil, clues.Wrap(err, "connecting to account").WithClues(ctx)
	}

	defer r.Close(ctx)

	backups, err := r.BackupsByTag(ctx, store.Service(service))
	if err != nil {
		return nil, clues.Wrap(err, "listing backups").WithClues(ctx)
	}

	var (
		deleted []string
		cutoff  = time.Now().Add(-time.Hour * 24 * time.Duration(deletionDays))
	)

	for _, backup := range backups {
		if backup.StartAndEndTime.CompletedAt.Before(cutoff) {
			if err := r.DeleteBackups(ctx, true, backup.ID.String()); err != nil {
				return nil, clues.Wrap(
					err,
					"deleting backup").
					With("backup_id", backup.ID).
					WithClues(ctx)
			}

			deleted = append(deleted, backup.ID.String())
			logAndPrint(ctx, "Deleted backup %s", backup.ID.String())
		}
	}

	return deleted, nil
}

// pitrListBackups connects to the repository at the given point in time and
// lists the backups for service. It then checks the list of backups contains
// the backups in backupIDs.
func pitrListBackups(
	ctx context.Context,
	service path.ServiceType,
	pitr time.Time,
	backupIDs []string,
) error {
	if len(backupIDs) == 0 {
		return nil
	}

	ctx = clues.Add(ctx, "pitr_time", pitr, "search_backups", backupIDs)

	// TODO(ashmrtn): This may be moved into CLI layer at some point when we add
	// flags for opening a repo at a point in time.
	cfg, err := config.GetConfigRepoDetails(ctx, true, true, nil)
	if err != nil {
		return clues.Wrap(err, "getting config info")
	}

	opts := utils.ControlWithConfig(cfg)
	opts.Repo.ViewTimestamp = &pitr

	r, err := repository.Connect(ctx, cfg.Account, cfg.Storage, cfg.RepoID, opts)
	if err != nil {
		return clues.Wrap(err, "connecting to repo").WithClues(ctx)
	}

	defer r.Close(ctx)

	backups, err := r.BackupsByTag(ctx, store.Service(service))
	if err != nil {
		return clues.Wrap(err, "listing backups").WithClues(ctx)
	}

	bups := map[string]struct{}{}

	for _, backup := range backups {
		bups[backup.ID.String()] = struct{}{}
	}

	ctx = clues.Add(ctx, "found_backups", maps.Keys(bups))

	for _, backupID := range backupIDs {
		if _, ok := bups[backupID]; !ok {
			return clues.New("looking for backup").
				With("search_backup_id", backupID).
				WithClues(ctx)
		}
	}

	return nil
}

func main() {
	var (
		service path.ServiceType
		cc      = cobra.Command{}
	)

	cc.SetContext(context.Background())

	if err := config.InitFunc(&cc, []string{}); err != nil {
		return
	}

	switch serviceName := os.Getenv("SERVICE"); serviceName {
	case "exchange":
		service = path.ExchangeService
	case "onedrive":
		service = path.OneDriveService
	case "sharepoint":
		service = path.SharePointService
	default:
		fatal(cc.Context(), "unknown service", nil)
	}

	ctx := clues.Add(cc.Context(), "service", service)

	days, err := strconv.Atoi(os.Getenv("DELETION_DAYS"))
	if err != nil {
		fatal(ctx, "invalid number of days provided", nil)
	}

	beforeDel := time.Now()

	backups, err := deleteBackups(ctx, service, days)
	if err != nil {
		fatal(ctx, "deleting backups", clues.Stack(err))
	}

	if err := pitrListBackups(ctx, service, beforeDel, backups); err != nil {
		fatal(ctx, "listing backups from point in time", clues.Stack(err))
	}
}

func fatal(ctx context.Context, msg string, err error) {
	logger.CtxErr(ctx, err).Error("test failure: " + msg)
	fmt.Println(msg+": ", err)
	os.Exit(1)
}

func logAndPrint(ctx context.Context, tmpl string, vs ...any) {
	logger.Ctx(ctx).Infof(tmpl, vs...)
	fmt.Printf(tmpl+"\n", vs...)
}
