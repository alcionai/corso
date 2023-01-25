package managedtenants
import (
    "errors"
)
// Provides operations to call the add method.
type ManagementProvider int

const (
    MICROSOFT_MANAGEMENTPROVIDER ManagementProvider = iota
    COMMUNITY_MANAGEMENTPROVIDER
    INDIRECTPROVIDER_MANAGEMENTPROVIDER
    SELF_MANAGEMENTPROVIDER
    UNKNOWNFUTUREVALUE_MANAGEMENTPROVIDER
)

func (i ManagementProvider) String() string {
    return []string{"microsoft", "community", "indirectProvider", "self", "unknownFutureValue"}[i]
}
func ParseManagementProvider(v string) (interface{}, error) {
    result := MICROSOFT_MANAGEMENTPROVIDER
    switch v {
        case "microsoft":
            result = MICROSOFT_MANAGEMENTPROVIDER
        case "community":
            result = COMMUNITY_MANAGEMENTPROVIDER
        case "indirectProvider":
            result = INDIRECTPROVIDER_MANAGEMENTPROVIDER
        case "self":
            result = SELF_MANAGEMENTPROVIDER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTPROVIDER
        default:
            return 0, errors.New("Unknown ManagementProvider value: " + v)
    }
    return &result, nil
}
func SerializeManagementProvider(values []ManagementProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
