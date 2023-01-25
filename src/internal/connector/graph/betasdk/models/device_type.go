package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceType int

const (
    // Desktop.
    DESKTOP_DEVICETYPE DeviceType = iota
    // WindowsRT.
    WINDOWSRT_DEVICETYPE
    // WinMO6.
    WINMO6_DEVICETYPE
    // Nokia.
    NOKIA_DEVICETYPE
    // Windows phone.
    WINDOWSPHONE_DEVICETYPE
    // Mac.
    MAC_DEVICETYPE
    // WinCE.
    WINCE_DEVICETYPE
    // WinEmbedded.
    WINEMBEDDED_DEVICETYPE
    // iPhone.
    IPHONE_DEVICETYPE
    // iPad.
    IPAD_DEVICETYPE
    // iPodTouch.
    IPOD_DEVICETYPE
    // Android.
    ANDROID_DEVICETYPE
    // iSocConsumer.
    ISOCCONSUMER_DEVICETYPE
    // Unix.
    UNIX_DEVICETYPE
    // Mac OS X client using built in MDM agent.
    MACMDM_DEVICETYPE
    // Representing the fancy Windows 10 goggles.
    HOLOLENS_DEVICETYPE
    // Surface HUB device.
    SURFACEHUB_DEVICETYPE
    // Android for work device.
    ANDROIDFORWORK_DEVICETYPE
    // Android enterprise device.
    ANDROIDENTERPRISE_DEVICETYPE
    // Windows 10x device.
    WINDOWS10X_DEVICETYPE
    // Android non Google managed device.
    ANDROIDNGMS_DEVICETYPE
    // ChromeOS device.
    CHROMEOS_DEVICETYPE
    // Linux device.
    LINUX_DEVICETYPE
    // Blackberry.
    BLACKBERRY_DEVICETYPE
    // Palm.
    PALM_DEVICETYPE
    // Represents that the device type is unknown.
    UNKNOWN_DEVICETYPE
    // Cloud PC device.
    CLOUDPC_DEVICETYPE
)

func (i DeviceType) String() string {
    return []string{"desktop", "windowsRT", "winMO6", "nokia", "windowsPhone", "mac", "winCE", "winEmbedded", "iPhone", "iPad", "iPod", "android", "iSocConsumer", "unix", "macMDM", "holoLens", "surfaceHub", "androidForWork", "androidEnterprise", "windows10x", "androidnGMS", "chromeOS", "linux", "blackberry", "palm", "unknown", "cloudPC"}[i]
}
func ParseDeviceType(v string) (interface{}, error) {
    result := DESKTOP_DEVICETYPE
    switch v {
        case "desktop":
            result = DESKTOP_DEVICETYPE
        case "windowsRT":
            result = WINDOWSRT_DEVICETYPE
        case "winMO6":
            result = WINMO6_DEVICETYPE
        case "nokia":
            result = NOKIA_DEVICETYPE
        case "windowsPhone":
            result = WINDOWSPHONE_DEVICETYPE
        case "mac":
            result = MAC_DEVICETYPE
        case "winCE":
            result = WINCE_DEVICETYPE
        case "winEmbedded":
            result = WINEMBEDDED_DEVICETYPE
        case "iPhone":
            result = IPHONE_DEVICETYPE
        case "iPad":
            result = IPAD_DEVICETYPE
        case "iPod":
            result = IPOD_DEVICETYPE
        case "android":
            result = ANDROID_DEVICETYPE
        case "iSocConsumer":
            result = ISOCCONSUMER_DEVICETYPE
        case "unix":
            result = UNIX_DEVICETYPE
        case "macMDM":
            result = MACMDM_DEVICETYPE
        case "holoLens":
            result = HOLOLENS_DEVICETYPE
        case "surfaceHub":
            result = SURFACEHUB_DEVICETYPE
        case "androidForWork":
            result = ANDROIDFORWORK_DEVICETYPE
        case "androidEnterprise":
            result = ANDROIDENTERPRISE_DEVICETYPE
        case "windows10x":
            result = WINDOWS10X_DEVICETYPE
        case "androidnGMS":
            result = ANDROIDNGMS_DEVICETYPE
        case "chromeOS":
            result = CHROMEOS_DEVICETYPE
        case "linux":
            result = LINUX_DEVICETYPE
        case "blackberry":
            result = BLACKBERRY_DEVICETYPE
        case "palm":
            result = PALM_DEVICETYPE
        case "unknown":
            result = UNKNOWN_DEVICETYPE
        case "cloudPC":
            result = CLOUDPC_DEVICETYPE
        default:
            return 0, errors.New("Unknown DeviceType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceType(values []DeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
