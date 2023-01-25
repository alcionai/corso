package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AssociatedAssignmentPayloadType int

const (
    // Invalid payload type
    UNKNOWN_ASSOCIATEDASSIGNMENTPAYLOADTYPE AssociatedAssignmentPayloadType = iota
    // Indicates that this filter is associated with a configuration or compliance policy payload type
    DEVICECONFIGURATIONANDCOMPLIANCE_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this assignment filter is associated with an application payload type
    APPLICATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this filter is associated with a Android Enterprise application payload type
    ANDROIDENTERPRISEAPP_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this filter is associated with an enrollment restriction or enrollment status page policy payload type
    ENROLLMENTCONFIGURATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this filter is associated with an Administrative Template policy payload type
    GROUPPOLICYCONFIGURATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this assignment filter is associated with Zero touch deployment Device Configuration Profile payload type
    ZEROTOUCHDEPLOYMENTDEVICECONFIGPROFILE_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this filter is associated with an Android Enterprise Configuration policy payload type
    ANDROIDENTERPRISECONFIGURATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this assignment filter is associated with Device Firmware Configuration Interface(DCFI) payload type
    DEVICEFIRMWARECONFIGURATIONINTERFACEPOLICY_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this filter is associated with a resource access policy (Wifi, VPN, Certificate) payload type
    RESOURCEACCESSPOLICY_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this filter is associated with a Win32 app payload type
    WIN32APP_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    // Indicates that this filter is associated with a configuration or compliance policy on Device Configuration v2 Infrastructure payload type
    DEVICEMANAGMENTCONFIGURATIONANDCOMPLIANCEPOLICY_ASSOCIATEDASSIGNMENTPAYLOADTYPE
)

func (i AssociatedAssignmentPayloadType) String() string {
    return []string{"unknown", "deviceConfigurationAndCompliance", "application", "androidEnterpriseApp", "enrollmentConfiguration", "groupPolicyConfiguration", "zeroTouchDeploymentDeviceConfigProfile", "androidEnterpriseConfiguration", "deviceFirmwareConfigurationInterfacePolicy", "resourceAccessPolicy", "win32app", "deviceManagmentConfigurationAndCompliancePolicy"}[i]
}
func ParseAssociatedAssignmentPayloadType(v string) (interface{}, error) {
    result := UNKNOWN_ASSOCIATEDASSIGNMENTPAYLOADTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "deviceConfigurationAndCompliance":
            result = DEVICECONFIGURATIONANDCOMPLIANCE_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "application":
            result = APPLICATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "androidEnterpriseApp":
            result = ANDROIDENTERPRISEAPP_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "enrollmentConfiguration":
            result = ENROLLMENTCONFIGURATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "groupPolicyConfiguration":
            result = GROUPPOLICYCONFIGURATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "zeroTouchDeploymentDeviceConfigProfile":
            result = ZEROTOUCHDEPLOYMENTDEVICECONFIGPROFILE_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "androidEnterpriseConfiguration":
            result = ANDROIDENTERPRISECONFIGURATION_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "deviceFirmwareConfigurationInterfacePolicy":
            result = DEVICEFIRMWARECONFIGURATIONINTERFACEPOLICY_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "resourceAccessPolicy":
            result = RESOURCEACCESSPOLICY_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "win32app":
            result = WIN32APP_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        case "deviceManagmentConfigurationAndCompliancePolicy":
            result = DEVICEMANAGMENTCONFIGURATIONANDCOMPLIANCEPOLICY_ASSOCIATEDASSIGNMENTPAYLOADTYPE
        default:
            return 0, errors.New("Unknown AssociatedAssignmentPayloadType value: " + v)
    }
    return &result, nil
}
func SerializeAssociatedAssignmentPayloadType(values []AssociatedAssignmentPayloadType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
