package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceMobileAppConfigurationState managed Device Mobile App Configuration State for a given device.
type ManagedDeviceMobileAppConfigurationState struct {
    Entity
    // The name of the policy for this policyBase
    displayName *string
    // Supported platform types for policies.
    platformType *PolicyPlatformType
    // Count of how many setting a policy holds
    settingCount *int32
    // The settingStates property
    settingStates []ManagedDeviceMobileAppConfigurationSettingStateable
    // The state property
    state *ComplianceStatus
    // User unique identifier, must be Guid
    userId *string
    // User Principal Name
    userPrincipalName *string
    // The version of the policy
    version *int32
}
// NewManagedDeviceMobileAppConfigurationState instantiates a new managedDeviceMobileAppConfigurationState and sets the default values.
func NewManagedDeviceMobileAppConfigurationState()(*ManagedDeviceMobileAppConfigurationState) {
    m := &ManagedDeviceMobileAppConfigurationState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceMobileAppConfigurationState(), nil
}
// GetDisplayName gets the displayName property value. The name of the policy for this policyBase
func (m *ManagedDeviceMobileAppConfigurationState) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceMobileAppConfigurationState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["settingCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCount(val)
        }
        return nil
    }
    res["settingStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationSettingStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationSettingStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationSettingStateable)
            }
            m.SetSettingStates(res)
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
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *ManagedDeviceMobileAppConfigurationState) GetPlatformType()(*PolicyPlatformType) {
    return m.platformType
}
// GetSettingCount gets the settingCount property value. Count of how many setting a policy holds
func (m *ManagedDeviceMobileAppConfigurationState) GetSettingCount()(*int32) {
    return m.settingCount
}
// GetSettingStates gets the settingStates property value. The settingStates property
func (m *ManagedDeviceMobileAppConfigurationState) GetSettingStates()([]ManagedDeviceMobileAppConfigurationSettingStateable) {
    return m.settingStates
}
// GetState gets the state property value. The state property
func (m *ManagedDeviceMobileAppConfigurationState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetUserId gets the userId property value. User unique identifier, must be Guid
func (m *ManagedDeviceMobileAppConfigurationState) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name
func (m *ManagedDeviceMobileAppConfigurationState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetVersion gets the version property value. The version of the policy
func (m *ManagedDeviceMobileAppConfigurationState) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *ManagedDeviceMobileAppConfigurationState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteInt32Value("settingCount", m.GetSettingCount())
        if err != nil {
            return err
        }
    }
    if m.GetSettingStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingStates()))
        for i, v := range m.GetSettingStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingStates", cast)
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
        err = writer.WriteStringValue("userId", m.GetUserId())
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
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the policy for this policyBase
func (m *ManagedDeviceMobileAppConfigurationState) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *ManagedDeviceMobileAppConfigurationState) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// SetSettingCount sets the settingCount property value. Count of how many setting a policy holds
func (m *ManagedDeviceMobileAppConfigurationState) SetSettingCount(value *int32)() {
    m.settingCount = value
}
// SetSettingStates sets the settingStates property value. The settingStates property
func (m *ManagedDeviceMobileAppConfigurationState) SetSettingStates(value []ManagedDeviceMobileAppConfigurationSettingStateable)() {
    m.settingStates = value
}
// SetState sets the state property value. The state property
func (m *ManagedDeviceMobileAppConfigurationState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetUserId sets the userId property value. User unique identifier, must be Guid
func (m *ManagedDeviceMobileAppConfigurationState) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name
func (m *ManagedDeviceMobileAppConfigurationState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetVersion sets the version property value. The version of the policy
func (m *ManagedDeviceMobileAppConfigurationState) SetVersion(value *int32)() {
    m.version = value
}
