package models
import (
    "errors"
)
// Provides operations to call the add method.
type WiFiAuthenticationMethod int

const (
    // Use an identity certificate for authentication.
    CERTIFICATE_WIFIAUTHENTICATIONMETHOD WiFiAuthenticationMethod = iota
    // Use username and password for authentication.
    USERNAMEANDPASSWORD_WIFIAUTHENTICATIONMETHOD
    // Use Derived Credential for authentication.
    DERIVEDCREDENTIAL_WIFIAUTHENTICATIONMETHOD
)

func (i WiFiAuthenticationMethod) String() string {
    return []string{"certificate", "usernameAndPassword", "derivedCredential"}[i]
}
func ParseWiFiAuthenticationMethod(v string) (interface{}, error) {
    result := CERTIFICATE_WIFIAUTHENTICATIONMETHOD
    switch v {
        case "certificate":
            result = CERTIFICATE_WIFIAUTHENTICATIONMETHOD
        case "usernameAndPassword":
            result = USERNAMEANDPASSWORD_WIFIAUTHENTICATIONMETHOD
        case "derivedCredential":
            result = DERIVEDCREDENTIAL_WIFIAUTHENTICATIONMETHOD
        default:
            return 0, errors.New("Unknown WiFiAuthenticationMethod value: " + v)
    }
    return &result, nil
}
func SerializeWiFiAuthenticationMethod(values []WiFiAuthenticationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
