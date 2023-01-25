package models
import (
    "errors"
)
// Provides operations to call the add method.
type MlClassificationMatchTolerance int

const (
    EXACT_MLCLASSIFICATIONMATCHTOLERANCE MlClassificationMatchTolerance = iota
    NEAR_MLCLASSIFICATIONMATCHTOLERANCE
)

func (i MlClassificationMatchTolerance) String() string {
    return []string{"exact", "near"}[i]
}
func ParseMlClassificationMatchTolerance(v string) (interface{}, error) {
    result := EXACT_MLCLASSIFICATIONMATCHTOLERANCE
    switch v {
        case "exact":
            result = EXACT_MLCLASSIFICATIONMATCHTOLERANCE
        case "near":
            result = NEAR_MLCLASSIFICATIONMATCHTOLERANCE
        default:
            return 0, errors.New("Unknown MlClassificationMatchTolerance value: " + v)
    }
    return &result, nil
}
func SerializeMlClassificationMatchTolerance(values []MlClassificationMatchTolerance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
