package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementConfigurationDeviceMode int

const (
    // No Device Mode specified
    NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE DeviceManagementConfigurationDeviceMode = iota
    // Device must be in kiosk mode for this setting to apply
    KIOSK_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
)

func (i DeviceManagementConfigurationDeviceMode) String() string {
    return []string{"none", "kiosk"}[i]
}
func ParseDeviceManagementConfigurationDeviceMode(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
        case "kiosk":
            result = KIOSK_DEVICEMANAGEMENTCONFIGURATIONDEVICEMODE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationDeviceMode value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationDeviceMode(values []DeviceManagementConfigurationDeviceMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
