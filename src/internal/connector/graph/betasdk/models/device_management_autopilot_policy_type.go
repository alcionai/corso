package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementAutopilotPolicyType int

const (
    UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE DeviceManagementAutopilotPolicyType = iota
    APPLICATION_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
    APPMODEL_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
    CONFIGURATIONPOLICY_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
)

func (i DeviceManagementAutopilotPolicyType) String() string {
    return []string{"unknown", "application", "appModel", "configurationPolicy"}[i]
}
func ParseDeviceManagementAutopilotPolicyType(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        case "application":
            result = APPLICATION_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        case "appModel":
            result = APPMODEL_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        case "configurationPolicy":
            result = CONFIGURATIONPOLICY_DEVICEMANAGEMENTAUTOPILOTPOLICYTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementAutopilotPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementAutopilotPolicyType(values []DeviceManagementAutopilotPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
