package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceAppManagementTaskStatus int

const (
    // State is undefined.
    UNKNOWN_DEVICEAPPMANAGEMENTTASKSTATUS DeviceAppManagementTaskStatus = iota
    // The task is ready for review.
    PENDING_DEVICEAPPMANAGEMENTTASKSTATUS
    // The task has been accepted and is being worked on.
    ACTIVE_DEVICEAPPMANAGEMENTTASKSTATUS
    // The work is complete.
    COMPLETED_DEVICEAPPMANAGEMENTTASKSTATUS
    // The task was rejected.
    REJECTED_DEVICEAPPMANAGEMENTTASKSTATUS
)

func (i DeviceAppManagementTaskStatus) String() string {
    return []string{"unknown", "pending", "active", "completed", "rejected"}[i]
}
func ParseDeviceAppManagementTaskStatus(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEAPPMANAGEMENTTASKSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEAPPMANAGEMENTTASKSTATUS
        case "pending":
            result = PENDING_DEVICEAPPMANAGEMENTTASKSTATUS
        case "active":
            result = ACTIVE_DEVICEAPPMANAGEMENTTASKSTATUS
        case "completed":
            result = COMPLETED_DEVICEAPPMANAGEMENTTASKSTATUS
        case "rejected":
            result = REJECTED_DEVICEAPPMANAGEMENTTASKSTATUS
        default:
            return 0, errors.New("Unknown DeviceAppManagementTaskStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceAppManagementTaskStatus(values []DeviceAppManagementTaskStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
