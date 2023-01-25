package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PersonalProfilePersonalPlayStoreMode int

const (
    // Not configured.
    NOTCONFIGURED_PERSONALPROFILEPERSONALPLAYSTOREMODE PersonalProfilePersonalPlayStoreMode = iota
    // Blocked Apps.
    BLOCKEDAPPS_PERSONALPROFILEPERSONALPLAYSTOREMODE
    // Allowed Apps.
    ALLOWEDAPPS_PERSONALPROFILEPERSONALPLAYSTOREMODE
)

func (i PersonalProfilePersonalPlayStoreMode) String() string {
    return []string{"notConfigured", "blockedApps", "allowedApps"}[i]
}
func ParsePersonalProfilePersonalPlayStoreMode(v string) (interface{}, error) {
    result := NOTCONFIGURED_PERSONALPROFILEPERSONALPLAYSTOREMODE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_PERSONALPROFILEPERSONALPLAYSTOREMODE
        case "blockedApps":
            result = BLOCKEDAPPS_PERSONALPROFILEPERSONALPLAYSTOREMODE
        case "allowedApps":
            result = ALLOWEDAPPS_PERSONALPROFILEPERSONALPLAYSTOREMODE
        default:
            return 0, errors.New("Unknown PersonalProfilePersonalPlayStoreMode value: " + v)
    }
    return &result, nil
}
func SerializePersonalProfilePersonalPlayStoreMode(values []PersonalProfilePersonalPlayStoreMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
