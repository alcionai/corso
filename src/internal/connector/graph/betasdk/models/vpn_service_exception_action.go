package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type VpnServiceExceptionAction int

const (
    // Make all traffic from that service go through the VPN
    FORCETRAFFICVIAVPN_VPNSERVICEEXCEPTIONACTION VpnServiceExceptionAction = iota
    // Allow the service outside of the VPN
    ALLOWTRAFFICOUTSIDE_VPNSERVICEEXCEPTIONACTION
    // Drop all traffic from the service
    DROPTRAFFIC_VPNSERVICEEXCEPTIONACTION
)

func (i VpnServiceExceptionAction) String() string {
    return []string{"forceTrafficViaVPN", "allowTrafficOutside", "dropTraffic"}[i]
}
func ParseVpnServiceExceptionAction(v string) (interface{}, error) {
    result := FORCETRAFFICVIAVPN_VPNSERVICEEXCEPTIONACTION
    switch v {
        case "forceTrafficViaVPN":
            result = FORCETRAFFICVIAVPN_VPNSERVICEEXCEPTIONACTION
        case "allowTrafficOutside":
            result = ALLOWTRAFFICOUTSIDE_VPNSERVICEEXCEPTIONACTION
        case "dropTraffic":
            result = DROPTRAFFIC_VPNSERVICEEXCEPTIONACTION
        default:
            return 0, errors.New("Unknown VpnServiceExceptionAction value: " + v)
    }
    return &result, nil
}
func SerializeVpnServiceExceptionAction(values []VpnServiceExceptionAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
