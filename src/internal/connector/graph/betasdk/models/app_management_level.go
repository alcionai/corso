package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AppManagementLevel int

const (
    // Unspecified
    UNSPECIFIED_APPMANAGEMENTLEVEL AppManagementLevel = iota
    // Unmanaged
    UNMANAGED_APPMANAGEMENTLEVEL
    // MDM
    MDM_APPMANAGEMENTLEVEL
    // Android Enterprise
    ANDROIDENTERPRISE_APPMANAGEMENTLEVEL
    // Android Enterprise dedicated devices with Azure AD Shared mode
    ANDROIDENTERPRISEDEDICATEDDEVICESWITHAZUREADSHAREDMODE_APPMANAGEMENTLEVEL
    // Android Open Source Project (AOSP) devices
    ANDROIDOPENSOURCEPROJECTUSERASSOCIATED_APPMANAGEMENTLEVEL
    // Android Open Source Project (AOSP) userless devices
    ANDROIDOPENSOURCEPROJECTUSERLESS_APPMANAGEMENTLEVEL
    // Place holder for evolvable enum
    UNKNOWNFUTUREVALUE_APPMANAGEMENTLEVEL
)

func (i AppManagementLevel) String() string {
    return []string{"unspecified", "unmanaged", "mdm", "androidEnterprise", "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode", "androidOpenSourceProjectUserAssociated", "androidOpenSourceProjectUserless", "unknownFutureValue"}[i]
}
func ParseAppManagementLevel(v string) (interface{}, error) {
    result := UNSPECIFIED_APPMANAGEMENTLEVEL
    switch v {
        case "unspecified":
            result = UNSPECIFIED_APPMANAGEMENTLEVEL
        case "unmanaged":
            result = UNMANAGED_APPMANAGEMENTLEVEL
        case "mdm":
            result = MDM_APPMANAGEMENTLEVEL
        case "androidEnterprise":
            result = ANDROIDENTERPRISE_APPMANAGEMENTLEVEL
        case "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode":
            result = ANDROIDENTERPRISEDEDICATEDDEVICESWITHAZUREADSHAREDMODE_APPMANAGEMENTLEVEL
        case "androidOpenSourceProjectUserAssociated":
            result = ANDROIDOPENSOURCEPROJECTUSERASSOCIATED_APPMANAGEMENTLEVEL
        case "androidOpenSourceProjectUserless":
            result = ANDROIDOPENSOURCEPROJECTUSERLESS_APPMANAGEMENTLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPMANAGEMENTLEVEL
        default:
            return 0, errors.New("Unknown AppManagementLevel value: " + v)
    }
    return &result, nil
}
func SerializeAppManagementLevel(values []AppManagementLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
