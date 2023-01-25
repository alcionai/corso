package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AccessPackageCustomExtensionHandlerStatus int

const (
    REQUESTSENT_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS AccessPackageCustomExtensionHandlerStatus = iota
    REQUESTRECEIVED_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
    UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
)

func (i AccessPackageCustomExtensionHandlerStatus) String() string {
    return []string{"requestSent", "requestReceived", "unknownFutureValue"}[i]
}
func ParseAccessPackageCustomExtensionHandlerStatus(v string) (interface{}, error) {
    result := REQUESTSENT_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
    switch v {
        case "requestSent":
            result = REQUESTSENT_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
        case "requestReceived":
            result = REQUESTRECEIVED_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
        default:
            return 0, errors.New("Unknown AccessPackageCustomExtensionHandlerStatus value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageCustomExtensionHandlerStatus(values []AccessPackageCustomExtensionHandlerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
