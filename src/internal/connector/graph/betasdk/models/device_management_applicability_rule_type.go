package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManagementApplicabilityRuleType int

const (
    // Include
    INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE DeviceManagementApplicabilityRuleType = iota
    // Exclude
    EXCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
)

func (i DeviceManagementApplicabilityRuleType) String() string {
    return []string{"include", "exclude"}[i]
}
func ParseDeviceManagementApplicabilityRuleType(v string) (interface{}, error) {
    result := INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
    switch v {
        case "include":
            result = INCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
        case "exclude":
            result = EXCLUDE_DEVICEMANAGEMENTAPPLICABILITYRULETYPE
        default:
            return 0, errors.New("Unknown DeviceManagementApplicabilityRuleType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementApplicabilityRuleType(values []DeviceManagementApplicabilityRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
