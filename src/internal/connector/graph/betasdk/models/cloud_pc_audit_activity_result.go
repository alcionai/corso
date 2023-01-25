package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcAuditActivityResult int

const (
    SUCCESS_CLOUDPCAUDITACTIVITYRESULT CloudPcAuditActivityResult = iota
    CLIENTERROR_CLOUDPCAUDITACTIVITYRESULT
    FAILURE_CLOUDPCAUDITACTIVITYRESULT
    TIMEOUT_CLOUDPCAUDITACTIVITYRESULT
    OTHER_CLOUDPCAUDITACTIVITYRESULT
)

func (i CloudPcAuditActivityResult) String() string {
    return []string{"success", "clientError", "failure", "timeout", "other"}[i]
}
func ParseCloudPcAuditActivityResult(v string) (interface{}, error) {
    result := SUCCESS_CLOUDPCAUDITACTIVITYRESULT
    switch v {
        case "success":
            result = SUCCESS_CLOUDPCAUDITACTIVITYRESULT
        case "clientError":
            result = CLIENTERROR_CLOUDPCAUDITACTIVITYRESULT
        case "failure":
            result = FAILURE_CLOUDPCAUDITACTIVITYRESULT
        case "timeout":
            result = TIMEOUT_CLOUDPCAUDITACTIVITYRESULT
        case "other":
            result = OTHER_CLOUDPCAUDITACTIVITYRESULT
        default:
            return 0, errors.New("Unknown CloudPcAuditActivityResult value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcAuditActivityResult(values []CloudPcAuditActivityResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
