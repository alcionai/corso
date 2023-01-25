package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SingleSignOnMode int

const (
    NONE_SINGLESIGNONMODE SingleSignOnMode = iota
    ONPREMISESKERBEROS_SINGLESIGNONMODE
    SAML_SINGLESIGNONMODE
    PINGHEADERBASED_SINGLESIGNONMODE
    AADHEADERBASED_SINGLESIGNONMODE
    OAUTHTOKEN_SINGLESIGNONMODE
    UNKNOWNFUTUREVALUE_SINGLESIGNONMODE
)

func (i SingleSignOnMode) String() string {
    return []string{"none", "onPremisesKerberos", "saml", "pingHeaderBased", "aadHeaderBased", "oAuthToken", "unknownFutureValue"}[i]
}
func ParseSingleSignOnMode(v string) (interface{}, error) {
    result := NONE_SINGLESIGNONMODE
    switch v {
        case "none":
            result = NONE_SINGLESIGNONMODE
        case "onPremisesKerberos":
            result = ONPREMISESKERBEROS_SINGLESIGNONMODE
        case "saml":
            result = SAML_SINGLESIGNONMODE
        case "pingHeaderBased":
            result = PINGHEADERBASED_SINGLESIGNONMODE
        case "aadHeaderBased":
            result = AADHEADERBASED_SINGLESIGNONMODE
        case "oAuthToken":
            result = OAUTHTOKEN_SINGLESIGNONMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SINGLESIGNONMODE
        default:
            return 0, errors.New("Unknown SingleSignOnMode value: " + v)
    }
    return &result, nil
}
func SerializeSingleSignOnMode(values []SingleSignOnMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
