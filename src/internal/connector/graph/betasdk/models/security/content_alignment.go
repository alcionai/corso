package security
import (
    "errors"
)
// Provides operations to call the add method.
type ContentAlignment int

const (
    LEFT_CONTENTALIGNMENT ContentAlignment = iota
    RIGHT_CONTENTALIGNMENT
    CENTER_CONTENTALIGNMENT
)

func (i ContentAlignment) String() string {
    return []string{"left", "right", "center"}[i]
}
func ParseContentAlignment(v string) (interface{}, error) {
    result := LEFT_CONTENTALIGNMENT
    switch v {
        case "left":
            result = LEFT_CONTENTALIGNMENT
        case "right":
            result = RIGHT_CONTENTALIGNMENT
        case "center":
            result = CENTER_CONTENTALIGNMENT
        default:
            return 0, errors.New("Unknown ContentAlignment value: " + v)
    }
    return &result, nil
}
func SerializeContentAlignment(values []ContentAlignment) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
