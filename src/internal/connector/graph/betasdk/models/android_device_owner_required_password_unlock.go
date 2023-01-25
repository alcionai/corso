package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidDeviceOwnerRequiredPasswordUnlock int

const (
    // Timeout period before strong authentication is required is set to the device's default.
    DEVICEDEFAULT_ANDROIDDEVICEOWNERREQUIREDPASSWORDUNLOCK AndroidDeviceOwnerRequiredPasswordUnlock = iota
    // Timeout period before strong authentication is required is set to 24 hours.
    DAILY_ANDROIDDEVICEOWNERREQUIREDPASSWORDUNLOCK
    // Unknown future value (reserved, not used right now)
    UNKOWNFUTUREVALUE_ANDROIDDEVICEOWNERREQUIREDPASSWORDUNLOCK
)

func (i AndroidDeviceOwnerRequiredPasswordUnlock) String() string {
    return []string{"deviceDefault", "daily", "unkownFutureValue"}[i]
}
func ParseAndroidDeviceOwnerRequiredPasswordUnlock(v string) (interface{}, error) {
    result := DEVICEDEFAULT_ANDROIDDEVICEOWNERREQUIREDPASSWORDUNLOCK
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDDEVICEOWNERREQUIREDPASSWORDUNLOCK
        case "daily":
            result = DAILY_ANDROIDDEVICEOWNERREQUIREDPASSWORDUNLOCK
        case "unkownFutureValue":
            result = UNKOWNFUTUREVALUE_ANDROIDDEVICEOWNERREQUIREDPASSWORDUNLOCK
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerRequiredPasswordUnlock value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerRequiredPasswordUnlock(values []AndroidDeviceOwnerRequiredPasswordUnlock) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
