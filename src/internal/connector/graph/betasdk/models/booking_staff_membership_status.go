package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type BookingStaffMembershipStatus int

const (
    ACTIVE_BOOKINGSTAFFMEMBERSHIPSTATUS BookingStaffMembershipStatus = iota
    PENDINGACCEPTANCE_BOOKINGSTAFFMEMBERSHIPSTATUS
    REJECTEDBYSTAFF_BOOKINGSTAFFMEMBERSHIPSTATUS
    UNKNOWNFUTUREVALUE_BOOKINGSTAFFMEMBERSHIPSTATUS
)

func (i BookingStaffMembershipStatus) String() string {
    return []string{"active", "pendingAcceptance", "rejectedByStaff", "unknownFutureValue"}[i]
}
func ParseBookingStaffMembershipStatus(v string) (interface{}, error) {
    result := ACTIVE_BOOKINGSTAFFMEMBERSHIPSTATUS
    switch v {
        case "active":
            result = ACTIVE_BOOKINGSTAFFMEMBERSHIPSTATUS
        case "pendingAcceptance":
            result = PENDINGACCEPTANCE_BOOKINGSTAFFMEMBERSHIPSTATUS
        case "rejectedByStaff":
            result = REJECTEDBYSTAFF_BOOKINGSTAFFMEMBERSHIPSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_BOOKINGSTAFFMEMBERSHIPSTATUS
        default:
            return 0, errors.New("Unknown BookingStaffMembershipStatus value: " + v)
    }
    return &result, nil
}
func SerializeBookingStaffMembershipStatus(values []BookingStaffMembershipStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
