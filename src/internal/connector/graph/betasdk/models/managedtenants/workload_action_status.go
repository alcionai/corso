package managedtenants
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WorkloadActionStatus int

const (
    TOADDRESS_WORKLOADACTIONSTATUS WorkloadActionStatus = iota
    COMPLETED_WORKLOADACTIONSTATUS
    ERROR_WORKLOADACTIONSTATUS
    TIMEOUT_WORKLOADACTIONSTATUS
    INPROGRESS_WORKLOADACTIONSTATUS
    UNKNOWNFUTUREVALUE_WORKLOADACTIONSTATUS
)

func (i WorkloadActionStatus) String() string {
    return []string{"toAddress", "completed", "error", "timeOut", "inProgress", "unknownFutureValue"}[i]
}
func ParseWorkloadActionStatus(v string) (interface{}, error) {
    result := TOADDRESS_WORKLOADACTIONSTATUS
    switch v {
        case "toAddress":
            result = TOADDRESS_WORKLOADACTIONSTATUS
        case "completed":
            result = COMPLETED_WORKLOADACTIONSTATUS
        case "error":
            result = ERROR_WORKLOADACTIONSTATUS
        case "timeOut":
            result = TIMEOUT_WORKLOADACTIONSTATUS
        case "inProgress":
            result = INPROGRESS_WORKLOADACTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WORKLOADACTIONSTATUS
        default:
            return 0, errors.New("Unknown WorkloadActionStatus value: " + v)
    }
    return &result, nil
}
func SerializeWorkloadActionStatus(values []WorkloadActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
