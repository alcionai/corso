package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EndpointSecurityConfigurationProfileType int

const (
    // Unknown.
    UNKNOWN_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE EndpointSecurityConfigurationProfileType = iota
    // Antivirus.
    ANTIVIRUS_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Windows Security.
    WINDOWSSECURITY_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // BitLocker.
    BITLOCKER_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // FileVault.
    FILEVAULT_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Firewall.
    FIREWALL_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Firewall rules.
    FIREWALLRULES_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Endpoint detection and response.
    ENDPOINTDETECTIONANDRESPONSE_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Device control.
    DEVICECONTROL_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // App and browser isolation.
    APPANDBROWSERISOLATION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Exploit protection.
    EXPLOITPROTECTION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Web protection.
    WEBPROTECTION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Application control.
    APPLICATIONCONTROL_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Attack surface reduction rules.
    ATTACKSURFACEREDUCTIONRULES_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    // Account protection.
    ACCOUNTPROTECTION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
)

func (i EndpointSecurityConfigurationProfileType) String() string {
    return []string{"unknown", "antivirus", "windowsSecurity", "bitLocker", "fileVault", "firewall", "firewallRules", "endpointDetectionAndResponse", "deviceControl", "appAndBrowserIsolation", "exploitProtection", "webProtection", "applicationControl", "attackSurfaceReductionRules", "accountProtection"}[i]
}
func ParseEndpointSecurityConfigurationProfileType(v string) (interface{}, error) {
    result := UNKNOWN_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
    switch v {
        case "unknown":
            result = UNKNOWN_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "antivirus":
            result = ANTIVIRUS_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "windowsSecurity":
            result = WINDOWSSECURITY_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "bitLocker":
            result = BITLOCKER_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "fileVault":
            result = FILEVAULT_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "firewall":
            result = FIREWALL_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "firewallRules":
            result = FIREWALLRULES_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "endpointDetectionAndResponse":
            result = ENDPOINTDETECTIONANDRESPONSE_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "deviceControl":
            result = DEVICECONTROL_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "appAndBrowserIsolation":
            result = APPANDBROWSERISOLATION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "exploitProtection":
            result = EXPLOITPROTECTION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "webProtection":
            result = WEBPROTECTION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "applicationControl":
            result = APPLICATIONCONTROL_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "attackSurfaceReductionRules":
            result = ATTACKSURFACEREDUCTIONRULES_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        case "accountProtection":
            result = ACCOUNTPROTECTION_ENDPOINTSECURITYCONFIGURATIONPROFILETYPE
        default:
            return 0, errors.New("Unknown EndpointSecurityConfigurationProfileType value: " + v)
    }
    return &result, nil
}
func SerializeEndpointSecurityConfigurationProfileType(values []EndpointSecurityConfigurationProfileType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
