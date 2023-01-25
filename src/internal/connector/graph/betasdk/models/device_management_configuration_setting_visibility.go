package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceManagementConfigurationSettingVisibility int

const (
    // Not visible
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY DeviceManagementConfigurationSettingVisibility = iota
    // Visible to setting catalog UX
    SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
    // Visible to template
    TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
)

func (i DeviceManagementConfigurationSettingVisibility) String() string {
    return []string{"none", "settingsCatalog", "template"}[i]
}
func ParseDeviceManagementConfigurationSettingVisibility(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        case "settingsCatalog":
            result = SETTINGSCATALOG_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        case "template":
            result = TEMPLATE_DEVICEMANAGEMENTCONFIGURATIONSETTINGVISIBILITY
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationSettingVisibility value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSettingVisibility(values []DeviceManagementConfigurationSettingVisibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
