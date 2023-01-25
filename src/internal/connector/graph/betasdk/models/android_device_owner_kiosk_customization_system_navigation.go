package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidDeviceOwnerKioskCustomizationSystemNavigation int

const (
    // Not configured; this value defaults to NAVIGATION_DISABLED.
    NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSYSTEMNAVIGATION AndroidDeviceOwnerKioskCustomizationSystemNavigation = iota
    // Home and overview buttons are enabled.
    NAVIGATIONENABLED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSYSTEMNAVIGATION
    //  Only the home button is enabled.
    HOMEBUTTONONLY_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSYSTEMNAVIGATION
)

func (i AndroidDeviceOwnerKioskCustomizationSystemNavigation) String() string {
    return []string{"notConfigured", "navigationEnabled", "homeButtonOnly"}[i]
}
func ParseAndroidDeviceOwnerKioskCustomizationSystemNavigation(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSYSTEMNAVIGATION
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSYSTEMNAVIGATION
        case "navigationEnabled":
            result = NAVIGATIONENABLED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSYSTEMNAVIGATION
        case "homeButtonOnly":
            result = HOMEBUTTONONLY_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSYSTEMNAVIGATION
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerKioskCustomizationSystemNavigation value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerKioskCustomizationSystemNavigation(values []AndroidDeviceOwnerKioskCustomizationSystemNavigation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
