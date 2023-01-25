package models
import (
    "errors"
)
// Provides operations to call the add method.
type TeamworkDeviceActivityState int

const (
    UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE TeamworkDeviceActivityState = iota
    BUSY_TEAMWORKDEVICEACTIVITYSTATE
    IDLE_TEAMWORKDEVICEACTIVITYSTATE
    UNAVAILABLE_TEAMWORKDEVICEACTIVITYSTATE
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICEACTIVITYSTATE
)

func (i TeamworkDeviceActivityState) String() string {
    return []string{"unknown", "busy", "idle", "unavailable", "unknownFutureValue"}[i]
}
func ParseTeamworkDeviceActivityState(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMWORKDEVICEACTIVITYSTATE
        case "busy":
            result = BUSY_TEAMWORKDEVICEACTIVITYSTATE
        case "idle":
            result = IDLE_TEAMWORKDEVICEACTIVITYSTATE
        case "unavailable":
            result = UNAVAILABLE_TEAMWORKDEVICEACTIVITYSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICEACTIVITYSTATE
        default:
            return 0, errors.New("Unknown TeamworkDeviceActivityState value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkDeviceActivityState(values []TeamworkDeviceActivityState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
