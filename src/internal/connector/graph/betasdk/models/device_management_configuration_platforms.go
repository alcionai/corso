package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceManagementConfigurationPlatforms int

const (
    // None.
    NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS DeviceManagementConfigurationPlatforms = iota
    // Android.
    ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // iOS.
    IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // MacOS.
    MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Windows 10 X.
    WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Windows 10.
    WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Linux.
    LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    // Sentinel member for cases where the client cannot handle the new enum values.
    UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
)

func (i DeviceManagementConfigurationPlatforms) String() string {
    return []string{"none", "android", "iOS", "macOS", "windows10X", "windows10", "linux", "unknownFutureValue"}[i]
}
func ParseDeviceManagementConfigurationPlatforms(v string) (interface{}, error) {
    result := NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
    switch v {
        case "none":
            result = NONE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "android":
            result = ANDROID_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "iOS":
            result = IOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "macOS":
            result = MACOS_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "windows10X":
            result = WINDOWS10X_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "windows10":
            result = WINDOWS10_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "linux":
            result = LINUX_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEMANAGEMENTCONFIGURATIONPLATFORMS
        default:
            return 0, errors.New("Unknown DeviceManagementConfigurationPlatforms value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementConfigurationPlatforms(values []DeviceManagementConfigurationPlatforms) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
