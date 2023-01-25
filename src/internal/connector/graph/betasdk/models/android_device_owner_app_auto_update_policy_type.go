package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AndroidDeviceOwnerAppAutoUpdatePolicyType int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE AndroidDeviceOwnerAppAutoUpdatePolicyType = iota
    // The user can control auto-updates.
    USERCHOICE_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
    // Apps are never auto-updated.
    NEVER_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
    // Apps are auto-updated over Wi-Fi only.
    WIFIONLY_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
    // Apps are auto-updated at any time. Data charges may apply.
    ALWAYS_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
)

func (i AndroidDeviceOwnerAppAutoUpdatePolicyType) String() string {
    return []string{"notConfigured", "userChoice", "never", "wiFiOnly", "always"}[i]
}
func ParseAndroidDeviceOwnerAppAutoUpdatePolicyType(v string) (interface{}, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
        case "userChoice":
            result = USERCHOICE_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
        case "never":
            result = NEVER_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
        case "wiFiOnly":
            result = WIFIONLY_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
        case "always":
            result = ALWAYS_ANDROIDDEVICEOWNERAPPAUTOUPDATEPOLICYTYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerAppAutoUpdatePolicyType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerAppAutoUpdatePolicyType(values []AndroidDeviceOwnerAppAutoUpdatePolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
