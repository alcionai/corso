package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DriverUpdateProfileApprovalType int

const (
    // This indicates a driver and firmware profile needs to be approved manually.
    MANUAL_DRIVERUPDATEPROFILEAPPROVALTYPE DriverUpdateProfileApprovalType = iota
    // This indicates a driver and firmware profile is approved automatically.
    AUTOMATIC_DRIVERUPDATEPROFILEAPPROVALTYPE
)

func (i DriverUpdateProfileApprovalType) String() string {
    return []string{"manual", "automatic"}[i]
}
func ParseDriverUpdateProfileApprovalType(v string) (interface{}, error) {
    result := MANUAL_DRIVERUPDATEPROFILEAPPROVALTYPE
    switch v {
        case "manual":
            result = MANUAL_DRIVERUPDATEPROFILEAPPROVALTYPE
        case "automatic":
            result = AUTOMATIC_DRIVERUPDATEPROFILEAPPROVALTYPE
        default:
            return 0, errors.New("Unknown DriverUpdateProfileApprovalType value: " + v)
    }
    return &result, nil
}
func SerializeDriverUpdateProfileApprovalType(values []DriverUpdateProfileApprovalType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
