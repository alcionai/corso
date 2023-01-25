package models
import (
    "errors"
)
// Provides operations to call the add method.
type KioskModeType int

const (
    // Not configured
    NOTCONFIGURED_KIOSKMODETYPE KioskModeType = iota
    // Run in single-app mode
    SINGLEAPPMODE_KIOSKMODETYPE
    // Run in multi-app mode
    MULTIAPPMODE_KIOSKMODETYPE
)

func (i KioskModeType) String() string {
    return []string{"notConfigured", "singleAppMode", "multiAppMode"}[i]
}
func ParseKioskModeType(v string) (interface{}, error) {
    result := NOTCONFIGURED_KIOSKMODETYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_KIOSKMODETYPE
        case "singleAppMode":
            result = SINGLEAPPMODE_KIOSKMODETYPE
        case "multiAppMode":
            result = MULTIAPPMODE_KIOSKMODETYPE
        default:
            return 0, errors.New("Unknown KioskModeType value: " + v)
    }
    return &result, nil
}
func SerializeKioskModeType(values []KioskModeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
