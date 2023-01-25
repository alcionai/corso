package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidDeviceOwnerKioskCustomizationStatusBar int

const (
    // Not configured; this value defaults to STATUS_BAR_UNSPECIFIED.
    NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSTATUSBAR AndroidDeviceOwnerKioskCustomizationStatusBar = iota
    // System info and notifications are shown on the status bar in kiosk mode.
    NOTIFICATIONSANDSYSTEMINFOENABLED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSTATUSBAR
    // Only system info is shown on the status bar in kiosk mode.
    SYSTEMINFOONLY_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSTATUSBAR
)

func (i AndroidDeviceOwnerKioskCustomizationStatusBar) String() string {
    return []string{"notConfigured", "notificationsAndSystemInfoEnabled", "systemInfoOnly"}[i]
}
func ParseAndroidDeviceOwnerKioskCustomizationStatusBar(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSTATUSBAR
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSTATUSBAR
        case "notificationsAndSystemInfoEnabled":
            result = NOTIFICATIONSANDSYSTEMINFOENABLED_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSTATUSBAR
        case "systemInfoOnly":
            result = SYSTEMINFOONLY_ANDROIDDEVICEOWNERKIOSKCUSTOMIZATIONSTATUSBAR
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerKioskCustomizationStatusBar value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerKioskCustomizationStatusBar(values []AndroidDeviceOwnerKioskCustomizationStatusBar) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
