package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceTypes int

const (
    // Desktop.
    DESKTOP_DEVICETYPES DeviceTypes = iota
    // WindowsRT.
    WINDOWSRT_DEVICETYPES
    // WinMO6.
    WINMO6_DEVICETYPES
    // Nokia.
    NOKIA_DEVICETYPES
    // Windows phone.
    WINDOWSPHONE_DEVICETYPES
    // Mac.
    MAC_DEVICETYPES
    // WinCE.
    WINCE_DEVICETYPES
    // WinEmbedded.
    WINEMBEDDED_DEVICETYPES
    // iPhone.
    IPHONE_DEVICETYPES
    // iPad.
    IPAD_DEVICETYPES
    // iPodTouch.
    IPOD_DEVICETYPES
    // Android.
    ANDROID_DEVICETYPES
    // iSocConsumer.
    ISOCCONSUMER_DEVICETYPES
    // Unix.
    UNIX_DEVICETYPES
    // Mac OS X client using built in MDM agent.
    MACMDM_DEVICETYPES
    // Representing the fancy Windows 10 goggles.
    HOLOLENS_DEVICETYPES
    // Surface HUB device.
    SURFACEHUB_DEVICETYPES
    // Android for work device.
    ANDROIDFORWORK_DEVICETYPES
    // Android enterprise device.
    ANDROIDENTERPRISE_DEVICETYPES
    // Blackberry.
    BLACKBERRY_DEVICETYPES
    // Palm.
    PALM_DEVICETYPES
    // Represents that the device type is unknown.
    UNKNOWN_DEVICETYPES
)

func (i DeviceTypes) String() string {
    return []string{"desktop", "windowsRT", "winMO6", "nokia", "windowsPhone", "mac", "winCE", "winEmbedded", "iPhone", "iPad", "iPod", "android", "iSocConsumer", "unix", "macMDM", "holoLens", "surfaceHub", "androidForWork", "androidEnterprise", "blackberry", "palm", "unknown"}[i]
}
func ParseDeviceTypes(v string) (interface{}, error) {
    result := DESKTOP_DEVICETYPES
    switch v {
        case "desktop":
            result = DESKTOP_DEVICETYPES
        case "windowsRT":
            result = WINDOWSRT_DEVICETYPES
        case "winMO6":
            result = WINMO6_DEVICETYPES
        case "nokia":
            result = NOKIA_DEVICETYPES
        case "windowsPhone":
            result = WINDOWSPHONE_DEVICETYPES
        case "mac":
            result = MAC_DEVICETYPES
        case "winCE":
            result = WINCE_DEVICETYPES
        case "winEmbedded":
            result = WINEMBEDDED_DEVICETYPES
        case "iPhone":
            result = IPHONE_DEVICETYPES
        case "iPad":
            result = IPAD_DEVICETYPES
        case "iPod":
            result = IPOD_DEVICETYPES
        case "android":
            result = ANDROID_DEVICETYPES
        case "iSocConsumer":
            result = ISOCCONSUMER_DEVICETYPES
        case "unix":
            result = UNIX_DEVICETYPES
        case "macMDM":
            result = MACMDM_DEVICETYPES
        case "holoLens":
            result = HOLOLENS_DEVICETYPES
        case "surfaceHub":
            result = SURFACEHUB_DEVICETYPES
        case "androidForWork":
            result = ANDROIDFORWORK_DEVICETYPES
        case "androidEnterprise":
            result = ANDROIDENTERPRISE_DEVICETYPES
        case "blackberry":
            result = BLACKBERRY_DEVICETYPES
        case "palm":
            result = PALM_DEVICETYPES
        case "unknown":
            result = UNKNOWN_DEVICETYPES
        default:
            return 0, errors.New("Unknown DeviceTypes value: " + v)
    }
    return &result, nil
}
func SerializeDeviceTypes(values []DeviceTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
