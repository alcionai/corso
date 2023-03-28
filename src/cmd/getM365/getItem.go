// getItem.go is a source file designed to retrieve an m365 object from an
// existing M365 account. Data displayed is representative of the current
// serialization abstraction versioning used by Microsoft Graph and stored by Corso.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/spf13/cobra"

	"github.com/alcionai/clues"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a M365ID item JSON",
	RunE:  handleGetCommand,
}

// Required inputs from user for command execution
var (
	tenant, user, m365ID, category string
)

// main function will produce the JSON String for a given m365 object of a
// user. Displayed Objects can be used as inputs for Mockable data
// Supports:
// - exchange (contacts, email, and events)
// Input: go run ./getItem.go     --user <user>
//
//	--m365ID <m365ID> --category <oneof: contacts, email, events>
func main() {
	ctx, _ := logger.SeedLevel(context.Background(), logger.Development)
	ctx = SetRootCmd(ctx, getCmd)

	defer logger.Flush(ctx)

	fs := getCmd.PersistentFlags()
	fs.StringVar(&user, "user", "", "m365 user id of M365 user")
	fs.StringVar(&tenant, "tenant", "",
		"m365 Tenant: m365 identifier for the tenant, not required if active in OS Environment")
	fs.StringVar(&m365ID, "m365ID", "", "m365 identifier for object to be created")
	fs.StringVar(&category, "category", "", "type of M365 data (contacts, email, events or files)") // files not supported

	cobra.CheckErr(getCmd.MarkPersistentFlagRequired("user"))
	cobra.CheckErr(getCmd.MarkPersistentFlagRequired("m365ID"))
	cobra.CheckErr(getCmd.MarkPersistentFlagRequired("category"))

	if err := getCmd.ExecuteContext(ctx); err != nil {
		logger.Flush(ctx)
		os.Exit(1)
	}
}

func handleGetCommand(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	_, creds, err := getGC(ctx)
	if err != nil {
		return err
	}

	err = runDisplayM365JSON(ctx, creds, user, m365ID, fault.New(true))
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Error displaying item: "+m365ID))
	}

	return nil
}

func runDisplayM365JSON(
	ctx context.Context,
	creds account.M365Config,
	user, itemID string,
	errs *fault.Bus,
) error {
	var (
		bs  []byte
		err error
		cat = path.ToCategoryType(category)
		sw  = kw.NewJsonSerializationWriter()
	)

	ac, err := api.NewClient(creds)
	if err != nil {
		return err
	}

	switch cat {
	case path.EmailCategory:
		bs, err = getItem(ctx, ac.Mail(), user, itemID, errs)
	case path.EventsCategory:
		bs, err = getItem(ctx, ac.Events(), user, itemID, errs)
	case path.ContactsCategory:
		bs, err = getItem(ctx, ac.Contacts(), user, itemID, errs)
	default:
		return fmt.Errorf("unable to process category: %s", cat)
	}

	if err != nil {
		return err
	}

	str := string(bs)

	err = sw.WriteStringValue("", &str)
	if err != nil {
		return clues.Wrap(err, "Error writing string value: "+itemID)
	}

	array, err := sw.GetSerializedContent()
	if err != nil {
		return clues.Wrap(err, "Error serializing item: "+itemID)
	}

	fmt.Println(string(array))

	return nil
}

type itemer interface {
	GetItem(
		ctx context.Context,
		user, itemID string,
		errs *fault.Bus,
	) (serialization.Parsable, *details.ExchangeInfo, error)
	Serialize(
		ctx context.Context,
		item serialization.Parsable,
		user, itemID string,
	) ([]byte, error)
}

func getItem(
	ctx context.Context,
	itm itemer,
	user, itemID string,
	errs *fault.Bus,
) ([]byte, error) {
	sp, _, err := itm.GetItem(ctx, user, itemID, errs)
	if err != nil {
		return nil, clues.Wrap(err, "getting item")
	}

	return itm.Serialize(ctx, sp, user, itemID)
}

//-------------------------------------------------------------------------------
// Helpers
//-------------------------------------------------------------------------------

func getGC(ctx context.Context) (*connector.GraphConnector, account.M365Config, error) {
	// get account info
	m365Cfg := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: common.First(tenant, os.Getenv(account.AzureTenantID)),
	}

	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return nil, m365Cfg, Only(ctx, clues.Wrap(err, "finding m365 account details"))
	}

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Users, errs)
	if err != nil {
		return nil, m365Cfg, Only(ctx, clues.Wrap(err, "connecting to graph API"))
	}

	return gc, m365Cfg, nil
}
