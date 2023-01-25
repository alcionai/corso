package ediscovery
import (
    "errors"
)
// Provides operations to call the add method.
type DataSourceContainerStatus int

const (
    ACTIVE_DATASOURCECONTAINERSTATUS DataSourceContainerStatus = iota
    RELEASED_DATASOURCECONTAINERSTATUS
    UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
)

func (i DataSourceContainerStatus) String() string {
    return []string{"Active", "Released", "UnknownFutureValue"}[i]
}
func ParseDataSourceContainerStatus(v string) (interface{}, error) {
    result := ACTIVE_DATASOURCECONTAINERSTATUS
    switch v {
        case "Active":
            result = ACTIVE_DATASOURCECONTAINERSTATUS
        case "Released":
            result = RELEASED_DATASOURCECONTAINERSTATUS
        case "UnknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
        default:
            return 0, errors.New("Unknown DataSourceContainerStatus value: " + v)
    }
    return &result, nil
}
func SerializeDataSourceContainerStatus(values []DataSourceContainerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
