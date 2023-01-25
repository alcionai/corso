package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidManagedAppSafetyNetDeviceAttestationType int

const (
    // no requirement set
    NONE_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE AndroidManagedAppSafetyNetDeviceAttestationType = iota
    // require that Android device passes SafetyNet Basic Integrity validation
    BASICINTEGRITY_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
    // require that Android device passes SafetyNet Basic Integrity and Device Certification validations
    BASICINTEGRITYANDDEVICECERTIFICATION_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
)

func (i AndroidManagedAppSafetyNetDeviceAttestationType) String() string {
    return []string{"none", "basicIntegrity", "basicIntegrityAndDeviceCertification"}[i]
}
func ParseAndroidManagedAppSafetyNetDeviceAttestationType(v string) (interface{}, error) {
    result := NONE_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
    switch v {
        case "none":
            result = NONE_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
        case "basicIntegrity":
            result = BASICINTEGRITY_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
        case "basicIntegrityAndDeviceCertification":
            result = BASICINTEGRITYANDDEVICECERTIFICATION_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
        default:
            return 0, errors.New("Unknown AndroidManagedAppSafetyNetDeviceAttestationType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedAppSafetyNetDeviceAttestationType(values []AndroidManagedAppSafetyNetDeviceAttestationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
