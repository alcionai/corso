package models
import (
    "errors"
)
// Provides operations to call the add method.
type InkAccessSetting int

const (
    // Not configured.
    NOTCONFIGURED_INKACCESSSETTING InkAccessSetting = iota
    // Enabled.
    ENABLED_INKACCESSSETTING
    // Disabled.
    DISABLED_INKACCESSSETTING
)

func (i InkAccessSetting) String() string {
    return []string{"notConfigured", "enabled", "disabled"}[i]
}
func ParseInkAccessSetting(v string) (interface{}, error) {
    result := NOTCONFIGURED_INKACCESSSETTING
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_INKACCESSSETTING
        case "enabled":
            result = ENABLED_INKACCESSSETTING
        case "disabled":
            result = DISABLED_INKACCESSSETTING
        default:
            return 0, errors.New("Unknown InkAccessSetting value: " + v)
    }
    return &result, nil
}
func SerializeInkAccessSetting(values []InkAccessSetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
