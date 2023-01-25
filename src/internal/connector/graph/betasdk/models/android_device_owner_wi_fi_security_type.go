package models
import (
    "errors"
)
// Provides operations to call the add method.
type AndroidDeviceOwnerWiFiSecurityType int

const (
    // Open (No Authentication).
    OPEN_ANDROIDDEVICEOWNERWIFISECURITYTYPE AndroidDeviceOwnerWiFiSecurityType = iota
    // WEP Encryption.
    WEP_ANDROIDDEVICEOWNERWIFISECURITYTYPE
    // WPA-Personal/WPA2-Personal.
    WPAPERSONAL_ANDROIDDEVICEOWNERWIFISECURITYTYPE
    // WPA-Enterprise/WPA2-Enterprise. Must use AndroidDeviceOwnerEnterpriseWifiConfiguration type to configure enterprise options.
    WPAENTERPRISE_ANDROIDDEVICEOWNERWIFISECURITYTYPE
)

func (i AndroidDeviceOwnerWiFiSecurityType) String() string {
    return []string{"open", "wep", "wpaPersonal", "wpaEnterprise"}[i]
}
func ParseAndroidDeviceOwnerWiFiSecurityType(v string) (interface{}, error) {
    result := OPEN_ANDROIDDEVICEOWNERWIFISECURITYTYPE
    switch v {
        case "open":
            result = OPEN_ANDROIDDEVICEOWNERWIFISECURITYTYPE
        case "wep":
            result = WEP_ANDROIDDEVICEOWNERWIFISECURITYTYPE
        case "wpaPersonal":
            result = WPAPERSONAL_ANDROIDDEVICEOWNERWIFISECURITYTYPE
        case "wpaEnterprise":
            result = WPAENTERPRISE_ANDROIDDEVICEOWNERWIFISECURITYTYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerWiFiSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerWiFiSecurityType(values []AndroidDeviceOwnerWiFiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
