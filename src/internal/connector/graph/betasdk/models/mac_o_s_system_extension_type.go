package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MacOSSystemExtensionType int

const (
    // Enables driver extensions.
    DRIVEREXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE MacOSSystemExtensionType = iota
    // Enables network extensions.
    NETWORKEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
    // Enables endpoint security extensions.
    ENDPOINTSECURITYEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
)

func (i MacOSSystemExtensionType) String() string {
    return []string{"driverExtensionsAllowed", "networkExtensionsAllowed", "endpointSecurityExtensionsAllowed"}[i]
}
func ParseMacOSSystemExtensionType(v string) (interface{}, error) {
    result := DRIVEREXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
    switch v {
        case "driverExtensionsAllowed":
            result = DRIVEREXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
        case "networkExtensionsAllowed":
            result = NETWORKEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
        case "endpointSecurityExtensionsAllowed":
            result = ENDPOINTSECURITYEXTENSIONSALLOWED_MACOSSYSTEMEXTENSIONTYPE
        default:
            return 0, errors.New("Unknown MacOSSystemExtensionType value: " + v)
    }
    return &result, nil
}
func SerializeMacOSSystemExtensionType(values []MacOSSystemExtensionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
