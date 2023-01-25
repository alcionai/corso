package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementConfigurationTemplateFamily int

const (
    // Default for Template Family when Policy is not linked to a Template
    NONE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY DeviceManagementConfigurationTemplateFamily = iota
    // Template Family for EndpointSecurityAntivirus that manages the discrete group of antivirus settings for managed devices
    ENDPOINTSECURITYANTIVIRUS_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for EndpointSecurityDiskEncryption that provides settings that are relevant for a devices built-in encryption  method, like FileVault or BitLocker
    ENDPOINTSECURITYDISKENCRYPTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for EndpointSecurityFirewall that helps configure a devices built-in firewall for device that run macOS and Windows 10
    ENDPOINTSECURITYFIREWALL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for EndpointSecurityEndpointDetectionAndResponse that facilitates management of the EDR settings and onboard devices to Microsoft Defender for Endpoint
    ENDPOINTSECURITYENDPOINTDETECTIONANDRESPONSE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for EndpointSecurityAttackSurfaceReduction that help reduce your attack surfaces, by minimizing the places where your organization is vulnerable to cyberthreats and attacks
    ENDPOINTSECURITYATTACKSURFACEREDUCTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for EndpointSecurityAccountProtection that facilitates protecting the identity and accounts of users
    ENDPOINTSECURITYACCOUNTPROTECTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for ApplicationControl that helps mitigate security threats by restricting the applications that users can run and the code that runs in the System Core (kernel)
    ENDPOINTSECURITYAPPLICATIONCONTROL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for EPM Elevation Rules
    ENDPOINTSECURITYENDPOINTPRIVILEGEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for EnrollmentConfiguration
    ENROLLMENTCONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for QuietTimeIndicates Template Family for all the Apps QuietTime policies and templates
    APPQUIETTIME_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Template Family for Baseline
    BASELINE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
)

func (i DeviceManagementConfigurationTemplateFamily) String() string {
    return []string{"none", "endpointSecurityAntivirus", "endpointSecurityDiskEncryption", "endpointSecurityFirewall", "endpointSecurityEndpointDetectionAndResponse", "endpointSecurityAttackSurfaceReduction", "endpointSecurityAccountProtection", "endpointSecurityApplicationControl", "endpointSecurityEndpointPrivilegeManagement", "enrollmentConfiguration", "appQuietTime", "baseline", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationTemplateFamily(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityAntivirus":
            result = ENDPOINTSECURITYANTIVIRUS_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityDiskEncryption":
            result = ENDPOINTSECURITYDISKENCRYPTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityFirewall":
            result = ENDPOINTSECURITYFIREWALL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityEndpointDetectionAndResponse":
            result = ENDPOINTSECURITYENDPOINTDETECTIONANDRESPONSE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityAttackSurfaceReduction":
            result = ENDPOINTSECURITYATTACKSURFACEREDUCTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityAccountProtection":
            result = ENDPOINTSECURITYACCOUNTPROTECTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityApplicationControl":
            result = ENDPOINTSECURITYAPPLICATIONCONTROL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "endpointSecurityEndpointPrivilegeManagement":
            result = ENDPOINTSECURITYENDPOINTPRIVILEGEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "enrollmentConfiguration":
            result = ENROLLMENTCONFIGURATION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "appQuietTime":
            result = APPQUIETTIME_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "baseline":
            result = BASELINE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationTemplateFamily value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationTemplateFamily(values []DeviceManagementConfigurationTemplateFamily) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
