package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Platform int

const (
    // Unknown device platform
    UNKNOWN_PLATFORM Platform = iota
    // IOS device platform
    IOS_PLATFORM
    // Android device platform
    ANDROID_PLATFORM
    // Windows device platform
    WINDOWS_PLATFORM
    // WindowsMobile device platform
    WINDOWSMOBILE_PLATFORM
    // Mac device platform
    MACOS_PLATFORM
)

func (i Platform) String() string {
    return []string{"unknown", "ios", "android", "windows", "windowsMobile", "macOS"}[i]
}
func ParsePlatform(v string) (interface{}, error) {
    result := UNKNOWN_PLATFORM
    switch v {
        case "unknown":
            result = UNKNOWN_PLATFORM
        case "ios":
            result = IOS_PLATFORM
        case "android":
            result = ANDROID_PLATFORM
        case "windows":
            result = WINDOWS_PLATFORM
        case "windowsMobile":
            result = WINDOWSMOBILE_PLATFORM
        case "macOS":
            result = MACOS_PLATFORM
        default:
            return 0, errors.New("Unknown Platform value: " + v)
    }
    return &result, nil
}
func SerializePlatform(values []Platform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
