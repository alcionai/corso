package models
import (
    "errors"
)
// Provides operations to call the add method.
type Windows10VpnConnectionType int

const (
    // Pulse Secure.
    PULSESECURE_WINDOWS10VPNCONNECTIONTYPE Windows10VpnConnectionType = iota
    // F5 Edge Client.
    F5EDGECLIENT_WINDOWS10VPNCONNECTIONTYPE
    // Dell SonicWALL Mobile Connection.
    DELLSONICWALLMOBILECONNECT_WINDOWS10VPNCONNECTIONTYPE
    // Check Point Capsule VPN.
    CHECKPOINTCAPSULEVPN_WINDOWS10VPNCONNECTIONTYPE
    // Automatic.
    AUTOMATIC_WINDOWS10VPNCONNECTIONTYPE
    // IKEv2.
    IKEV2_WINDOWS10VPNCONNECTIONTYPE
    // L2TP.
    L2TP_WINDOWS10VPNCONNECTIONTYPE
    // PPTP.
    PPTP_WINDOWS10VPNCONNECTIONTYPE
    // Citrix.
    CITRIX_WINDOWS10VPNCONNECTIONTYPE
    // Palo Alto Networks GlobalProtect.
    PALOALTOGLOBALPROTECT_WINDOWS10VPNCONNECTIONTYPE
    // Cisco AnyConnect
    CISCOANYCONNECT_WINDOWS10VPNCONNECTIONTYPE
    // Sentinel member for cases where the client cannot handle the new enum values.
    UNKNOWNFUTUREVALUE_WINDOWS10VPNCONNECTIONTYPE
    // Microsoft Tunnel connection type
    MICROSOFTTUNNEL_WINDOWS10VPNCONNECTIONTYPE
)

func (i Windows10VpnConnectionType) String() string {
    return []string{"pulseSecure", "f5EdgeClient", "dellSonicWallMobileConnect", "checkPointCapsuleVpn", "automatic", "ikEv2", "l2tp", "pptp", "citrix", "paloAltoGlobalProtect", "ciscoAnyConnect", "unknownFutureValue", "microsoftTunnel"}[i]
}
func ParseWindows10VpnConnectionType(v string) (interface{}, error) {
    result := PULSESECURE_WINDOWS10VPNCONNECTIONTYPE
    switch v {
        case "pulseSecure":
            result = PULSESECURE_WINDOWS10VPNCONNECTIONTYPE
        case "f5EdgeClient":
            result = F5EDGECLIENT_WINDOWS10VPNCONNECTIONTYPE
        case "dellSonicWallMobileConnect":
            result = DELLSONICWALLMOBILECONNECT_WINDOWS10VPNCONNECTIONTYPE
        case "checkPointCapsuleVpn":
            result = CHECKPOINTCAPSULEVPN_WINDOWS10VPNCONNECTIONTYPE
        case "automatic":
            result = AUTOMATIC_WINDOWS10VPNCONNECTIONTYPE
        case "ikEv2":
            result = IKEV2_WINDOWS10VPNCONNECTIONTYPE
        case "l2tp":
            result = L2TP_WINDOWS10VPNCONNECTIONTYPE
        case "pptp":
            result = PPTP_WINDOWS10VPNCONNECTIONTYPE
        case "citrix":
            result = CITRIX_WINDOWS10VPNCONNECTIONTYPE
        case "paloAltoGlobalProtect":
            result = PALOALTOGLOBALPROTECT_WINDOWS10VPNCONNECTIONTYPE
        case "ciscoAnyConnect":
            result = CISCOANYCONNECT_WINDOWS10VPNCONNECTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWS10VPNCONNECTIONTYPE
        case "microsoftTunnel":
            result = MICROSOFTTUNNEL_WINDOWS10VPNCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown Windows10VpnConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeWindows10VpnConnectionType(values []Windows10VpnConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
