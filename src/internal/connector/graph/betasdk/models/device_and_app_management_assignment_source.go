package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceAndAppManagementAssignmentSource int

const (
    // Direct indicates a direct assignment.
    DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE DeviceAndAppManagementAssignmentSource = iota
    // PolicySets indicates assignment was made via PolicySet assignment.
    POLICYSETS_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
)

func (i DeviceAndAppManagementAssignmentSource) String() string {
    return []string{"direct", "policySets"}[i]
}
func ParseDeviceAndAppManagementAssignmentSource(v string) (interface{}, error) {
    result := DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
    switch v {
        case "direct":
            result = DIRECT_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
        case "policySets":
            result = POLICYSETS_DEVICEANDAPPMANAGEMENTASSIGNMENTSOURCE
        default:
            return 0, errors.New("Unknown DeviceAndAppManagementAssignmentSource value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAndAppManagementAssignmentSource(values []DeviceAndAppManagementAssignmentSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
