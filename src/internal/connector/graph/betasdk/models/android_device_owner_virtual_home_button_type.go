package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidDeviceOwnerVirtualHomeButtonType int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERVIRTUALHOMEBUTTONTYPE AndroidDeviceOwnerVirtualHomeButtonType = iota
    // Swipe-up for home button.
    SWIPEUP_ANDROIDDEVICEOWNERVIRTUALHOMEBUTTONTYPE
    // Floating home button.
    FLOATING_ANDROIDDEVICEOWNERVIRTUALHOMEBUTTONTYPE
)

func (i AndroidDeviceOwnerVirtualHomeButtonType) String() string {
    return []string{"notConfigured", "swipeUp", "floating"}[i]
}
func ParseAndroidDeviceOwnerVirtualHomeButtonType(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERVIRTUALHOMEBUTTONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERVIRTUALHOMEBUTTONTYPE
        case "swipeUp":
            result = SWIPEUP_ANDROIDDEVICEOWNERVIRTUALHOMEBUTTONTYPE
        case "floating":
            result = FLOATING_ANDROIDDEVICEOWNERVIRTUALHOMEBUTTONTYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerVirtualHomeButtonType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerVirtualHomeButtonType(values []AndroidDeviceOwnerVirtualHomeButtonType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
