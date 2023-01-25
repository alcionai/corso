package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RecommendationPriority int

const (
    LOW_RECOMMENDATIONPRIORITY RecommendationPriority = iota
    MEDIUM_RECOMMENDATIONPRIORITY
    HIGH_RECOMMENDATIONPRIORITY
)

func (i RecommendationPriority) String() string {
    return []string{"low", "medium", "high"}[i]
}
func ParseRecommendationPriority(v string) (interface{}, error) {
    result := LOW_RECOMMENDATIONPRIORITY
    switch v {
        case "low":
            result = LOW_RECOMMENDATIONPRIORITY
        case "medium":
            result = MEDIUM_RECOMMENDATIONPRIORITY
        case "high":
            result = HIGH_RECOMMENDATIONPRIORITY
        default:
            return 0, errors.New("Unknown RecommendationPriority value: " + v)
    }
    return &result, nil
}
func SerializeRecommendationPriority(values []RecommendationPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
