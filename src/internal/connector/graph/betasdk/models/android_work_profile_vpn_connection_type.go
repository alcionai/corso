package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidWorkProfileVpnConnectionType int

const (
    // Cisco AnyConnect.
    CISCOANYCONNECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE AndroidWorkProfileVpnConnectionType = iota
    // Pulse Secure.
    PULSESECURE_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // F5 Edge Client.
    F5EDGECLIENT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // Dell SonicWALL Mobile Connection.
    DELLSONICWALLMOBILECONNECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // Check Point Capsule VPN.
    CHECKPOINTCAPSULEVPN_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // Citrix
    CITRIX_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // Palo Alto Networks GlobalProtect.
    PALOALTOGLOBALPROTECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // Microsoft Tunnel.
    MICROSOFTTUNNEL_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // NetMotion Mobility.
    NETMOTIONMOBILITY_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    // Microsoft Protect.
    MICROSOFTPROTECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
)

func (i AndroidWorkProfileVpnConnectionType) String() string {
    return []string{"ciscoAnyConnect", "pulseSecure", "f5EdgeClient", "dellSonicWallMobileConnect", "checkPointCapsuleVpn", "citrix", "paloAltoGlobalProtect", "microsoftTunnel", "netMotionMobility", "microsoftProtect"}[i]
}
func ParseAndroidWorkProfileVpnConnectionType(v string) (interface{}, error) {
    result := CISCOANYCONNECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
    switch v {
        case "ciscoAnyConnect":
            result = CISCOANYCONNECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "pulseSecure":
            result = PULSESECURE_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "f5EdgeClient":
            result = F5EDGECLIENT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "dellSonicWallMobileConnect":
            result = DELLSONICWALLMOBILECONNECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "checkPointCapsuleVpn":
            result = CHECKPOINTCAPSULEVPN_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "citrix":
            result = CITRIX_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "paloAltoGlobalProtect":
            result = PALOALTOGLOBALPROTECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "microsoftTunnel":
            result = MICROSOFTTUNNEL_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "netMotionMobility":
            result = NETMOTIONMOBILITY_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        case "microsoftProtect":
            result = MICROSOFTPROTECT_ANDROIDWORKPROFILEVPNCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown AndroidWorkProfileVpnConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidWorkProfileVpnConnectionType(values []AndroidWorkProfileVpnConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
