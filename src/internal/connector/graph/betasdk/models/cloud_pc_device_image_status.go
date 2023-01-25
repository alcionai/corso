package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcDeviceImageStatus int

const (
    PENDING_CLOUDPCDEVICEIMAGESTATUS CloudPcDeviceImageStatus = iota
    READY_CLOUDPCDEVICEIMAGESTATUS
    FAILED_CLOUDPCDEVICEIMAGESTATUS
)

func (i CloudPcDeviceImageStatus) String() string {
    return []string{"pending", "ready", "failed"}[i]
}
func ParseCloudPcDeviceImageStatus(v string) (interface{}, error) {
    result := PENDING_CLOUDPCDEVICEIMAGESTATUS
    switch v {
        case "pending":
            result = PENDING_CLOUDPCDEVICEIMAGESTATUS
        case "ready":
            result = READY_CLOUDPCDEVICEIMAGESTATUS
        case "failed":
            result = FAILED_CLOUDPCDEVICEIMAGESTATUS
        default:
            return 0, errors.New("Unknown CloudPcDeviceImageStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDeviceImageStatus(values []CloudPcDeviceImageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
