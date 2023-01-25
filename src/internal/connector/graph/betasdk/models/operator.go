package models
import (
    "errors"
)
// Provides operations to manage the columns property of the microsoft.graph.site entity.
type Operator int

const (
    // None operator.
    NONE_OPERATOR Operator = iota
    // And operator.
    AND_OPERATOR
    // Or operator.
    OR_OPERATOR
    // IsEquals operator.
    ISEQUALS_OPERATOR
    // NotEquals operator.
    NOTEQUALS_OPERATOR
    // GreaterThan operator.
    GREATERTHAN_OPERATOR
    // LessThan operator.
    LESSTHAN_OPERATOR
    // Between operator.
    BETWEEN_OPERATOR
    // NotBetween operator.
    NOTBETWEEN_OPERATOR
    // GreaterEquals operator.
    GREATEREQUALS_OPERATOR
    // LessEquals operator.
    LESSEQUALS_OPERATOR
    // DayTimeBetween operator.
    DAYTIMEBETWEEN_OPERATOR
    // BeginsWith operator.
    BEGINSWITH_OPERATOR
    // NotBeginsWith operator.
    NOTBEGINSWITH_OPERATOR
    // EndsWith operator.
    ENDSWITH_OPERATOR
    // NotEndsWith operator.
    NOTENDSWITH_OPERATOR
    // Contains operator.
    CONTAINS_OPERATOR
    // NotContains operator.
    NOTCONTAINS_OPERATOR
    // AllOf operator.
    ALLOF_OPERATOR
    // OneOf operator.
    ONEOF_OPERATOR
    // NoneOf operator.
    NONEOF_OPERATOR
    // SetEquals operator.
    SETEQUALS_OPERATOR
    // OrderedSetEquals operator.
    ORDEREDSETEQUALS_OPERATOR
    // SubsetOf operator.
    SUBSETOF_OPERATOR
    // ExcludesAll operator.
    EXCLUDESALL_OPERATOR
)

func (i Operator) String() string {
    return []string{"none", "and", "or", "isEquals", "notEquals", "greaterThan", "lessThan", "between", "notBetween", "greaterEquals", "lessEquals", "dayTimeBetween", "beginsWith", "notBeginsWith", "endsWith", "notEndsWith", "contains", "notContains", "allOf", "oneOf", "noneOf", "setEquals", "orderedSetEquals", "subsetOf", "excludesAll"}[i]
}
func ParseOperator(v string) (interface{}, error) {
    result := NONE_OPERATOR
    switch v {
        case "none":
            result = NONE_OPERATOR
        case "and":
            result = AND_OPERATOR
        case "or":
            result = OR_OPERATOR
        case "isEquals":
            result = ISEQUALS_OPERATOR
        case "notEquals":
            result = NOTEQUALS_OPERATOR
        case "greaterThan":
            result = GREATERTHAN_OPERATOR
        case "lessThan":
            result = LESSTHAN_OPERATOR
        case "between":
            result = BETWEEN_OPERATOR
        case "notBetween":
            result = NOTBETWEEN_OPERATOR
        case "greaterEquals":
            result = GREATEREQUALS_OPERATOR
        case "lessEquals":
            result = LESSEQUALS_OPERATOR
        case "dayTimeBetween":
            result = DAYTIMEBETWEEN_OPERATOR
        case "beginsWith":
            result = BEGINSWITH_OPERATOR
        case "notBeginsWith":
            result = NOTBEGINSWITH_OPERATOR
        case "endsWith":
            result = ENDSWITH_OPERATOR
        case "notEndsWith":
            result = NOTENDSWITH_OPERATOR
        case "contains":
            result = CONTAINS_OPERATOR
        case "notContains":
            result = NOTCONTAINS_OPERATOR
        case "allOf":
            result = ALLOF_OPERATOR
        case "oneOf":
            result = ONEOF_OPERATOR
        case "noneOf":
            result = NONEOF_OPERATOR
        case "setEquals":
            result = SETEQUALS_OPERATOR
        case "orderedSetEquals":
            result = ORDEREDSETEQUALS_OPERATOR
        case "subsetOf":
            result = SUBSETOF_OPERATOR
        case "excludesAll":
            result = EXCLUDESALL_OPERATOR
        default:
            return 0, errors.New("Unknown Operator value: " + v)
    }
    return &result, nil
}
func SerializeOperator(values []Operator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
