package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsSModeConfiguration int

const (
    // This option will remove all restrictions to unlock S mode - default
    NORESTRICTION_WINDOWSSMODECONFIGURATION WindowsSModeConfiguration = iota
    // This option will block the user to unlock the device from S mode
    BLOCK_WINDOWSSMODECONFIGURATION
    // This option will unlock the device from S mode
    UNLOCK_WINDOWSSMODECONFIGURATION
)

func (i WindowsSModeConfiguration) String() string {
    return []string{"noRestriction", "block", "unlock"}[i]
}
func ParseWindowsSModeConfiguration(v string) (interface{}, error) {
    result := NORESTRICTION_WINDOWSSMODECONFIGURATION
    switch v {
        case "noRestriction":
            result = NORESTRICTION_WINDOWSSMODECONFIGURATION
        case "block":
            result = BLOCK_WINDOWSSMODECONFIGURATION
        case "unlock":
            result = UNLOCK_WINDOWSSMODECONFIGURATION
        default:
            return 0, errors.New("Unknown WindowsSModeConfiguration value: " + v)
    }
    return &result, nil
}
func SerializeWindowsSModeConfiguration(values []WindowsSModeConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
