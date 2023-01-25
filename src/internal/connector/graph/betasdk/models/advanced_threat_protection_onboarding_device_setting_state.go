package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdvancedThreatProtectionOnboardingDeviceSettingState aTP onboarding State for a given device.
type AdvancedThreatProtectionOnboardingDeviceSettingState struct {
    Entity
    // The DateTime when device compliance grace period expires
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Device Id that is being reported
    deviceId *string
    // The device model that is being reported
    deviceModel *string
    // The Device Name that is being reported
    deviceName *string
    // Device type.
    platformType *DeviceType
    // The setting class name and property name.
    setting *string
    // The Setting Name that is being reported
    settingName *string
    // The state property
    state *ComplianceStatus
    // The User email address that is being reported
    userEmail *string
    // The user Id that is being reported
    userId *string
    // The User Name that is being reported
    userName *string
    // The User PrincipalName that is being reported
    userPrincipalName *string
}
// NewAdvancedThreatProtectionOnboardingDeviceSettingState instantiates a new advancedThreatProtectionOnboardingDeviceSettingState and sets the default values.
func NewAdvancedThreatProtectionOnboardingDeviceSettingState()(*AdvancedThreatProtectionOnboardingDeviceSettingState) {
    m := &AdvancedThreatProtectionOnboardingDeviceSettingState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdvancedThreatProtectionOnboardingDeviceSettingState(), nil
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceGracePeriodExpirationDateTime
}
// GetDeviceId gets the deviceId property value. The Device Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceModel gets the deviceModel property value. The device model that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetDeviceName gets the deviceName property value. The Device Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceGracePeriodExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*DeviceType))
        }
        return nil
    }
    res["setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["userEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetPlatformType gets the platformType property value. Device type.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetPlatformType()(*DeviceType) {
    return m.platformType
}
// GetSetting gets the setting property value. The setting class name and property name.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetSetting()(*string) {
    return m.setting
}
// GetSettingName gets the settingName property value. The Setting Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetSettingName()(*string) {
    return m.settingName
}
// GetState gets the state property value. The state property
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetUserEmail gets the userEmail property value. The User email address that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserEmail()(*string) {
    return m.userEmail
}
// GetUserId gets the userId property value. The user Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserId()(*string) {
    return m.userId
}
// GetUserName gets the userName property value. The User Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. The User PrincipalName that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("complianceGracePeriodExpirationDateTime", m.GetComplianceGracePeriodExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// SetDeviceId sets the deviceId property value. The Device Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceModel sets the deviceModel property value. The device model that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetDeviceName sets the deviceName property value. The Device Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetPlatformType sets the platformType property value. Device type.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetPlatformType(value *DeviceType)() {
    m.platformType = value
}
// SetSetting sets the setting property value. The setting class name and property name.
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetSetting(value *string)() {
    m.setting = value
}
// SetSettingName sets the settingName property value. The Setting Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetSettingName(value *string)() {
    m.settingName = value
}
// SetState sets the state property value. The state property
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetUserEmail sets the userEmail property value. The User email address that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserEmail(value *string)() {
    m.userEmail = value
}
// SetUserId sets the userId property value. The user Id that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserId(value *string)() {
    m.userId = value
}
// SetUserName sets the userName property value. The User Name that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The User PrincipalName that is being reported
func (m *AdvancedThreatProtectionOnboardingDeviceSettingState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
