package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AccountStatus int

const (
    UNKNOWN_ACCOUNTSTATUS AccountStatus = iota
    STAGED_ACCOUNTSTATUS
    ACTIVE_ACCOUNTSTATUS
    SUSPENDED_ACCOUNTSTATUS
    DELETED_ACCOUNTSTATUS
    UNKNOWNFUTUREVALUE_ACCOUNTSTATUS
)

func (i AccountStatus) String() string {
    return []string{"unknown", "staged", "active", "suspended", "deleted", "unknownFutureValue"}[i]
}
func ParseAccountStatus(v string) (interface{}, error) {
    result := UNKNOWN_ACCOUNTSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_ACCOUNTSTATUS
        case "staged":
            result = STAGED_ACCOUNTSTATUS
        case "active":
            result = ACTIVE_ACCOUNTSTATUS
        case "suspended":
            result = SUSPENDED_ACCOUNTSTATUS
        case "deleted":
            result = DELETED_ACCOUNTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCOUNTSTATUS
        default:
            return 0, errors.New("Unknown AccountStatus value: " + v)
    }
    return &result, nil
}
func SerializeAccountStatus(values []AccountStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
