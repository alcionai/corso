package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidEasEmailProfileConfiguration 
type AndroidEasEmailProfileConfiguration struct {
    DeviceConfiguration
    // Exchange ActiveSync account name, displayed to users as name of EAS (this) profile.
    accountName *string
    // Exchange Active Sync authentication method.
    authenticationMethod *EasAuthenticationMethod
    // Custom domain name value used while generating an email profile before installing on the device.
    customDomainName *string
    // Possible values for email sync duration.
    durationOfEmailToSync *EmailSyncDuration
    // Possible values for username source or email source.
    emailAddressSource *UserEmailSource
    // Possible values for email sync schedule.
    emailSyncSchedule *EmailSyncSchedule
    // Exchange location (URL) that the native mail app connects to.
    hostName *string
    // Identity certificate.
    identityCertificate AndroidCertificateProfileBaseable
    // Indicates whether or not to use S/MIME certificate.
    requireSmime *bool
    // Indicates whether or not to use SSL.
    requireSsl *bool
    // S/MIME signing certificate.
    smimeSigningCertificate AndroidCertificateProfileBaseable
    // Toggles syncing the calendar. If set to false calendar is turned off on the device.
    syncCalendar *bool
    // Toggles syncing contacts. If set to false contacts are turned off on the device.
    syncContacts *bool
    // Toggles syncing notes. If set to false notes are turned off on the device.
    syncNotes *bool
    // Toggles syncing tasks. If set to false tasks are turned off on the device.
    syncTasks *bool
    // UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
    userDomainNameSource *DomainNameSource
    // Android username source.
    usernameSource *AndroidUsernameSource
}
// NewAndroidEasEmailProfileConfiguration instantiates a new AndroidEasEmailProfileConfiguration and sets the default values.
func NewAndroidEasEmailProfileConfiguration()(*AndroidEasEmailProfileConfiguration) {
    m := &AndroidEasEmailProfileConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidEasEmailProfileConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidEasEmailProfileConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidEasEmailProfileConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidEasEmailProfileConfiguration(), nil
}
// GetAccountName gets the accountName property value. Exchange ActiveSync account name, displayed to users as name of EAS (this) profile.
func (m *AndroidEasEmailProfileConfiguration) GetAccountName()(*string) {
    return m.accountName
}
// GetAuthenticationMethod gets the authenticationMethod property value. Exchange Active Sync authentication method.
func (m *AndroidEasEmailProfileConfiguration) GetAuthenticationMethod()(*EasAuthenticationMethod) {
    return m.authenticationMethod
}
// GetCustomDomainName gets the customDomainName property value. Custom domain name value used while generating an email profile before installing on the device.
func (m *AndroidEasEmailProfileConfiguration) GetCustomDomainName()(*string) {
    return m.customDomainName
}
// GetDurationOfEmailToSync gets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidEasEmailProfileConfiguration) GetDurationOfEmailToSync()(*EmailSyncDuration) {
    return m.durationOfEmailToSync
}
// GetEmailAddressSource gets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidEasEmailProfileConfiguration) GetEmailAddressSource()(*UserEmailSource) {
    return m.emailAddressSource
}
// GetEmailSyncSchedule gets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *AndroidEasEmailProfileConfiguration) GetEmailSyncSchedule()(*EmailSyncSchedule) {
    return m.emailSyncSchedule
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidEasEmailProfileConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
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
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEasAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*EasAuthenticationMethod))
        }
        return nil
    }
    res["customDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDomainName(val)
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
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidCertificateProfileBaseable))
        }
        return nil
    }
    res["requireSmime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireSmime(val)
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
    res["smimeSigningCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmimeSigningCertificate(val.(AndroidCertificateProfileBaseable))
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
    res["syncNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncNotes(val)
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
    res["userDomainNameSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDomainNameSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDomainNameSource(val.(*DomainNameSource))
        }
        return nil
    }
    res["usernameSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidUsernameSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameSource(val.(*AndroidUsernameSource))
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. Exchange location (URL) that the native mail app connects to.
func (m *AndroidEasEmailProfileConfiguration) GetHostName()(*string) {
    return m.hostName
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate.
func (m *AndroidEasEmailProfileConfiguration) GetIdentityCertificate()(AndroidCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetRequireSmime gets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *AndroidEasEmailProfileConfiguration) GetRequireSmime()(*bool) {
    return m.requireSmime
}
// GetRequireSsl gets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidEasEmailProfileConfiguration) GetRequireSsl()(*bool) {
    return m.requireSsl
}
// GetSmimeSigningCertificate gets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *AndroidEasEmailProfileConfiguration) GetSmimeSigningCertificate()(AndroidCertificateProfileBaseable) {
    return m.smimeSigningCertificate
}
// GetSyncCalendar gets the syncCalendar property value. Toggles syncing the calendar. If set to false calendar is turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncCalendar()(*bool) {
    return m.syncCalendar
}
// GetSyncContacts gets the syncContacts property value. Toggles syncing contacts. If set to false contacts are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncContacts()(*bool) {
    return m.syncContacts
}
// GetSyncNotes gets the syncNotes property value. Toggles syncing notes. If set to false notes are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncNotes()(*bool) {
    return m.syncNotes
}
// GetSyncTasks gets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) GetSyncTasks()(*bool) {
    return m.syncTasks
}
// GetUserDomainNameSource gets the userDomainNameSource property value. UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
func (m *AndroidEasEmailProfileConfiguration) GetUserDomainNameSource()(*DomainNameSource) {
    return m.userDomainNameSource
}
// GetUsernameSource gets the usernameSource property value. Android username source.
func (m *AndroidEasEmailProfileConfiguration) GetUsernameSource()(*AndroidUsernameSource) {
    return m.usernameSource
}
// Serialize serializes information the current object
func (m *AndroidEasEmailProfileConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDomainName", m.GetCustomDomainName())
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
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireSmime", m.GetRequireSmime())
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
        err = writer.WriteObjectValue("smimeSigningCertificate", m.GetSmimeSigningCertificate())
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
        err = writer.WriteBoolValue("syncNotes", m.GetSyncNotes())
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
    if m.GetUserDomainNameSource() != nil {
        cast := (*m.GetUserDomainNameSource()).String()
        err = writer.WriteStringValue("userDomainNameSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsernameSource() != nil {
        cast := (*m.GetUsernameSource()).String()
        err = writer.WriteStringValue("usernameSource", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountName sets the accountName property value. Exchange ActiveSync account name, displayed to users as name of EAS (this) profile.
func (m *AndroidEasEmailProfileConfiguration) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. Exchange Active Sync authentication method.
func (m *AndroidEasEmailProfileConfiguration) SetAuthenticationMethod(value *EasAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetCustomDomainName sets the customDomainName property value. Custom domain name value used while generating an email profile before installing on the device.
func (m *AndroidEasEmailProfileConfiguration) SetCustomDomainName(value *string)() {
    m.customDomainName = value
}
// SetDurationOfEmailToSync sets the durationOfEmailToSync property value. Possible values for email sync duration.
func (m *AndroidEasEmailProfileConfiguration) SetDurationOfEmailToSync(value *EmailSyncDuration)() {
    m.durationOfEmailToSync = value
}
// SetEmailAddressSource sets the emailAddressSource property value. Possible values for username source or email source.
func (m *AndroidEasEmailProfileConfiguration) SetEmailAddressSource(value *UserEmailSource)() {
    m.emailAddressSource = value
}
// SetEmailSyncSchedule sets the emailSyncSchedule property value. Possible values for email sync schedule.
func (m *AndroidEasEmailProfileConfiguration) SetEmailSyncSchedule(value *EmailSyncSchedule)() {
    m.emailSyncSchedule = value
}
// SetHostName sets the hostName property value. Exchange location (URL) that the native mail app connects to.
func (m *AndroidEasEmailProfileConfiguration) SetHostName(value *string)() {
    m.hostName = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate.
func (m *AndroidEasEmailProfileConfiguration) SetIdentityCertificate(value AndroidCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetRequireSmime sets the requireSmime property value. Indicates whether or not to use S/MIME certificate.
func (m *AndroidEasEmailProfileConfiguration) SetRequireSmime(value *bool)() {
    m.requireSmime = value
}
// SetRequireSsl sets the requireSsl property value. Indicates whether or not to use SSL.
func (m *AndroidEasEmailProfileConfiguration) SetRequireSsl(value *bool)() {
    m.requireSsl = value
}
// SetSmimeSigningCertificate sets the smimeSigningCertificate property value. S/MIME signing certificate.
func (m *AndroidEasEmailProfileConfiguration) SetSmimeSigningCertificate(value AndroidCertificateProfileBaseable)() {
    m.smimeSigningCertificate = value
}
// SetSyncCalendar sets the syncCalendar property value. Toggles syncing the calendar. If set to false calendar is turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncCalendar(value *bool)() {
    m.syncCalendar = value
}
// SetSyncContacts sets the syncContacts property value. Toggles syncing contacts. If set to false contacts are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncContacts(value *bool)() {
    m.syncContacts = value
}
// SetSyncNotes sets the syncNotes property value. Toggles syncing notes. If set to false notes are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncNotes(value *bool)() {
    m.syncNotes = value
}
// SetSyncTasks sets the syncTasks property value. Toggles syncing tasks. If set to false tasks are turned off on the device.
func (m *AndroidEasEmailProfileConfiguration) SetSyncTasks(value *bool)() {
    m.syncTasks = value
}
// SetUserDomainNameSource sets the userDomainNameSource property value. UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
func (m *AndroidEasEmailProfileConfiguration) SetUserDomainNameSource(value *DomainNameSource)() {
    m.userDomainNameSource = value
}
// SetUsernameSource sets the usernameSource property value. Android username source.
func (m *AndroidEasEmailProfileConfiguration) SetUsernameSource(value *AndroidUsernameSource)() {
    m.usernameSource = value
}
