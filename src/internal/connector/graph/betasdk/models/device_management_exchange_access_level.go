package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementExchangeAccessLevel int

const (
    // No device access rule has been configured in Exchange.
    NONE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL DeviceManagementExchangeAccessLevel = iota
    // Allow the device access to Exchange.
    ALLOW_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
    // Block the device from accessing Exchange.
    BLOCK_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
    // Quarantine the device in Exchange.
    QUARANTINE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
)

func (i DeviceManagementExchangeAccessLevel) String() string {
    return []string{"none", "allow", "block", "quarantine"}[i]
}
func ParseDeviceManagementExchangeAccessLevel(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        case "allow":
            result = ALLOW_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        case "block":
            result = BLOCK_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        case "quarantine":
            result = QUARANTINE_DEVICEMANAGEMENTEXCHANGEACCESSLEVEL
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeAccessLevel value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeAccessLevel(values []DeviceManagementExchangeAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
