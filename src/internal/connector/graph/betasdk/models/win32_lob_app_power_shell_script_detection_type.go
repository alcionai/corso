package models
import (
    "errors"
)
// Provides operations to call the add method.
type Win32LobAppPowerShellScriptDetectionType int

const (
    // Not configured.
    NOTCONFIGURED_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE Win32LobAppPowerShellScriptDetectionType = iota
    // Output data type is string.
    STRING_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
    // Output data type is date time.
    DATETIME_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
    // Output data type is integer.
    INTEGER_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
    // Output data type is float.
    FLOAT_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
    // Output data type is version.
    VERSION_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
    // Output data type is boolean.
    BOOLEAN_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
)

func (i Win32LobAppPowerShellScriptDetectionType) String() string {
    return []string{"notConfigured", "string", "dateTime", "integer", "float", "version", "boolean"}[i]
}
func ParseWin32LobAppPowerShellScriptDetectionType(v string) (interface{}, error) {
    result := NOTCONFIGURED_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
        case "string":
            result = STRING_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
        case "dateTime":
            result = DATETIME_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
        case "integer":
            result = INTEGER_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
        case "float":
            result = FLOAT_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
        case "version":
            result = VERSION_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
        case "boolean":
            result = BOOLEAN_WIN32LOBAPPPOWERSHELLSCRIPTDETECTIONTYPE
        default:
            return 0, errors.New("Unknown Win32LobAppPowerShellScriptDetectionType value: " + v)
    }
    return &result, nil
}
func SerializeWin32LobAppPowerShellScriptDetectionType(values []Win32LobAppPowerShellScriptDetectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
