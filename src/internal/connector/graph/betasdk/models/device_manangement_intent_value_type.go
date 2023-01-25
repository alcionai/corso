package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceManangementIntentValueType int

const (
    // The setting value is an integer
    INTEGER_DEVICEMANANGEMENTINTENTVALUETYPE DeviceManangementIntentValueType = iota
    // The setting value is a boolean
    BOOLEAN_DEVICEMANANGEMENTINTENTVALUETYPE
    // The setting value is a string
    STRING_DEVICEMANANGEMENTINTENTVALUETYPE
    // The setting value is a complex object
    COMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
    // The setting value is a collection
    COLLECTION_DEVICEMANANGEMENTINTENTVALUETYPE
    // The setting value is an abstract complex object
    ABSTRACTCOMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
)

func (i DeviceManangementIntentValueType) String() string {
    return []string{"integer", "boolean", "string", "complex", "collection", "abstractComplex"}[i]
}
func ParseDeviceManangementIntentValueType(v string) (interface{}, error) {
    result := INTEGER_DEVICEMANANGEMENTINTENTVALUETYPE
    switch v {
        case "integer":
            result = INTEGER_DEVICEMANANGEMENTINTENTVALUETYPE
        case "boolean":
            result = BOOLEAN_DEVICEMANANGEMENTINTENTVALUETYPE
        case "string":
            result = STRING_DEVICEMANANGEMENTINTENTVALUETYPE
        case "complex":
            result = COMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
        case "collection":
            result = COLLECTION_DEVICEMANANGEMENTINTENTVALUETYPE
        case "abstractComplex":
            result = ABSTRACTCOMPLEX_DEVICEMANANGEMENTINTENTVALUETYPE
        default:
            return 0, errors.New("Unknown DeviceManangementIntentValueType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManangementIntentValueType(values []DeviceManangementIntentValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
