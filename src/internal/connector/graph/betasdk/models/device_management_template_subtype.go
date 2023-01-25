package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManagementTemplateSubtype int

const (
    // Template has no subtype
    NONE_DEVICEMANAGEMENTTEMPLATESUBTYPE DeviceManagementTemplateSubtype = iota
    // Endpoint security firewall subtype
    FIREWALL_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security disk encryption subtype
    DISKENCRYPTION_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security attack surface reduction subtype
    ATTACKSURFACEREDUCTION_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security endpoint detection and response subtype
    ENDPOINTDETECTIONREPONSE_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security account protection subtype
    ACCOUNTPROTECTION_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security anitivirus subtype
    ANTIVIRUS_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security firewall shared app subtype
    FIREWALLSHAREDAPPLIST_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security firewall shared ip range list subtype
    FIREWALLSHAREDIPLIST_DEVICEMANAGEMENTTEMPLATESUBTYPE
    // Endpoint security firewall shared port range list subtype
    FIREWALLSHAREDPORTLIST_DEVICEMANAGEMENTTEMPLATESUBTYPE
)

func (i DeviceManagementTemplateSubtype) String() string {
    return []string{"none", "firewall", "diskEncryption", "attackSurfaceReduction", "endpointDetectionReponse", "accountProtection", "antivirus", "firewallSharedAppList", "firewallSharedIpList", "firewallSharedPortlist"}[i]
}
func ParseDeviceManagementTemplateSubtype(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTTEMPLATESUBTYPE
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "firewall":
            result = FIREWALL_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "diskEncryption":
            result = DISKENCRYPTION_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "attackSurfaceReduction":
            result = ATTACKSURFACEREDUCTION_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "endpointDetectionReponse":
            result = ENDPOINTDETECTIONREPONSE_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "accountProtection":
            result = ACCOUNTPROTECTION_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "antivirus":
            result = ANTIVIRUS_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "firewallSharedAppList":
            result = FIREWALLSHAREDAPPLIST_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "firewallSharedIpList":
            result = FIREWALLSHAREDIPLIST_DEVICEMANAGEMENTTEMPLATESUBTYPE
        case "firewallSharedPortlist":
            result = FIREWALLSHAREDPORTLIST_DEVICEMANAGEMENTTEMPLATESUBTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementTemplateSubtype value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementTemplateSubtype(values []DeviceManagementTemplateSubtype) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
