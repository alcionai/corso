package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PlannerExternalTaskSourceDisplayType int

const (
    NONE_PLANNEREXTERNALTASKSOURCEDISPLAYTYPE PlannerExternalTaskSourceDisplayType = iota
    DEFAULT_ESCAPED_PLANNEREXTERNALTASKSOURCEDISPLAYTYPE
    UNKNOWNFUTUREVALUE_PLANNEREXTERNALTASKSOURCEDISPLAYTYPE
)

func (i PlannerExternalTaskSourceDisplayType) String() string {
    return []string{"none", "default", "unknownFutureValue"}[i]
}
func ParsePlannerExternalTaskSourceDisplayType(v string) (interface{}, error) {
    result := NONE_PLANNEREXTERNALTASKSOURCEDISPLAYTYPE
    switch v {
        case "none":
            result = NONE_PLANNEREXTERNALTASKSOURCEDISPLAYTYPE
        case "default":
            result = DEFAULT_ESCAPED_PLANNEREXTERNALTASKSOURCEDISPLAYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNEREXTERNALTASKSOURCEDISPLAYTYPE
        default:
            return 0, errors.New("Unknown PlannerExternalTaskSourceDisplayType value: " + v)
    }
    return &result, nil
}
func SerializePlannerExternalTaskSourceDisplayType(values []PlannerExternalTaskSourceDisplayType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
