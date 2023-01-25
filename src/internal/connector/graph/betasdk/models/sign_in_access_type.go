package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SignInAccessType int

const (
    NONE_SIGNINACCESSTYPE SignInAccessType = iota
    B2BCOLLABORATION_SIGNINACCESSTYPE
    B2BDIRECTCONNECT_SIGNINACCESSTYPE
    MICROSOFTSUPPORT_SIGNINACCESSTYPE
    SERVICEPROVIDER_SIGNINACCESSTYPE
    UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE
)

func (i SignInAccessType) String() string {
    return []string{"none", "b2bCollaboration", "b2bDirectConnect", "microsoftSupport", "serviceProvider", "unknownFutureValue"}[i]
}
func ParseSignInAccessType(v string) (interface{}, error) {
    result := NONE_SIGNINACCESSTYPE
    switch v {
        case "none":
            result = NONE_SIGNINACCESSTYPE
        case "b2bCollaboration":
            result = B2BCOLLABORATION_SIGNINACCESSTYPE
        case "b2bDirectConnect":
            result = B2BDIRECTCONNECT_SIGNINACCESSTYPE
        case "microsoftSupport":
            result = MICROSOFTSUPPORT_SIGNINACCESSTYPE
        case "serviceProvider":
            result = SERVICEPROVIDER_SIGNINACCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIGNINACCESSTYPE
        default:
            return 0, errors.New("Unknown SignInAccessType value: " + v)
    }
    return &result, nil
}
func SerializeSignInAccessType(values []SignInAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
