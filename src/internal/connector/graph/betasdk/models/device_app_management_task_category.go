package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceAppManagementTaskCategory int

const (
    // Unknown source.
    UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY DeviceAppManagementTaskCategory = iota
    // Windows Defender ATP Threat & Vulnerability Management.
    ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY
)

func (i DeviceAppManagementTaskCategory) String() string {
    return []string{"unknown", "advancedThreatProtection"}[i]
}
func ParseDeviceAppManagementTaskCategory(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEAPPMANAGEMENTTASKCATEGORY
        case "advancedThreatProtection":
            result = ADVANCEDTHREATPROTECTION_DEVICEAPPMANAGEMENTTASKCATEGORY
        default:
            return 0, errors.New("Unknown DeviceAppManagementTaskCategory value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAppManagementTaskCategory(values []DeviceAppManagementTaskCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
