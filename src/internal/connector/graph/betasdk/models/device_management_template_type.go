package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManagementTemplateType int

const (
    // Security baseline template
    SECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE DeviceManagementTemplateType = iota
    // Specialized devices template
    SPECIALIZEDDEVICES_DEVICEMANAGEMENTTEMPLATETYPE
    // Advanced Threat Protection security baseline template
    ADVANCEDTHREATPROTECTIONSECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
    // Device configuration template
    DEVICECONFIGURATION_DEVICEMANAGEMENTTEMPLATETYPE
    // Custom admin defined template
    CUSTOM_DEVICEMANAGEMENTTEMPLATETYPE
    // Templates containing specific security focused settings
    SECURITYTEMPLATE_DEVICEMANAGEMENTTEMPLATETYPE
    // Microsoft Edge security baseline template
    MICROSOFTEDGESECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
    // Microsoft Office 365 ProPlus security baseline template
    MICROSOFTOFFICE365PROPLUSSECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
    // Device compliance template
    DEVICECOMPLIANCE_DEVICEMANAGEMENTTEMPLATETYPE
    // Device Configuration for Microsoft Office 365 settings
    DEVICECONFIGURATIONFOROFFICE365_DEVICEMANAGEMENTTEMPLATETYPE
    // Windows 365 security baseline template
    CLOUDPC_DEVICEMANAGEMENTTEMPLATETYPE
    // Firewall Shared Object templates for reference settings
    FIREWALLSHAREDSETTINGS_DEVICEMANAGEMENTTEMPLATETYPE
)

func (i DeviceManagementTemplateType) String() string {
    return []string{"securityBaseline", "specializedDevices", "advancedThreatProtectionSecurityBaseline", "deviceConfiguration", "custom", "securityTemplate", "microsoftEdgeSecurityBaseline", "microsoftOffice365ProPlusSecurityBaseline", "deviceCompliance", "deviceConfigurationForOffice365", "cloudPC", "firewallSharedSettings"}[i]
}
func ParseDeviceManagementTemplateType(v string) (interface{}, error) {
    result := SECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
    switch v {
        case "securityBaseline":
            result = SECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
        case "specializedDevices":
            result = SPECIALIZEDDEVICES_DEVICEMANAGEMENTTEMPLATETYPE
        case "advancedThreatProtectionSecurityBaseline":
            result = ADVANCEDTHREATPROTECTIONSECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
        case "deviceConfiguration":
            result = DEVICECONFIGURATION_DEVICEMANAGEMENTTEMPLATETYPE
        case "custom":
            result = CUSTOM_DEVICEMANAGEMENTTEMPLATETYPE
        case "securityTemplate":
            result = SECURITYTEMPLATE_DEVICEMANAGEMENTTEMPLATETYPE
        case "microsoftEdgeSecurityBaseline":
            result = MICROSOFTEDGESECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
        case "microsoftOffice365ProPlusSecurityBaseline":
            result = MICROSOFTOFFICE365PROPLUSSECURITYBASELINE_DEVICEMANAGEMENTTEMPLATETYPE
        case "deviceCompliance":
            result = DEVICECOMPLIANCE_DEVICEMANAGEMENTTEMPLATETYPE
        case "deviceConfigurationForOffice365":
            result = DEVICECONFIGURATIONFOROFFICE365_DEVICEMANAGEMENTTEMPLATETYPE
        case "cloudPC":
            result = CLOUDPC_DEVICEMANAGEMENTTEMPLATETYPE
        case "firewallSharedSettings":
            result = FIREWALLSHAREDSETTINGS_DEVICEMANAGEMENTTEMPLATETYPE
        default:
            return 0, errors.New("Unknown DeviceManagementTemplateType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementTemplateType(values []DeviceManagementTemplateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
