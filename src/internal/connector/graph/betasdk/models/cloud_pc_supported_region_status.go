package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcSupportedRegionStatus int

const (
    AVAILABLE_CLOUDPCSUPPORTEDREGIONSTATUS CloudPcSupportedRegionStatus = iota
    RESTRICTED_CLOUDPCSUPPORTEDREGIONSTATUS
    UNAVAILABLE_CLOUDPCSUPPORTEDREGIONSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCSUPPORTEDREGIONSTATUS
)

func (i CloudPcSupportedRegionStatus) String() string {
    return []string{"available", "restricted", "unavailable", "unknownFutureValue"}[i]
}
func ParseCloudPcSupportedRegionStatus(v string) (interface{}, error) {
    result := AVAILABLE_CLOUDPCSUPPORTEDREGIONSTATUS
    switch v {
        case "available":
            result = AVAILABLE_CLOUDPCSUPPORTEDREGIONSTATUS
        case "restricted":
            result = RESTRICTED_CLOUDPCSUPPORTEDREGIONSTATUS
        case "unavailable":
            result = UNAVAILABLE_CLOUDPCSUPPORTEDREGIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSUPPORTEDREGIONSTATUS
        default:
            return 0, errors.New("Unknown CloudPcSupportedRegionStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcSupportedRegionStatus(values []CloudPcSupportedRegionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
