package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RecommendationCategory int

const (
    IDENTITYBESTPRACTICE_RECOMMENDATIONCATEGORY RecommendationCategory = iota
    IDENTITYSECURESCORE_RECOMMENDATIONCATEGORY
    UNKNOWNFUTUREVALUE_RECOMMENDATIONCATEGORY
)

func (i RecommendationCategory) String() string {
    return []string{"identityBestPractice", "identitySecureScore", "unknownFutureValue"}[i]
}
func ParseRecommendationCategory(v string) (interface{}, error) {
    result := IDENTITYBESTPRACTICE_RECOMMENDATIONCATEGORY
    switch v {
        case "identityBestPractice":
            result = IDENTITYBESTPRACTICE_RECOMMENDATIONCATEGORY
        case "identitySecureScore":
            result = IDENTITYSECURESCORE_RECOMMENDATIONCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RECOMMENDATIONCATEGORY
        default:
            return 0, errors.New("Unknown RecommendationCategory value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationCategory(values []RecommendationCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
