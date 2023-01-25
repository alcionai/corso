package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManagementAutopilotPolicyComplianceStatus int

const (
    UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS DeviceManagementAutopilotPolicyComplianceStatus = iota
    COMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    INSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    NOTCOMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    NOTINSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    ERROR_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
)

func (i DeviceManagementAutopilotPolicyComplianceStatus) String() string {
    return []string{"unknown", "compliant", "installed", "notCompliant", "notInstalled", "error"}[i]
}
func ParseDeviceManagementAutopilotPolicyComplianceStatus(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "compliant":
            result = COMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "installed":
            result = INSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "notCompliant":
            result = NOTCOMPLIANT_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "notInstalled":
            result = NOTINSTALLED_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        case "error":
            result = ERROR_DEVICEMANAGEMENTAUTOPILOTPOLICYCOMPLIANCESTATUS
        default:
            return 0, errors.New("Unknown DeviceManagementAutopilotPolicyComplianceStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementAutopilotPolicyComplianceStatus(values []DeviceManagementAutopilotPolicyComplianceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
