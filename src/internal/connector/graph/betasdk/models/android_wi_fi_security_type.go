package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidWiFiSecurityType int

const (
    // Open (No Authentication).
    OPEN_ANDROIDWIFISECURITYTYPE AndroidWiFiSecurityType = iota
    // WPA-Enterprise. Must use AndroidEnterpriseWifiConfiguration type to configure enterprise options.
    WPAENTERPRISE_ANDROIDWIFISECURITYTYPE
    // WPA2-Enterprise. Must use AndroidEnterpriseWifiConfiguration type to configure enterprise options.
    WPA2ENTERPRISE_ANDROIDWIFISECURITYTYPE
)

func (i AndroidWiFiSecurityType) String() string {
    return []string{"open", "wpaEnterprise", "wpa2Enterprise"}[i]
}
func ParseAndroidWiFiSecurityType(v string) (interface{}, error) {
    result := OPEN_ANDROIDWIFISECURITYTYPE
    switch v {
        case "open":
            result = OPEN_ANDROIDWIFISECURITYTYPE
        case "wpaEnterprise":
            result = WPAENTERPRISE_ANDROIDWIFISECURITYTYPE
        case "wpa2Enterprise":
            result = WPA2ENTERPRISE_ANDROIDWIFISECURITYTYPE
        default:
            return 0, errors.New("Unknown AndroidWiFiSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidWiFiSecurityType(values []AndroidWiFiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
