package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type GroupPolicySettingScope int

const (
    // Device scope unknown
    UNKNOWN_GROUPPOLICYSETTINGSCOPE GroupPolicySettingScope = iota
    // Device scope
    DEVICE_GROUPPOLICYSETTINGSCOPE
    // User scope
    USER_GROUPPOLICYSETTINGSCOPE
)

func (i GroupPolicySettingScope) String() string {
    return []string{"unknown", "device", "user"}[i]
}
func ParseGroupPolicySettingScope(v string) (interface{}, error) {
    result := UNKNOWN_GROUPPOLICYSETTINGSCOPE
    switch v {
        case "unknown":
            result = UNKNOWN_GROUPPOLICYSETTINGSCOPE
        case "device":
            result = DEVICE_GROUPPOLICYSETTINGSCOPE
        case "user":
            result = USER_GROUPPOLICYSETTINGSCOPE
        default:
            return 0, errors.New("Unknown GroupPolicySettingScope value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicySettingScope(values []GroupPolicySettingScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
