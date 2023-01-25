package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidForWorkEnrollmentTarget int

const (
    NONE_ANDROIDFORWORKENROLLMENTTARGET AndroidForWorkEnrollmentTarget = iota
    ALL_ANDROIDFORWORKENROLLMENTTARGET
    TARGETED_ANDROIDFORWORKENROLLMENTTARGET
    TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDFORWORKENROLLMENTTARGET
)

func (i AndroidForWorkEnrollmentTarget) String() string {
    return []string{"none", "all", "targeted", "targetedAsEnrollmentRestrictions"}[i]
}
func ParseAndroidForWorkEnrollmentTarget(v string) (interface{}, error) {
    result := NONE_ANDROIDFORWORKENROLLMENTTARGET
    switch v {
        case "none":
            result = NONE_ANDROIDFORWORKENROLLMENTTARGET
        case "all":
            result = ALL_ANDROIDFORWORKENROLLMENTTARGET
        case "targeted":
            result = TARGETED_ANDROIDFORWORKENROLLMENTTARGET
        case "targetedAsEnrollmentRestrictions":
            result = TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDFORWORKENROLLMENTTARGET
        default:
            return 0, errors.New("Unknown AndroidForWorkEnrollmentTarget value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkEnrollmentTarget(values []AndroidForWorkEnrollmentTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
