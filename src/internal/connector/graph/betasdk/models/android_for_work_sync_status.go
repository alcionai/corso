package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AndroidForWorkSyncStatus int

const (
    SUCCESS_ANDROIDFORWORKSYNCSTATUS AndroidForWorkSyncStatus = iota
    CREDENTIALSNOTVALID_ANDROIDFORWORKSYNCSTATUS
    ANDROIDFORWORKAPIERROR_ANDROIDFORWORKSYNCSTATUS
    MANAGEMENTSERVICEERROR_ANDROIDFORWORKSYNCSTATUS
    UNKNOWNERROR_ANDROIDFORWORKSYNCSTATUS
    NONE_ANDROIDFORWORKSYNCSTATUS
)

func (i AndroidForWorkSyncStatus) String() string {
    return []string{"success", "credentialsNotValid", "androidForWorkApiError", "managementServiceError", "unknownError", "none"}[i]
}
func ParseAndroidForWorkSyncStatus(v string) (interface{}, error) {
    result := SUCCESS_ANDROIDFORWORKSYNCSTATUS
    switch v {
        case "success":
            result = SUCCESS_ANDROIDFORWORKSYNCSTATUS
        case "credentialsNotValid":
            result = CREDENTIALSNOTVALID_ANDROIDFORWORKSYNCSTATUS
        case "androidForWorkApiError":
            result = ANDROIDFORWORKAPIERROR_ANDROIDFORWORKSYNCSTATUS
        case "managementServiceError":
            result = MANAGEMENTSERVICEERROR_ANDROIDFORWORKSYNCSTATUS
        case "unknownError":
            result = UNKNOWNERROR_ANDROIDFORWORKSYNCSTATUS
        case "none":
            result = NONE_ANDROIDFORWORKSYNCSTATUS
        default:
            return 0, errors.New("Unknown AndroidForWorkSyncStatus value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkSyncStatus(values []AndroidForWorkSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
