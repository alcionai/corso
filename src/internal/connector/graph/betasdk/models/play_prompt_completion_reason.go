package models
import (
    "errors"
)
// Provides operations to call the add method.
type PlayPromptCompletionReason int

const (
    UNKNOWN_PLAYPROMPTCOMPLETIONREASON PlayPromptCompletionReason = iota
    COMPLETEDSUCCESSFULLY_PLAYPROMPTCOMPLETIONREASON
    MEDIAOPERATIONCANCELED_PLAYPROMPTCOMPLETIONREASON
    UNKNOWNFUTUREVALUE_PLAYPROMPTCOMPLETIONREASON
)

func (i PlayPromptCompletionReason) String() string {
    return []string{"unknown", "completedSuccessfully", "mediaOperationCanceled", "unknownFutureValue"}[i]
}
func ParsePlayPromptCompletionReason(v string) (interface{}, error) {
    result := UNKNOWN_PLAYPROMPTCOMPLETIONREASON
    switch v {
        case "unknown":
            result = UNKNOWN_PLAYPROMPTCOMPLETIONREASON
        case "completedSuccessfully":
            result = COMPLETEDSUCCESSFULLY_PLAYPROMPTCOMPLETIONREASON
        case "mediaOperationCanceled":
            result = MEDIAOPERATIONCANCELED_PLAYPROMPTCOMPLETIONREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLAYPROMPTCOMPLETIONREASON
        default:
            return 0, errors.New("Unknown PlayPromptCompletionReason value: " + v)
    }
    return &result, nil
}
func SerializePlayPromptCompletionReason(values []PlayPromptCompletionReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
