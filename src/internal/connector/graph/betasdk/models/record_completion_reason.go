package models
import (
    "errors"
)
// Provides operations to call the add method.
type RecordCompletionReason int

const (
    OPERATIONCANCELED_RECORDCOMPLETIONREASON RecordCompletionReason = iota
    STOPTONEDETECTED_RECORDCOMPLETIONREASON
    MAXRECORDDURATIONREACHED_RECORDCOMPLETIONREASON
    INITIALSILENCETIMEOUT_RECORDCOMPLETIONREASON
    MAXSILENCETIMEOUT_RECORDCOMPLETIONREASON
    PLAYPROMPTFAILED_RECORDCOMPLETIONREASON
    PLAYBEEPFAILED_RECORDCOMPLETIONREASON
    MEDIARECEIVETIMEOUT_RECORDCOMPLETIONREASON
    UNSPECIFIEDERROR_RECORDCOMPLETIONREASON
)

func (i RecordCompletionReason) String() string {
    return []string{"operationCanceled", "stopToneDetected", "maxRecordDurationReached", "initialSilenceTimeout", "maxSilenceTimeout", "playPromptFailed", "playBeepFailed", "mediaReceiveTimeout", "unspecifiedError"}[i]
}
func ParseRecordCompletionReason(v string) (interface{}, error) {
    result := OPERATIONCANCELED_RECORDCOMPLETIONREASON
    switch v {
        case "operationCanceled":
            result = OPERATIONCANCELED_RECORDCOMPLETIONREASON
        case "stopToneDetected":
            result = STOPTONEDETECTED_RECORDCOMPLETIONREASON
        case "maxRecordDurationReached":
            result = MAXRECORDDURATIONREACHED_RECORDCOMPLETIONREASON
        case "initialSilenceTimeout":
            result = INITIALSILENCETIMEOUT_RECORDCOMPLETIONREASON
        case "maxSilenceTimeout":
            result = MAXSILENCETIMEOUT_RECORDCOMPLETIONREASON
        case "playPromptFailed":
            result = PLAYPROMPTFAILED_RECORDCOMPLETIONREASON
        case "playBeepFailed":
            result = PLAYBEEPFAILED_RECORDCOMPLETIONREASON
        case "mediaReceiveTimeout":
            result = MEDIARECEIVETIMEOUT_RECORDCOMPLETIONREASON
        case "unspecifiedError":
            result = UNSPECIFIEDERROR_RECORDCOMPLETIONREASON
        default:
            return 0, errors.New("Unknown RecordCompletionReason value: " + v)
    }
    return &result, nil
}
func SerializeRecordCompletionReason(values []RecordCompletionReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
