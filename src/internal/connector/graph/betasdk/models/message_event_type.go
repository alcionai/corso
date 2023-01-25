package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MessageEventType int

const (
    RECEIVED_MESSAGEEVENTTYPE MessageEventType = iota
    SENT_MESSAGEEVENTTYPE
    DELIVERED_MESSAGEEVENTTYPE
    FAILED_MESSAGEEVENTTYPE
    PROCESSINGFAILED_MESSAGEEVENTTYPE
    DISTRIBUTIONGROUPEXPANDED_MESSAGEEVENTTYPE
    SUBMITTED_MESSAGEEVENTTYPE
    DELAYED_MESSAGEEVENTTYPE
    REDIRECTED_MESSAGEEVENTTYPE
    RESOLVED_MESSAGEEVENTTYPE
    DROPPED_MESSAGEEVENTTYPE
    RECIPIENTSADDED_MESSAGEEVENTTYPE
    MALWAREDETECTED_MESSAGEEVENTTYPE
    MALWAREDETECTEDINMESSAGE_MESSAGEEVENTTYPE
    MALWAREDETECTEDINATTACHMENT_MESSAGEEVENTTYPE
    TTZAPPED_MESSAGEEVENTTYPE
    TTDELIVERED_MESSAGEEVENTTYPE
    SPAMDETECTED_MESSAGEEVENTTYPE
    TRANSPORTRULETRIGGERED_MESSAGEEVENTTYPE
    DLPRULETRIGGERED_MESSAGEEVENTTYPE
    JOURNALED_MESSAGEEVENTTYPE
    UNKNOWNFUTUREVALUE_MESSAGEEVENTTYPE
)

func (i MessageEventType) String() string {
    return []string{"received", "sent", "delivered", "failed", "processingFailed", "distributionGroupExpanded", "submitted", "delayed", "redirected", "resolved", "dropped", "recipientsAdded", "malwareDetected", "malwareDetectedInMessage", "malwareDetectedInAttachment", "ttZapped", "ttDelivered", "spamDetected", "transportRuleTriggered", "dlpRuleTriggered", "journaled", "unknownFutureValue"}[i]
}
func ParseMessageEventType(v string) (interface{}, error) {
    result := RECEIVED_MESSAGEEVENTTYPE
    switch v {
        case "received":
            result = RECEIVED_MESSAGEEVENTTYPE
        case "sent":
            result = SENT_MESSAGEEVENTTYPE
        case "delivered":
            result = DELIVERED_MESSAGEEVENTTYPE
        case "failed":
            result = FAILED_MESSAGEEVENTTYPE
        case "processingFailed":
            result = PROCESSINGFAILED_MESSAGEEVENTTYPE
        case "distributionGroupExpanded":
            result = DISTRIBUTIONGROUPEXPANDED_MESSAGEEVENTTYPE
        case "submitted":
            result = SUBMITTED_MESSAGEEVENTTYPE
        case "delayed":
            result = DELAYED_MESSAGEEVENTTYPE
        case "redirected":
            result = REDIRECTED_MESSAGEEVENTTYPE
        case "resolved":
            result = RESOLVED_MESSAGEEVENTTYPE
        case "dropped":
            result = DROPPED_MESSAGEEVENTTYPE
        case "recipientsAdded":
            result = RECIPIENTSADDED_MESSAGEEVENTTYPE
        case "malwareDetected":
            result = MALWAREDETECTED_MESSAGEEVENTTYPE
        case "malwareDetectedInMessage":
            result = MALWAREDETECTEDINMESSAGE_MESSAGEEVENTTYPE
        case "malwareDetectedInAttachment":
            result = MALWAREDETECTEDINATTACHMENT_MESSAGEEVENTTYPE
        case "ttZapped":
            result = TTZAPPED_MESSAGEEVENTTYPE
        case "ttDelivered":
            result = TTDELIVERED_MESSAGEEVENTTYPE
        case "spamDetected":
            result = SPAMDETECTED_MESSAGEEVENTTYPE
        case "transportRuleTriggered":
            result = TRANSPORTRULETRIGGERED_MESSAGEEVENTTYPE
        case "dlpRuleTriggered":
            result = DLPRULETRIGGERED_MESSAGEEVENTTYPE
        case "journaled":
            result = JOURNALED_MESSAGEEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MESSAGEEVENTTYPE
        default:
            return 0, errors.New("Unknown MessageEventType value: " + v)
    }
    return &result, nil
}
func SerializeMessageEventType(values []MessageEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
