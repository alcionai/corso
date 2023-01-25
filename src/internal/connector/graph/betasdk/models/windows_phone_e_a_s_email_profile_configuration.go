package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhoneEASEmailProfileConfiguration 
type WindowsPhoneEASEmailProfileConfiguration struct {
    EasEmailProfileConfigurationBase
    // Account name.
    accountName *string
    // Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
    applyOnlyToWindowsPhone81 *bool
    // Possible values for email sync duration.
    durationOfEmailToSync *EmailSyncDuration
    // Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
    emailAddressSource *UserEmailSource
    // Possible values for email sync schedule.
    emailSyncSchedule *EmailSyncSchedule
    // Exchange location that (URL) that the native mail app connects to.
    hostName *string
    // Indicates whether or not to use SSL.
    requireSsl *bool
    // Whether or not to sync the calendar.
    syncCalendar *bool
    // Whether or not to sync contacts.
    syncContacts *bool
    // Whether or not to sync tasks.
    syncTasks *bool
}
// NewWindowsPhoneEASEmailProfileConfiguration instantiates a new WindowsPhoneEASEmailProfileConfiguration and sets the default values.
func NewWindowsPhoneEASEmailProfileConfiguration()(*WindowsPhoneEASEmailProfileConfiguration) {
    m := &WindowsPhoneEASEmailProfileConfiguration{
        EasEmailProfileConfigurationBase: *NewEasEmailProfileConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhoneEASEmailProfileConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsPhoneEASEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhoneEASEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhoneEASEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Account name.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetAccountName()(*string) {
    return m.accountName
}
// GetApplyOnlyToWindowsPhone81 gets the applyOnlyToWindowsPhone81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetApplyOnlyToWindowsPhone81()(*bool) {
    return m.applyOnlyToWindowsPhone81
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    return m.durationOfEmailToSync
}
// GetEmailAddressSource gets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
    return m.emailAddressSource
}
// GetEmailSyncSchedule gets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetEmailSyncSchedule()(*EmailSyncSchedule) {
    return m.emailSyncSchedule
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhoneEASEmailProfileConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EasEmailProfileConfigurationBase.GetFieldDeserializers()
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["applyOnlyToWindowsPhone81"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyOnlyToWindowsPhone81(val)
        }
        return nil
    }
    res["durationOfEmailToSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailSyncDuration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationOfEmailToSync(val.(*EmailSyncDuration))
        }
        return nil
    }
    res["emailAddressSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserEmailSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddressSource(val.(*UserEmailSource))
        }
        return nil
    }
    res["emailSyncSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailSyncSchedule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSyncSchedule(val.(*EmailSyncSchedule))
        }
        return nil
    }
    res["hostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    res["requireSsl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSsl(val)
        }
        return nil
    }
    res["syncCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncCalendar(val)
        }
        return nil
    }
    res["syncContacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncContacts(val)
        }
        return nil
    }
    res["syncTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncTasks(val)
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetHostName()(*string) {
    return m.hostName
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetRequireSsl()(*bool) {
    return m.requireSsl
}
// GetSyncCalendar gets the syncCalendar property value. Whether or not to sync the calendar.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetSyncCalendar()(*bool) {
    return m.syncCalendar
}
// GetSyncContacts gets the syncContacts property value. Whether or not to sync contacts.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetSyncContacts()(*bool) {
    return m.syncContacts
}
// GetSyncTasks gets the syncTasks property value. Whether or not to sync tasks.
func (m *WindowsPhoneEASEmailProfileConfiguration) GetSyncTasks()(*bool) {
    return m.syncTasks
}
// Serialize serializes information the current object
func (m *WindowsPhoneEASEmailProfileConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EasEmailProfileConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    if m.GetDurationOfEmailToSync() != nil {
        cast := (*m.GetDurationOfEmailToSync()).String()
        err = writer.WriteStringValue("durationOfEmailToSync", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailAddressSource() != nil {
        cast := (*m.GetEmailAddressSource()).String()
        err = writer.WriteStringValue("emailAddressSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailSyncSchedule() != nil {
        cast := (*m.GetEmailSyncSchedule()).String()
        err = writer.WriteStringValue("emailSyncSchedule", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hostName", m.GetHostName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireSsl", m.GetRequireSsl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncCalendar", m.GetSyncCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncContacts", m.GetSyncContacts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("syncTasks", m.GetSyncTasks())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountName sets the accountName property value. Account name.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetAccountName(value *string)() {
    m.accountName = value
}
// SetApplyOnlyToWindowsPhone81 sets the applyOnlyToWindowsPhone81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetApplyOnlyToWindowsPhone81(value *bool)() {
    m.applyOnlyToWindowsPhone81 = value
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    m.durationOfEmailToSync = value
}
// SetEmailAddressSource sets the emailAddressSource property value. Email attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: userPrincipalName, primarySmtpAddress.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    m.emailAddressSource = value
}
// SetEmailSyncSchedule sets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetEmailSyncSchedule(value *EmailSyncSchedule)() {
    m.emailSyncSchedule = value
}
// SetHostName sets the hostName property value. Exchange location that (URL) that the native mail app connects to.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetHostName(value *string)() {
    m.hostName = value
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    m.requireSsl = value
}
// SetSyncCalendar sets the syncCalendar property value. Whether or not to sync the calendar.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetSyncCalendar(value *bool)() {
    m.syncCalendar = value
}
// SetSyncContacts sets the syncContacts property value. Whether or not to sync contacts.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetSyncContacts(value *bool)() {
    m.syncContacts = value
}
// SetSyncTasks sets the syncTasks property value. Whether or not to sync tasks.
func (m *WindowsPhoneEASEmailProfileConfiguration) SetSyncTasks(value *bool)() {
    m.syncTasks = value
}
