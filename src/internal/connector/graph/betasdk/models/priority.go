package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type Priority int

const (
    NONE_PRIORITY Priority = iota
    HIGH_PRIORITY
    LOW_PRIORITY
)

func (i Priority) String() string {
    return []string{"None", "High", "Low"}[i]
}
func ParsePriority(v string) (interface{}, error) {
    result := NONE_PRIORITY
    switch v {
        case "None":
            result = NONE_PRIORITY
        case "High":
            result = HIGH_PRIORITY
        case "Low":
            result = LOW_PRIORITY
        default:
            return 0, errors.New("Unknown Priority value: " + v)
    }
    return &result, nil
}
func SerializePriority(values []Priority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
