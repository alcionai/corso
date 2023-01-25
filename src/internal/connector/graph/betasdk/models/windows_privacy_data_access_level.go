package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsPrivacyDataAccessLevel int

const (
    // No access level specified, no intents. Device may behave either as in UserInControl or ForceAllow. It may depend on the privacy data been accessed, Windows versions and other factors.
    NOTCONFIGURED_WINDOWSPRIVACYDATAACCESSLEVEL WindowsPrivacyDataAccessLevel = iota
    // Apps will be allowed to access the specified privacy data.
    FORCEALLOW_WINDOWSPRIVACYDATAACCESSLEVEL
    // Apps will be denied to access specified privacy data.
    FORCEDENY_WINDOWSPRIVACYDATAACCESSLEVEL
    // Users will be prompted when apps try to access specified privacy data.
    USERINCONTROL_WINDOWSPRIVACYDATAACCESSLEVEL
)

func (i WindowsPrivacyDataAccessLevel) String() string {
    return []string{"notConfigured", "forceAllow", "forceDeny", "userInControl"}[i]
}
func ParseWindowsPrivacyDataAccessLevel(v string) (interface{}, error) {
    result := NOTCONFIGURED_WINDOWSPRIVACYDATAACCESSLEVEL
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWSPRIVACYDATAACCESSLEVEL
        case "forceAllow":
            result = FORCEALLOW_WINDOWSPRIVACYDATAACCESSLEVEL
        case "forceDeny":
            result = FORCEDENY_WINDOWSPRIVACYDATAACCESSLEVEL
        case "userInControl":
            result = USERINCONTROL_WINDOWSPRIVACYDATAACCESSLEVEL
        default:
            return 0, errors.New("Unknown WindowsPrivacyDataAccessLevel value: " + v)
    }
    return &result, nil
}
func SerializeWindowsPrivacyDataAccessLevel(values []WindowsPrivacyDataAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
