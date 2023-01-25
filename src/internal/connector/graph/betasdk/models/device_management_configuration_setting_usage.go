package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementConfigurationSettingUsage int

const (
    // No setting type specified
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE DeviceManagementConfigurationSettingUsage = iota
    // Configuration setting
    CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
    // Compliance setting
    COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
)

func (i DeviceManagementConfigurationSettingUsage) String() string {
    return []string{"none", "configuration", "compliance"}[i]
}
func ParseDeviceManagementConfigurationSettingUsage(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        case "configuration":
            result = CONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        case "compliance":
            result = COMPLIANCE_DEVICEMANAGEMENTCONFIGURATIONSETTINGUSAGE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationSettingUsage value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSettingUsage(values []DeviceManagementConfigurationSettingUsage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
