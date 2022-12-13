// getItem.go is a source file designed to retrieve an m365 object from an
// existing M365 account. Data displayed is representative of the current
// serialization abstraction versioning used by Microsoft Graph and stored by Corso.

package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
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

	gc, err := getGC(ctx)
	if err != nil {
		return err
	}

	err = runDisplayM365JSON(
		ctx,
		gc)
	if err != nil {
		return Only(ctx, errors.Wrapf(err, "unable to create mock from M365: %s", m365ID))
	}

	return nil
}

func runDisplayM365JSON(
	ctx context.Context,
	gs graph.Servicer,
) error {
	var (
		get           exchange.GraphRetrievalFunc
		serializeFunc exchange.GraphSerializeFunc
		cat           = graph.StringToPathCategory(category)
	)

	switch cat {
	case path.EmailCategory, path.EventsCategory, path.ContactsCategory:
		get, serializeFunc = exchange.GetQueryAndSerializeFunc(exchange.CategoryToOptionIdentifier(cat))
	default:
		return fmt.Errorf("unable to process category: %s", cat)
	}

	channel := make(chan data.Stream, 1)

	sw := kw.NewJsonSerializationWriter()

	response, err := get(ctx, gs, user, m365ID)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	// First return is the number of bytes that were serialized. Ignored
	_, err = serializeFunc(ctx, gs.Client(), sw, channel, response, user)
	close(channel)

	if err != nil {
		return err
	}

	for item := range channel {
		buf := &bytes.Buffer{}

		_, err := buf.ReadFrom(item.ToReader())
		if err != nil {
			return errors.Wrapf(err, "unable to parse given data: %s", m365ID)
		}

		byteArray := buf.Bytes()
		newValue := string(byteArray)

		err = sw.WriteStringValue("", &newValue)
		if err != nil {
			return errors.Wrapf(err, "unable to %s to string value", m365ID)
		}

		array, err := sw.GetSerializedContent()
		if err != nil {
			return errors.Wrapf(err, "unable to serialize new value from M365:%s", m365ID)
		}

		fmt.Println(string(array))

		//lint:ignore SA4004 only expecting one item
		return nil
	}

	// This should never happen
	return errors.New("m365 object not serialized")
}

//-------------------------------------------------------------------------------
// Helpers
//-------------------------------------------------------------------------------

func getGC(ctx context.Context) (*connector.GraphConnector, error) {
	// get account info
	m365Cfg := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: common.First(tenant, os.Getenv(account.AzureTenantID)),
	}

	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return nil, Only(ctx, errors.Wrap(err, "finding m365 account details"))
	}

	gc, err := connector.NewGraphConnector(ctx, acct, connector.Users)
	if err != nil {
		return nil, Only(ctx, errors.Wrap(err, "connecting to graph API"))
	}

	return gc, nil
}
