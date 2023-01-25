package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcConnectivityStatus int

const (
    UNKNOWN_CLOUDPCCONNECTIVITYSTATUS CloudPcConnectivityStatus = iota
    AVAILABLE_CLOUDPCCONNECTIVITYSTATUS
    AVAILABLEWITHWARNING_CLOUDPCCONNECTIVITYSTATUS
    UNAVAILABLE_CLOUDPCCONNECTIVITYSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIVITYSTATUS
)

func (i CloudPcConnectivityStatus) String() string {
    return []string{"unknown", "available", "availableWithWarning", "unavailable", "unknownFutureValue"}[i]
}
func ParseCloudPcConnectivityStatus(v string) (interface{}, error) {
    result := UNKNOWN_CLOUDPCCONNECTIVITYSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDPCCONNECTIVITYSTATUS
        case "available":
            result = AVAILABLE_CLOUDPCCONNECTIVITYSTATUS
        case "availableWithWarning":
            result = AVAILABLEWITHWARNING_CLOUDPCCONNECTIVITYSTATUS
        case "unavailable":
            result = UNAVAILABLE_CLOUDPCCONNECTIVITYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIVITYSTATUS
        default:
            return 0, errors.New("Unknown CloudPcConnectivityStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcConnectivityStatus(values []CloudPcConnectivityStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
