package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidDeviceOwnerKioskModeIconSize int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE AndroidDeviceOwnerKioskModeIconSize = iota
    // Smallest icon size.
    SMALLEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Small icon size.
    SMALL_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Regular icon size.
    REGULAR_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Large icon size.
    LARGE_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    // Largest icon size.
    LARGEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
)

func (i AndroidDeviceOwnerKioskModeIconSize) String() string {
    return []string{"notConfigured", "smallest", "small", "regular", "large", "largest"}[i]
}
func ParseAndroidDeviceOwnerKioskModeIconSize(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "smallest":
            result = SMALLEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "small":
            result = SMALL_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "regular":
            result = REGULAR_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "large":
            result = LARGE_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        case "largest":
            result = LARGEST_ANDROIDDEVICEOWNERKIOSKMODEICONSIZE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerKioskModeIconSize value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerKioskModeIconSize(values []AndroidDeviceOwnerKioskModeIconSize) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
