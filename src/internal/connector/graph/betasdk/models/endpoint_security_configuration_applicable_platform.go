package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EndpointSecurityConfigurationApplicablePlatform int

const (
    // Unknown.
    UNKNOWN_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM EndpointSecurityConfigurationApplicablePlatform = iota
    // MacOS.
    MACOS_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
    // Windows 10 and later.
    WINDOWS10ANDLATER_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
    // Windows 10 and Windows Server.
    WINDOWS10ANDWINDOWSSERVER_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
)

func (i EndpointSecurityConfigurationApplicablePlatform) String() string {
    return []string{"unknown", "macOS", "windows10AndLater", "windows10AndWindowsServer"}[i]
}
func ParseEndpointSecurityConfigurationApplicablePlatform(v string) (interface{}, error) {
    result := UNKNOWN_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
    switch v {
        case "unknown":
            result = UNKNOWN_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
        case "macOS":
            result = MACOS_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
        case "windows10AndLater":
            result = WINDOWS10ANDLATER_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
        case "windows10AndWindowsServer":
            result = WINDOWS10ANDWINDOWSSERVER_ENDPOINTSECURITYCONFIGURATIONAPPLICABLEPLATFORM
        default:
            return 0, errors.New("Unknown EndpointSecurityConfigurationApplicablePlatform value: " + v)
    }
    return &result, nil
}
func SerializeEndpointSecurityConfigurationApplicablePlatform(values []EndpointSecurityConfigurationApplicablePlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
