package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type OverrideOption int

const (
    NOTALLOWED_OVERRIDEOPTION OverrideOption = iota
    ALLOWFALSEPOSITIVEOVERRIDE_OVERRIDEOPTION
    ALLOWWITHJUSTIFICATION_OVERRIDEOPTION
    ALLOWWITHOUTJUSTIFICATION_OVERRIDEOPTION
)

func (i OverrideOption) String() string {
    return []string{"notAllowed", "allowFalsePositiveOverride", "allowWithJustification", "allowWithoutJustification"}[i]
}
func ParseOverrideOption(v string) (interface{}, error) {
    result := NOTALLOWED_OVERRIDEOPTION
    switch v {
        case "notAllowed":
            result = NOTALLOWED_OVERRIDEOPTION
        case "allowFalsePositiveOverride":
            result = ALLOWFALSEPOSITIVEOVERRIDE_OVERRIDEOPTION
        case "allowWithJustification":
            result = ALLOWWITHJUSTIFICATION_OVERRIDEOPTION
        case "allowWithoutJustification":
            result = ALLOWWITHOUTJUSTIFICATION_OVERRIDEOPTION
        default:
            return 0, errors.New("Unknown OverrideOption value: " + v)
    }
    return &result, nil
}
func SerializeOverrideOption(values []OverrideOption) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
