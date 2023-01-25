package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DriverApprovalStatus int

const (
    // This indicates a driver needs IT admin's review.
    NEEDSREVIEW_DRIVERAPPROVALSTATUS DriverApprovalStatus = iota
    // This indicates IT admin has declined a driver.
    DECLINED_DRIVERAPPROVALSTATUS
    // This indicates IT admin has approved a driver.
    APPROVED_DRIVERAPPROVALSTATUS
    // This indicates IT admin has suspended a driver.
    SUSPENDED_DRIVERAPPROVALSTATUS
)

func (i DriverApprovalStatus) String() string {
    return []string{"needsReview", "declined", "approved", "suspended"}[i]
}
func ParseDriverApprovalStatus(v string) (interface{}, error) {
    result := NEEDSREVIEW_DRIVERAPPROVALSTATUS
    switch v {
        case "needsReview":
            result = NEEDSREVIEW_DRIVERAPPROVALSTATUS
        case "declined":
            result = DECLINED_DRIVERAPPROVALSTATUS
        case "approved":
            result = APPROVED_DRIVERAPPROVALSTATUS
        case "suspended":
            result = SUSPENDED_DRIVERAPPROVALSTATUS
        default:
            return 0, errors.New("Unknown DriverApprovalStatus value: " + v)
    }
    return &result, nil
}
func SerializeDriverApprovalStatus(values []DriverApprovalStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
