package security
import (
    "errors"
)
// Provides operations to call the add method.
type PolicyStatus int

const (
    PENDING_POLICYSTATUS PolicyStatus = iota
    ERROR_POLICYSTATUS
    SUCCESS_POLICYSTATUS
    UNKNOWNFUTUREVALUE_POLICYSTATUS
)

func (i PolicyStatus) String() string {
    return []string{"pending", "error", "success", "unknownFutureValue"}[i]
}
func ParsePolicyStatus(v string) (interface{}, error) {
    result := PENDING_POLICYSTATUS
    switch v {
        case "pending":
            result = PENDING_POLICYSTATUS
        case "error":
            result = ERROR_POLICYSTATUS
        case "success":
            result = SUCCESS_POLICYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_POLICYSTATUS
        default:
            return 0, errors.New("Unknown PolicyStatus value: " + v)
    }
    return &result, nil
}
func SerializePolicyStatus(values []PolicyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
