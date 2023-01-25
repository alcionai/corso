package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidManagedStoreAccountEnrollmentTarget int

const (
    NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET AndroidManagedStoreAccountEnrollmentTarget = iota
    ALL_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    TARGETED_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
)

func (i AndroidManagedStoreAccountEnrollmentTarget) String() string {
    return []string{"none", "all", "targeted", "targetedAsEnrollmentRestrictions"}[i]
}
func ParseAndroidManagedStoreAccountEnrollmentTarget(v string) (interface{}, error) {
    result := NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    switch v {
        case "none":
            result = NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        case "all":
            result = ALL_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        case "targeted":
            result = TARGETED_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        case "targetedAsEnrollmentRestrictions":
            result = TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAccountEnrollmentTarget value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAccountEnrollmentTarget(values []AndroidManagedStoreAccountEnrollmentTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
