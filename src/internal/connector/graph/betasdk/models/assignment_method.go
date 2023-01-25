package models
import (
    "errors"
)
// Provides operations to call the evaluateApplication method.
type AssignmentMethod int

const (
    STANDARD_ASSIGNMENTMETHOD AssignmentMethod = iota
    PRIVILEGED_ASSIGNMENTMETHOD
    AUTO_ASSIGNMENTMETHOD
)

func (i AssignmentMethod) String() string {
    return []string{"standard", "privileged", "auto"}[i]
}
func ParseAssignmentMethod(v string) (interface{}, error) {
    result := STANDARD_ASSIGNMENTMETHOD
    switch v {
        case "standard":
            result = STANDARD_ASSIGNMENTMETHOD
        case "privileged":
            result = PRIVILEGED_ASSIGNMENTMETHOD
        case "auto":
            result = AUTO_ASSIGNMENTMETHOD
        default:
            return 0, errors.New("Unknown AssignmentMethod value: " + v)
    }
    return &result, nil
}
func SerializeAssignmentMethod(values []AssignmentMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
