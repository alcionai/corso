package api

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/require"
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
