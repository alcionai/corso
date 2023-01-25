package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidDeviceOwnerEnrollmentProfileType int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE AndroidDeviceOwnerEnrollmentProfileType = iota
    // Dedicated device.
    DEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
    // Fully managed.
    FULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
)

func (i AndroidDeviceOwnerEnrollmentProfileType) String() string {
    return []string{"notConfigured", "dedicatedDevice", "fullyManaged"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentProfileType(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
        case "dedicatedDevice":
            result = DEDICATEDDEVICE_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
        case "fullyManaged":
            result = FULLYMANAGED_ANDROIDDEVICEOWNERENROLLMENTPROFILETYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerEnrollmentProfileType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerEnrollmentProfileType(values []AndroidDeviceOwnerEnrollmentProfileType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
