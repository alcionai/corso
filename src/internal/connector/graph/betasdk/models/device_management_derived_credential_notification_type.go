package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceManagementDerivedCredentialNotificationType int

const (
    // None
    NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE DeviceManagementDerivedCredentialNotificationType = iota
    // Company Portal
    COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
    // Email
    EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
)

func (i DeviceManagementDerivedCredentialNotificationType) String() string {
    return []string{"none", "companyPortal", "email"}[i]
}
func ParseDeviceManagementDerivedCredentialNotificationType(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
        case "companyPortal":
            result = COMPANYPORTAL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
        case "email":
            result = EMAIL_DEVICEMANAGEMENTDERIVEDCREDENTIALNOTIFICATIONTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementDerivedCredentialNotificationType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementDerivedCredentialNotificationType(values []DeviceManagementDerivedCredentialNotificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
