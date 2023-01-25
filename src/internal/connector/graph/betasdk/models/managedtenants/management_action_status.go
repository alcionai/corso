package managedtenants
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ManagementActionStatus int

const (
    TOADDRESS_MANAGEMENTACTIONSTATUS ManagementActionStatus = iota
    COMPLETED_MANAGEMENTACTIONSTATUS
    ERROR_MANAGEMENTACTIONSTATUS
    TIMEOUT_MANAGEMENTACTIONSTATUS
    INPROGRESS_MANAGEMENTACTIONSTATUS
    PLANNED_MANAGEMENTACTIONSTATUS
    RESOLVEDBY3RDPARTY_MANAGEMENTACTIONSTATUS
    RESOLVEDTHROUGHALTERNATEMITIGATION_MANAGEMENTACTIONSTATUS
    RISKACCEPTED_MANAGEMENTACTIONSTATUS
    UNKNOWNFUTUREVALUE_MANAGEMENTACTIONSTATUS
)

func (i ManagementActionStatus) String() string {
    return []string{"toAddress", "completed", "error", "timeOut", "inProgress", "planned", "resolvedBy3rdParty", "resolvedThroughAlternateMitigation", "riskAccepted", "unknownFutureValue"}[i]
}
func ParseManagementActionStatus(v string) (interface{}, error) {
    result := TOADDRESS_MANAGEMENTACTIONSTATUS
    switch v {
        case "toAddress":
            result = TOADDRESS_MANAGEMENTACTIONSTATUS
        case "completed":
            result = COMPLETED_MANAGEMENTACTIONSTATUS
        case "error":
            result = ERROR_MANAGEMENTACTIONSTATUS
        case "timeOut":
            result = TIMEOUT_MANAGEMENTACTIONSTATUS
        case "inProgress":
            result = INPROGRESS_MANAGEMENTACTIONSTATUS
        case "planned":
            result = PLANNED_MANAGEMENTACTIONSTATUS
        case "resolvedBy3rdParty":
            result = RESOLVEDBY3RDPARTY_MANAGEMENTACTIONSTATUS
        case "resolvedThroughAlternateMitigation":
            result = RESOLVEDTHROUGHALTERNATEMITIGATION_MANAGEMENTACTIONSTATUS
        case "riskAccepted":
            result = RISKACCEPTED_MANAGEMENTACTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTACTIONSTATUS
        default:
            return 0, errors.New("Unknown ManagementActionStatus value: " + v)
    }
    return &result, nil
}
func SerializeManagementActionStatus(values []ManagementActionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
