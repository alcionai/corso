package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceAppManagementTaskPriority int

const (
    // No priority set.
    NONE_DEVICEAPPMANAGEMENTTASKPRIORITY DeviceAppManagementTaskPriority = iota
    // High priority.
    HIGH_DEVICEAPPMANAGEMENTTASKPRIORITY
    // Low priority.
    LOW_DEVICEAPPMANAGEMENTTASKPRIORITY
)

func (i DeviceAppManagementTaskPriority) String() string {
    return []string{"none", "high", "low"}[i]
}
func ParseDeviceAppManagementTaskPriority(v string) (interface{}, error) {
    result := NONE_DEVICEAPPMANAGEMENTTASKPRIORITY
    switch v {
        case "none":
            result = NONE_DEVICEAPPMANAGEMENTTASKPRIORITY
        case "high":
            result = HIGH_DEVICEAPPMANAGEMENTTASKPRIORITY
        case "low":
            result = LOW_DEVICEAPPMANAGEMENTTASKPRIORITY
        default:
            return 0, errors.New("Unknown DeviceAppManagementTaskPriority value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAppManagementTaskPriority(values []DeviceAppManagementTaskPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
