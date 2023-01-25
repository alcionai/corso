package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AnalyticsActivityType int

const (
    EMAIL_ANALYTICSACTIVITYTYPE AnalyticsActivityType = iota
    MEETING_ANALYTICSACTIVITYTYPE
    FOCUS_ANALYTICSACTIVITYTYPE
    CHAT_ANALYTICSACTIVITYTYPE
    CALL_ANALYTICSACTIVITYTYPE
)

func (i AnalyticsActivityType) String() string {
    return []string{"Email", "Meeting", "Focus", "Chat", "Call"}[i]
}
func ParseAnalyticsActivityType(v string) (interface{}, error) {
    result := EMAIL_ANALYTICSACTIVITYTYPE
    switch v {
        case "Email":
            result = EMAIL_ANALYTICSACTIVITYTYPE
        case "Meeting":
            result = MEETING_ANALYTICSACTIVITYTYPE
        case "Focus":
            result = FOCUS_ANALYTICSACTIVITYTYPE
        case "Chat":
            result = CHAT_ANALYTICSACTIVITYTYPE
        case "Call":
            result = CALL_ANALYTICSACTIVITYTYPE
        default:
            return 0, errors.New("Unknown AnalyticsActivityType value: " + v)
    }
    return &result, nil
}
func SerializeAnalyticsActivityType(values []AnalyticsActivityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
