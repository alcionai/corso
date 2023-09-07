package api_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

// ---------------------------------------------------------------------------
// Intercepting calls with Gock
// ---------------------------------------------------------------------------

const graphAPIHostURL = "https://graph.microsoft.com"

func v1APIURLPath(parts ...string) string {
	return strings.Join(append([]string{"/v1.0"}, parts...), "/")
}

func interceptV1Path(pathParts ...string) *gock.Request {
	return gock.New(graphAPIHostURL).Get(v1APIURLPath(pathParts...))
}

func odErr(code string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func odErrMsg(code, message string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	merr.SetMessage(&message)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func parseableToMap(t *testing.T, thing serialization.Parsable) map[string]any {
	sw := kjson.NewJsonSerializationWriter()

	err := sw.WriteObjectValue("", thing)
	require.NoError(t, err, "serialize")

	content, err := sw.GetSerializedContent()
	require.NoError(t, err, "deserialize")

	var out map[string]any
	err = json.Unmarshal([]byte(content), &out)
	require.NoError(t, err, "unmarshall")

	return out
}

// ---------------------------------------------------------------------------
// Suite Setup
// ---------------------------------------------------------------------------

type ids struct {
	id                string
	driveID           string
	driveRootFolderID string
	testContainerID   string
}

type intgTesterSetup struct {
	ac     api.Client
	gockAC api.Client
	user   ids
	site   ids
	group  ids
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	a := tconfig.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	its.gockAC, err = mock.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	// user drive

	its.user.id = tconfig.M365UserID(t)

	userDrive, err := its.ac.Users().GetDefaultDrive(ctx, its.user.id)
	require.NoError(t, err, clues.ToCore(err))

	its.user.driveID = ptr.Val(userDrive.GetId())

	userDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.user.driveID)
	require.NoError(t, err, clues.ToCore(err))

	its.user.driveRootFolderID = ptr.Val(userDriveRootFolder.GetId())

	// site

	its.site.id = tconfig.M365SiteID(t)

	siteDrive, err := its.ac.Sites().GetDefaultDrive(ctx, its.site.id)
	require.NoError(t, err, clues.ToCore(err))

	its.site.driveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.site.driveID)
	require.NoError(t, err, clues.ToCore(err))

	its.site.driveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	// groups/teams

	// use of the TeamID is intentional here, so that we are assured
	// the group has full usage of the teams api.
	its.group.id = tconfig.M365TeamID(t)

	channel, err := its.ac.Channels().
		GetChannelByName(
			ctx,
			its.group.id,
			"Test")
	require.NoError(t, err, clues.ToCore(err))
	require.Equal(t, "Test", ptr.Val(channel.GetDisplayName()))

	its.group.testContainerID = ptr.Val(channel.GetId())

	return its
}
