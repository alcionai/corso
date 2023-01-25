package models
import (
    "errors"
)
// Provides operations to call the add method.
type WiFiSecurityType int

const (
    // Open (No Authentication).
    OPEN_WIFISECURITYTYPE WiFiSecurityType = iota
    // WPA-Personal.
    WPAPERSONAL_WIFISECURITYTYPE
    // WPA-Enterprise. Must use IOSEnterpriseWifiConfiguration type to configure enterprise options.
    WPAENTERPRISE_WIFISECURITYTYPE
    // WEP Encryption.
    WEP_WIFISECURITYTYPE
    // WPA2-Personal.
    WPA2PERSONAL_WIFISECURITYTYPE
    // WPA2-Enterprise. Must use WindowsWifiEnterpriseEAPConfiguration type to configure enterprise options.
    WPA2ENTERPRISE_WIFISECURITYTYPE
)

func (i WiFiSecurityType) String() string {
    return []string{"open", "wpaPersonal", "wpaEnterprise", "wep", "wpa2Personal", "wpa2Enterprise"}[i]
}
func ParseWiFiSecurityType(v string) (interface{}, error) {
    result := OPEN_WIFISECURITYTYPE
    switch v {
        case "open":
            result = OPEN_WIFISECURITYTYPE
        case "wpaPersonal":
            result = WPAPERSONAL_WIFISECURITYTYPE
        case "wpaEnterprise":
            result = WPAENTERPRISE_WIFISECURITYTYPE
        case "wep":
            result = WEP_WIFISECURITYTYPE
        case "wpa2Personal":
            result = WPA2PERSONAL_WIFISECURITYTYPE
        case "wpa2Enterprise":
            result = WPA2ENTERPRISE_WIFISECURITYTYPE
        default:
            return 0, errors.New("Unknown WiFiSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeWiFiSecurityType(values []WiFiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
