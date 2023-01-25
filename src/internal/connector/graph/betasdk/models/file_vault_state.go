package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type FileVaultState int

const (
    // FileVault State Success
    SUCCESS_FILEVAULTSTATE FileVaultState = iota
    // FileVault has been enabled by user and is not being managed by policy
    DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE
    // FileVault policy is successfully installed but user has not started encryption
    USERDEFERREDENCRYPTION_FILEVAULTSTATE
    // FileVault recovery key escrow is not enabled
    ESCROWNOTENABLED_FILEVAULTSTATE
)

func (i FileVaultState) String() string {
    return []string{"success", "driveEncryptedByUser", "userDeferredEncryption", "escrowNotEnabled"}[i]
}
func ParseFileVaultState(v string) (interface{}, error) {
    result := SUCCESS_FILEVAULTSTATE
    switch v {
        case "success":
            result = SUCCESS_FILEVAULTSTATE
        case "driveEncryptedByUser":
            result = DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE
        case "userDeferredEncryption":
            result = USERDEFERREDENCRYPTION_FILEVAULTSTATE
        case "escrowNotEnabled":
            result = ESCROWNOTENABLED_FILEVAULTSTATE
        default:
            return 0, errors.New("Unknown FileVaultState value: " + v)
    }
    return &result, nil
}
func SerializeFileVaultState(values []FileVaultState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
