package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AuthenticationEventType int

const (
    TOKENISSUANCESTART_AUTHENTICATIONEVENTTYPE AuthenticationEventType = iota
    PAGERENDERSTART_AUTHENTICATIONEVENTTYPE
    UNKNOWNFUTUREVALUE_AUTHENTICATIONEVENTTYPE
)

func (i AuthenticationEventType) String() string {
    return []string{"tokenIssuanceStart", "pageRenderStart", "unknownFutureValue"}[i]
}
func ParseAuthenticationEventType(v string) (interface{}, error) {
    result := TOKENISSUANCESTART_AUTHENTICATIONEVENTTYPE
    switch v {
        case "tokenIssuanceStart":
            result = TOKENISSUANCESTART_AUTHENTICATIONEVENTTYPE
        case "pageRenderStart":
            result = PAGERENDERSTART_AUTHENTICATIONEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONEVENTTYPE
        default:
            return 0, errors.New("Unknown AuthenticationEventType value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationEventType(values []AuthenticationEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
