package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkSettings 
type AndroidForWorkSettings struct {
    Entity
    // Bind status of the tenant with the Google EMM API
    bindStatus *AndroidForWorkBindStatus
    // Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
    deviceOwnerManagementEnabled *bool
    // Android for Work device management targeting type for the account
    enrollmentTarget *AndroidForWorkEnrollmentTarget
    // Last completion time for app sync
    lastAppSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Sync status of the tenant with the Google EMM API
    lastAppSyncStatus *AndroidForWorkSyncStatus
    // Last modification time for Android for Work settings
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Organization name used when onboarding Android for Work
    ownerOrganizationName *string
    // Owner UPN that created the enterprise
    ownerUserPrincipalName *string
    // Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
    targetGroupIds []string
}
// NewAndroidForWorkSettings instantiates a new androidForWorkSettings and sets the default values.
func NewAndroidForWorkSettings()(*AndroidForWorkSettings) {
    m := &AndroidForWorkSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidForWorkSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkSettings(), nil
}
// GetBindStatus gets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidForWorkSettings) GetBindStatus()(*AndroidForWorkBindStatus) {
    return m.bindStatus
}
// GetDeviceOwnerManagementEnabled gets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidForWorkSettings) GetDeviceOwnerManagementEnabled()(*bool) {
    return m.deviceOwnerManagementEnabled
}
// GetEnrollmentTarget gets the enrollmentTarget property value. Android for Work device management targeting type for the account
func (m *AndroidForWorkSettings) GetEnrollmentTarget()(*AndroidForWorkEnrollmentTarget) {
    return m.enrollmentTarget
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bindStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkBindStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindStatus(val.(*AndroidForWorkBindStatus))
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
        val, err := n.GetEnumValue(ParseAndroidForWorkEnrollmentTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTarget(val.(*AndroidForWorkEnrollmentTarget))
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
        val, err := n.GetEnumValue(ParseAndroidForWorkSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncStatus(val.(*AndroidForWorkSyncStatus))
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
func (m *AndroidForWorkSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastAppSyncDateTime
}
// GetLastAppSyncStatus gets the lastAppSyncStatus property value. Sync status of the tenant with the Google EMM API
func (m *AndroidForWorkSettings) GetLastAppSyncStatus()(*AndroidForWorkSyncStatus) {
    return m.lastAppSyncStatus
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification time for Android for Work settings
func (m *AndroidForWorkSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOwnerOrganizationName gets the ownerOrganizationName property value. Organization name used when onboarding Android for Work
func (m *AndroidForWorkSettings) GetOwnerOrganizationName()(*string) {
    return m.ownerOrganizationName
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidForWorkSettings) GetOwnerUserPrincipalName()(*string) {
    return m.ownerUserPrincipalName
}
// GetTargetGroupIds gets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidForWorkSettings) GetTargetGroupIds()([]string) {
    return m.targetGroupIds
}
// Serialize serializes information the current object
func (m *AndroidForWorkSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBindStatus() != nil {
        cast := (*m.GetBindStatus()).String()
        err = writer.WriteStringValue("bindStatus", &cast)
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
// SetBindStatus sets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidForWorkSettings) SetBindStatus(value *AndroidForWorkBindStatus)() {
    m.bindStatus = value
}
// SetDeviceOwnerManagementEnabled sets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidForWorkSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    m.deviceOwnerManagementEnabled = value
}
// SetEnrollmentTarget sets the enrollmentTarget property value. Android for Work device management targeting type for the account
func (m *AndroidForWorkSettings) SetEnrollmentTarget(value *AndroidForWorkEnrollmentTarget)() {
    m.enrollmentTarget = value
}
// SetLastAppSyncDateTime sets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidForWorkSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastAppSyncDateTime = value
}
// SetLastAppSyncStatus sets the lastAppSyncStatus property value. Sync status of the tenant with the Google EMM API
func (m *AndroidForWorkSettings) SetLastAppSyncStatus(value *AndroidForWorkSyncStatus)() {
    m.lastAppSyncStatus = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification time for Android for Work settings
func (m *AndroidForWorkSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOwnerOrganizationName sets the ownerOrganizationName property value. Organization name used when onboarding Android for Work
func (m *AndroidForWorkSettings) SetOwnerOrganizationName(value *string)() {
    m.ownerOrganizationName = value
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidForWorkSettings) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
// SetTargetGroupIds sets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidForWorkSettings) SetTargetGroupIds(value []string)() {
    m.targetGroupIds = value
}
