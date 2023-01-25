package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type OpenIdConnectResponseMode int

const (
    FORM_POST_OPENIDCONNECTRESPONSEMODE OpenIdConnectResponseMode = iota
    QUERY_OPENIDCONNECTRESPONSEMODE
    UNKNOWNFUTUREVALUE_OPENIDCONNECTRESPONSEMODE
)

func (i OpenIdConnectResponseMode) String() string {
    return []string{"form_post", "query", "unknownFutureValue"}[i]
}
func ParseOpenIdConnectResponseMode(v string) (interface{}, error) {
    result := FORM_POST_OPENIDCONNECTRESPONSEMODE
    switch v {
        case "form_post":
            result = FORM_POST_OPENIDCONNECTRESPONSEMODE
        case "query":
            result = QUERY_OPENIDCONNECTRESPONSEMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OPENIDCONNECTRESPONSEMODE
        default:
            return 0, errors.New("Unknown OpenIdConnectResponseMode value: " + v)
    }
    return &result, nil
}
func SerializeOpenIdConnectResponseMode(values []OpenIdConnectResponseMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
