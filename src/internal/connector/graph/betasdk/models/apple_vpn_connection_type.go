package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AppleVpnConnectionType int

const (
    // Cisco AnyConnect.
    CISCOANYCONNECT_APPLEVPNCONNECTIONTYPE AppleVpnConnectionType = iota
    // Pulse Secure.
    PULSESECURE_APPLEVPNCONNECTIONTYPE
    // F5 Edge Client.
    F5EDGECLIENT_APPLEVPNCONNECTIONTYPE
    // Dell SonicWALL Mobile Connection.
    DELLSONICWALLMOBILECONNECT_APPLEVPNCONNECTIONTYPE
    // Check Point Capsule VPN.
    CHECKPOINTCAPSULEVPN_APPLEVPNCONNECTIONTYPE
    // Custom VPN.
    CUSTOMVPN_APPLEVPNCONNECTIONTYPE
    // Cisco (IPSec).
    CISCOIPSEC_APPLEVPNCONNECTIONTYPE
    // Citrix.
    CITRIX_APPLEVPNCONNECTIONTYPE
    // Cisco AnyConnect V2.
    CISCOANYCONNECTV2_APPLEVPNCONNECTIONTYPE
    // Palo Alto Networks GlobalProtect.
    PALOALTOGLOBALPROTECT_APPLEVPNCONNECTIONTYPE
    // Zscaler Private Access.
    ZSCALERPRIVATEACCESS_APPLEVPNCONNECTIONTYPE
    // F5 Access 2018.
    F5ACCESS2018_APPLEVPNCONNECTIONTYPE
    // Citrix Sso.
    CITRIXSSO_APPLEVPNCONNECTIONTYPE
    // Palo Alto Networks GlobalProtect V2.
    PALOALTOGLOBALPROTECTV2_APPLEVPNCONNECTIONTYPE
    // IKEv2.
    IKEV2_APPLEVPNCONNECTIONTYPE
    // AlwaysOn.
    ALWAYSON_APPLEVPNCONNECTIONTYPE
    // Microsoft Tunnel.
    MICROSOFTTUNNEL_APPLEVPNCONNECTIONTYPE
    // NetMotion Mobility.
    NETMOTIONMOBILITY_APPLEVPNCONNECTIONTYPE
    // Microsoft Protect.
    MICROSOFTPROTECT_APPLEVPNCONNECTIONTYPE
)

func (i AppleVpnConnectionType) String() string {
    return []string{"ciscoAnyConnect", "pulseSecure", "f5EdgeClient", "dellSonicWallMobileConnect", "checkPointCapsuleVpn", "customVpn", "ciscoIPSec", "citrix", "ciscoAnyConnectV2", "paloAltoGlobalProtect", "zscalerPrivateAccess", "f5Access2018", "citrixSso", "paloAltoGlobalProtectV2", "ikEv2", "alwaysOn", "microsoftTunnel", "netMotionMobility", "microsoftProtect"}[i]
}
func ParseAppleVpnConnectionType(v string) (interface{}, error) {
    result := CISCOANYCONNECT_APPLEVPNCONNECTIONTYPE
    switch v {
        case "ciscoAnyConnect":
            result = CISCOANYCONNECT_APPLEVPNCONNECTIONTYPE
        case "pulseSecure":
            result = PULSESECURE_APPLEVPNCONNECTIONTYPE
        case "f5EdgeClient":
            result = F5EDGECLIENT_APPLEVPNCONNECTIONTYPE
        case "dellSonicWallMobileConnect":
            result = DELLSONICWALLMOBILECONNECT_APPLEVPNCONNECTIONTYPE
        case "checkPointCapsuleVpn":
            result = CHECKPOINTCAPSULEVPN_APPLEVPNCONNECTIONTYPE
        case "customVpn":
            result = CUSTOMVPN_APPLEVPNCONNECTIONTYPE
        case "ciscoIPSec":
            result = CISCOIPSEC_APPLEVPNCONNECTIONTYPE
        case "citrix":
            result = CITRIX_APPLEVPNCONNECTIONTYPE
        case "ciscoAnyConnectV2":
            result = CISCOANYCONNECTV2_APPLEVPNCONNECTIONTYPE
        case "paloAltoGlobalProtect":
            result = PALOALTOGLOBALPROTECT_APPLEVPNCONNECTIONTYPE
        case "zscalerPrivateAccess":
            result = ZSCALERPRIVATEACCESS_APPLEVPNCONNECTIONTYPE
        case "f5Access2018":
            result = F5ACCESS2018_APPLEVPNCONNECTIONTYPE
        case "citrixSso":
            result = CITRIXSSO_APPLEVPNCONNECTIONTYPE
        case "paloAltoGlobalProtectV2":
            result = PALOALTOGLOBALPROTECTV2_APPLEVPNCONNECTIONTYPE
        case "ikEv2":
            result = IKEV2_APPLEVPNCONNECTIONTYPE
        case "alwaysOn":
            result = ALWAYSON_APPLEVPNCONNECTIONTYPE
        case "microsoftTunnel":
            result = MICROSOFTTUNNEL_APPLEVPNCONNECTIONTYPE
        case "netMotionMobility":
            result = NETMOTIONMOBILITY_APPLEVPNCONNECTIONTYPE
        case "microsoftProtect":
            result = MICROSOFTPROTECT_APPLEVPNCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown AppleVpnConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeAppleVpnConnectionType(values []AppleVpnConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
