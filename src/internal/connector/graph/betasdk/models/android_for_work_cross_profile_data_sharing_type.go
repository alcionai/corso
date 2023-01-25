package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidForWorkCrossProfileDataSharingType int

const (
    // Device default value, no intent.
    DEVICEDEFAULT_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE AndroidForWorkCrossProfileDataSharingType = iota
    // Prevent any sharing.
    PREVENTANY_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
    // Allow data sharing request from personal profile to work profile.
    ALLOWPERSONALTOWORK_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
    // No restrictions on sharing.
    NORESTRICTIONS_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
)

func (i AndroidForWorkCrossProfileDataSharingType) String() string {
    return []string{"deviceDefault", "preventAny", "allowPersonalToWork", "noRestrictions"}[i]
}
func ParseAndroidForWorkCrossProfileDataSharingType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
        case "preventAny":
            result = PREVENTANY_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
        case "allowPersonalToWork":
            result = ALLOWPERSONALTOWORK_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
        case "noRestrictions":
            result = NORESTRICTIONS_ANDROIDFORWORKCROSSPROFILEDATASHARINGTYPE
        default:
            return 0, errors.New("Unknown AndroidForWorkCrossProfileDataSharingType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkCrossProfileDataSharingType(values []AndroidForWorkCrossProfileDataSharingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
