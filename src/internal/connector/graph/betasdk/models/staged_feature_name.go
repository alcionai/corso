package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type StagedFeatureName int

const (
    PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME StagedFeatureName = iota
    SEAMLESSSSO_STAGEDFEATURENAME
    PASSWORDHASHSYNC_STAGEDFEATURENAME
    EMAILASALTERNATEID_STAGEDFEATURENAME
    UNKNOWNFUTUREVALUE_STAGEDFEATURENAME
    CERTIFICATEBASEDAUTHENTICATION_STAGEDFEATURENAME
)

func (i StagedFeatureName) String() string {
    return []string{"passthroughAuthentication", "seamlessSso", "passwordHashSync", "emailAsAlternateId", "unknownFutureValue", "certificateBasedAuthentication"}[i]
}
func ParseStagedFeatureName(v string) (interface{}, error) {
    result := PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME
    switch v {
        case "passthroughAuthentication":
            result = PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME
        case "seamlessSso":
            result = SEAMLESSSSO_STAGEDFEATURENAME
        case "passwordHashSync":
            result = PASSWORDHASHSYNC_STAGEDFEATURENAME
        case "emailAsAlternateId":
            result = EMAILASALTERNATEID_STAGEDFEATURENAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_STAGEDFEATURENAME
        case "certificateBasedAuthentication":
            result = CERTIFICATEBASEDAUTHENTICATION_STAGEDFEATURENAME
        default:
            return 0, errors.New("Unknown StagedFeatureName value: " + v)
    }
    return &result, nil
}
func SerializeStagedFeatureName(values []StagedFeatureName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
