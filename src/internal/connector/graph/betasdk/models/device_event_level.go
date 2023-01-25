package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceEventLevel int

const (
    // Indicates that the device event level is none.
    NONE_DEVICEEVENTLEVEL DeviceEventLevel = iota
    // Indicates that the device event level is verbose.
    VERBOSE_DEVICEEVENTLEVEL
    // Indicates that the device event level is information.
    INFORMATION_DEVICEEVENTLEVEL
    // Indicates that the device event level is warning.
    WARNING_DEVICEEVENTLEVEL
    // Indicates that the device event level is error.
    ERROR_DEVICEEVENTLEVEL
    // Indicates that the device event level is critical.
    CRITICAL_DEVICEEVENTLEVEL
    // Placeholder value for future expansion.
    UNKNOWNFUTUREVALUE_DEVICEEVENTLEVEL
)

func (i DeviceEventLevel) String() string {
    return []string{"none", "verbose", "information", "warning", "error", "critical", "unknownFutureValue"}[i]
}
func ParseDeviceEventLevel(v string) (interface{}, error) {
    result := NONE_DEVICEEVENTLEVEL
    switch v {
        case "none":
            result = NONE_DEVICEEVENTLEVEL
        case "verbose":
            result = VERBOSE_DEVICEEVENTLEVEL
        case "information":
            result = INFORMATION_DEVICEEVENTLEVEL
        case "warning":
            result = WARNING_DEVICEEVENTLEVEL
        case "error":
            result = ERROR_DEVICEEVENTLEVEL
        case "critical":
            result = CRITICAL_DEVICEEVENTLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEEVENTLEVEL
        default:
            return 0, errors.New("Unknown DeviceEventLevel value: " + v)
    }
    return &result, nil
}
func SerializeDeviceEventLevel(values []DeviceEventLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
