package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidDeviceOwnerBatteryPluggedMode int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE AndroidDeviceOwnerBatteryPluggedMode = iota
    // Power source is an AC charger.
    AC_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
    // Power source is a USB port.
    USB_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
    // Power source is wireless.
    WIRELESS_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
)

func (i AndroidDeviceOwnerBatteryPluggedMode) String() string {
    return []string{"notConfigured", "ac", "usb", "wireless"}[i]
}
func ParseAndroidDeviceOwnerBatteryPluggedMode(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        case "ac":
            result = AC_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        case "usb":
            result = USB_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        case "wireless":
            result = WIRELESS_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerBatteryPluggedMode value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerBatteryPluggedMode(values []AndroidDeviceOwnerBatteryPluggedMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
