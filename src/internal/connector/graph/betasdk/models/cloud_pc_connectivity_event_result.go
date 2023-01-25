package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcConnectivityEventResult int

const (
    UNKNOWN_CLOUDPCCONNECTIVITYEVENTRESULT CloudPcConnectivityEventResult = iota
    SUCCESS_CLOUDPCCONNECTIVITYEVENTRESULT
    FAILURE_CLOUDPCCONNECTIVITYEVENTRESULT
    UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIVITYEVENTRESULT
)

func (i CloudPcConnectivityEventResult) String() string {
    return []string{"unknown", "success", "failure", "unknownFutureValue"}[i]
}
func ParseCloudPcConnectivityEventResult(v string) (interface{}, error) {
    result := UNKNOWN_CLOUDPCCONNECTIVITYEVENTRESULT
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDPCCONNECTIVITYEVENTRESULT
        case "success":
            result = SUCCESS_CLOUDPCCONNECTIVITYEVENTRESULT
        case "failure":
            result = FAILURE_CLOUDPCCONNECTIVITYEVENTRESULT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIVITYEVENTRESULT
        default:
            return 0, errors.New("Unknown CloudPcConnectivityEventResult value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcConnectivityEventResult(values []CloudPcConnectivityEventResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
