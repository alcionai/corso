package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AssignmentFilterPayloadType int

const (
    // NotSet
    NOTSET_ASSIGNMENTFILTERPAYLOADTYPE AssignmentFilterPayloadType = iota
    // EnrollmentRestrictions
    ENROLLMENTRESTRICTIONS_ASSIGNMENTFILTERPAYLOADTYPE
)

func (i AssignmentFilterPayloadType) String() string {
    return []string{"notSet", "enrollmentRestrictions"}[i]
}
func ParseAssignmentFilterPayloadType(v string) (interface{}, error) {
    result := NOTSET_ASSIGNMENTFILTERPAYLOADTYPE
    switch v {
        case "notSet":
            result = NOTSET_ASSIGNMENTFILTERPAYLOADTYPE
        case "enrollmentRestrictions":
            result = ENROLLMENTRESTRICTIONS_ASSIGNMENTFILTERPAYLOADTYPE
        default:
            return 0, errors.New("Unknown AssignmentFilterPayloadType value: " + v)
    }
    return &result, nil
}
func SerializeAssignmentFilterPayloadType(values []AssignmentFilterPayloadType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
