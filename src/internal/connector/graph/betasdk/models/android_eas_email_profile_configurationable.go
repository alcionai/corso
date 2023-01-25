package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidEasEmailProfileConfigurationable 
type AndroidEasEmailProfileConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetAuthenticationMethod()(*EasAuthenticationMethod)
    GetCustomDomainName()(*string)
    GetDurationOfEmailToSync()(*EmailSyncDuration)
    GetEmailAddressSource()(*UserEmailSource)
    GetEmailSyncSchedule()(*EmailSyncSchedule)
    GetHostName()(*string)
    GetIdentityCertificate()(AndroidCertificateProfileBaseable)
    GetRequireSmime()(*bool)
    GetRequireSsl()(*bool)
    GetSmimeSigningCertificate()(AndroidCertificateProfileBaseable)
    GetSyncCalendar()(*bool)
    GetSyncContacts()(*bool)
    GetSyncNotes()(*bool)
    GetSyncTasks()(*bool)
    GetUserDomainNameSource()(*DomainNameSource)
    GetUsernameSource()(*AndroidUsernameSource)
    SetAccountName(value *string)()
    SetAuthenticationMethod(value *EasAuthenticationMethod)()
    SetCustomDomainName(value *string)()
    SetDurationOfEmailToSync(value *EmailSyncDuration)()
    SetEmailAddressSource(value *UserEmailSource)()
    SetEmailSyncSchedule(value *EmailSyncSchedule)()
    SetHostName(value *string)()
    SetIdentityCertificate(value AndroidCertificateProfileBaseable)()
    SetRequireSmime(value *bool)()
    SetRequireSsl(value *bool)()
    SetSmimeSigningCertificate(value AndroidCertificateProfileBaseable)()
    SetSyncCalendar(value *bool)()
    SetSyncContacts(value *bool)()
    SetSyncNotes(value *bool)()
    SetSyncTasks(value *bool)()
    SetUserDomainNameSource(value *DomainNameSource)()
    SetUsernameSource(value *AndroidUsernameSource)()
}
