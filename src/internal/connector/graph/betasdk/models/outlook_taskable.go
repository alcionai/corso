package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTaskable 
type OutlookTaskable interface {
    OutlookItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTo()(*string)
    GetAttachments()([]Attachmentable)
    GetBody()(ItemBodyable)
    GetCompletedDateTime()(DateTimeTimeZoneable)
    GetDueDateTime()(DateTimeTimeZoneable)
    GetHasAttachments()(*bool)
    GetImportance()(*Importance)
    GetIsReminderOn()(*bool)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetOwner()(*string)
    GetParentFolderId()(*string)
    GetRecurrence()(PatternedRecurrenceable)
    GetReminderDateTime()(DateTimeTimeZoneable)
    GetSensitivity()(*Sensitivity)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*TaskStatus)
    GetSubject()(*string)
    SetAssignedTo(value *string)()
    SetAttachments(value []Attachmentable)()
    SetBody(value ItemBodyable)()
    SetCompletedDateTime(value DateTimeTimeZoneable)()
    SetDueDateTime(value DateTimeTimeZoneable)()
    SetHasAttachments(value *bool)()
    SetImportance(value *Importance)()
    SetIsReminderOn(value *bool)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetOwner(value *string)()
    SetParentFolderId(value *string)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetReminderDateTime(value DateTimeTimeZoneable)()
    SetSensitivity(value *Sensitivity)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *TaskStatus)()
    SetSubject(value *string)()
}
