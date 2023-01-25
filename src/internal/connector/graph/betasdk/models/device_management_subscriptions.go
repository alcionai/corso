package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceManagementSubscriptions int

const (
    // None
    NONE_DEVICEMANAGEMENTSUBSCRIPTIONS DeviceManagementSubscriptions = iota
    // Microsoft Intune Subscription
    INTUNE_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Office365 Subscription
    OFFICE365_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Microsoft Intune Premium Subscription
    INTUNEPREMIUM_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Microsoft Intune for Education Subscription
    INTUNE_EDU_DEVICEMANAGEMENTSUBSCRIPTIONS
    // Microsoft Intune for Small Businesses Subscription
    INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS
)

func (i DeviceManagementSubscriptions) String() string {
    return []string{"none", "intune", "office365", "intunePremium", "intune_EDU", "intune_SMB"}[i]
}
func ParseDeviceManagementSubscriptions(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTSUBSCRIPTIONS
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "intune":
            result = INTUNE_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "office365":
            result = OFFICE365_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "intunePremium":
            result = INTUNEPREMIUM_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "intune_EDU":
            result = INTUNE_EDU_DEVICEMANAGEMENTSUBSCRIPTIONS
        case "intune_SMB":
            result = INTUNE_SMB_DEVICEMANAGEMENTSUBSCRIPTIONS
        default:
            return 0, errors.New("Unknown DeviceManagementSubscriptions value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementSubscriptions(values []DeviceManagementSubscriptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
