package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Win32LobAppRegistryDetectionType int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPREGISTRYDETECTIONTYPE Win32LobAppRegistryDetectionType = iota
    // The specified registry key or value exists.
    EXISTS_WIN32LOBAPPREGISTRYDETECTIONTYPE
    // The specified registry key or value does not exist.
    DOESNOTEXIST_WIN32LOBAPPREGISTRYDETECTIONTYPE
    // String value type.
    STRING_WIN32LOBAPPREGISTRYDETECTIONTYPE
    // Integer value type.
    INTEGER_WIN32LOBAPPREGISTRYDETECTIONTYPE
    // Version value type.
    VERSION_WIN32LOBAPPREGISTRYDETECTIONTYPE
)

func (i Win32LobAppRegistryDetectionType) String() string {
    return []string{"notConfigured", "exists", "doesNotExist", "string", "integer", "version"}[i]
}
func ParseWin32LobAppRegistryDetectionType(v string) (interface{}, error) {
    result := NOTCONFIGURED_WIN32LOBAPPREGISTRYDETECTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPREGISTRYDETECTIONTYPE
        case "exists":
            result = EXISTS_WIN32LOBAPPREGISTRYDETECTIONTYPE
        case "doesNotExist":
            result = DOESNOTEXIST_WIN32LOBAPPREGISTRYDETECTIONTYPE
        case "string":
            result = STRING_WIN32LOBAPPREGISTRYDETECTIONTYPE
        case "integer":
            result = INTEGER_WIN32LOBAPPREGISTRYDETECTIONTYPE
        case "version":
            result = VERSION_WIN32LOBAPPREGISTRYDETECTIONTYPE
        default:
            return 0, errors.New("Unknown Win32LobAppRegistryDetectionType value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppRegistryDetectionType(values []Win32LobAppRegistryDetectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
