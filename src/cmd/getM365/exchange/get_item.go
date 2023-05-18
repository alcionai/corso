// get_item.go is a source file designed to retrieve an m365 object from an
// existing M365 account. Data displayed is representative of the current
// serialization abstraction versioning used by Microsoft Graph and stored by Corso.

package exchange

import (
	"context"
	"fmt"
	"os"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// Required inputs from user for command execution
var (
	user, tenant, m365ID, category string
)

func AddCommands(parent *cobra.Command) {
	exCmd := &cobra.Command{
		Use:   "exchange",
		Short: "Get an M365ID item JSON",
		RunE:  handleExchangeCmd,
	}

	fs := exCmd.PersistentFlags()
	fs.StringVar(&m365ID, "id", "", "m365 identifier for object")
	fs.StringVar(&category, "category", "", "type of M365 data (contacts, email, events)")
	fs.StringVar(&user, "user", "", "m365 user id of M365 user")
	fs.StringVar(&tenant, "tenant", "", "m365 identifier for the tenant")

	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("user"))
	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("id"))
	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("category"))

	parent.AddCommand(exCmd)
}

func handleExchangeCmd(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	tid := common.First(tenant, os.Getenv(account.AzureTenantID))

	ctx := clues.Add(
		cmd.Context(),
		"item_id", m365ID,
		"resource_owner", user,
		"tenant", tid)

	creds := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: tid,
	}

	err := runDisplayM365JSON(ctx, creds, user, m365ID, fault.New(true))
	if err != nil {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		return clues.Wrap(err, "getting item")
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
		bs, err = getItem(ctx, ac.Mail(), user, itemID, true, errs)
	case path.EventsCategory:
		bs, err = getItem(ctx, ac.Events(), user, itemID, true, errs)
	case path.ContactsCategory:
		bs, err = getItem(ctx, ac.Contacts(), user, itemID, true, errs)
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
		immutableID bool,
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
	immutableIDs bool,
	errs *fault.Bus,
) ([]byte, error) {
	sp, _, err := itm.GetItem(ctx, user, itemID, immutableIDs, errs)
	if err != nil {
		return nil, clues.Wrap(err, "getting item")
	}

	return itm.Serialize(ctx, sp, user, itemID)
}
