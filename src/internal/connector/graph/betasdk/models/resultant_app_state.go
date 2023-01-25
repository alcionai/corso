package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ResultantAppState int

const (
    // The application is not applicable.
    NOTAPPLICABLE_RESULTANTAPPSTATE ResultantAppState = iota
    // The application is installed with no errors.
    INSTALLED_RESULTANTAPPSTATE
    // The application failed to install.
    FAILED_RESULTANTAPPSTATE
    // The application is not installed.
    NOTINSTALLED_RESULTANTAPPSTATE
    // The application failed to uninstall.
    UNINSTALLFAILED_RESULTANTAPPSTATE
    // The installation of the application is in progress.
    PENDINGINSTALL_RESULTANTAPPSTATE
    // The status of the application is unknown.
    UNKNOWN_RESULTANTAPPSTATE
)

func (i ResultantAppState) String() string {
    return []string{"notApplicable", "installed", "failed", "notInstalled", "uninstallFailed", "pendingInstall", "unknown"}[i]
}
func ParseResultantAppState(v string) (interface{}, error) {
    result := NOTAPPLICABLE_RESULTANTAPPSTATE
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_RESULTANTAPPSTATE
        case "installed":
            result = INSTALLED_RESULTANTAPPSTATE
        case "failed":
            result = FAILED_RESULTANTAPPSTATE
        case "notInstalled":
            result = NOTINSTALLED_RESULTANTAPPSTATE
        case "uninstallFailed":
            result = UNINSTALLFAILED_RESULTANTAPPSTATE
        case "pendingInstall":
            result = PENDINGINSTALL_RESULTANTAPPSTATE
        case "unknown":
            result = UNKNOWN_RESULTANTAPPSTATE
        default:
            return 0, errors.New("Unknown ResultantAppState value: " + v)
    }
    return &result, nil
}
func SerializeResultantAppState(values []ResultantAppState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
