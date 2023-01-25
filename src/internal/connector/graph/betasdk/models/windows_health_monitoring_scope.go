package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsHealthMonitoringScope int

const (
    // Undefined
    UNDEFINED_WINDOWSHEALTHMONITORINGSCOPE WindowsHealthMonitoringScope = iota
    // Basic events for windows device health monitoring
    HEALTHMONITORING_WINDOWSHEALTHMONITORINGSCOPE
    // Boot performance events
    BOOTPERFORMANCE_WINDOWSHEALTHMONITORINGSCOPE
    // Windows updates events
    WINDOWSUPDATES_WINDOWSHEALTHMONITORINGSCOPE
)

func (i WindowsHealthMonitoringScope) String() string {
    return []string{"undefined", "healthMonitoring", "bootPerformance", "windowsUpdates"}[i]
}
func ParseWindowsHealthMonitoringScope(v string) (interface{}, error) {
    result := UNDEFINED_WINDOWSHEALTHMONITORINGSCOPE
    switch v {
        case "undefined":
            result = UNDEFINED_WINDOWSHEALTHMONITORINGSCOPE
        case "healthMonitoring":
            result = HEALTHMONITORING_WINDOWSHEALTHMONITORINGSCOPE
        case "bootPerformance":
            result = BOOTPERFORMANCE_WINDOWSHEALTHMONITORINGSCOPE
        case "windowsUpdates":
            result = WINDOWSUPDATES_WINDOWSHEALTHMONITORINGSCOPE
        default:
            return 0, errors.New("Unknown WindowsHealthMonitoringScope value: " + v)
    }
    return &result, nil
}
func SerializeWindowsHealthMonitoringScope(values []WindowsHealthMonitoringScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
