package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementConfigurationTechnologies int

const (
    // Setting cannot be deployed through any channel
    NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES DeviceManagementConfigurationTechnologies = iota
    // Setting can be deployed through the MDM channel
    MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed through the Windows10XManagement channel
    WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed through the ConfigManager channel
    CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed through the AppleRemoteManagement channel
    APPLEREMOTEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed through the SENSE agent channel
    MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed through the Exchange Online agent channel
    EXCHANGEONLINE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed through the Linux Mdm channel
    LINUXMDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed through device enrollment.
    ENROLLMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Setting can be deployed using the Endpoint privilege management channel
    ENDPOINTPRIVILEGEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
)

func (i DeviceManagementConfigurationTechnologies) String() string {
    return []string{"none", "mdm", "windows10XManagement", "configManager", "appleRemoteManagement", "microsoftSense", "exchangeOnline", "linuxMdm", "enrollment", "endpointPrivilegeManagement", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationTechnologies(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "mdm":
            result = MDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "windows10XManagement":
            result = WINDOWS10XMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "configManager":
            result = CONFIGMANAGER_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "appleRemoteManagement":
            result = APPLEREMOTEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "microsoftSense":
            result = MICROSOFTSENSE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "exchangeOnline":
            result = EXCHANGEONLINE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "linuxMdm":
            result = LINUXMDM_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "enrollment":
            result = ENROLLMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "endpointPrivilegeManagement":
            result = ENDPOINTPRIVILEGEMANAGEMENT_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONTECHNOLOGIES
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationTechnologies value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationTechnologies(values []DeviceManagementConfigurationTechnologies) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
