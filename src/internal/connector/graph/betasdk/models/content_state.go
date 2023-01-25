package models
import (
    "errors"
)
// Provides operations to call the evaluateClassificationResults method.
type ContentState int

const (
    REST_CONTENTSTATE ContentState = iota
    MOTION_CONTENTSTATE
    USE_CONTENTSTATE
)

func (i ContentState) String() string {
    return []string{"rest", "motion", "use"}[i]
}
func ParseContentState(v string) (interface{}, error) {
    result := REST_CONTENTSTATE
    switch v {
        case "rest":
            result = REST_CONTENTSTATE
        case "motion":
            result = MOTION_CONTENTSTATE
        case "use":
            result = USE_CONTENTSTATE
        default:
            return 0, errors.New("Unknown ContentState value: " + v)
    }
    return &result, nil
}
func SerializeContentState(values []ContentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
