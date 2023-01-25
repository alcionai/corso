package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceScopeParameter int

const (
    // Device Scope parameter is not set
    NONE_DEVICESCOPEPARAMETER DeviceScopeParameter = iota
    // use Scope Tag as parameter for the device scope configuration.
    SCOPETAG_DEVICESCOPEPARAMETER
    // Placeholder value for future expansion.
    UNKNOWNFUTUREVALUE_DEVICESCOPEPARAMETER
)

func (i DeviceScopeParameter) String() string {
    return []string{"none", "scopeTag", "unknownFutureValue"}[i]
}
func ParseDeviceScopeParameter(v string) (interface{}, error) {
    result := NONE_DEVICESCOPEPARAMETER
    switch v {
        case "none":
            result = NONE_DEVICESCOPEPARAMETER
        case "scopeTag":
            result = SCOPETAG_DEVICESCOPEPARAMETER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPEPARAMETER
        default:
            return 0, errors.New("Unknown DeviceScopeParameter value: " + v)
    }
    return &result, nil
}
func SerializeDeviceScopeParameter(values []DeviceScopeParameter) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
