package models
import (
    "errors"
)
// Provides operations to call the add method.
type MacOSSoftwareUpdateDelayPolicy int

const (
    // Software update delays will not be enforced.
    NONE_MACOSSOFTWAREUPDATEDELAYPOLICY MacOSSoftwareUpdateDelayPolicy = iota
    // Force delays for OS software updates.
    DELAYOSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
    // Force delays for app software updates.
    DELAYAPPUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
    // Sentinel member for cases where the client cannot handle the new enum values.
    UNKNOWNFUTUREVALUE_MACOSSOFTWAREUPDATEDELAYPOLICY
    // Force delays for major OS software updates.
    DELAYMAJOROSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
)

func (i MacOSSoftwareUpdateDelayPolicy) String() string {
    return []string{"none", "delayOSUpdateVisibility", "delayAppUpdateVisibility", "unknownFutureValue", "delayMajorOsUpdateVisibility"}[i]
}
func ParseMacOSSoftwareUpdateDelayPolicy(v string) (interface{}, error) {
    result := NONE_MACOSSOFTWAREUPDATEDELAYPOLICY
    switch v {
        case "none":
            result = NONE_MACOSSOFTWAREUPDATEDELAYPOLICY
        case "delayOSUpdateVisibility":
            result = DELAYOSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
        case "delayAppUpdateVisibility":
            result = DELAYAPPUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MACOSSOFTWAREUPDATEDELAYPOLICY
        case "delayMajorOsUpdateVisibility":
            result = DELAYMAJOROSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
        default:
            return 0, errors.New("Unknown MacOSSoftwareUpdateDelayPolicy value: " + v)
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateDelayPolicy(values []MacOSSoftwareUpdateDelayPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
