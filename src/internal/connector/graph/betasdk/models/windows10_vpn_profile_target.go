package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Windows10VpnProfileTarget int

const (
    // User targeted VPN profile.
    USER_WINDOWS10VPNPROFILETARGET Windows10VpnProfileTarget = iota
    // Device targeted VPN profile.
    DEVICE_WINDOWS10VPNPROFILETARGET
    // AutoPilot Device targeted VPN profile.
    AUTOPILOTDEVICE_WINDOWS10VPNPROFILETARGET
)

func (i Windows10VpnProfileTarget) String() string {
    return []string{"user", "device", "autoPilotDevice"}[i]
}
func ParseWindows10VpnProfileTarget(v string) (interface{}, error) {
    result := USER_WINDOWS10VPNPROFILETARGET
    switch v {
        case "user":
            result = USER_WINDOWS10VPNPROFILETARGET
        case "device":
            result = DEVICE_WINDOWS10VPNPROFILETARGET
        case "autoPilotDevice":
            result = AUTOPILOTDEVICE_WINDOWS10VPNPROFILETARGET
        default:
            return 0, errors.New("Unknown Windows10VpnProfileTarget value: " + v)
    }
    return &result, nil
}
func SerializeWindows10VpnProfileTarget(values []Windows10VpnProfileTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
