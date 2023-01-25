package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type KeyStorageProviderOption int

const (
    // Import to Trusted Platform Module (TPM) KSP if present, otherwise import to Software KSP.
    USETPMKSPOTHERWISEUSESOFTWAREKSP_KEYSTORAGEPROVIDEROPTION KeyStorageProviderOption = iota
    // Import to Trusted Platform Module (TPM) KSP if present, otherwise fail.
    USETPMKSPOTHERWISEFAIL_KEYSTORAGEPROVIDEROPTION
    // Import to Passport for work KSP if available, otherwise fail.
    USEPASSPORTFORWORKKSPOTHERWISEFAIL_KEYSTORAGEPROVIDEROPTION
    // Import to Software KSP.
    USESOFTWAREKSP_KEYSTORAGEPROVIDEROPTION
)

func (i KeyStorageProviderOption) String() string {
    return []string{"useTpmKspOtherwiseUseSoftwareKsp", "useTpmKspOtherwiseFail", "usePassportForWorkKspOtherwiseFail", "useSoftwareKsp"}[i]
}
func ParseKeyStorageProviderOption(v string) (interface{}, error) {
    result := USETPMKSPOTHERWISEUSESOFTWAREKSP_KEYSTORAGEPROVIDEROPTION
    switch v {
        case "useTpmKspOtherwiseUseSoftwareKsp":
            result = USETPMKSPOTHERWISEUSESOFTWAREKSP_KEYSTORAGEPROVIDEROPTION
        case "useTpmKspOtherwiseFail":
            result = USETPMKSPOTHERWISEFAIL_KEYSTORAGEPROVIDEROPTION
        case "usePassportForWorkKspOtherwiseFail":
            result = USEPASSPORTFORWORKKSPOTHERWISEFAIL_KEYSTORAGEPROVIDEROPTION
        case "useSoftwareKsp":
            result = USESOFTWAREKSP_KEYSTORAGEPROVIDEROPTION
        default:
            return 0, errors.New("Unknown KeyStorageProviderOption value: " + v)
    }
    return &result, nil
}
func SerializeKeyStorageProviderOption(values []KeyStorageProviderOption) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
