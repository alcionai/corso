// get_item.go is a source file designed to retrieve an m365 object from an
// existing M365 account. Data displayed is representative of the current
// serialization abstraction versioning used by Microsoft Graph and stored by Corso.

package onedrive

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const downloadURLKey = "@microsoft.graph.downloadUrl"

// Required inputs from user for command execution
var (
	user, tenant, m365ID string
)

func AddCommands(parent *cobra.Command) {
	exCmd := &cobra.Command{
		Use:   "onedrive",
		Short: "Get an M365ID item",
		RunE:  handleOneDriveCmd,
	}

	fs := exCmd.PersistentFlags()
	fs.StringVar(&m365ID, "id", "", "m365 identifier for object")
	fs.StringVar(&user, "user", "", "m365 user id of M365 user")
	fs.StringVar(&tenant, "tenant", "", "m365 identifier for the tenant")

	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("user"))
	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("id"))

	parent.AddCommand(exCmd)
}

func handleOneDriveCmd(cmd *cobra.Command, args []string) error {
	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	tid := str.First(tenant, os.Getenv(account.AzureTenantID))

	ctx := clues.Add(
		cmd.Context(),
		"item_id", m365ID,
		"resource_owner", user,
		"tenant", tid)

	// get account info
	creds := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: tid,
	}

	// todo: swap to drive api client, when finished.
	adpt, err := graph.CreateAdapter(tid, creds.AzureClientID, creds.AzureClientSecret)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "creating graph adapter"))
	}

	svc := graph.NewService(adpt)
	gr := graph.NewNoTimeoutHTTPWrapper()

	err = runDisplayM365JSON(ctx, svc, gr, creds, user, m365ID)
	if err != nil {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		return Only(ctx, clues.Wrap(err, "getting item"))
	}

	return nil
}

type itemData struct {
	Size int `json:"size"`
}

type itemPrintable struct {
	Info        json.RawMessage `json:"info"`
	Permissions json.RawMessage `json:"permissions"`
	Data        itemData        `json:"data"`
}

func (i itemPrintable) MinimumPrintable() any {
	return i
}

func runDisplayM365JSON(
	ctx context.Context,
	srv graph.Servicer,
	gr graph.Requester,
	creds account.M365Config,
	userID, itemID string,
) error {
	drive, err := api.GetUsersDefaultDrive(ctx, srv, userID)
	if err != nil {
		return err
	}

	driveID := ptr.Val(drive.GetId())

	it := itemPrintable{}

	item, err := api.GetDriveItem(ctx, srv, driveID, itemID)
	if err != nil {
		return err
	}

	if item != nil {
		content, err := getDriveItemContent(ctx, gr, item)
		if err != nil {
			return err
		}

		// We could get size from item.GetSize(), but the
		// getDriveItemContent call is to ensure that we are able to
		// download the file.
		it.Data.Size = len(content)
	}

	sInfo, err := serializeObject(item)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(sInfo), &it.Info)
	if err != nil {
		return err
	}

	perms, err := api.GetItemPermission(ctx, srv, driveID, itemID)
	if err != nil {
		return err
	}

	sPerms, err := serializeObject(perms)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(sPerms), &it.Permissions)
	if err != nil {
		return err
	}

	PrettyJSON(ctx, it)

	return nil
}

func serializeObject(data serialization.Parsable) (string, error) {
	sw := kjson.NewJsonSerializationWriter()

	err := sw.WriteObjectValue("", data)
	if err != nil {
		return "", clues.Wrap(err, "writing serializing info")
	}

	content, err := sw.GetSerializedContent()
	if err != nil {
		return "", clues.Wrap(err, "getting serializing info")
	}

	return string(content), err
}

func getDriveItemContent(
	ctx context.Context,
	gr graph.Requester,
	item models.DriveItemable,
) ([]byte, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return nil, clues.New("retrieving download url")
	}

	resp, err := gr.Request(ctx, http.MethodGet, *url, nil, nil)
	if err != nil {
		return nil, clues.New("downloading item").With("error", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, clues.New("read downloaded item").With("error", err)
	}

	return content, nil
}
