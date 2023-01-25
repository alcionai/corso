package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceConfigAssignmentIntent int

const (
    // Ensure that the configuration profile is applied to the devices in the assignment.
    APPLY_DEVICECONFIGASSIGNMENTINTENT DeviceConfigAssignmentIntent = iota
    // Ensure that the configuration profile is removed from devices that have previously installed the configuration profile.
    REMOVE_DEVICECONFIGASSIGNMENTINTENT
)

func (i DeviceConfigAssignmentIntent) String() string {
    return []string{"apply", "remove"}[i]
}
func ParseDeviceConfigAssignmentIntent(v string) (interface{}, error) {
    result := APPLY_DEVICECONFIGASSIGNMENTINTENT
    switch v {
        case "apply":
            result = APPLY_DEVICECONFIGASSIGNMENTINTENT
        case "remove":
            result = REMOVE_DEVICECONFIGASSIGNMENTINTENT
        default:
            return 0, errors.New("Unknown DeviceConfigAssignmentIntent value: " + v)
    }
    return &result, nil
}
func SerializeDeviceConfigAssignmentIntent(values []DeviceConfigAssignmentIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
