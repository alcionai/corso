package models
import (
    "errors"
)
// Provides operations to call the add method.
type WiredNetworkAuthenticationType int

const (
    // None
    NONE_WIREDNETWORKAUTHENTICATIONTYPE WiredNetworkAuthenticationType = iota
    // User Authentication
    USER_WIREDNETWORKAUTHENTICATIONTYPE
    // Machine Authentication
    MACHINE_WIREDNETWORKAUTHENTICATIONTYPE
    // Machine or User Authentication
    MACHINEORUSER_WIREDNETWORKAUTHENTICATIONTYPE
    // Guest Authentication
    GUEST_WIREDNETWORKAUTHENTICATIONTYPE
    // Sentinel member for cases where the client cannot handle the new enum values.
    UNKNOWNFUTUREVALUE_WIREDNETWORKAUTHENTICATIONTYPE
)

func (i WiredNetworkAuthenticationType) String() string {
    return []string{"none", "user", "machine", "machineOrUser", "guest", "unknownFutureValue"}[i]
}
func ParseWiredNetworkAuthenticationType(v string) (interface{}, error) {
    result := NONE_WIREDNETWORKAUTHENTICATIONTYPE
    switch v {
        case "none":
            result = NONE_WIREDNETWORKAUTHENTICATIONTYPE
        case "user":
            result = USER_WIREDNETWORKAUTHENTICATIONTYPE
        case "machine":
            result = MACHINE_WIREDNETWORKAUTHENTICATIONTYPE
        case "machineOrUser":
            result = MACHINEORUSER_WIREDNETWORKAUTHENTICATIONTYPE
        case "guest":
            result = GUEST_WIREDNETWORKAUTHENTICATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WIREDNETWORKAUTHENTICATIONTYPE
        default:
            return 0, errors.New("Unknown WiredNetworkAuthenticationType value: " + v)
    }
    return &result, nil
}
func SerializeWiredNetworkAuthenticationType(values []WiredNetworkAuthenticationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
