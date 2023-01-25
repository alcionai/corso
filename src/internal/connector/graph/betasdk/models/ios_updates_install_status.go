package models
import (
    "errors"
)
// Provides operations to call the add method.
type IosUpdatesInstallStatus int

const (
    UPDATESCANFAILED_IOSUPDATESINSTALLSTATUS IosUpdatesInstallStatus = iota
    DEVICEOSHIGHERTHANDESIREDOSVERSION_IOSUPDATESINSTALLSTATUS
    UPDATEERROR_IOSUPDATESINSTALLSTATUS
    SHAREDDEVICEUSERLOGGEDINERROR_IOSUPDATESINSTALLSTATUS
    NOTSUPPORTEDOPERATION_IOSUPDATESINSTALLSTATUS
    INSTALLFAILED_IOSUPDATESINSTALLSTATUS
    INSTALLPHONECALLINPROGRESS_IOSUPDATESINSTALLSTATUS
    INSTALLINSUFFICIENTPOWER_IOSUPDATESINSTALLSTATUS
    INSTALLINSUFFICIENTSPACE_IOSUPDATESINSTALLSTATUS
    INSTALLING_IOSUPDATESINSTALLSTATUS
    DOWNLOADINSUFFICIENTNETWORK_IOSUPDATESINSTALLSTATUS
    DOWNLOADINSUFFICIENTPOWER_IOSUPDATESINSTALLSTATUS
    DOWNLOADINSUFFICIENTSPACE_IOSUPDATESINSTALLSTATUS
    DOWNLOADREQUIRESCOMPUTER_IOSUPDATESINSTALLSTATUS
    DOWNLOADFAILED_IOSUPDATESINSTALLSTATUS
    DOWNLOADING_IOSUPDATESINSTALLSTATUS
    TIMEOUT_IOSUPDATESINSTALLSTATUS
    MDMCLIENTCRASHED_IOSUPDATESINSTALLSTATUS
    SUCCESS_IOSUPDATESINSTALLSTATUS
    AVAILABLE_IOSUPDATESINSTALLSTATUS
    IDLE_IOSUPDATESINSTALLSTATUS
    UNKNOWN_IOSUPDATESINSTALLSTATUS
)

func (i IosUpdatesInstallStatus) String() string {
    return []string{"updateScanFailed", "deviceOsHigherThanDesiredOsVersion", "updateError", "sharedDeviceUserLoggedInError", "notSupportedOperation", "installFailed", "installPhoneCallInProgress", "installInsufficientPower", "installInsufficientSpace", "installing", "downloadInsufficientNetwork", "downloadInsufficientPower", "downloadInsufficientSpace", "downloadRequiresComputer", "downloadFailed", "downloading", "timeout", "mdmClientCrashed", "success", "available", "idle", "unknown"}[i]
}
func ParseIosUpdatesInstallStatus(v string) (interface{}, error) {
    result := UPDATESCANFAILED_IOSUPDATESINSTALLSTATUS
    switch v {
        case "updateScanFailed":
            result = UPDATESCANFAILED_IOSUPDATESINSTALLSTATUS
        case "deviceOsHigherThanDesiredOsVersion":
            result = DEVICEOSHIGHERTHANDESIREDOSVERSION_IOSUPDATESINSTALLSTATUS
        case "updateError":
            result = UPDATEERROR_IOSUPDATESINSTALLSTATUS
        case "sharedDeviceUserLoggedInError":
            result = SHAREDDEVICEUSERLOGGEDINERROR_IOSUPDATESINSTALLSTATUS
        case "notSupportedOperation":
            result = NOTSUPPORTEDOPERATION_IOSUPDATESINSTALLSTATUS
        case "installFailed":
            result = INSTALLFAILED_IOSUPDATESINSTALLSTATUS
        case "installPhoneCallInProgress":
            result = INSTALLPHONECALLINPROGRESS_IOSUPDATESINSTALLSTATUS
        case "installInsufficientPower":
            result = INSTALLINSUFFICIENTPOWER_IOSUPDATESINSTALLSTATUS
        case "installInsufficientSpace":
            result = INSTALLINSUFFICIENTSPACE_IOSUPDATESINSTALLSTATUS
        case "installing":
            result = INSTALLING_IOSUPDATESINSTALLSTATUS
        case "downloadInsufficientNetwork":
            result = DOWNLOADINSUFFICIENTNETWORK_IOSUPDATESINSTALLSTATUS
        case "downloadInsufficientPower":
            result = DOWNLOADINSUFFICIENTPOWER_IOSUPDATESINSTALLSTATUS
        case "downloadInsufficientSpace":
            result = DOWNLOADINSUFFICIENTSPACE_IOSUPDATESINSTALLSTATUS
        case "downloadRequiresComputer":
            result = DOWNLOADREQUIRESCOMPUTER_IOSUPDATESINSTALLSTATUS
        case "downloadFailed":
            result = DOWNLOADFAILED_IOSUPDATESINSTALLSTATUS
        case "downloading":
            result = DOWNLOADING_IOSUPDATESINSTALLSTATUS
        case "timeout":
            result = TIMEOUT_IOSUPDATESINSTALLSTATUS
        case "mdmClientCrashed":
            result = MDMCLIENTCRASHED_IOSUPDATESINSTALLSTATUS
        case "success":
            result = SUCCESS_IOSUPDATESINSTALLSTATUS
        case "available":
            result = AVAILABLE_IOSUPDATESINSTALLSTATUS
        case "idle":
            result = IDLE_IOSUPDATESINSTALLSTATUS
        case "unknown":
            result = UNKNOWN_IOSUPDATESINSTALLSTATUS
        default:
            return 0, errors.New("Unknown IosUpdatesInstallStatus value: " + v)
    }
    return &result, nil
}
func SerializeIosUpdatesInstallStatus(values []IosUpdatesInstallStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
