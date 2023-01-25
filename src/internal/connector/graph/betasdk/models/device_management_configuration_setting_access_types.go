package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementConfigurationSettingAccessTypes int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES DeviceManagementConfigurationSettingAccessTypes = iota
    ADD_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    COPY_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    DELETE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    GET_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    REPLACE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
)

func (i DeviceManagementConfigurationSettingAccessTypes) String() string {
    return []string{"none", "add", "copy", "delete", "get", "replace", "execute"}[i]
}
func ParseDeviceManagementConfigurationSettingAccessTypes(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "add":
            result = ADD_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "copy":
            result = COPY_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "delete":
            result = DELETE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "get":
            result = GET_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "replace":
            result = REPLACE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        case "execute":
            result = EXECUTE_DEVICEMANAGEMENTCONFIGURATIONSETTINGACCESSTYPES
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationSettingAccessTypes value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationSettingAccessTypes(values []DeviceManagementConfigurationSettingAccessTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
