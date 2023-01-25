package models
import (
    "errors"
)
// Provides operations to call the add method.
type GroupPolicyOperationStatus int

const (
    // Group Policy unknown operation status.
    UNKNOWN_GROUPPOLICYOPERATIONSTATUS GroupPolicyOperationStatus = iota
    // Group Policy in progress operation status.
    INPROGRESS_GROUPPOLICYOPERATIONSTATUS
    // Group Policy successful operation status.
    SUCCESS_GROUPPOLICYOPERATIONSTATUS
    // Group Policy failed operation status.
    FAILED_GROUPPOLICYOPERATIONSTATUS
)

func (i GroupPolicyOperationStatus) String() string {
    return []string{"unknown", "inProgress", "success", "failed"}[i]
}
func ParseGroupPolicyOperationStatus(v string) (interface{}, error) {
    result := UNKNOWN_GROUPPOLICYOPERATIONSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_GROUPPOLICYOPERATIONSTATUS
        case "inProgress":
            result = INPROGRESS_GROUPPOLICYOPERATIONSTATUS
        case "success":
            result = SUCCESS_GROUPPOLICYOPERATIONSTATUS
        case "failed":
            result = FAILED_GROUPPOLICYOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown GroupPolicyOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyOperationStatus(values []GroupPolicyOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
