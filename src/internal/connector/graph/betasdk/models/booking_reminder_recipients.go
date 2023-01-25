package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type BookingReminderRecipients int

const (
    ALLATTENDEES_BOOKINGREMINDERRECIPIENTS BookingReminderRecipients = iota
    STAFF_BOOKINGREMINDERRECIPIENTS
    CUSTOMER_BOOKINGREMINDERRECIPIENTS
)

func (i BookingReminderRecipients) String() string {
    return []string{"allAttendees", "staff", "customer"}[i]
}
func ParseBookingReminderRecipients(v string) (interface{}, error) {
    result := ALLATTENDEES_BOOKINGREMINDERRECIPIENTS
    switch v {
        case "allAttendees":
            result = ALLATTENDEES_BOOKINGREMINDERRECIPIENTS
        case "staff":
            result = STAFF_BOOKINGREMINDERRECIPIENTS
        case "customer":
            result = CUSTOMER_BOOKINGREMINDERRECIPIENTS
        default:
            return 0, errors.New("Unknown BookingReminderRecipients value: " + v)
    }
    return &result, nil
}
func SerializeBookingReminderRecipients(values []BookingReminderRecipients) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
