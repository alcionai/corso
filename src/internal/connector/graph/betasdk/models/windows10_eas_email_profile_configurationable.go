package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EasEmailProfileConfigurationable 
type Windows10EasEmailProfileConfigurationable interface {
    EasEmailProfileConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEmailAddressSource()(*UserEmailSource)
    GetEmailSyncSchedule()(*EmailSyncSchedule)
    GetHostName()(*string)
    GetRequireSsl()(*bool)
    GetSyncCalendar()(*bool)
    GetSyncContacts()(*bool)
    GetSyncTasks()(*bool)
    SetAccountName(value *string)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetEmailSyncSchedule(value *EmailSyncSchedule)()
    SetHostName(value *string)()
    SetRequireSsl(value *bool)()
    SetSyncCalendar(value *bool)()
    SetSyncContacts(value *bool)()
    SetSyncTasks(value *bool)()
}
