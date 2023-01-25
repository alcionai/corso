package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PrivilegedAccessGroupAssignmentType int

const (
    ASSIGNED_PRIVILEGEDACCESSGROUPASSIGNMENTTYPE PrivilegedAccessGroupAssignmentType = iota
    ACTIVATED_PRIVILEGEDACCESSGROUPASSIGNMENTTYPE
    UNKNOWNFUTUREVALUE_PRIVILEGEDACCESSGROUPASSIGNMENTTYPE
)

func (i PrivilegedAccessGroupAssignmentType) String() string {
    return []string{"assigned", "activated", "unknownFutureValue"}[i]
}
func ParsePrivilegedAccessGroupAssignmentType(v string) (interface{}, error) {
    result := ASSIGNED_PRIVILEGEDACCESSGROUPASSIGNMENTTYPE
    switch v {
        case "assigned":
            result = ASSIGNED_PRIVILEGEDACCESSGROUPASSIGNMENTTYPE
        case "activated":
            result = ACTIVATED_PRIVILEGEDACCESSGROUPASSIGNMENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRIVILEGEDACCESSGROUPASSIGNMENTTYPE
        default:
            return 0, errors.New("Unknown PrivilegedAccessGroupAssignmentType value: " + v)
    }
    return &result, nil
}
func SerializePrivilegedAccessGroupAssignmentType(values []PrivilegedAccessGroupAssignmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
