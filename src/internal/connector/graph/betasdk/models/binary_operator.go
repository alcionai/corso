package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type BinaryOperator int

const (
    OR_BINARYOPERATOR BinaryOperator = iota
    AND_BINARYOPERATOR
)

func (i BinaryOperator) String() string {
    return []string{"or", "and"}[i]
}
func ParseBinaryOperator(v string) (interface{}, error) {
    result := OR_BINARYOPERATOR
    switch v {
        case "or":
            result = OR_BINARYOPERATOR
        case "and":
            result = AND_BINARYOPERATOR
        default:
            return 0, errors.New("Unknown BinaryOperator value: " + v)
    }
    return &result, nil
}
func SerializeBinaryOperator(values []BinaryOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
