package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AdvancedBitLockerState int

const (
    // Advanced BitLocker State Success
    SUCCESS_ADVANCEDBITLOCKERSTATE AdvancedBitLockerState = iota
    // User never gave consent for Encryption
    NOUSERCONSENT_ADVANCEDBITLOCKERSTATE
    // Un-protected OS Volume was detected
    OSVOLUMEUNPROTECTED_ADVANCEDBITLOCKERSTATE
    // TPM not used for protection of OS volume, but is required by policy
    OSVOLUMETPMREQUIRED_ADVANCEDBITLOCKERSTATE
    // TPM only protection not used for OS volume, but is required by policy
    OSVOLUMETPMONLYREQUIRED_ADVANCEDBITLOCKERSTATE
    // TPM+PIN protection not used for OS volume, but is required by policy
    OSVOLUMETPMPINREQUIRED_ADVANCEDBITLOCKERSTATE
    // TPM+Startup Key protection not used for OS volume, but is required by policy
    OSVOLUMETPMSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE
    // TPM+PIN+Startup Key not used for OS volume, but is required by policy
    OSVOLUMETPMPINSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE
    // Encryption method of OS Volume is different than that set by policy
    OSVOLUMEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE
    // Recovery key backup failed
    RECOVERYKEYBACKUPFAILED_ADVANCEDBITLOCKERSTATE
    // Fixed Drive not encrypted
    FIXEDDRIVENOTENCRYPTED_ADVANCEDBITLOCKERSTATE
    // Encryption method of Fixed Drive is different than that set by policy
    FIXEDDRIVEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE
    // Logged on user is non-admin. This requires “AllowStandardUserEncryption” policy set to 1
    LOGGEDONUSERNONADMIN_ADVANCEDBITLOCKERSTATE
    // WinRE is not configured
    WINDOWSRECOVERYENVIRONMENTNOTCONFIGURED_ADVANCEDBITLOCKERSTATE
    // TPM is not available for BitLocker. This means TPM is not present, or TPM unavailable registry override is set or host OS is on portable/rome-able drive
    TPMNOTAVAILABLE_ADVANCEDBITLOCKERSTATE
    // TPM is not ready for BitLocker
    TPMNOTREADY_ADVANCEDBITLOCKERSTATE
    // Network not available. This is required for recovery key backup. This is reported for Drive Encryption capable devices
    NETWORKERROR_ADVANCEDBITLOCKERSTATE
)

func (i AdvancedBitLockerState) String() string {
    return []string{"success", "noUserConsent", "osVolumeUnprotected", "osVolumeTpmRequired", "osVolumeTpmOnlyRequired", "osVolumeTpmPinRequired", "osVolumeTpmStartupKeyRequired", "osVolumeTpmPinStartupKeyRequired", "osVolumeEncryptionMethodMismatch", "recoveryKeyBackupFailed", "fixedDriveNotEncrypted", "fixedDriveEncryptionMethodMismatch", "loggedOnUserNonAdmin", "windowsRecoveryEnvironmentNotConfigured", "tpmNotAvailable", "tpmNotReady", "networkError"}[i]
}
func ParseAdvancedBitLockerState(v string) (interface{}, error) {
    result := SUCCESS_ADVANCEDBITLOCKERSTATE
    switch v {
        case "success":
            result = SUCCESS_ADVANCEDBITLOCKERSTATE
        case "noUserConsent":
            result = NOUSERCONSENT_ADVANCEDBITLOCKERSTATE
        case "osVolumeUnprotected":
            result = OSVOLUMEUNPROTECTED_ADVANCEDBITLOCKERSTATE
        case "osVolumeTpmRequired":
            result = OSVOLUMETPMREQUIRED_ADVANCEDBITLOCKERSTATE
        case "osVolumeTpmOnlyRequired":
            result = OSVOLUMETPMONLYREQUIRED_ADVANCEDBITLOCKERSTATE
        case "osVolumeTpmPinRequired":
            result = OSVOLUMETPMPINREQUIRED_ADVANCEDBITLOCKERSTATE
        case "osVolumeTpmStartupKeyRequired":
            result = OSVOLUMETPMSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE
        case "osVolumeTpmPinStartupKeyRequired":
            result = OSVOLUMETPMPINSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE
        case "osVolumeEncryptionMethodMismatch":
            result = OSVOLUMEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE
        case "recoveryKeyBackupFailed":
            result = RECOVERYKEYBACKUPFAILED_ADVANCEDBITLOCKERSTATE
        case "fixedDriveNotEncrypted":
            result = FIXEDDRIVENOTENCRYPTED_ADVANCEDBITLOCKERSTATE
        case "fixedDriveEncryptionMethodMismatch":
            result = FIXEDDRIVEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE
        case "loggedOnUserNonAdmin":
            result = LOGGEDONUSERNONADMIN_ADVANCEDBITLOCKERSTATE
        case "windowsRecoveryEnvironmentNotConfigured":
            result = WINDOWSRECOVERYENVIRONMENTNOTCONFIGURED_ADVANCEDBITLOCKERSTATE
        case "tpmNotAvailable":
            result = TPMNOTAVAILABLE_ADVANCEDBITLOCKERSTATE
        case "tpmNotReady":
            result = TPMNOTREADY_ADVANCEDBITLOCKERSTATE
        case "networkError":
            result = NETWORKERROR_ADVANCEDBITLOCKERSTATE
        default:
            return 0, errors.New("Unknown AdvancedBitLockerState value: " + v)
    }
    return &result, nil
}
func SerializeAdvancedBitLockerState(values []AdvancedBitLockerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
