package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettings 
type AndroidManagedStoreAccountEnterpriseSettings struct {
    Entity
    // Company codes for AndroidManagedStoreAccountEnterpriseSettings
    androidDeviceOwnerFullyManagedEnrollmentEnabled *bool
    // Bind status of the tenant with the Google EMM API
    bindStatus *AndroidManagedStoreAccountBindStatus
    // Company codes for AndroidManagedStoreAccountEnterpriseSettings
    companyCodes []AndroidEnrollmentCompanyCodeable
    // Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
    deviceOwnerManagementEnabled *bool
    // Android for Work device management targeting type for the account
    enrollmentTarget *AndroidManagedStoreAccountEnrollmentTarget
    // Last completion time for app sync
    lastAppSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Sync status of the tenant with the Google EMM API
    lastAppSyncStatus *AndroidManagedStoreAccountAppSyncStatus
    // Last modification time for Android enterprise settings
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Initial scope tags for MGP apps
    managedGooglePlayInitialScopeTagIds []string
    // Organization name used when onboarding Android Enterprise
    ownerOrganizationName *string
    // Owner UPN that created the enterprise
    ownerUserPrincipalName *string
    // Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
    targetGroupIds []string
}
// NewAndroidManagedStoreAccountEnterpriseSettings instantiates a new androidManagedStoreAccountEnterpriseSettings and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettings) {
    m := &AndroidManagedStoreAccountEnterpriseSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettings(), nil
}
// GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled gets the androidDeviceOwnerFullyManagedEnrollmentEnabled property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled()(*bool) {
    return m.androidDeviceOwnerFullyManagedEnrollmentEnabled
}
// GetBindStatus gets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetBindStatus()(*AndroidManagedStoreAccountBindStatus) {
    return m.bindStatus
}
// GetCompanyCodes gets the companyCodes property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetCompanyCodes()([]AndroidEnrollmentCompanyCodeable) {
    return m.companyCodes
}
// GetDeviceOwnerManagementEnabled gets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetDeviceOwnerManagementEnabled()(*bool) {
    return m.deviceOwnerManagementEnabled
}
// GetEnrollmentTarget gets the enrollmentTarget property value. Android for Work device management targeting type for the account
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetEnrollmentTarget()(*AndroidManagedStoreAccountEnrollmentTarget) {
    return m.enrollmentTarget
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidDeviceOwnerFullyManagedEnrollmentEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(val)
        }
        return nil
    }
    res["bindStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountBindStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindStatus(val.(*AndroidManagedStoreAccountBindStatus))
        }
        return nil
    }
    res["companyCodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidEnrollmentCompanyCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidEnrollmentCompanyCodeable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidEnrollmentCompanyCodeable)
            }
            m.SetCompanyCodes(res)
        }
        return nil
    }
    res["deviceOwnerManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnerManagementEnabled(val)
        }
        return nil
    }
    res["enrollmentTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountEnrollmentTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTarget(val.(*AndroidManagedStoreAccountEnrollmentTarget))
        }
        return nil
    }
    res["lastAppSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncDateTime(val)
        }
        return nil
    }
    res["lastAppSyncStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountAppSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncStatus(val.(*AndroidManagedStoreAccountAppSyncStatus))
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
    res["managedGooglePlayInitialScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedGooglePlayInitialScopeTagIds(res)
        }
        return nil
    }
    res["ownerOrganizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerOrganizationName(val)
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    res["targetGroupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTargetGroupIds(res)
        }
        return nil
    }
    return res
}
// GetLastAppSyncDateTime gets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastAppSyncDateTime
}
// GetLastAppSyncStatus gets the lastAppSyncStatus property value. Sync status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncStatus()(*AndroidManagedStoreAccountAppSyncStatus) {
    return m.lastAppSyncStatus
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification time for Android enterprise settings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetManagedGooglePlayInitialScopeTagIds gets the managedGooglePlayInitialScopeTagIds property value. Initial scope tags for MGP apps
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetManagedGooglePlayInitialScopeTagIds()([]string) {
    return m.managedGooglePlayInitialScopeTagIds
}
// GetOwnerOrganizationName gets the ownerOrganizationName property value. Organization name used when onboarding Android Enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerOrganizationName()(*string) {
    return m.ownerOrganizationName
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerUserPrincipalName()(*string) {
    return m.ownerUserPrincipalName
}
// GetTargetGroupIds gets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetTargetGroupIds()([]string) {
    return m.targetGroupIds
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("androidDeviceOwnerFullyManagedEnrollmentEnabled", m.GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetBindStatus() != nil {
        cast := (*m.GetBindStatus()).String()
        err = writer.WriteStringValue("bindStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompanyCodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyCodes()))
        for i, v := range m.GetCompanyCodes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("companyCodes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceOwnerManagementEnabled", m.GetDeviceOwnerManagementEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentTarget() != nil {
        cast := (*m.GetEnrollmentTarget()).String()
        err = writer.WriteStringValue("enrollmentTarget", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastAppSyncDateTime", m.GetLastAppSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLastAppSyncStatus() != nil {
        cast := (*m.GetLastAppSyncStatus()).String()
        err = writer.WriteStringValue("lastAppSyncStatus", &cast)
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
    if m.GetManagedGooglePlayInitialScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("managedGooglePlayInitialScopeTagIds", m.GetManagedGooglePlayInitialScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerOrganizationName", m.GetOwnerOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetTargetGroupIds() != nil {
        err = writer.WriteCollectionOfStringValues("targetGroupIds", m.GetTargetGroupIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled sets the androidDeviceOwnerFullyManagedEnrollmentEnabled property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(value *bool)() {
    m.androidDeviceOwnerFullyManagedEnrollmentEnabled = value
}
// SetBindStatus sets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetBindStatus(value *AndroidManagedStoreAccountBindStatus)() {
    m.bindStatus = value
}
// SetCompanyCodes sets the companyCodes property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetCompanyCodes(value []AndroidEnrollmentCompanyCodeable)() {
    m.companyCodes = value
}
// SetDeviceOwnerManagementEnabled sets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    m.deviceOwnerManagementEnabled = value
}
// SetEnrollmentTarget sets the enrollmentTarget property value. Android for Work device management targeting type for the account
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetEnrollmentTarget(value *AndroidManagedStoreAccountEnrollmentTarget)() {
    m.enrollmentTarget = value
}
// SetLastAppSyncDateTime sets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastAppSyncDateTime = value
}
// SetLastAppSyncStatus sets the lastAppSyncStatus property value. Sync status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncStatus(value *AndroidManagedStoreAccountAppSyncStatus)() {
    m.lastAppSyncStatus = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification time for Android enterprise settings
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetManagedGooglePlayInitialScopeTagIds sets the managedGooglePlayInitialScopeTagIds property value. Initial scope tags for MGP apps
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetManagedGooglePlayInitialScopeTagIds(value []string)() {
    m.managedGooglePlayInitialScopeTagIds = value
}
// SetOwnerOrganizationName sets the ownerOrganizationName property value. Organization name used when onboarding Android Enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerOrganizationName(value *string)() {
    m.ownerOrganizationName = value
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
// SetTargetGroupIds sets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetTargetGroupIds(value []string)() {
    m.targetGroupIds = value
}
