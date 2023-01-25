package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceManagementConfigurationAzureAdTrustType int

const (
    // No AAD Trust Type specified
    NONE_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE DeviceManagementConfigurationAzureAdTrustType = iota
    // AAD Joined Trust Type
    AZUREADJOINED_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
    // AddWorkAccount
    ADDWORKACCOUNT_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
    // MDM only
    MDMONLY_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
)

func (i DeviceManagementConfigurationAzureAdTrustType) String() string {
    return []string{"none", "azureAdJoined", "addWorkAccount", "mdmOnly"}[i]
}
func ParseDeviceManagementConfigurationAzureAdTrustType(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        case "azureAdJoined":
            result = AZUREADJOINED_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        case "addWorkAccount":
            result = ADDWORKACCOUNT_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        case "mdmOnly":
            result = MDMONLY_DEVICEMANAGEMENTCONFIGURATIONAZUREADTRUSTTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationAzureAdTrustType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationAzureAdTrustType(values []DeviceManagementConfigurationAzureAdTrustType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
