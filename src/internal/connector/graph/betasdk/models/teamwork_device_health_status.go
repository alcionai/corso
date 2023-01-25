package models
import (
    "errors"
)
// Provides operations to call the add method.
type TeamworkDeviceHealthStatus int

const (
    UNKNOWN_TEAMWORKDEVICEHEALTHSTATUS TeamworkDeviceHealthStatus = iota
    OFFLINE_TEAMWORKDEVICEHEALTHSTATUS
    CRITICAL_TEAMWORKDEVICEHEALTHSTATUS
    NONURGENT_TEAMWORKDEVICEHEALTHSTATUS
    HEALTHY_TEAMWORKDEVICEHEALTHSTATUS
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICEHEALTHSTATUS
)

func (i TeamworkDeviceHealthStatus) String() string {
    return []string{"unknown", "offline", "critical", "nonUrgent", "healthy", "unknownFutureValue"}[i]
}
func ParseTeamworkDeviceHealthStatus(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKDEVICEHEALTHSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMWORKDEVICEHEALTHSTATUS
        case "offline":
            result = OFFLINE_TEAMWORKDEVICEHEALTHSTATUS
        case "critical":
            result = CRITICAL_TEAMWORKDEVICEHEALTHSTATUS
        case "nonUrgent":
            result = NONURGENT_TEAMWORKDEVICEHEALTHSTATUS
        case "healthy":
            result = HEALTHY_TEAMWORKDEVICEHEALTHSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICEHEALTHSTATUS
        default:
            return 0, errors.New("Unknown TeamworkDeviceHealthStatus value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkDeviceHealthStatus(values []TeamworkDeviceHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
