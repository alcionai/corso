package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type MailboxRecipientType int

const (
    UNKNOWN_MAILBOXRECIPIENTTYPE MailboxRecipientType = iota
    USER_MAILBOXRECIPIENTTYPE
    LINKED_MAILBOXRECIPIENTTYPE
    SHARED_MAILBOXRECIPIENTTYPE
    ROOM_MAILBOXRECIPIENTTYPE
    EQUIPMENT_MAILBOXRECIPIENTTYPE
    OTHERS_MAILBOXRECIPIENTTYPE
)

func (i MailboxRecipientType) String() string {
    return []string{"unknown", "user", "linked", "shared", "room", "equipment", "others"}[i]
}
func ParseMailboxRecipientType(v string) (interface{}, error) {
    result := UNKNOWN_MAILBOXRECIPIENTTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_MAILBOXRECIPIENTTYPE
        case "user":
            result = USER_MAILBOXRECIPIENTTYPE
        case "linked":
            result = LINKED_MAILBOXRECIPIENTTYPE
        case "shared":
            result = SHARED_MAILBOXRECIPIENTTYPE
        case "room":
            result = ROOM_MAILBOXRECIPIENTTYPE
        case "equipment":
            result = EQUIPMENT_MAILBOXRECIPIENTTYPE
        case "others":
            result = OTHERS_MAILBOXRECIPIENTTYPE
        default:
            return 0, errors.New("Unknown MailboxRecipientType value: " + v)
    }
    return &result, nil
}
func SerializeMailboxRecipientType(values []MailboxRecipientType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
