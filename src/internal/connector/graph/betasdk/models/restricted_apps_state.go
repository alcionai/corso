package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RestrictedAppsState int

const (
    // Prohibited apps
    PROHIBITEDAPPS_RESTRICTEDAPPSSTATE RestrictedAppsState = iota
    // Not approved apps
    NOTAPPROVEDAPPS_RESTRICTEDAPPSSTATE
)

func (i RestrictedAppsState) String() string {
    return []string{"prohibitedApps", "notApprovedApps"}[i]
}
func ParseRestrictedAppsState(v string) (interface{}, error) {
    result := PROHIBITEDAPPS_RESTRICTEDAPPSSTATE
    switch v {
        case "prohibitedApps":
            result = PROHIBITEDAPPS_RESTRICTEDAPPSSTATE
        case "notApprovedApps":
            result = NOTAPPROVEDAPPS_RESTRICTEDAPPSSTATE
        default:
            return 0, errors.New("Unknown RestrictedAppsState value: " + v)
    }
    return &result, nil
}
func SerializeRestrictedAppsState(values []RestrictedAppsState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
