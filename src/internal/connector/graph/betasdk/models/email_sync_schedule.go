package models
import (
    "errors"
)
// Provides operations to call the add method.
type EmailSyncSchedule int

const (
    // User Defined, default value, no intent.
    USERDEFINED_EMAILSYNCSCHEDULE EmailSyncSchedule = iota
    // Sync as messages arrive.
    ASMESSAGESARRIVE_EMAILSYNCSCHEDULE
    // Sync manually.
    MANUAL_EMAILSYNCSCHEDULE
    // Sync every fifteen minutes.
    FIFTEENMINUTES_EMAILSYNCSCHEDULE
    // Sync every thirty minutes.
    THIRTYMINUTES_EMAILSYNCSCHEDULE
    // Sync every sixty minutes.
    SIXTYMINUTES_EMAILSYNCSCHEDULE
    // Sync based on my usage.
    BASEDONMYUSAGE_EMAILSYNCSCHEDULE
)

func (i EmailSyncSchedule) String() string {
    return []string{"userDefined", "asMessagesArrive", "manual", "fifteenMinutes", "thirtyMinutes", "sixtyMinutes", "basedOnMyUsage"}[i]
}
func ParseEmailSyncSchedule(v string) (interface{}, error) {
    result := USERDEFINED_EMAILSYNCSCHEDULE
    switch v {
        case "userDefined":
            result = USERDEFINED_EMAILSYNCSCHEDULE
        case "asMessagesArrive":
            result = ASMESSAGESARRIVE_EMAILSYNCSCHEDULE
        case "manual":
            result = MANUAL_EMAILSYNCSCHEDULE
        case "fifteenMinutes":
            result = FIFTEENMINUTES_EMAILSYNCSCHEDULE
        case "thirtyMinutes":
            result = THIRTYMINUTES_EMAILSYNCSCHEDULE
        case "sixtyMinutes":
            result = SIXTYMINUTES_EMAILSYNCSCHEDULE
        case "basedOnMyUsage":
            result = BASEDONMYUSAGE_EMAILSYNCSCHEDULE
        default:
            return 0, errors.New("Unknown EmailSyncSchedule value: " + v)
    }
    return &result, nil
}
func SerializeEmailSyncSchedule(values []EmailSyncSchedule) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
