package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type VpnProviderType int

const (
    // Tunnel traffic is not explicitly configured.
    NOTCONFIGURED_VPNPROVIDERTYPE VpnProviderType = iota
    // Tunnel traffic at the application layer.
    APPPROXY_VPNPROVIDERTYPE
    // Tunnel traffic at the IP layer.
    PACKETTUNNEL_VPNPROVIDERTYPE
)

func (i VpnProviderType) String() string {
    return []string{"notConfigured", "appProxy", "packetTunnel"}[i]
}
func ParseVpnProviderType(v string) (interface{}, error) {
    result := NOTCONFIGURED_VPNPROVIDERTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_VPNPROVIDERTYPE
        case "appProxy":
            result = APPPROXY_VPNPROVIDERTYPE
        case "packetTunnel":
            result = PACKETTUNNEL_VPNPROVIDERTYPE
        default:
            return 0, errors.New("Unknown VpnProviderType value: " + v)
    }
    return &result, nil
}
func SerializeVpnProviderType(values []VpnProviderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
