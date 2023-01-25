package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UserExperienceAnalyticsAnomalySeverity int

const (
    // Indicates the anomaly is of high severity.
    HIGH_USEREXPERIENCEANALYTICSANOMALYSEVERITY UserExperienceAnalyticsAnomalySeverity = iota
    // Indicates the anomaly is of medium severity.
    MEDIUM_USEREXPERIENCEANALYTICSANOMALYSEVERITY
    // Indicates the anomaly is of low severity.
    LOW_USEREXPERIENCEANALYTICSANOMALYSEVERITY
    // Indicates the anomaly is of informational severity.
    INFORMATIONAL_USEREXPERIENCEANALYTICSANOMALYSEVERITY
    // Indicates the severity of anomaly is undefined.
    OTHER_USEREXPERIENCEANALYTICSANOMALYSEVERITY
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYSEVERITY
)

func (i UserExperienceAnalyticsAnomalySeverity) String() string {
    return []string{"high", "medium", "low", "informational", "other", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsAnomalySeverity(v string) (interface{}, error) {
    result := HIGH_USEREXPERIENCEANALYTICSANOMALYSEVERITY
    switch v {
        case "high":
            result = HIGH_USEREXPERIENCEANALYTICSANOMALYSEVERITY
        case "medium":
            result = MEDIUM_USEREXPERIENCEANALYTICSANOMALYSEVERITY
        case "low":
            result = LOW_USEREXPERIENCEANALYTICSANOMALYSEVERITY
        case "informational":
            result = INFORMATIONAL_USEREXPERIENCEANALYTICSANOMALYSEVERITY
        case "other":
            result = OTHER_USEREXPERIENCEANALYTICSANOMALYSEVERITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYSEVERITY
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsAnomalySeverity value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsAnomalySeverity(values []UserExperienceAnalyticsAnomalySeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
