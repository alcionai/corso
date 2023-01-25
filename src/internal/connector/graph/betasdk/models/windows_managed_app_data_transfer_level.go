package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsManagedAppDataTransferLevel int

const (
    // All apps.
    ALLAPPS_WINDOWSMANAGEDAPPDATATRANSFERLEVEL WindowsManagedAppDataTransferLevel = iota
    // No apps.
    NONE_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
)

func (i WindowsManagedAppDataTransferLevel) String() string {
    return []string{"allApps", "none"}[i]
}
func ParseWindowsManagedAppDataTransferLevel(v string) (interface{}, error) {
    result := ALLAPPS_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
    switch v {
        case "allApps":
            result = ALLAPPS_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
        case "none":
            result = NONE_WINDOWSMANAGEDAPPDATATRANSFERLEVEL
        default:
            return 0, errors.New("Unknown WindowsManagedAppDataTransferLevel value: " + v)
    }
    return &result, nil
}
func SerializeWindowsManagedAppDataTransferLevel(values []WindowsManagedAppDataTransferLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
