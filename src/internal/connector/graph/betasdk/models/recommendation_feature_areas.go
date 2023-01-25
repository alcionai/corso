package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RecommendationFeatureAreas int

const (
    USERS_RECOMMENDATIONFEATUREAREAS RecommendationFeatureAreas = iota
    GROUPS_RECOMMENDATIONFEATUREAREAS
    DEVICES_RECOMMENDATIONFEATUREAREAS
    APPLICATIONS_RECOMMENDATIONFEATUREAREAS
    ACCESSREVIEWS_RECOMMENDATIONFEATUREAREAS
    CONDITIONALACCESS_RECOMMENDATIONFEATUREAREAS
    GOVERNANCE_RECOMMENDATIONFEATUREAREAS
    UNKNOWNFUTUREVALUE_RECOMMENDATIONFEATUREAREAS
)

func (i RecommendationFeatureAreas) String() string {
    return []string{"users", "groups", "devices", "applications", "accessReviews", "conditionalAccess", "governance", "unknownFutureValue"}[i]
}
func ParseRecommendationFeatureAreas(v string) (interface{}, error) {
    result := USERS_RECOMMENDATIONFEATUREAREAS
    switch v {
        case "users":
            result = USERS_RECOMMENDATIONFEATUREAREAS
        case "groups":
            result = GROUPS_RECOMMENDATIONFEATUREAREAS
        case "devices":
            result = DEVICES_RECOMMENDATIONFEATUREAREAS
        case "applications":
            result = APPLICATIONS_RECOMMENDATIONFEATUREAREAS
        case "accessReviews":
            result = ACCESSREVIEWS_RECOMMENDATIONFEATUREAREAS
        case "conditionalAccess":
            result = CONDITIONALACCESS_RECOMMENDATIONFEATUREAREAS
        case "governance":
            result = GOVERNANCE_RECOMMENDATIONFEATUREAREAS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RECOMMENDATIONFEATUREAREAS
        default:
            return 0, errors.New("Unknown RecommendationFeatureAreas value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationFeatureAreas(values []RecommendationFeatureAreas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
