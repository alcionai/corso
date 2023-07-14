package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/store"
)

func main() {
	var (
		service path.ServiceType
		cc      cobra.Command
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

	r, _, _, err := utils.GetAccountAndConnect(cc.Context(), service, nil)
	if err != nil {
		fatal(cc.Context(), "unable to connect account", err)
	}

	defer r.Close(cc.Context())

	backups, err := r.BackupsByTag(cc.Context(), store.Service(service))
	if err != nil {
		fatal(cc.Context(), "unable to find backups", err)
	}

	days, err := strconv.Atoi(os.Getenv("DELETION_DAYS"))
	if err != nil {
		fatal(cc.Context(), "invalid no of days provided", nil)
	}

	for _, backup := range backups {
		if backup.StartAndEndTime.CompletedAt.Before(time.Now().AddDate(0, 0, -days)) {
			err := r.DeleteBackup(cc.Context(), backup.ID.String())
			if err != nil {
				fatal(cc.Context(), "deleting backup", err)
			}

			logAndPrint(cc.Context(), "Deleted backup %s", backup.ID.String())
		}
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
