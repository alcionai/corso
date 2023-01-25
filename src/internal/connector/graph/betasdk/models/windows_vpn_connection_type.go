package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsVpnConnectionType int

const (
    // Pulse Secure.
    PULSESECURE_WINDOWSVPNCONNECTIONTYPE WindowsVpnConnectionType = iota
    // F5 Edge Client.
    F5EDGECLIENT_WINDOWSVPNCONNECTIONTYPE
    // Dell SonicWALL Mobile Connection.
    DELLSONICWALLMOBILECONNECT_WINDOWSVPNCONNECTIONTYPE
    // Check Point Capsule VPN.
    CHECKPOINTCAPSULEVPN_WINDOWSVPNCONNECTIONTYPE
)

func (i WindowsVpnConnectionType) String() string {
    return []string{"pulseSecure", "f5EdgeClient", "dellSonicWallMobileConnect", "checkPointCapsuleVpn"}[i]
}
func ParseWindowsVpnConnectionType(v string) (interface{}, error) {
    result := PULSESECURE_WINDOWSVPNCONNECTIONTYPE
    switch v {
        case "pulseSecure":
            result = PULSESECURE_WINDOWSVPNCONNECTIONTYPE
        case "f5EdgeClient":
            result = F5EDGECLIENT_WINDOWSVPNCONNECTIONTYPE
        case "dellSonicWallMobileConnect":
            result = DELLSONICWALLMOBILECONNECT_WINDOWSVPNCONNECTIONTYPE
        case "checkPointCapsuleVpn":
            result = CHECKPOINTCAPSULEVPN_WINDOWSVPNCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown WindowsVpnConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsVpnConnectionType(values []WindowsVpnConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
