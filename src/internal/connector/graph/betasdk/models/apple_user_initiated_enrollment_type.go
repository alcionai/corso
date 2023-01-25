package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AppleUserInitiatedEnrollmentType int

const (
    // Unknown enrollment type
    UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE AppleUserInitiatedEnrollmentType = iota
    // Device enrollment type
    DEVICE_APPLEUSERINITIATEDENROLLMENTTYPE
    // User enrollment type
    USER_APPLEUSERINITIATEDENROLLMENTTYPE
)

func (i AppleUserInitiatedEnrollmentType) String() string {
    return []string{"unknown", "device", "user"}[i]
}
func ParseAppleUserInitiatedEnrollmentType(v string) (interface{}, error) {
    result := UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_APPLEUSERINITIATEDENROLLMENTTYPE
        case "device":
            result = DEVICE_APPLEUSERINITIATEDENROLLMENTTYPE
        case "user":
            result = USER_APPLEUSERINITIATEDENROLLMENTTYPE
        default:
            return 0, errors.New("Unknown AppleUserInitiatedEnrollmentType value: " + v)
    }
    return &result, nil
}
func SerializeAppleUserInitiatedEnrollmentType(values []AppleUserInitiatedEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
