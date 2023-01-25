package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MessageStatus int

const (
    GETTINGSTATUS_MESSAGESTATUS MessageStatus = iota
    PENDING_MESSAGESTATUS
    FAILED_MESSAGESTATUS
    DELIVERED_MESSAGESTATUS
    EXPANDED_MESSAGESTATUS
    QUARANTINED_MESSAGESTATUS
    FILTEREDASSPAM_MESSAGESTATUS
    UNKNOWNFUTUREVALUE_MESSAGESTATUS
)

func (i MessageStatus) String() string {
    return []string{"gettingStatus", "pending", "failed", "delivered", "expanded", "quarantined", "filteredAsSpam", "unknownFutureValue"}[i]
}
func ParseMessageStatus(v string) (interface{}, error) {
    result := GETTINGSTATUS_MESSAGESTATUS
    switch v {
        case "gettingStatus":
            result = GETTINGSTATUS_MESSAGESTATUS
        case "pending":
            result = PENDING_MESSAGESTATUS
        case "failed":
            result = FAILED_MESSAGESTATUS
        case "delivered":
            result = DELIVERED_MESSAGESTATUS
        case "expanded":
            result = EXPANDED_MESSAGESTATUS
        case "quarantined":
            result = QUARANTINED_MESSAGESTATUS
        case "filteredAsSpam":
            result = FILTEREDASSPAM_MESSAGESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MESSAGESTATUS
        default:
            return 0, errors.New("Unknown MessageStatus value: " + v)
    }
    return &result, nil
}
func SerializeMessageStatus(values []MessageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
