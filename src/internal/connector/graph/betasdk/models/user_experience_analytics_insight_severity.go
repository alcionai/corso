package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UserExperienceAnalyticsInsightSeverity int

const (
    NONE_USEREXPERIENCEANALYTICSINSIGHTSEVERITY UserExperienceAnalyticsInsightSeverity = iota
    INFORMATIONAL_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
    WARNING_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
    ERROR_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
)

func (i UserExperienceAnalyticsInsightSeverity) String() string {
    return []string{"none", "informational", "warning", "error"}[i]
}
func ParseUserExperienceAnalyticsInsightSeverity(v string) (interface{}, error) {
    result := NONE_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
    switch v {
        case "none":
            result = NONE_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        case "informational":
            result = INFORMATIONAL_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        case "warning":
            result = WARNING_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        case "error":
            result = ERROR_USEREXPERIENCEANALYTICSINSIGHTSEVERITY
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsInsightSeverity value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsInsightSeverity(values []UserExperienceAnalyticsInsightSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
