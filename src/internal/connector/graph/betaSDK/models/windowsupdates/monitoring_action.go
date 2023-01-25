package windowsupdates
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MonitoringAction int

const (
    ALERTERROR_MONITORINGACTION MonitoringAction = iota
    PAUSEDEPLOYMENT_MONITORINGACTION
    UNKNOWNFUTUREVALUE_MONITORINGACTION
)

func (i MonitoringAction) String() string {
    return []string{"alertError", "pauseDeployment", "unknownFutureValue"}[i]
}
func ParseMonitoringAction(v string) (interface{}, error) {
    result := ALERTERROR_MONITORINGACTION
    switch v {
        case "alertError":
            result = ALERTERROR_MONITORINGACTION
        case "pauseDeployment":
            result = PAUSEDEPLOYMENT_MONITORINGACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MONITORINGACTION
        default:
            return 0, errors.New("Unknown MonitoringAction value: " + v)
    }
    return &result, nil
}
func SerializeMonitoringAction(values []MonitoringAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
