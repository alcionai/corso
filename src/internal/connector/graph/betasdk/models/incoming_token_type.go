package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type IncomingTokenType int

const (
    NONE_INCOMINGTOKENTYPE IncomingTokenType = iota
    PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE
    SAML11_INCOMINGTOKENTYPE
    SAML20_INCOMINGTOKENTYPE
    UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE
    REMOTEDESKTOPTOKEN_INCOMINGTOKENTYPE
)

func (i IncomingTokenType) String() string {
    return []string{"none", "primaryRefreshToken", "saml11", "saml20", "unknownFutureValue", "remoteDesktopToken"}[i]
}
func ParseIncomingTokenType(v string) (interface{}, error) {
    result := NONE_INCOMINGTOKENTYPE
    switch v {
        case "none":
            result = NONE_INCOMINGTOKENTYPE
        case "primaryRefreshToken":
            result = PRIMARYREFRESHTOKEN_INCOMINGTOKENTYPE
        case "saml11":
            result = SAML11_INCOMINGTOKENTYPE
        case "saml20":
            result = SAML20_INCOMINGTOKENTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INCOMINGTOKENTYPE
        case "remoteDesktopToken":
            result = REMOTEDESKTOPTOKEN_INCOMINGTOKENTYPE
        default:
            return 0, errors.New("Unknown IncomingTokenType value: " + v)
    }
    return &result, nil
}
func SerializeIncomingTokenType(values []IncomingTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
