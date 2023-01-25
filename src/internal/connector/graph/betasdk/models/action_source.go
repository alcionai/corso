package models
import (
    "errors"
)
// Provides operations to manage the columns property of the microsoft.graph.site entity.
type ActionSource int

const (
    MANUAL_ACTIONSOURCE ActionSource = iota
    AUTOMATIC_ACTIONSOURCE
    RECOMMENDED_ACTIONSOURCE
    DEFAULT_ESCAPED_ACTIONSOURCE
)

func (i ActionSource) String() string {
    return []string{"manual", "automatic", "recommended", "default"}[i]
}
func ParseActionSource(v string) (interface{}, error) {
    result := MANUAL_ACTIONSOURCE
    switch v {
        case "manual":
            result = MANUAL_ACTIONSOURCE
        case "automatic":
            result = AUTOMATIC_ACTIONSOURCE
        case "recommended":
            result = RECOMMENDED_ACTIONSOURCE
        case "default":
            result = DEFAULT_ESCAPED_ACTIONSOURCE
        default:
            return 0, errors.New("Unknown ActionSource value: " + v)
    }
    return &result, nil
}
func SerializeActionSource(values []ActionSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
