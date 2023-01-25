package managedtenants
import (
    "errors"
)
// Provides operations to call the add method.
type DelegatedPrivilegeStatus int

const (
    NONE_DELEGATEDPRIVILEGESTATUS DelegatedPrivilegeStatus = iota
    DELEGATEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
    UNKNOWNFUTUREVALUE_DELEGATEDPRIVILEGESTATUS
    GRANULARDELEGATEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
    DELEGATEDANDGRANULARDELEGETEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
)

func (i DelegatedPrivilegeStatus) String() string {
    return []string{"none", "delegatedAdminPrivileges", "unknownFutureValue", "granularDelegatedAdminPrivileges", "delegatedAndGranularDelegetedAdminPrivileges"}[i]
}
func ParseDelegatedPrivilegeStatus(v string) (interface{}, error) {
    result := NONE_DELEGATEDPRIVILEGESTATUS
    switch v {
        case "none":
            result = NONE_DELEGATEDPRIVILEGESTATUS
        case "delegatedAdminPrivileges":
            result = DELEGATEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DELEGATEDPRIVILEGESTATUS
        case "granularDelegatedAdminPrivileges":
            result = GRANULARDELEGATEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
        case "delegatedAndGranularDelegetedAdminPrivileges":
            result = DELEGATEDANDGRANULARDELEGETEDADMINPRIVILEGES_DELEGATEDPRIVILEGESTATUS
        default:
            return 0, errors.New("Unknown DelegatedPrivilegeStatus value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedPrivilegeStatus(values []DelegatedPrivilegeStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
