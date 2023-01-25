package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AttributeFlowBehavior int

const (
    FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR AttributeFlowBehavior = iota
    FLOWALWAYS_ATTRIBUTEFLOWBEHAVIOR
)

func (i AttributeFlowBehavior) String() string {
    return []string{"FlowWhenChanged", "FlowAlways"}[i]
}
func ParseAttributeFlowBehavior(v string) (interface{}, error) {
    result := FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR
    switch v {
        case "FlowWhenChanged":
            result = FLOWWHENCHANGED_ATTRIBUTEFLOWBEHAVIOR
        case "FlowAlways":
            result = FLOWALWAYS_ATTRIBUTEFLOWBEHAVIOR
        default:
            return 0, errors.New("Unknown AttributeFlowBehavior value: " + v)
    }
    return &result, nil
}
func SerializeAttributeFlowBehavior(values []AttributeFlowBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
