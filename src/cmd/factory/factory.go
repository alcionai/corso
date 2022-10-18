package main

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "Generate all types of m365 folders",
	RunE:  handleFactoryRoot,
}

var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Generate exchange data",
	RunE:  handleExchangeFactory,
}

var oneDriveCmd = &cobra.Command{
	Use:   "onedrive",
	Short: "Generate onedrive data",
	RunE:  handleOneDriveFactory,
}

var (
	count       int
	destination string
	tenant      string
	user        string
)

// TODO: ErrGenerating       = errors.New("not all items were successfully generated")

var ErrNotYetImplemeted = errors.New("not yet implemented")

// ------------------------------------------------------------------------------------------
// CLI command handlers
// ------------------------------------------------------------------------------------------

func main() {
	ctx, _ := logger.SeedLevel(context.Background(), logger.Development)
	ctx = SetRootCmd(ctx, factoryCmd)

	defer logger.Flush(ctx)

	// persistent flags that are common to all use cases
	fs := factoryCmd.PersistentFlags()
	fs.StringVar(&tenant, "tenant", "", "m365 tenant containing the user")
	fs.StringVar(&user, "user", "", "m365 user owning the new data")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("user"))
	fs.IntVar(&count, "count", 0, "count of items to produce")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("count"))
	fs.StringVar(&destination, "destination", "", "destination of the new data (will create as needed)")
	cobra.CheckErr(factoryCmd.MarkPersistentFlagRequired("destination"))

	factoryCmd.AddCommand(exchangeCmd)
	addExchangeCommands(exchangeCmd)
	factoryCmd.AddCommand(oneDriveCmd)
	addOneDriveCommands(oneDriveCmd)

	if err := factoryCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func handleFactoryRoot(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)
	return cmd.Help()
}

func handleExchangeFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)
	return cmd.Help()
}

func handleOneDriveFactory(cmd *cobra.Command, args []string) error {
	Err(cmd.Context(), ErrNotYetImplemeted)
	return cmd.Help()
}

// ------------------------------------------------------------------------------------------
// Common Helpers
// ------------------------------------------------------------------------------------------

func getGCAndVerifyUser(ctx context.Context, userID string) (*connector.GraphConnector, string, error) {
	tid := common.First(tenant, os.Getenv(account.AzureTenantID))

	// get account info
	m365Cfg := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: tid,
	}

	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return nil, "", errors.Wrap(err, "finding m365 account details")
	}

	// build a graph connector
	gc, err := connector.NewGraphConnector(ctx, acct)
	if err != nil {
		return nil, "", errors.Wrap(err, "connecting to graph api")
	}

	if _, ok := gc.Users[user]; !ok {
		return nil, "", errors.New("user not found within tenant")
	}

	return gc, tid, nil
}

type item struct {
	name string
	data []byte
}

type collection struct {
	// Elements (in order) for the path representing this collection. Should
	// only contain elements after the prefix that corso uses for the path. For
	// example, a collection for the Inbox folder in exchange mail would just be
	// "Inbox".
	pathElements []string
	category     path.CategoryType
	items        []item
}

func buildCollections(
	service path.ServiceType,
	tenant, user string,
	dest control.RestoreDestination,
	colls []collection,
) ([]data.Collection, error) {
	collections := make([]data.Collection, 0, len(colls))

	for _, c := range colls {
		pth, err := toDataLayerPath(
			service,
			tenant,
			user,
			c.category,
			c.pathElements,
			false,
		)
		if err != nil {
			return nil, err
		}

		mc := mockconnector.NewMockExchangeCollection(pth, len(c.items))

		for i := 0; i < len(c.items); i++ {
			mc.Names[i] = c.items[i].name
			mc.Data[i] = c.items[i].data
		}

		collections = append(collections, mc)
	}

	return collections, nil
}

func toDataLayerPath(
	service path.ServiceType,
	tenant, user string,
	category path.CategoryType,
	elements []string,
	isItem bool,
) (path.Path, error) {
	pb := path.Builder{}.Append(elements...)

	switch service {
	case path.ExchangeService:
		return pb.ToDataLayerExchangePathForCategory(tenant, user, category, isItem)
	case path.OneDriveService:
		return pb.ToDataLayerOneDrivePath(tenant, user, isItem)
	}

	return nil, errors.Errorf("unknown service %s", service.String())
}
