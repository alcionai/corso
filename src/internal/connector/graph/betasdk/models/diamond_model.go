package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DiamondModel int

const (
    UNKNOWN_DIAMONDMODEL DiamondModel = iota
    ADVERSARY_DIAMONDMODEL
    CAPABILITY_DIAMONDMODEL
    INFRASTRUCTURE_DIAMONDMODEL
    VICTIM_DIAMONDMODEL
    UNKNOWNFUTUREVALUE_DIAMONDMODEL
)

func (i DiamondModel) String() string {
    return []string{"unknown", "adversary", "capability", "infrastructure", "victim", "unknownFutureValue"}[i]
}
func ParseDiamondModel(v string) (interface{}, error) {
    result := UNKNOWN_DIAMONDMODEL
    switch v {
        case "unknown":
            result = UNKNOWN_DIAMONDMODEL
        case "adversary":
            result = ADVERSARY_DIAMONDMODEL
        case "capability":
            result = CAPABILITY_DIAMONDMODEL
        case "infrastructure":
            result = INFRASTRUCTURE_DIAMONDMODEL
        case "victim":
            result = VICTIM_DIAMONDMODEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DIAMONDMODEL
        default:
            return 0, errors.New("Unknown DiamondModel value: " + v)
    }
    return &result, nil
}
func SerializeDiamondModel(values []DiamondModel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
