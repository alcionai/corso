package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepOnboardingSetting the depOnboardingSetting represents an instance of the Apple DEP service being onboarded to Intune. The onboarded service instance manages an Apple Token used to synchronize data between Apple and Intune.
type DepOnboardingSetting struct {
    Entity
    // The Apple ID used to obtain the current token.
    appleIdentifier *string
    // Consent granted for data sharing with Apple Dep Service
    dataSharingConsentGranted *bool
    // Default iOS Enrollment Profile
    defaultIosEnrollmentProfile DepIOSEnrollmentProfileable
    // Default MacOs Enrollment Profile
    defaultMacOsEnrollmentProfile DepMacOSEnrollmentProfileable
    // The enrollment profiles.
    enrollmentProfiles []EnrollmentProfileable
    // The imported Apple device identities.
    importedAppleDeviceIdentities []ImportedAppleDeviceIdentityable
    // When the service was onboarded.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // When the service last syned with Intune
    lastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Error code reported by Apple during last dep sync.
    lastSyncErrorCode *int32
    // When Intune last requested a sync.
    lastSyncTriggeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // Whether or not the Dep token sharing is enabled with the School Data Sync service.
    shareTokenWithSchoolDataSyncService *bool
    // Gets synced device count
    syncedDeviceCount *int32
    // When the token will expire.
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Friendly Name for Dep Token
    tokenName *string
    // The tokenType property
    tokenType *DepTokenType
}
// NewDepOnboardingSetting instantiates a new depOnboardingSetting and sets the default values.
func NewDepOnboardingSetting()(*DepOnboardingSetting) {
    m := &DepOnboardingSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDepOnboardingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepOnboardingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepOnboardingSetting(), nil
}
// GetAppleIdentifier gets the appleIdentifier property value. The Apple ID used to obtain the current token.
func (m *DepOnboardingSetting) GetAppleIdentifier()(*string) {
    return m.appleIdentifier
}
// GetDataSharingConsentGranted gets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
func (m *DepOnboardingSetting) GetDataSharingConsentGranted()(*bool) {
    return m.dataSharingConsentGranted
}
// GetDefaultIosEnrollmentProfile gets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultIosEnrollmentProfile()(DepIOSEnrollmentProfileable) {
    return m.defaultIosEnrollmentProfile
}
// GetDefaultMacOsEnrollmentProfile gets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultMacOsEnrollmentProfile()(DepMacOSEnrollmentProfileable) {
    return m.defaultMacOsEnrollmentProfile
}
// GetEnrollmentProfiles gets the enrollmentProfiles property value. The enrollment profiles.
func (m *DepOnboardingSetting) GetEnrollmentProfiles()([]EnrollmentProfileable) {
    return m.enrollmentProfiles
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepOnboardingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleIdentifier(val)
        }
        return nil
    }
    res["dataSharingConsentGranted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSharingConsentGranted(val)
        }
        return nil
    }
    res["defaultIosEnrollmentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDepIOSEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultIosEnrollmentProfile(val.(DepIOSEnrollmentProfileable))
        }
        return nil
    }
    res["defaultMacOsEnrollmentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDepMacOSEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultMacOsEnrollmentProfile(val.(DepMacOSEnrollmentProfileable))
        }
        return nil
    }
    res["enrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnrollmentProfileable, len(val))
            for i, v := range val {
                res[i] = v.(EnrollmentProfileable)
            }
            m.SetEnrollmentProfiles(res)
        }
        return nil
    }
    res["importedAppleDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedAppleDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedAppleDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(ImportedAppleDeviceIdentityable)
            }
            m.SetImportedAppleDeviceIdentities(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["lastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["lastSyncErrorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncErrorCode(val)
        }
        return nil
    }
    res["lastSyncTriggeredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncTriggeredDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["shareTokenWithSchoolDataSyncService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareTokenWithSchoolDataSyncService(val)
        }
        return nil
    }
    res["syncedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncedDeviceCount(val)
        }
        return nil
    }
    res["tokenExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpirationDateTime(val)
        }
        return nil
    }
    res["tokenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenName(val)
        }
        return nil
    }
    res["tokenType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDepTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenType(val.(*DepTokenType))
        }
        return nil
    }
    return res
}
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. The imported Apple device identities.
func (m *DepOnboardingSetting) GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentityable) {
    return m.importedAppleDeviceIdentities
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the service was onboarded.
func (m *DepOnboardingSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLastSuccessfulSyncDateTime gets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
func (m *DepOnboardingSetting) GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSuccessfulSyncDateTime
}
// GetLastSyncErrorCode gets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
func (m *DepOnboardingSetting) GetLastSyncErrorCode()(*int32) {
    return m.lastSyncErrorCode
}
// GetLastSyncTriggeredDateTime gets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
func (m *DepOnboardingSetting) GetLastSyncTriggeredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncTriggeredDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DepOnboardingSetting) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetShareTokenWithSchoolDataSyncService gets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
func (m *DepOnboardingSetting) GetShareTokenWithSchoolDataSyncService()(*bool) {
    return m.shareTokenWithSchoolDataSyncService
}
// GetSyncedDeviceCount gets the syncedDeviceCount property value. Gets synced device count
func (m *DepOnboardingSetting) GetSyncedDeviceCount()(*int32) {
    return m.syncedDeviceCount
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. When the token will expire.
func (m *DepOnboardingSetting) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.tokenExpirationDateTime
}
// GetTokenName gets the tokenName property value. Friendly Name for Dep Token
func (m *DepOnboardingSetting) GetTokenName()(*string) {
    return m.tokenName
}
// GetTokenType gets the tokenType property value. The tokenType property
func (m *DepOnboardingSetting) GetTokenType()(*DepTokenType) {
    return m.tokenType
}
// Serialize serializes information the current object
func (m *DepOnboardingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appleIdentifier", m.GetAppleIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataSharingConsentGranted", m.GetDataSharingConsentGranted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultIosEnrollmentProfile", m.GetDefaultIosEnrollmentProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultMacOsEnrollmentProfile", m.GetDefaultMacOsEnrollmentProfile())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnrollmentProfiles()))
        for i, v := range m.GetEnrollmentProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("enrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedAppleDeviceIdentities()))
        for i, v := range m.GetImportedAppleDeviceIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("importedAppleDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSuccessfulSyncDateTime", m.GetLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("lastSyncErrorCode", m.GetLastSyncErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncTriggeredDateTime", m.GetLastSyncTriggeredDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shareTokenWithSchoolDataSyncService", m.GetShareTokenWithSchoolDataSyncService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("syncedDeviceCount", m.GetSyncedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("tokenExpirationDateTime", m.GetTokenExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenName", m.GetTokenName())
        if err != nil {
            return err
        }
    }
    if m.GetTokenType() != nil {
        cast := (*m.GetTokenType()).String()
        err = writer.WriteStringValue("tokenType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppleIdentifier sets the appleIdentifier property value. The Apple ID used to obtain the current token.
func (m *DepOnboardingSetting) SetAppleIdentifier(value *string)() {
    m.appleIdentifier = value
}
// SetDataSharingConsentGranted sets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
func (m *DepOnboardingSetting) SetDataSharingConsentGranted(value *bool)() {
    m.dataSharingConsentGranted = value
}
// SetDefaultIosEnrollmentProfile sets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
func (m *DepOnboardingSetting) SetDefaultIosEnrollmentProfile(value DepIOSEnrollmentProfileable)() {
    m.defaultIosEnrollmentProfile = value
}
// SetDefaultMacOsEnrollmentProfile sets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
func (m *DepOnboardingSetting) SetDefaultMacOsEnrollmentProfile(value DepMacOSEnrollmentProfileable)() {
    m.defaultMacOsEnrollmentProfile = value
}
// SetEnrollmentProfiles sets the enrollmentProfiles property value. The enrollment profiles.
func (m *DepOnboardingSetting) SetEnrollmentProfiles(value []EnrollmentProfileable)() {
    m.enrollmentProfiles = value
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. The imported Apple device identities.
func (m *DepOnboardingSetting) SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentityable)() {
    m.importedAppleDeviceIdentities = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the service was onboarded.
func (m *DepOnboardingSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastSuccessfulSyncDateTime sets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
func (m *DepOnboardingSetting) SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSuccessfulSyncDateTime = value
}
// SetLastSyncErrorCode sets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
func (m *DepOnboardingSetting) SetLastSyncErrorCode(value *int32)() {
    m.lastSyncErrorCode = value
}
// SetLastSyncTriggeredDateTime sets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
func (m *DepOnboardingSetting) SetLastSyncTriggeredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncTriggeredDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DepOnboardingSetting) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetShareTokenWithSchoolDataSyncService sets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
func (m *DepOnboardingSetting) SetShareTokenWithSchoolDataSyncService(value *bool)() {
    m.shareTokenWithSchoolDataSyncService = value
}
// SetSyncedDeviceCount sets the syncedDeviceCount property value. Gets synced device count
func (m *DepOnboardingSetting) SetSyncedDeviceCount(value *int32)() {
    m.syncedDeviceCount = value
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. When the token will expire.
func (m *DepOnboardingSetting) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenExpirationDateTime = value
}
// SetTokenName sets the tokenName property value. Friendly Name for Dep Token
func (m *DepOnboardingSetting) SetTokenName(value *string)() {
    m.tokenName = value
}
// SetTokenType sets the tokenType property value. The tokenType property
func (m *DepOnboardingSetting) SetTokenType(value *DepTokenType)() {
    m.tokenType = value
}
