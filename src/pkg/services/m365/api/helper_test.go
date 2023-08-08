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
	odErr.SetError(merr)

	return odErr
}

func odErrMsg(code, message string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	merr.SetMessage(&message)
	odErr.SetError(merr)

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

type intgTesterSetup struct {
	ac                    api.Client
	gockAC                api.Client
	userID                string
	userDriveID           string
	userDriveRootFolderID string
	siteID                string
	siteDriveID           string
	siteDriveRootFolderID string
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	a := tconfig.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.ac, err = api.NewClient(creds, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	its.gockAC, err = mock.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	// user drive

	its.userID = tconfig.M365UserID(t)

	userDrive, err := its.ac.Users().GetDefaultDrive(ctx, its.userID)
	require.NoError(t, err, clues.ToCore(err))

	its.userDriveID = ptr.Val(userDrive.GetId())

	userDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.userDriveID)
	require.NoError(t, err, clues.ToCore(err))

	its.userDriveRootFolderID = ptr.Val(userDriveRootFolder.GetId())

	its.siteID = tconfig.M365SiteID(t)

	// site

	siteDrive, err := its.ac.Sites().GetDefaultDrive(ctx, its.siteID)
	require.NoError(t, err, clues.ToCore(err))

	its.siteDriveID = ptr.Val(siteDrive.GetId())

	siteDriveRootFolder, err := its.ac.Drives().GetRootFolder(ctx, its.siteDriveID)
	require.NoError(t, err, clues.ToCore(err))

	its.siteDriveRootFolderID = ptr.Val(siteDriveRootFolder.GetId())

	return its
}
