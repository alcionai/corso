package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type UserExperienceAnalyticsAnomalyState int

const (
    // Indicates the state of anomaly is new.
    NEW_USEREXPERIENCEANALYTICSANOMALYSTATE UserExperienceAnalyticsAnomalyState = iota
    // Indicates the state of anomaly is active.
    ACTIVE_USEREXPERIENCEANALYTICSANOMALYSTATE
    // Indicates the state of anomaly is disabled.
    DISABLED_USEREXPERIENCEANALYTICSANOMALYSTATE
    // Indicates the state of anomaly is removed.
    REMOVED_USEREXPERIENCEANALYTICSANOMALYSTATE
    // Indicates the state of anomaly is undefined.
    OTHER_USEREXPERIENCEANALYTICSANOMALYSTATE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYSTATE
)

func (i UserExperienceAnalyticsAnomalyState) String() string {
    return []string{"new", "active", "disabled", "removed", "other", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsAnomalyState(v string) (interface{}, error) {
    result := NEW_USEREXPERIENCEANALYTICSANOMALYSTATE
    switch v {
        case "new":
            result = NEW_USEREXPERIENCEANALYTICSANOMALYSTATE
        case "active":
            result = ACTIVE_USEREXPERIENCEANALYTICSANOMALYSTATE
        case "disabled":
            result = DISABLED_USEREXPERIENCEANALYTICSANOMALYSTATE
        case "removed":
            result = REMOVED_USEREXPERIENCEANALYTICSANOMALYSTATE
        case "other":
            result = OTHER_USEREXPERIENCEANALYTICSANOMALYSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYSTATE
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsAnomalyState value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsAnomalyState(values []UserExperienceAnalyticsAnomalyState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
