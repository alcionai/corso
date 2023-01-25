package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EnrollmentRestrictionPlatformType int

const (
    // Indicates that the enrollment configuration applies to all platforms
    ALLPLATFORMS_ENROLLMENTRESTRICTIONPLATFORMTYPE EnrollmentRestrictionPlatformType = iota
    // Indicates that the enrollment configuration applies only to iOS/iPadOS devices
    IOS_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Indicates that the enrollment configuration applies only to Windows devices
    WINDOWS_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Indicates that the enrollment configuration applies only to Windows Phone devices
    WINDOWSPHONE_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Indicates that the enrollment configuration applies only to Android devices
    ANDROID_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Indicates that the enrollment configuration applies only to Android for Work devices
    ANDROIDFORWORK_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Indicates that the enrollment configuration applies only to macOS devices
    MAC_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Indicates that the enrollment configuration applies only to Linux devices
    LINUX_ENROLLMENTRESTRICTIONPLATFORMTYPE
    // Evolvable enumeration sentinel value. Do not use
    UNKNOWNFUTUREVALUE_ENROLLMENTRESTRICTIONPLATFORMTYPE
)

func (i EnrollmentRestrictionPlatformType) String() string {
    return []string{"allPlatforms", "ios", "windows", "windowsPhone", "android", "androidForWork", "mac", "linux", "unknownFutureValue"}[i]
}
func ParseEnrollmentRestrictionPlatformType(v string) (interface{}, error) {
    result := ALLPLATFORMS_ENROLLMENTRESTRICTIONPLATFORMTYPE
    switch v {
        case "allPlatforms":
            result = ALLPLATFORMS_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "ios":
            result = IOS_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "windows":
            result = WINDOWS_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "windowsPhone":
            result = WINDOWSPHONE_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "android":
            result = ANDROID_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "androidForWork":
            result = ANDROIDFORWORK_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "mac":
            result = MAC_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "linux":
            result = LINUX_ENROLLMENTRESTRICTIONPLATFORMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENROLLMENTRESTRICTIONPLATFORMTYPE
        default:
            return 0, errors.New("Unknown EnrollmentRestrictionPlatformType value: " + v)
    }
    return &result, nil
}
func SerializeEnrollmentRestrictionPlatformType(values []EnrollmentRestrictionPlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
