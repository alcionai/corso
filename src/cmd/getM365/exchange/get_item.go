// get_item.go is a source file designed to retrieve an m365 object from an
// existing M365 account. Data displayed is representative of the current
// serialization abstraction versioning used by Microsoft Graph and stored by Corso.

package exchange

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/utils"
	gutils "github.com/alcionai/corso/src/cmd/getM365/utils"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// Required inputs from user for command execution
var (
	user, tenant, m365ID, category string
)

func AddCommands(parent *cobra.Command) {
	exCmd := &cobra.Command{
		Use:   "exchange",
		Short: "Get a M365ID item JSON",
		RunE:  handleExchangeCmd,
	}

	fs := exCmd.PersistentFlags()
	fs.StringVar(&m365ID, "m365ID", "", "m365 identifier for object")
	fs.StringVar(&category, "category", "", "type of M365 data (contacts, email, events)")
	fs.StringVar(&user, "user", "", "m365 user id of M365 user")
	fs.StringVar(&tenant, "tenant", "", "m365 Tenant: m365 identifier for the tenant, not required if active in OS Environment")

	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("user"))
	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("m365ID"))
	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("category"))

	parent.AddCommand(exCmd)
}

func handleExchangeCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	_, creds, err := gutils.GetGC(ctx, tenant)
	if err != nil {
		return err
	}

	err = runDisplayM365JSON(ctx, creds, user, m365ID, fault.New(true))
	if err != nil {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true
		return errors.Wrapf(err, "unable to get from M365: %s", m365ID)
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
