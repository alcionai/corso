package security
import (
    "errors"
)
// Provides operations to call the add method.
type TenantAllowBlockListEntryType int

const (
    URL_TENANTALLOWBLOCKLISTENTRYTYPE TenantAllowBlockListEntryType = iota
    FILEHASH_TENANTALLOWBLOCKLISTENTRYTYPE
    SENDER_TENANTALLOWBLOCKLISTENTRYTYPE
    RECIPIENT_TENANTALLOWBLOCKLISTENTRYTYPE
    UNKNOWNFUTUREVALUE_TENANTALLOWBLOCKLISTENTRYTYPE
)

func (i TenantAllowBlockListEntryType) String() string {
    return []string{"url", "fileHash", "sender", "recipient", "unknownFutureValue"}[i]
}
func ParseTenantAllowBlockListEntryType(v string) (interface{}, error) {
    result := URL_TENANTALLOWBLOCKLISTENTRYTYPE
    switch v {
        case "url":
            result = URL_TENANTALLOWBLOCKLISTENTRYTYPE
        case "fileHash":
            result = FILEHASH_TENANTALLOWBLOCKLISTENTRYTYPE
        case "sender":
            result = SENDER_TENANTALLOWBLOCKLISTENTRYTYPE
        case "recipient":
            result = RECIPIENT_TENANTALLOWBLOCKLISTENTRYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TENANTALLOWBLOCKLISTENTRYTYPE
        default:
            return 0, errors.New("Unknown TenantAllowBlockListEntryType value: " + v)
    }
    return &result, nil
}
func SerializeTenantAllowBlockListEntryType(values []TenantAllowBlockListEntryType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
