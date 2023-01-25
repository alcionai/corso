package devicemanagement
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AlertStatusType int

const (
    ACTIVE_ALERTSTATUSTYPE AlertStatusType = iota
    RESOLVED_ALERTSTATUSTYPE
    UNKNOWNFUTUREVALUE_ALERTSTATUSTYPE
)

func (i AlertStatusType) String() string {
    return []string{"active", "resolved", "unknownFutureValue"}[i]
}
func ParseAlertStatusType(v string) (interface{}, error) {
    result := ACTIVE_ALERTSTATUSTYPE
    switch v {
        case "active":
            result = ACTIVE_ALERTSTATUSTYPE
        case "resolved":
            result = RESOLVED_ALERTSTATUSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTSTATUSTYPE
        default:
            return 0, errors.New("Unknown AlertStatusType value: " + v)
    }
    return &result, nil
}
func SerializeAlertStatusType(values []AlertStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
