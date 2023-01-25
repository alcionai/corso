package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WiredNetworkAuthenticationMethod int

const (
    // Use an identity certificate for authentication.
    CERTIFICATE_WIREDNETWORKAUTHENTICATIONMETHOD WiredNetworkAuthenticationMethod = iota
    // Use username and password for authentication.
    USERNAMEANDPASSWORD_WIREDNETWORKAUTHENTICATIONMETHOD
    // Use Derived Credential for authentication.
    DERIVEDCREDENTIAL_WIREDNETWORKAUTHENTICATIONMETHOD
    // Sentinel member for cases where the client cannot handle the new enum values.
    UNKNOWNFUTUREVALUE_WIREDNETWORKAUTHENTICATIONMETHOD
)

func (i WiredNetworkAuthenticationMethod) String() string {
    return []string{"certificate", "usernameAndPassword", "derivedCredential", "unknownFutureValue"}[i]
}
func ParseWiredNetworkAuthenticationMethod(v string) (interface{}, error) {
    result := CERTIFICATE_WIREDNETWORKAUTHENTICATIONMETHOD
    switch v {
        case "certificate":
            result = CERTIFICATE_WIREDNETWORKAUTHENTICATIONMETHOD
        case "usernameAndPassword":
            result = USERNAMEANDPASSWORD_WIREDNETWORKAUTHENTICATIONMETHOD
        case "derivedCredential":
            result = DERIVEDCREDENTIAL_WIREDNETWORKAUTHENTICATIONMETHOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WIREDNETWORKAUTHENTICATIONMETHOD
        default:
            return 0, errors.New("Unknown WiredNetworkAuthenticationMethod value: " + v)
    }
    return &result, nil
}
func SerializeWiredNetworkAuthenticationMethod(values []WiredNetworkAuthenticationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
