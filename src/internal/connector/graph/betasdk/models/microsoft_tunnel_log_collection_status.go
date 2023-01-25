package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MicrosoftTunnelLogCollectionStatus int

const (
    // Indicates that the log collection is in progress
    PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS MicrosoftTunnelLogCollectionStatus = iota
    // Indicates that the log collection is completed
    COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    // Indicates that the log collection has failed
    FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    // Placeholder value for future expansion enums
    UNKNOWNFUTUREVALUE_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
)

func (i MicrosoftTunnelLogCollectionStatus) String() string {
    return []string{"pending", "completed", "failed", "unknownFutureValue"}[i]
}
func ParseMicrosoftTunnelLogCollectionStatus(v string) (interface{}, error) {
    result := PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
    switch v {
        case "pending":
            result = PENDING_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "completed":
            result = COMPLETED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "failed":
            result = FAILED_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MICROSOFTTUNNELLOGCOLLECTIONSTATUS
        default:
            return 0, errors.New("Unknown MicrosoftTunnelLogCollectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftTunnelLogCollectionStatus(values []MicrosoftTunnelLogCollectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
