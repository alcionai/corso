package models
import (
    "errors"
)
// Provides operations to call the add method.
type EmailSyncDuration int

const (
    // User Defined, default value, no intent.
    USERDEFINED_EMAILSYNCDURATION EmailSyncDuration = iota
    // Sync one day of email.
    ONEDAY_EMAILSYNCDURATION
    // Sync three days of email.
    THREEDAYS_EMAILSYNCDURATION
    // Sync one week of email.
    ONEWEEK_EMAILSYNCDURATION
    // Sync two weeks of email.
    TWOWEEKS_EMAILSYNCDURATION
    // Sync one month of email.
    ONEMONTH_EMAILSYNCDURATION
    // Sync an unlimited duration of email.
    UNLIMITED_EMAILSYNCDURATION
)

func (i EmailSyncDuration) String() string {
    return []string{"userDefined", "oneDay", "threeDays", "oneWeek", "twoWeeks", "oneMonth", "unlimited"}[i]
}
func ParseEmailSyncDuration(v string) (interface{}, error) {
    result := USERDEFINED_EMAILSYNCDURATION
    switch v {
        case "userDefined":
            result = USERDEFINED_EMAILSYNCDURATION
        case "oneDay":
            result = ONEDAY_EMAILSYNCDURATION
        case "threeDays":
            result = THREEDAYS_EMAILSYNCDURATION
        case "oneWeek":
            result = ONEWEEK_EMAILSYNCDURATION
        case "twoWeeks":
            result = TWOWEEKS_EMAILSYNCDURATION
        case "oneMonth":
            result = ONEMONTH_EMAILSYNCDURATION
        case "unlimited":
            result = UNLIMITED_EMAILSYNCDURATION
        default:
            return 0, errors.New("Unknown EmailSyncDuration value: " + v)
    }
    return &result, nil
}
func SerializeEmailSyncDuration(values []EmailSyncDuration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
