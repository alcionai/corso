package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TeamworkConnectionStatus int

const (
    UNKNOWN_TEAMWORKCONNECTIONSTATUS TeamworkConnectionStatus = iota
    CONNECTED_TEAMWORKCONNECTIONSTATUS
    DISCONNECTED_TEAMWORKCONNECTIONSTATUS
    UNKNOWNFUTUREVALUE_TEAMWORKCONNECTIONSTATUS
)

func (i TeamworkConnectionStatus) String() string {
    return []string{"unknown", "connected", "disconnected", "unknownFutureValue"}[i]
}
func ParseTeamworkConnectionStatus(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKCONNECTIONSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_TEAMWORKCONNECTIONSTATUS
        case "connected":
            result = CONNECTED_TEAMWORKCONNECTIONSTATUS
        case "disconnected":
            result = DISCONNECTED_TEAMWORKCONNECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMWORKCONNECTIONSTATUS
        default:
            return 0, errors.New("Unknown TeamworkConnectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkConnectionStatus(values []TeamworkConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
