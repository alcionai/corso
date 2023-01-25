package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceLicensingStatus int

const (
    // Default. Set to unknown when status cannot be determined.
    UNKNOWN_DEVICELICENSINGSTATUS DeviceLicensingStatus = iota
    // This status is set when the license refresh is started.
    LICENSEREFRESHSTARTED_DEVICELICENSINGSTATUS
    // This status is set when the license refresh is pending.
    LICENSEREFRESHPENDING_DEVICELICENSINGSTATUS
    // This status is set when the device is not joined to Azure Active Directory.
    DEVICEISNOTAZUREACTIVEDIRECTORYJOINED_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft device identity is being verified.
    VERIFYINGMICROSOFTDEVICEIDENTITY_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft device identity verification fails.
    DEVICEIDENTITYVERIFICATIONFAILED_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft account identity is being verified.
    VERIFYINGMIROSOFTACCOUNTIDENTITY_DEVICELICENSINGSTATUS
    // This status is set when the Microsoft account identity verification fails.
    MIROSOFTACCOUNTVERIFICATIONFAILED_DEVICELICENSINGSTATUS
    // This status is set when the device license is being acquired.
    ACQUIRINGDEVICELICENSE_DEVICELICENSINGSTATUS
    // This status is set when the device license is being refreshed.
    REFRESHINGDEVICELICENSE_DEVICELICENSINGSTATUS
    // This status is set when the device license refresh succeeds.
    DEVICELICENSEREFRESHSUCCEED_DEVICELICENSINGSTATUS
    // This status is set when the device license refresh fails.
    DEVICELICENSEREFRESHFAILED_DEVICELICENSINGSTATUS
    // This status is set when the device license is being removed.
    REMOVINGDEVICELICENSE_DEVICELICENSINGSTATUS
    // This status is set when the device license removing succeeds.
    DEVICELICENSEREMOVESUCCEED_DEVICELICENSINGSTATUS
    // This status is set when the device license removing fails.
    DEVICELICENSEREMOVEFAILED_DEVICELICENSINGSTATUS
    // This is put here as a place holder for future extension.
    UNKNOWNFUTUREVALUE_DEVICELICENSINGSTATUS
)

func (i DeviceLicensingStatus) String() string {
    return []string{"unknown", "licenseRefreshStarted", "licenseRefreshPending", "deviceIsNotAzureActiveDirectoryJoined", "verifyingMicrosoftDeviceIdentity", "deviceIdentityVerificationFailed", "verifyingMirosoftAccountIdentity", "mirosoftAccountVerificationFailed", "acquiringDeviceLicense", "refreshingDeviceLicense", "deviceLicenseRefreshSucceed", "deviceLicenseRefreshFailed", "removingDeviceLicense", "deviceLicenseRemoveSucceed", "deviceLicenseRemoveFailed", "unknownFutureValue"}[i]
}
func ParseDeviceLicensingStatus(v string) (interface{}, error) {
    result := UNKNOWN_DEVICELICENSINGSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICELICENSINGSTATUS
        case "licenseRefreshStarted":
            result = LICENSEREFRESHSTARTED_DEVICELICENSINGSTATUS
        case "licenseRefreshPending":
            result = LICENSEREFRESHPENDING_DEVICELICENSINGSTATUS
        case "deviceIsNotAzureActiveDirectoryJoined":
            result = DEVICEISNOTAZUREACTIVEDIRECTORYJOINED_DEVICELICENSINGSTATUS
        case "verifyingMicrosoftDeviceIdentity":
            result = VERIFYINGMICROSOFTDEVICEIDENTITY_DEVICELICENSINGSTATUS
        case "deviceIdentityVerificationFailed":
            result = DEVICEIDENTITYVERIFICATIONFAILED_DEVICELICENSINGSTATUS
        case "verifyingMirosoftAccountIdentity":
            result = VERIFYINGMIROSOFTACCOUNTIDENTITY_DEVICELICENSINGSTATUS
        case "mirosoftAccountVerificationFailed":
            result = MIROSOFTACCOUNTVERIFICATIONFAILED_DEVICELICENSINGSTATUS
        case "acquiringDeviceLicense":
            result = ACQUIRINGDEVICELICENSE_DEVICELICENSINGSTATUS
        case "refreshingDeviceLicense":
            result = REFRESHINGDEVICELICENSE_DEVICELICENSINGSTATUS
        case "deviceLicenseRefreshSucceed":
            result = DEVICELICENSEREFRESHSUCCEED_DEVICELICENSINGSTATUS
        case "deviceLicenseRefreshFailed":
            result = DEVICELICENSEREFRESHFAILED_DEVICELICENSINGSTATUS
        case "removingDeviceLicense":
            result = REMOVINGDEVICELICENSE_DEVICELICENSINGSTATUS
        case "deviceLicenseRemoveSucceed":
            result = DEVICELICENSEREMOVESUCCEED_DEVICELICENSINGSTATUS
        case "deviceLicenseRemoveFailed":
            result = DEVICELICENSEREMOVEFAILED_DEVICELICENSINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICELICENSINGSTATUS
        default:
            return 0, errors.New("Unknown DeviceLicensingStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceLicensingStatus(values []DeviceLicensingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
