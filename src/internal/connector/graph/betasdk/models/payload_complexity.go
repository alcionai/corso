package models
import (
    "errors"
)
// Provides operations to call the add method.
type PayloadComplexity int

const (
    UNKNOWN_PAYLOADCOMPLEXITY PayloadComplexity = iota
    LOW_PAYLOADCOMPLEXITY
    MEDIUM_PAYLOADCOMPLEXITY
    HIGH_PAYLOADCOMPLEXITY
    UNKNOWNFUTUREVALUE_PAYLOADCOMPLEXITY
)

func (i PayloadComplexity) String() string {
    return []string{"unknown", "low", "medium", "high", "unknownFutureValue"}[i]
}
func ParsePayloadComplexity(v string) (interface{}, error) {
    result := UNKNOWN_PAYLOADCOMPLEXITY
    switch v {
        case "unknown":
            result = UNKNOWN_PAYLOADCOMPLEXITY
        case "low":
            result = LOW_PAYLOADCOMPLEXITY
        case "medium":
            result = MEDIUM_PAYLOADCOMPLEXITY
        case "high":
            result = HIGH_PAYLOADCOMPLEXITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PAYLOADCOMPLEXITY
        default:
            return 0, errors.New("Unknown PayloadComplexity value: " + v)
    }
    return &result, nil
}
func SerializePayloadComplexity(values []PayloadComplexity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
