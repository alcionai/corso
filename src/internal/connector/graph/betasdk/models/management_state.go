package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ManagementState int

const (
    // The device is under management
    MANAGED_MANAGEMENTSTATE ManagementState = iota
    // A retire command is occuring on the device and in the process of unenrolling from management
    RETIREPENDING_MANAGEMENTSTATE
    // Retire command failed on the device
    RETIREFAILED_MANAGEMENTSTATE
    // A wipe command is occuring on the device and in the process of unenrolling from management
    WIPEPENDING_MANAGEMENTSTATE
    // Wipe command failed on the device
    WIPEFAILED_MANAGEMENTSTATE
    // The device is unhealthy.
    UNHEALTHY_MANAGEMENTSTATE
    // A delete command is occuring on the device 
    DELETEPENDING_MANAGEMENTSTATE
    // A retire command was issued for the device
    RETIREISSUED_MANAGEMENTSTATE
    // A wipe command was issued for the device
    WIPEISSUED_MANAGEMENTSTATE
    // A wipe command for this device has been canceled
    WIPECANCELED_MANAGEMENTSTATE
    // A retire command for this device has been canceled
    RETIRECANCELED_MANAGEMENTSTATE
    // The device is discovered but not fully enrolled.
    DISCOVERED_MANAGEMENTSTATE
)

func (i ManagementState) String() string {
    return []string{"managed", "retirePending", "retireFailed", "wipePending", "wipeFailed", "unhealthy", "deletePending", "retireIssued", "wipeIssued", "wipeCanceled", "retireCanceled", "discovered"}[i]
}
func ParseManagementState(v string) (interface{}, error) {
    result := MANAGED_MANAGEMENTSTATE
    switch v {
        case "managed":
            result = MANAGED_MANAGEMENTSTATE
        case "retirePending":
            result = RETIREPENDING_MANAGEMENTSTATE
        case "retireFailed":
            result = RETIREFAILED_MANAGEMENTSTATE
        case "wipePending":
            result = WIPEPENDING_MANAGEMENTSTATE
        case "wipeFailed":
            result = WIPEFAILED_MANAGEMENTSTATE
        case "unhealthy":
            result = UNHEALTHY_MANAGEMENTSTATE
        case "deletePending":
            result = DELETEPENDING_MANAGEMENTSTATE
        case "retireIssued":
            result = RETIREISSUED_MANAGEMENTSTATE
        case "wipeIssued":
            result = WIPEISSUED_MANAGEMENTSTATE
        case "wipeCanceled":
            result = WIPECANCELED_MANAGEMENTSTATE
        case "retireCanceled":
            result = RETIRECANCELED_MANAGEMENTSTATE
        case "discovered":
            result = DISCOVERED_MANAGEMENTSTATE
        default:
            return 0, errors.New("Unknown ManagementState value: " + v)
    }
    return &result, nil
}
func SerializeManagementState(values []ManagementState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
