package models
import (
    "errors"
)
// Provides operations to call the add method.
type AospDeviceOwnerWiFiSecurityType int

const (
    // Open (No Authentication).
    OPEN_AOSPDEVICEOWNERWIFISECURITYTYPE AospDeviceOwnerWiFiSecurityType = iota
    // WEP Encryption.
    WEP_AOSPDEVICEOWNERWIFISECURITYTYPE
    // WPA-Personal/WPA2-Personal.
    WPAPERSONAL_AOSPDEVICEOWNERWIFISECURITYTYPE
    // WPA-Enterprise/WPA2-Enterprise. Must use AOSPDeviceOwnerEnterpriseWifiConfiguration type to configure enterprise options.
    WPAENTERPRISE_AOSPDEVICEOWNERWIFISECURITYTYPE
)

func (i AospDeviceOwnerWiFiSecurityType) String() string {
    return []string{"open", "wep", "wpaPersonal", "wpaEnterprise"}[i]
}
func ParseAospDeviceOwnerWiFiSecurityType(v string) (interface{}, error) {
    result := OPEN_AOSPDEVICEOWNERWIFISECURITYTYPE
    switch v {
        case "open":
            result = OPEN_AOSPDEVICEOWNERWIFISECURITYTYPE
        case "wep":
            result = WEP_AOSPDEVICEOWNERWIFISECURITYTYPE
        case "wpaPersonal":
            result = WPAPERSONAL_AOSPDEVICEOWNERWIFISECURITYTYPE
        case "wpaEnterprise":
            result = WPAENTERPRISE_AOSPDEVICEOWNERWIFISECURITYTYPE
        default:
            return 0, errors.New("Unknown AospDeviceOwnerWiFiSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeAospDeviceOwnerWiFiSecurityType(values []AospDeviceOwnerWiFiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
