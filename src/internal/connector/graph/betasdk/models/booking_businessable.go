package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingBusinessable 
type BookingBusinessable interface {
    BookingNamedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PhysicalAddressable)
    GetAppointments()([]BookingAppointmentable)
    GetBusinessHours()([]BookingWorkHoursable)
    GetBusinessType()(*string)
    GetCalendarView()([]BookingAppointmentable)
    GetCustomers()([]BookingCustomerable)
    GetCustomQuestions()([]BookingCustomQuestionable)
    GetDefaultCurrencyIso()(*string)
    GetEmail()(*string)
    GetIsPublished()(*bool)
    GetLanguageTag()(*string)
    GetPhone()(*string)
    GetPublicUrl()(*string)
    GetSchedulingPolicy()(BookingSchedulingPolicyable)
    GetServices()([]BookingServiceable)
    GetStaffMembers()([]BookingStaffMemberable)
    GetWebSiteUrl()(*string)
    SetAddress(value PhysicalAddressable)()
    SetAppointments(value []BookingAppointmentable)()
    SetBusinessHours(value []BookingWorkHoursable)()
    SetBusinessType(value *string)()
    SetCalendarView(value []BookingAppointmentable)()
    SetCustomers(value []BookingCustomerable)()
    SetCustomQuestions(value []BookingCustomQuestionable)()
    SetDefaultCurrencyIso(value *string)()
    SetEmail(value *string)()
    SetIsPublished(value *bool)()
    SetLanguageTag(value *string)()
    SetPhone(value *string)()
    SetPublicUrl(value *string)()
    SetSchedulingPolicy(value BookingSchedulingPolicyable)()
    SetServices(value []BookingServiceable)()
    SetStaffMembers(value []BookingStaffMemberable)()
    SetWebSiteUrl(value *string)()
}
