package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Windows10DeviceModeType int

const (
    // Standard Configuration
    STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE Windows10DeviceModeType = iota
    // S Mode Configuration
    SMODECONFIGURATION_WINDOWS10DEVICEMODETYPE
)

func (i Windows10DeviceModeType) String() string {
    return []string{"standardConfiguration", "sModeConfiguration"}[i]
}
func ParseWindows10DeviceModeType(v string) (interface{}, error) {
    result := STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE
    switch v {
        case "standardConfiguration":
            result = STANDARDCONFIGURATION_WINDOWS10DEVICEMODETYPE
        case "sModeConfiguration":
            result = SMODECONFIGURATION_WINDOWS10DEVICEMODETYPE
        default:
            return 0, errors.New("Unknown Windows10DeviceModeType value: " + v)
    }
    return &result, nil
}
func SerializeWindows10DeviceModeType(values []Windows10DeviceModeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
