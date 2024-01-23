package testdata

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

func ODataErr(code string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	// graph sdk expects the message to be available
	merr.SetMessage(&code)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func ODataErrWithMsg(code, message string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	merr.SetMessage(&message)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func ODataErrWithStatus(status int, code string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	// graph sdk expects the message to be available
	merr.SetMessage(&code)
	odErr.SetErrorEscaped(merr)
	odErr.SetStatusCode(status)

	return odErr
}

func ODataInner(innerCode string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	inerr := odataerrors.NewInnerError()
	merr := odataerrors.NewMainError()

	inerr.SetAdditionalData(map[string]any{
		"code": innerCode,
	})
	inerr.SetClientRequestId(ptr.To(uuid.NewString()))
	inerr.SetOdataType(ptr.To("@odata.type"))
	inerr.SetRequestId(ptr.To("req-id"))

	merr.SetInnerError(inerr)
	merr.SetCode(ptr.To("main code"))
	merr.SetMessage(ptr.To("main message"))

	odErr.SetErrorEscaped(merr)

	return odErr
}

func ODataErrWithAPIResponse(
	code string,
	respCode int,
) *odataerrors.ODataError {
	odErr := ODataErr(code)

	// TODO(pandeyabs): Expand this function to set response headers as well.
	odErr.SetStatusCode(respCode)

	return odErr
}

func ParseableToMap(t *testing.T, thing serialization.Parsable) map[string]any {
	sw := kjson.NewJsonSerializationWriter()

	err := sw.WriteObjectValue("", thing)
	require.NoError(t, err, "serialize parsable")

	content, err := sw.GetSerializedContent()
	require.NoError(t, err, "deserialize parsable")

	var out map[string]any
	err = json.Unmarshal([]byte(content), &out)
	require.NoError(t, err, "unmarshal parsable")

	return out
}

func ParseableToReader(t *testing.T, thing serialization.Parsable) (int64, io.ReadCloser) {
	sw := kjson.NewJsonSerializationWriter()

	err := sw.WriteObjectValue("", thing)
	require.NoError(t, err, "serialize parsable")

	content, err := sw.GetSerializedContent()
	require.NoError(t, err, "deserialize parsable")

	return int64(len(content)), io.NopCloser(bytes.NewReader(content))
}
