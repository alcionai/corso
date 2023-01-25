package models
import (
    "errors"
)
// Provides operations to call the add method.
type AuthenticationMethodFeature int

const (
    SSPRREGISTERED_AUTHENTICATIONMETHODFEATURE AuthenticationMethodFeature = iota
    SSPRENABLED_AUTHENTICATIONMETHODFEATURE
    SSPRCAPABLE_AUTHENTICATIONMETHODFEATURE
    PASSWORDLESSCAPABLE_AUTHENTICATIONMETHODFEATURE
    MFACAPABLE_AUTHENTICATIONMETHODFEATURE
)

func (i AuthenticationMethodFeature) String() string {
    return []string{"ssprRegistered", "ssprEnabled", "ssprCapable", "passwordlessCapable", "mfaCapable"}[i]
}
func ParseAuthenticationMethodFeature(v string) (interface{}, error) {
    result := SSPRREGISTERED_AUTHENTICATIONMETHODFEATURE
    switch v {
        case "ssprRegistered":
            result = SSPRREGISTERED_AUTHENTICATIONMETHODFEATURE
        case "ssprEnabled":
            result = SSPRENABLED_AUTHENTICATIONMETHODFEATURE
        case "ssprCapable":
            result = SSPRCAPABLE_AUTHENTICATIONMETHODFEATURE
        case "passwordlessCapable":
            result = PASSWORDLESSCAPABLE_AUTHENTICATIONMETHODFEATURE
        case "mfaCapable":
            result = MFACAPABLE_AUTHENTICATIONMETHODFEATURE
        default:
            return 0, errors.New("Unknown AuthenticationMethodFeature value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationMethodFeature(values []AuthenticationMethodFeature) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
