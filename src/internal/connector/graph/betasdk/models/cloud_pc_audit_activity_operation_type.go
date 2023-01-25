package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcAuditActivityOperationType int

const (
    CREATE_CLOUDPCAUDITACTIVITYOPERATIONTYPE CloudPcAuditActivityOperationType = iota
    DELETE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
    PATCH_CLOUDPCAUDITACTIVITYOPERATIONTYPE
    OTHER_CLOUDPCAUDITACTIVITYOPERATIONTYPE
)

func (i CloudPcAuditActivityOperationType) String() string {
    return []string{"create", "delete", "patch", "other"}[i]
}
func ParseCloudPcAuditActivityOperationType(v string) (interface{}, error) {
    result := CREATE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
    switch v {
        case "create":
            result = CREATE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        case "delete":
            result = DELETE_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        case "patch":
            result = PATCH_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        case "other":
            result = OTHER_CLOUDPCAUDITACTIVITYOPERATIONTYPE
        default:
            return 0, errors.New("Unknown CloudPcAuditActivityOperationType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditActivityOperationType(values []CloudPcAuditActivityOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
