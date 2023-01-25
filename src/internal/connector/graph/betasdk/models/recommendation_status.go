package models
import (
    "errors"
)
// Provides operations to call the add method.
type RecommendationStatus int

const (
    ACTIVE_RECOMMENDATIONSTATUS RecommendationStatus = iota
    COMPLETEDBYSYSTEM_RECOMMENDATIONSTATUS
    COMPLETEDBYUSER_RECOMMENDATIONSTATUS
    DISMISSED_RECOMMENDATIONSTATUS
    POSTPONED_RECOMMENDATIONSTATUS
    UNKNOWNFUTUREVALUE_RECOMMENDATIONSTATUS
)

func (i RecommendationStatus) String() string {
    return []string{"active", "completedBySystem", "completedByUser", "dismissed", "postponed", "unknownFutureValue"}[i]
}
func ParseRecommendationStatus(v string) (interface{}, error) {
    result := ACTIVE_RECOMMENDATIONSTATUS
    switch v {
        case "active":
            result = ACTIVE_RECOMMENDATIONSTATUS
        case "completedBySystem":
            result = COMPLETEDBYSYSTEM_RECOMMENDATIONSTATUS
        case "completedByUser":
            result = COMPLETEDBYUSER_RECOMMENDATIONSTATUS
        case "dismissed":
            result = DISMISSED_RECOMMENDATIONSTATUS
        case "postponed":
            result = POSTPONED_RECOMMENDATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RECOMMENDATIONSTATUS
        default:
            return 0, errors.New("Unknown RecommendationStatus value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationStatus(values []RecommendationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
