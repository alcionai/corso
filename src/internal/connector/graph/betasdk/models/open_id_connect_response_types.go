package models
import (
    "errors"
)
// Provides operations to call the add method.
type OpenIdConnectResponseTypes int

const (
    CODE_OPENIDCONNECTRESPONSETYPES OpenIdConnectResponseTypes = iota
    ID_TOKEN_OPENIDCONNECTRESPONSETYPES
    TOKEN_OPENIDCONNECTRESPONSETYPES
)

func (i OpenIdConnectResponseTypes) String() string {
    return []string{"code", "id_token", "token"}[i]
}
func ParseOpenIdConnectResponseTypes(v string) (interface{}, error) {
    result := CODE_OPENIDCONNECTRESPONSETYPES
    switch v {
        case "code":
            result = CODE_OPENIDCONNECTRESPONSETYPES
        case "id_token":
            result = ID_TOKEN_OPENIDCONNECTRESPONSETYPES
        case "token":
            result = TOKEN_OPENIDCONNECTRESPONSETYPES
        default:
            return 0, errors.New("Unknown OpenIdConnectResponseTypes value: " + v)
    }
    return &result, nil
}
func SerializeOpenIdConnectResponseTypes(values []OpenIdConnectResponseTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
