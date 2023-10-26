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
	gmock "github.com/alcionai/corso/src/internal/m365/graph/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// Gockable client
// ---------------------------------------------------------------------------

// GockClient produces a new exchange api client that can be
// mocked using gock.
func gockClient(creds account.M365Config) (api.Client, error) {
	s, err := gmock.NewService(creds)
	if err != nil {
		return api.Client{}, err
	}

	li, err := gmock.NewService(creds, graph.NoTimeout())
	if err != nil {
		return api.Client{}, err
	}

	return api.Client{
		Credentials: creds,
		Stable:      s,
		LargeItem:   li,
	}, nil
}

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
	merr.SetMessage(&code) // sdk expect message to be available
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

func requireParseableToMap(t *testing.T, thing serialization.Parsable) map[string]any {
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
	email             string
	driveID           string
	driveRootFolderID string
	testContainerID   string
}

type intgTesterSetup struct {
	ac           api.Client
	gockAC       api.Client
	user         ids
	site         ids
	group        ids
	nonTeamGroup ids // group which does not have an associated team
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

	its.gockAC, err = gockClient(creds)
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
	its.group.email = tconfig.M365TeamEmail(t)

	its.nonTeamGroup.id = tconfig.M365GroupID(t)

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
