package security
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EventStatusType int

const (
    PENDING_EVENTSTATUSTYPE EventStatusType = iota
    ERROR_EVENTSTATUSTYPE
    SUCCESS_EVENTSTATUSTYPE
    NOTAVALIABLE_EVENTSTATUSTYPE
    UNKNOWNFUTUREVALUE_EVENTSTATUSTYPE
)

func (i EventStatusType) String() string {
    return []string{"pending", "error", "success", "notAvaliable", "unknownFutureValue"}[i]
}
func ParseEventStatusType(v string) (interface{}, error) {
    result := PENDING_EVENTSTATUSTYPE
    switch v {
        case "pending":
            result = PENDING_EVENTSTATUSTYPE
        case "error":
            result = ERROR_EVENTSTATUSTYPE
        case "success":
            result = SUCCESS_EVENTSTATUSTYPE
        case "notAvaliable":
            result = NOTAVALIABLE_EVENTSTATUSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EVENTSTATUSTYPE
        default:
            return 0, errors.New("Unknown EventStatusType value: " + v)
    }
    return &result, nil
}
func SerializeEventStatusType(values []EventStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
