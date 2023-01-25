package windowsupdates
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MonitoringSignal int

const (
    ROLLBACK_MONITORINGSIGNAL MonitoringSignal = iota
    UNKNOWNFUTUREVALUE_MONITORINGSIGNAL
)

func (i MonitoringSignal) String() string {
    return []string{"rollback", "unknownFutureValue"}[i]
}
func ParseMonitoringSignal(v string) (interface{}, error) {
    result := ROLLBACK_MONITORINGSIGNAL
    switch v {
        case "rollback":
            result = ROLLBACK_MONITORINGSIGNAL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MONITORINGSIGNAL
        default:
            return 0, errors.New("Unknown MonitoringSignal value: " + v)
    }
    return &result, nil
}
func SerializeMonitoringSignal(values []MonitoringSignal) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
