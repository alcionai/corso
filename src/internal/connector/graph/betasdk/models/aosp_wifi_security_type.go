package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AospWifiSecurityType int

const (
    // No security type.
    NONE_AOSPWIFISECURITYTYPE AospWifiSecurityType = iota
    // WPA-Pre-shared-key
    WPA_AOSPWIFISECURITYTYPE
    // WEP-Pre-shared-key
    WEP_AOSPWIFISECURITYTYPE
)

func (i AospWifiSecurityType) String() string {
    return []string{"none", "wpa", "wep"}[i]
}
func ParseAospWifiSecurityType(v string) (interface{}, error) {
    result := NONE_AOSPWIFISECURITYTYPE
    switch v {
        case "none":
            result = NONE_AOSPWIFISECURITYTYPE
        case "wpa":
            result = WPA_AOSPWIFISECURITYTYPE
        case "wep":
            result = WEP_AOSPWIFISECURITYTYPE
        default:
            return 0, errors.New("Unknown AospWifiSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeAospWifiSecurityType(values []AospWifiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
