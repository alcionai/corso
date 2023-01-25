package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsDefenderTamperProtectionOptions int

const (
    // Not Configured
    NOTCONFIGURED_WINDOWSDEFENDERTAMPERPROTECTIONOPTIONS WindowsDefenderTamperProtectionOptions = iota
    // Enable windows defender tamper protection
    ENABLE_WINDOWSDEFENDERTAMPERPROTECTIONOPTIONS
    // Disable windows defender tamper protection
    DISABLE_WINDOWSDEFENDERTAMPERPROTECTIONOPTIONS
)

func (i WindowsDefenderTamperProtectionOptions) String() string {
    return []string{"notConfigured", "enable", "disable"}[i]
}
func ParseWindowsDefenderTamperProtectionOptions(v string) (interface{}, error) {
    result := NOTCONFIGURED_WINDOWSDEFENDERTAMPERPROTECTIONOPTIONS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWSDEFENDERTAMPERPROTECTIONOPTIONS
        case "enable":
            result = ENABLE_WINDOWSDEFENDERTAMPERPROTECTIONOPTIONS
        case "disable":
            result = DISABLE_WINDOWSDEFENDERTAMPERPROTECTIONOPTIONS
        default:
            return 0, errors.New("Unknown WindowsDefenderTamperProtectionOptions value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDefenderTamperProtectionOptions(values []WindowsDefenderTamperProtectionOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
