package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceComplianceScriptRulOperator int

const (
    // None operator.
    NONE_DEVICECOMPLIANCESCRIPTRULOPERATOR DeviceComplianceScriptRulOperator = iota
    // And operator.
    AND_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // Or operator.
    OR_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // IsEquals operator.
    ISEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // NotEquals operator.
    NOTEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // GreaterThan operator.
    GREATERTHAN_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // LessThan operator.
    LESSTHAN_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // Between operator.
    BETWEEN_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // NotBetween operator.
    NOTBETWEEN_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // GreaterEquals operator.
    GREATEREQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // LessEquals operator.
    LESSEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // DayTimeBetween operator.
    DAYTIMEBETWEEN_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // BeginsWith operator.
    BEGINSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // NotBeginsWith operator.
    NOTBEGINSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // EndsWith operator.
    ENDSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // NotEndsWith operator.
    NOTENDSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // Contains operator.
    CONTAINS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // NotContains operator.
    NOTCONTAINS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // AllOf operator.
    ALLOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // OneOf operator.
    ONEOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // NoneOf operator.
    NONEOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // SetEquals operator.
    SETEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // OrderedSetEquals operator.
    ORDEREDSETEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // SubsetOf operator.
    SUBSETOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
    // ExcludesAll operator.
    EXCLUDESALL_DEVICECOMPLIANCESCRIPTRULOPERATOR
)

func (i DeviceComplianceScriptRulOperator) String() string {
    return []string{"none", "and", "or", "isEquals", "notEquals", "greaterThan", "lessThan", "between", "notBetween", "greaterEquals", "lessEquals", "dayTimeBetween", "beginsWith", "notBeginsWith", "endsWith", "notEndsWith", "contains", "notContains", "allOf", "oneOf", "noneOf", "setEquals", "orderedSetEquals", "subsetOf", "excludesAll"}[i]
}
func ParseDeviceComplianceScriptRulOperator(v string) (interface{}, error) {
    result := NONE_DEVICECOMPLIANCESCRIPTRULOPERATOR
    switch v {
        case "none":
            result = NONE_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "and":
            result = AND_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "or":
            result = OR_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "isEquals":
            result = ISEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "notEquals":
            result = NOTEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "greaterThan":
            result = GREATERTHAN_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "lessThan":
            result = LESSTHAN_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "between":
            result = BETWEEN_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "notBetween":
            result = NOTBETWEEN_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "greaterEquals":
            result = GREATEREQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "lessEquals":
            result = LESSEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "dayTimeBetween":
            result = DAYTIMEBETWEEN_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "beginsWith":
            result = BEGINSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "notBeginsWith":
            result = NOTBEGINSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "endsWith":
            result = ENDSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "notEndsWith":
            result = NOTENDSWITH_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "contains":
            result = CONTAINS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "notContains":
            result = NOTCONTAINS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "allOf":
            result = ALLOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "oneOf":
            result = ONEOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "noneOf":
            result = NONEOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "setEquals":
            result = SETEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "orderedSetEquals":
            result = ORDEREDSETEQUALS_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "subsetOf":
            result = SUBSETOF_DEVICECOMPLIANCESCRIPTRULOPERATOR
        case "excludesAll":
            result = EXCLUDESALL_DEVICECOMPLIANCESCRIPTRULOPERATOR
        default:
            return 0, errors.New("Unknown DeviceComplianceScriptRulOperator value: " + v)
    }
    return &result, nil
}
func SerializeDeviceComplianceScriptRulOperator(values []DeviceComplianceScriptRulOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
