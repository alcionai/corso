package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Alignment int

const (
    LEFT_ALIGNMENT Alignment = iota
    RIGHT_ALIGNMENT
    CENTER_ALIGNMENT
)

func (i Alignment) String() string {
    return []string{"left", "right", "center"}[i]
}
func ParseAlignment(v string) (interface{}, error) {
    result := LEFT_ALIGNMENT
    switch v {
        case "left":
            result = LEFT_ALIGNMENT
        case "right":
            result = RIGHT_ALIGNMENT
        case "center":
            result = CENTER_ALIGNMENT
        default:
            return 0, errors.New("Unknown Alignment value: " + v)
    }
    return &result, nil
}
func SerializeAlignment(values []Alignment) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
