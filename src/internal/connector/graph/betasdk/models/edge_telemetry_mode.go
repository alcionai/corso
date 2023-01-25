package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EdgeTelemetryMode int

const (
    // Default – No telemetry data collected or sent
    NOTCONFIGURED_EDGETELEMETRYMODE EdgeTelemetryMode = iota
    // Allow sending intranet history only: Only send browsing history data for intranet sites
    INTRANET_EDGETELEMETRYMODE
    // Allow sending internet history only: Only send browsing history data for internet sites
    INTERNET_EDGETELEMETRYMODE
    // Allow sending both intranet and internet history: Send browsing history data for intranet and internet sites
    INTRANETANDINTERNET_EDGETELEMETRYMODE
)

func (i EdgeTelemetryMode) String() string {
    return []string{"notConfigured", "intranet", "internet", "intranetAndInternet"}[i]
}
func ParseEdgeTelemetryMode(v string) (interface{}, error) {
    result := NOTCONFIGURED_EDGETELEMETRYMODE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_EDGETELEMETRYMODE
        case "intranet":
            result = INTRANET_EDGETELEMETRYMODE
        case "internet":
            result = INTERNET_EDGETELEMETRYMODE
        case "intranetAndInternet":
            result = INTRANETANDINTERNET_EDGETELEMETRYMODE
        default:
            return 0, errors.New("Unknown EdgeTelemetryMode value: " + v)
    }
    return &result, nil
}
func SerializeEdgeTelemetryMode(values []EdgeTelemetryMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
