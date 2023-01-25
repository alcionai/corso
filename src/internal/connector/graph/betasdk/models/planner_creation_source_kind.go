package models
import (
    "errors"
)
// Provides operations to call the add method.
type PlannerCreationSourceKind int

const (
    NONE_PLANNERCREATIONSOURCEKIND PlannerCreationSourceKind = iota
    EXTERNAL_PLANNERCREATIONSOURCEKIND
    PUBLICATION_PLANNERCREATIONSOURCEKIND
    UNKNOWNFUTUREVALUE_PLANNERCREATIONSOURCEKIND
)

func (i PlannerCreationSourceKind) String() string {
    return []string{"none", "external", "publication", "unknownFutureValue"}[i]
}
func ParsePlannerCreationSourceKind(v string) (interface{}, error) {
    result := NONE_PLANNERCREATIONSOURCEKIND
    switch v {
        case "none":
            result = NONE_PLANNERCREATIONSOURCEKIND
        case "external":
            result = EXTERNAL_PLANNERCREATIONSOURCEKIND
        case "publication":
            result = PUBLICATION_PLANNERCREATIONSOURCEKIND
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERCREATIONSOURCEKIND
        default:
            return 0, errors.New("Unknown PlannerCreationSourceKind value: " + v)
    }
    return &result, nil
}
func SerializePlannerCreationSourceKind(values []PlannerCreationSourceKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
