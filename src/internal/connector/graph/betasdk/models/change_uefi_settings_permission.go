package models
import (
    "errors"
)
// Provides operations to call the add method.
type ChangeUefiSettingsPermission int

const (
    // Device default value, no intent.
    NOTCONFIGUREDONLY_CHANGEUEFISETTINGSPERMISSION ChangeUefiSettingsPermission = iota
    // Prevent change of UEFI setting permission
    NONE_CHANGEUEFISETTINGSPERMISSION
)

func (i ChangeUefiSettingsPermission) String() string {
    return []string{"notConfiguredOnly", "none"}[i]
}
func ParseChangeUefiSettingsPermission(v string) (interface{}, error) {
    result := NOTCONFIGUREDONLY_CHANGEUEFISETTINGSPERMISSION
    switch v {
        case "notConfiguredOnly":
            result = NOTCONFIGUREDONLY_CHANGEUEFISETTINGSPERMISSION
        case "none":
            result = NONE_CHANGEUEFISETTINGSPERMISSION
        default:
            return 0, errors.New("Unknown ChangeUefiSettingsPermission value: " + v)
    }
    return &result, nil
}
func SerializeChangeUefiSettingsPermission(values []ChangeUefiSettingsPermission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
