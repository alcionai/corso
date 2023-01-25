package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DelegatedAdminRelationshipOperationType int

const (
    DELEGATEDADMINACCESSASSIGNMENTUPDATE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE DelegatedAdminRelationshipOperationType = iota
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
)

func (i DelegatedAdminRelationshipOperationType) String() string {
    return []string{"delegatedAdminAccessAssignmentUpdate", "unknownFutureValue"}[i]
}
func ParseDelegatedAdminRelationshipOperationType(v string) (interface{}, error) {
    result := DELEGATEDADMINACCESSASSIGNMENTUPDATE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
    switch v {
        case "delegatedAdminAccessAssignmentUpdate":
            result = DELEGATEDADMINACCESSASSIGNMENTUPDATE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPOPERATIONTYPE
        default:
            return 0, errors.New("Unknown DelegatedAdminRelationshipOperationType value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminRelationshipOperationType(values []DelegatedAdminRelationshipOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
