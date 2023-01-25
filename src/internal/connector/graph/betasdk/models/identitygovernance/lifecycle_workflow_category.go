package identitygovernance
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type LifecycleWorkflowCategory int

const (
    JOINER_LIFECYCLEWORKFLOWCATEGORY LifecycleWorkflowCategory = iota
    LEAVER_LIFECYCLEWORKFLOWCATEGORY
    UNKNOWNFUTUREVALUE_LIFECYCLEWORKFLOWCATEGORY
)

func (i LifecycleWorkflowCategory) String() string {
    return []string{"joiner", "leaver", "unknownFutureValue"}[i]
}
func ParseLifecycleWorkflowCategory(v string) (interface{}, error) {
    result := JOINER_LIFECYCLEWORKFLOWCATEGORY
    switch v {
        case "joiner":
            result = JOINER_LIFECYCLEWORKFLOWCATEGORY
        case "leaver":
            result = LEAVER_LIFECYCLEWORKFLOWCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LIFECYCLEWORKFLOWCATEGORY
        default:
            return 0, errors.New("Unknown LifecycleWorkflowCategory value: " + v)
    }
    return &result, nil
}
func SerializeLifecycleWorkflowCategory(values []LifecycleWorkflowCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
