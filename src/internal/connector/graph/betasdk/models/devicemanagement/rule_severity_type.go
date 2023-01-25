package devicemanagement
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RuleSeverityType int

const (
    UNKNOWN_RULESEVERITYTYPE RuleSeverityType = iota
    INFORMATIONAL_RULESEVERITYTYPE
    WARNING_RULESEVERITYTYPE
    CRITICAL_RULESEVERITYTYPE
    UNKNOWNFUTUREVALUE_RULESEVERITYTYPE
)

func (i RuleSeverityType) String() string {
    return []string{"unknown", "informational", "warning", "critical", "unknownFutureValue"}[i]
}
func ParseRuleSeverityType(v string) (interface{}, error) {
    result := UNKNOWN_RULESEVERITYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_RULESEVERITYTYPE
        case "informational":
            result = INFORMATIONAL_RULESEVERITYTYPE
        case "warning":
            result = WARNING_RULESEVERITYTYPE
        case "critical":
            result = CRITICAL_RULESEVERITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RULESEVERITYTYPE
        default:
            return 0, errors.New("Unknown RuleSeverityType value: " + v)
    }
    return &result, nil
}
func SerializeRuleSeverityType(values []RuleSeverityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
