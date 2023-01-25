package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ApprovalState int

const (
    PENDING_APPROVALSTATE ApprovalState = iota
    APPROVED_APPROVALSTATE
    DENIED_APPROVALSTATE
    ABORTED_APPROVALSTATE
    CANCELED_APPROVALSTATE
)

func (i ApprovalState) String() string {
    return []string{"pending", "approved", "denied", "aborted", "canceled"}[i]
}
func ParseApprovalState(v string) (interface{}, error) {
    result := PENDING_APPROVALSTATE
    switch v {
        case "pending":
            result = PENDING_APPROVALSTATE
        case "approved":
            result = APPROVED_APPROVALSTATE
        case "denied":
            result = DENIED_APPROVALSTATE
        case "aborted":
            result = ABORTED_APPROVALSTATE
        case "canceled":
            result = CANCELED_APPROVALSTATE
        default:
            return 0, errors.New("Unknown ApprovalState value: " + v)
    }
    return &result, nil
}
func SerializeApprovalState(values []ApprovalState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
