package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type MicrosoftTunnelServerHealthStatus int

const (
    // Indicates that the health status of the server is unknown. This occurs when no health status has been reported, for example when the server is initialized, but has not yet been evaluated for its health.
    UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS MicrosoftTunnelServerHealthStatus = iota
    // Indicates that the health status of the server is healthy. This should be the normal operational health status.
    HEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // Indicates that the health status of the server is unhealthy. This is normally a transient condition that will last up to 5 minutes. If the server cannot be remediated while reporting unhealthy state, the health state will change to 'warning'. If it can be remediated, the health state will return to 'healthy'.
    UNHEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // Indicates a warning based on the Tunnel Gateway server's CPU usage, memory usage, latency, TLS certificate, version
    WARNING_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // Indicates the server state is offline
    OFFLINE_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // Indicates the upgrade in progress during the upgrade cycle of when Intune begins upgrading servers, one server at a time
    UPGRADEINPROGRESS_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // Indicates the failure of the upgrade during the upgrade cycle of when Intune begins upgrading servers, one server at a time
    UPGRADEFAILED_MICROSOFTTUNNELSERVERHEALTHSTATUS
    // Evolvable enumeration sentinel value. Do not use enums.
    UNKNOWNFUTUREVALUE_MICROSOFTTUNNELSERVERHEALTHSTATUS
)

func (i MicrosoftTunnelServerHealthStatus) String() string {
    return []string{"unknown", "healthy", "unhealthy", "warning", "offline", "upgradeInProgress", "upgradeFailed", "unknownFutureValue"}[i]
}
func ParseMicrosoftTunnelServerHealthStatus(v string) (interface{}, error) {
    result := UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "healthy":
            result = HEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "unhealthy":
            result = UNHEALTHY_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "warning":
            result = WARNING_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "offline":
            result = OFFLINE_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "upgradeInProgress":
            result = UPGRADEINPROGRESS_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "upgradeFailed":
            result = UPGRADEFAILED_MICROSOFTTUNNELSERVERHEALTHSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MICROSOFTTUNNELSERVERHEALTHSTATUS
        default:
            return 0, errors.New("Unknown MicrosoftTunnelServerHealthStatus value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftTunnelServerHealthStatus(values []MicrosoftTunnelServerHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
