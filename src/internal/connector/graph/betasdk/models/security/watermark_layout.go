package security
import (
    "errors"
)
// Provides operations to call the add method.
type WatermarkLayout int

const (
    HORIZONTAL_WATERMARKLAYOUT WatermarkLayout = iota
    DIAGONAL_WATERMARKLAYOUT
)

func (i WatermarkLayout) String() string {
    return []string{"horizontal", "diagonal"}[i]
}
func ParseWatermarkLayout(v string) (interface{}, error) {
    result := HORIZONTAL_WATERMARKLAYOUT
    switch v {
        case "horizontal":
            result = HORIZONTAL_WATERMARKLAYOUT
        case "diagonal":
            result = DIAGONAL_WATERMARKLAYOUT
        default:
            return 0, errors.New("Unknown WatermarkLayout value: " + v)
    }
    return &result, nil
}
func SerializeWatermarkLayout(values []WatermarkLayout) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
