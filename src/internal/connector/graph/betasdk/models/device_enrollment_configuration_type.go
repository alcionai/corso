package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeviceEnrollmentConfigurationType int

const (
    // Default. Set to unknown if the configuration type cannot be determined.
    UNKNOWN_DEVICEENROLLMENTCONFIGURATIONTYPE DeviceEnrollmentConfigurationType = iota
    // Indicates that configuration is of type limit which refers to number of devices a user is allowed to enroll.
    LIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type platform restriction which refers to types of devices a user is allowed to enroll.
    PLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type Windows Hello which refers to authentication method devices would use.
    WINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type default limit which refers to types of devices a user is allowed to enroll by default.
    DEFAULTLIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type default platform restriction which refers to types of devices a user is allowed to enroll by default.
    DEFAULTPLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type default Windows Hello which refers to authentication method devices would use by default.
    DEFAULTWINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type default Enrollment status page which refers to startup page displayed during OOBE in Autopilot devices by default.
    DEFAULTWINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type Enrollment status page which refers to startup page displayed during OOBE in Autopilot devices.
    WINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type Comanagement Authority which refers to policies applied to Co-Managed devices.
    DEVICECOMANAGEMENTAUTHORITYCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type single platform restriction which refers to types of devices a user is allowed to enroll.
    SINGLEPLATFORMRESTRICTION_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Unknown future value
    UNKNOWNFUTUREVALUE_DEVICEENROLLMENTCONFIGURATIONTYPE
    // Indicates that configuration is of type Enrollment Notification which refers to types of notification a user receives during enrollment.
    ENROLLMENTNOTIFICATIONSCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
)

func (i DeviceEnrollmentConfigurationType) String() string {
    return []string{"unknown", "limit", "platformRestrictions", "windowsHelloForBusiness", "defaultLimit", "defaultPlatformRestrictions", "defaultWindowsHelloForBusiness", "defaultWindows10EnrollmentCompletionPageConfiguration", "windows10EnrollmentCompletionPageConfiguration", "deviceComanagementAuthorityConfiguration", "singlePlatformRestriction", "unknownFutureValue", "enrollmentNotificationsConfiguration"}[i]
}
func ParseDeviceEnrollmentConfigurationType(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEENROLLMENTCONFIGURATIONTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "limit":
            result = LIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "platformRestrictions":
            result = PLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "windowsHelloForBusiness":
            result = WINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "defaultLimit":
            result = DEFAULTLIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "defaultPlatformRestrictions":
            result = DEFAULTPLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "defaultWindowsHelloForBusiness":
            result = DEFAULTWINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "defaultWindows10EnrollmentCompletionPageConfiguration":
            result = DEFAULTWINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "windows10EnrollmentCompletionPageConfiguration":
            result = WINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "deviceComanagementAuthorityConfiguration":
            result = DEVICECOMANAGEMENTAUTHORITYCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "singlePlatformRestriction":
            result = SINGLEPLATFORMRESTRICTION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "enrollmentNotificationsConfiguration":
            result = ENROLLMENTNOTIFICATIONSCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        default:
            return 0, errors.New("Unknown DeviceEnrollmentConfigurationType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceEnrollmentConfigurationType(values []DeviceEnrollmentConfigurationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
