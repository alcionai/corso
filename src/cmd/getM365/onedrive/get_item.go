// get_item.go is a source file designed to retrieve an m365 object from an
// existing M365 account. Data displayed is representative of the current
// serialization abstraction versioning used by Microsoft Graph and stored by Corso.

package onedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/cli/utils"
	gutils "github.com/alcionai/corso/src/cmd/getM365/utils"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/api"
	"github.com/alcionai/corso/src/pkg/account"
)

const downloadURLKey = "@microsoft.graph.downloadUrl"

// Required inputs from user for command execution
var (
	user, tenant, m365ID string
)

func AddCommands(parent *cobra.Command) {
	exCmd := &cobra.Command{
		Use:   "onedrive",
		Short: "Get a M365ID item",
		RunE:  handleOneDriveCmd,
	}

	fs := exCmd.PersistentFlags()
	fs.StringVar(&m365ID, "m365ID", "", "m365 identifier for object")
	fs.StringVar(&user, "user", "", "m365 user id of M365 user")
	fs.StringVar(&tenant, "tenant", "", "m365 identifier for the tenant")

	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("user"))
	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("tenant"))
	cobra.CheckErr(exCmd.MarkPersistentFlagRequired("m365ID"))

	parent.AddCommand(exCmd)
}

func handleOneDriveCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if utils.HasNoFlagsAndShownHelp(cmd) {
		return nil
	}

	gc, creds, err := gutils.GetGC(ctx, tenant)
	if err != nil {
		return err
	}

	err = runDisplayM365JSON(ctx, gc, creds, user, m365ID)
	if err != nil {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		return errors.Wrapf(err, "unable to fetch from M365: %s", m365ID)
	}

	return nil
}

type itemData struct {
	Size int `json:"size"`
}

type item struct {
	Info        models.DriveItemable
	Permissions models.PermissionCollectionResponseable
	Data        itemData
}

func runDisplayM365JSON(
	ctx context.Context,
	gc *connector.GraphConnector,
	creds account.M365Config,
	user, itemID string,
) error {
	driveID, err := getDriveID(ctx, gc.Service, user)
	if err != nil {
		return err
	}

	it := item{}

	item, err := api.GetDriveItem(ctx, gc.Service, driveID, itemID)
	if err != nil {
		return err
	}

	it.Info = item

	perms, err := api.GetItemPermission(ctx, gc.Service, driveID, itemID)
	if err != nil {
		return err
	}

	it.Permissions = perms

	if item != nil {
		content, err := getDriveItemContent(item)
		if err != nil {
			return err
		}

		it.Data.Size = len(content)
	}

	out, err := serialize(it)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}

func serialize(data item) (string, error) {
	var (
		info  = "{}"
		perms = "{}"
		err   error
	)

	if data.Info != nil {
		info, err = serializeObject(data.Info)
		if err != nil {
			return "", err
		}
	}

	if data.Permissions != nil {
		perms, err = serializeObject(data.Permissions)
		if err != nil {
			return "", err
		}
	}

	file, err := json.Marshal(data.Data)
	if err != nil {
		return "", err
	}

	return `{"info":` + info + `, "permissions":` + perms + `, "data":` + string(file) + `}`, nil
}

func serializeObject(data serialization.Parsable) (string, error) {
	sw := kw.NewJsonSerializationWriter()

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

func getDriveID(
	ctx context.Context,
	service graph.Servicer,
	userID string,
) (string, error) {
	//revive:enable:context-as-argument
	d, err := service.Client().UsersById(userID).Drive().Get(ctx, nil)
	if err != nil {
		err = graph.Wrap(ctx, err, "retrieving drive")
		return "", err
	}

	id := ptr.Val(d.GetId())

	return id, nil
}

func getDriveItemContent(item models.DriveItemable) ([]byte, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return nil, clues.New("get download url")
	}

	req, err := http.NewRequest(http.MethodGet, *url, nil)
	if err != nil {
		return nil, clues.New("create download request").With("error", err)
	}

	hc := graph.HTTPClient(graph.NoTimeout())

	resp, err := hc.Do(req)
	if err != nil {
		return nil, clues.New("download item").With("error", err)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, clues.New("read downloaded item").With("error", err)
	}

	return content, nil
}
