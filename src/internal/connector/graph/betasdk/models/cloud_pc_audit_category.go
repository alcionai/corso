package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcAuditCategory int

const (
    CLOUDPC_CLOUDPCAUDITCATEGORY CloudPcAuditCategory = iota
    OTHER_CLOUDPCAUDITCATEGORY
)

func (i CloudPcAuditCategory) String() string {
    return []string{"cloudPC", "other"}[i]
}
func ParseCloudPcAuditCategory(v string) (interface{}, error) {
    result := CLOUDPC_CLOUDPCAUDITCATEGORY
    switch v {
        case "cloudPC":
            result = CLOUDPC_CLOUDPCAUDITCATEGORY
        case "other":
            result = OTHER_CLOUDPCAUDITCATEGORY
        default:
            return 0, errors.New("Unknown CloudPcAuditCategory value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditCategory(values []CloudPcAuditCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
