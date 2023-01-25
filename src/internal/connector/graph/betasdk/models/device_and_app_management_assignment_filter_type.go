package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceAndAppManagementAssignmentFilterType int

const (
    // Default value. Do not use.
    NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE DeviceAndAppManagementAssignmentFilterType = iota
    // Indicates in-filter, rule matching will offer the payload to devices.
    INCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
    // Indicates out-filter, rule matching will not offer the payload to devices.
    EXCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
)

func (i DeviceAndAppManagementAssignmentFilterType) String() string {
    return []string{"none", "include", "exclude"}[i]
}
func ParseDeviceAndAppManagementAssignmentFilterType(v string) (interface{}, error) {
    result := NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
    switch v {
        case "none":
            result = NONE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
        case "include":
            result = INCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
        case "exclude":
            result = EXCLUDE_DEVICEANDAPPMANAGEMENTASSIGNMENTFILTERTYPE
        default:
            return 0, errors.New("Unknown DeviceAndAppManagementAssignmentFilterType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAndAppManagementAssignmentFilterType(values []DeviceAndAppManagementAssignmentFilterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
