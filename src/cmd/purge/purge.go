package main

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/cli/utils"
	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/exchange"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/credentials"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge m365 data",
	RunE:  doFolderPurge,
}

var (
	before string
	user   string
	tenant string
	prefix string
)

func doFolderPurge(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	// get account info
	m365Cfg := account.M365Config{
		M365:     credentials.GetM365(),
		TenantID: common.First(tenant, os.Getenv(account.TenantID)),
	}
	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return Only(errors.Wrap(err, "finding m365 account details"))
	}

	// build a graph connector
	gc, err := connector.NewGraphConnector(acct)
	if err != nil {
		return Only(errors.Wrap(err, "connecting to graph api"))
	}

	// get them folders
	mfs, err := exchange.GetAllMailFolders(gc.Service(), user, prefix)
	if err != nil {
		return Only(errors.Wrap(err, "retrieving mail folders"))
	}

	// format the time input
	beforeTime := time.Now().UTC()
	if len(before) > 0 {
		beforeTime, err = common.ParseTime(before)
		if err != nil {
			return Only(errors.Wrap(err, "parsing before flag to time"))
		}
	}
	stLen := len(common.SimpleDateTimeFormat)

	// delete files
	for _, mf := range mfs {

		// compare the folder time to the deletion boundary time first
		var shouldDelete bool
		dnLen := len(mf.DisplayName)
		if dnLen > stLen {
			dnSuff := mf.DisplayName[dnLen-stLen:]
			dnTime, err := common.ParseTime(dnSuff)
			if err != nil {
				Info(errors.Wrapf(err, "Error: deleting folder [%s]", mf.DisplayName))
				continue
			}
			shouldDelete = dnTime.Before(beforeTime)
		}

		if !shouldDelete {
			continue
		}

		Info("Deleting folder: ", mf.DisplayName)
		err = exchange.DeleteMailFolder(gc.Service(), user, mf.ID)
		if err != nil {
			Info(errors.Wrapf(err, "Error: deleting folder [%s]", mf.DisplayName))
		}
	}

	return nil
}

func main() {
	fs := purgeCmd.Flags()
	fs.StringVar(&before, "before", "", "folders older than this date are deleted.  (default: now in UTC)")
	fs.StringVar(&user, "user", "", "m365 user id whose folders will be deleted")
	cobra.CheckErr(purgeCmd.MarkFlagRequired("user"))
	fs.StringVar(&tenant, "tenant", "", "m365 tenant containing the user")
	fs.StringVar(&prefix, "prefix", "", "filters mail folders by displayName prefix")
	cobra.CheckErr(purgeCmd.MarkFlagRequired("prefix"))

	if err := purgeCmd.Execute(); err != nil {
		Info("Error: ", err.Error())
		os.Exit(1)
	}
}
