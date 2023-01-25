package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsKioskAppType int

const (
    // Unknown.
    UNKNOWN_WINDOWSKIOSKAPPTYPE WindowsKioskAppType = iota
    // Store App.
    STORE_WINDOWSKIOSKAPPTYPE
    // Desktop App.
    DESKTOP_WINDOWSKIOSKAPPTYPE
    // Input by AUMID.
    AUMID_WINDOWSKIOSKAPPTYPE
)

func (i WindowsKioskAppType) String() string {
    return []string{"unknown", "store", "desktop", "aumId"}[i]
}
func ParseWindowsKioskAppType(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSKIOSKAPPTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSKIOSKAPPTYPE
        case "store":
            result = STORE_WINDOWSKIOSKAPPTYPE
        case "desktop":
            result = DESKTOP_WINDOWSKIOSKAPPTYPE
        case "aumId":
            result = AUMID_WINDOWSKIOSKAPPTYPE
        default:
            return 0, errors.New("Unknown WindowsKioskAppType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsKioskAppType(values []WindowsKioskAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
