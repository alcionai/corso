package models
import (
    "errors"
)
// Provides operations to call the add method.
type WarrantyType int

const (
    // Unknown warranty type
    UNKNOWN_WARRANTYTYPE WarrantyType = iota
    // Manufacturer warranty
    MANUFACTURER_WARRANTYTYPE
    // Contractual warranty
    CONTRACTUAL_WARRANTYTYPE
    // Unknown future value
    UNKNOWNFUTUREVALUE_WARRANTYTYPE
)

func (i WarrantyType) String() string {
    return []string{"unknown", "manufacturer", "contractual", "unknownFutureValue"}[i]
}
func ParseWarrantyType(v string) (interface{}, error) {
    result := UNKNOWN_WARRANTYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_WARRANTYTYPE
        case "manufacturer":
            result = MANUFACTURER_WARRANTYTYPE
        case "contractual":
            result = CONTRACTUAL_WARRANTYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WARRANTYTYPE
        default:
            return 0, errors.New("Unknown WarrantyType value: " + v)
    }
    return &result, nil
}
func SerializeWarrantyType(values []WarrantyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
