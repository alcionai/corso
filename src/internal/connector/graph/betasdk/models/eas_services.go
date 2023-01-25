package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EasServices int

const (
    NONE_EASSERVICES EasServices = iota
    // Enables synchronization of calendars.
    CALENDARS_EASSERVICES
    // Enables synchronization of contacts.
    CONTACTS_EASSERVICES
    // Enables synchronization of email.
    EMAIL_EASSERVICES
    // Enables synchronization of notes.
    NOTES_EASSERVICES
    // Enables synchronization of reminders.
    REMINDERS_EASSERVICES
)

func (i EasServices) String() string {
    return []string{"none", "calendars", "contacts", "email", "notes", "reminders"}[i]
}
func ParseEasServices(v string) (interface{}, error) {
    result := NONE_EASSERVICES
    switch v {
        case "none":
            result = NONE_EASSERVICES
        case "calendars":
            result = CALENDARS_EASSERVICES
        case "contacts":
            result = CONTACTS_EASSERVICES
        case "email":
            result = EMAIL_EASSERVICES
        case "notes":
            result = NOTES_EASSERVICES
        case "reminders":
            result = REMINDERS_EASSERVICES
        default:
            return 0, errors.New("Unknown EasServices value: " + v)
    }
    return &result, nil
}
func SerializeEasServices(values []EasServices) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
