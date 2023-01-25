package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AuthenticationAppEvaluation int

const (
    SUCCESS_AUTHENTICATIONAPPEVALUATION AuthenticationAppEvaluation = iota
    FAILURE_AUTHENTICATIONAPPEVALUATION
    UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPEVALUATION
)

func (i AuthenticationAppEvaluation) String() string {
    return []string{"success", "failure", "unknownFutureValue"}[i]
}
func ParseAuthenticationAppEvaluation(v string) (interface{}, error) {
    result := SUCCESS_AUTHENTICATIONAPPEVALUATION
    switch v {
        case "success":
            result = SUCCESS_AUTHENTICATIONAPPEVALUATION
        case "failure":
            result = FAILURE_AUTHENTICATIONAPPEVALUATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPEVALUATION
        default:
            return 0, errors.New("Unknown AuthenticationAppEvaluation value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationAppEvaluation(values []AuthenticationAppEvaluation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
