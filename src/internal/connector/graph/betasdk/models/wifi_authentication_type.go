package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WifiAuthenticationType int

const (
    // None
    NONE_WIFIAUTHENTICATIONTYPE WifiAuthenticationType = iota
    // User Authentication
    USER_WIFIAUTHENTICATIONTYPE
    // Machine Authentication
    MACHINE_WIFIAUTHENTICATIONTYPE
    // Machine or User Authentication
    MACHINEORUSER_WIFIAUTHENTICATIONTYPE
    // Guest Authentication
    GUEST_WIFIAUTHENTICATIONTYPE
)

func (i WifiAuthenticationType) String() string {
    return []string{"none", "user", "machine", "machineOrUser", "guest"}[i]
}
func ParseWifiAuthenticationType(v string) (interface{}, error) {
    result := NONE_WIFIAUTHENTICATIONTYPE
    switch v {
        case "none":
            result = NONE_WIFIAUTHENTICATIONTYPE
        case "user":
            result = USER_WIFIAUTHENTICATIONTYPE
        case "machine":
            result = MACHINE_WIFIAUTHENTICATIONTYPE
        case "machineOrUser":
            result = MACHINEORUSER_WIFIAUTHENTICATIONTYPE
        case "guest":
            result = GUEST_WIFIAUTHENTICATIONTYPE
        default:
            return 0, errors.New("Unknown WifiAuthenticationType value: " + v)
    }
    return &result, nil
}
func SerializeWifiAuthenticationType(values []WifiAuthenticationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
