package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineSettingState the security baseline compliance state of a setting for a device
type SecurityBaselineSettingState struct {
    Entity
    // The policies that contribute to this setting instance
    contributingPolicies []SecurityBaselineContributingPolicyable
    // The error code if the setting is in error state
    errorCode *string
    // The setting category id which this setting belongs to
    settingCategoryId *string
    // The setting category name which this setting belongs to
    settingCategoryName *string
    // The setting id guid
    settingId *string
    // The setting name that is being reported
    settingName *string
    // The policies that contribute to this setting instance
    sourcePolicies []SettingSourceable
    // Security Baseline Compliance State
    state *SecurityBaselineComplianceState
}
// NewSecurityBaselineSettingState instantiates a new securityBaselineSettingState and sets the default values.
func NewSecurityBaselineSettingState()(*SecurityBaselineSettingState) {
    m := &SecurityBaselineSettingState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityBaselineSettingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineSettingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineSettingState(), nil
}
// GetContributingPolicies gets the contributingPolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) GetContributingPolicies()([]SecurityBaselineContributingPolicyable) {
    return m.contributingPolicies
}
// GetErrorCode gets the errorCode property value. The error code if the setting is in error state
func (m *SecurityBaselineSettingState) GetErrorCode()(*string) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineSettingState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contributingPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineContributingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineContributingPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityBaselineContributingPolicyable)
            }
            m.SetContributingPolicies(res)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["settingCategoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategoryId(val)
        }
        return nil
    }
    res["settingCategoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategoryName(val)
        }
        return nil
    }
    res["settingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingId(val)
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
    res["sourcePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSourceable, len(val))
            for i, v := range val {
                res[i] = v.(SettingSourceable)
            }
            m.SetSourcePolicies(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityBaselineComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*SecurityBaselineComplianceState))
        }
        return nil
    }
    return res
}
// GetSettingCategoryId gets the settingCategoryId property value. The setting category id which this setting belongs to
func (m *SecurityBaselineSettingState) GetSettingCategoryId()(*string) {
    return m.settingCategoryId
}
// GetSettingCategoryName gets the settingCategoryName property value. The setting category name which this setting belongs to
func (m *SecurityBaselineSettingState) GetSettingCategoryName()(*string) {
    return m.settingCategoryName
}
// GetSettingId gets the settingId property value. The setting id guid
func (m *SecurityBaselineSettingState) GetSettingId()(*string) {
    return m.settingId
}
// GetSettingName gets the settingName property value. The setting name that is being reported
func (m *SecurityBaselineSettingState) GetSettingName()(*string) {
    return m.settingName
}
// GetSourcePolicies gets the sourcePolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) GetSourcePolicies()([]SettingSourceable) {
    return m.sourcePolicies
}
// GetState gets the state property value. Security Baseline Compliance State
func (m *SecurityBaselineSettingState) GetState()(*SecurityBaselineComplianceState) {
    return m.state
}
// Serialize serializes information the current object
func (m *SecurityBaselineSettingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContributingPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContributingPolicies()))
        for i, v := range m.GetContributingPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("contributingPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingCategoryId", m.GetSettingCategoryId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingCategoryName", m.GetSettingCategoryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingId", m.GetSettingId())
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
    if m.GetSourcePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSourcePolicies()))
        for i, v := range m.GetSourcePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sourcePolicies", cast)
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
    return nil
}
// SetContributingPolicies sets the contributingPolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) SetContributingPolicies(value []SecurityBaselineContributingPolicyable)() {
    m.contributingPolicies = value
}
// SetErrorCode sets the errorCode property value. The error code if the setting is in error state
func (m *SecurityBaselineSettingState) SetErrorCode(value *string)() {
    m.errorCode = value
}
// SetSettingCategoryId sets the settingCategoryId property value. The setting category id which this setting belongs to
func (m *SecurityBaselineSettingState) SetSettingCategoryId(value *string)() {
    m.settingCategoryId = value
}
// SetSettingCategoryName sets the settingCategoryName property value. The setting category name which this setting belongs to
func (m *SecurityBaselineSettingState) SetSettingCategoryName(value *string)() {
    m.settingCategoryName = value
}
// SetSettingId sets the settingId property value. The setting id guid
func (m *SecurityBaselineSettingState) SetSettingId(value *string)() {
    m.settingId = value
}
// SetSettingName sets the settingName property value. The setting name that is being reported
func (m *SecurityBaselineSettingState) SetSettingName(value *string)() {
    m.settingName = value
}
// SetSourcePolicies sets the sourcePolicies property value. The policies that contribute to this setting instance
func (m *SecurityBaselineSettingState) SetSourcePolicies(value []SettingSourceable)() {
    m.sourcePolicies = value
}
// SetState sets the state property value. Security Baseline Compliance State
func (m *SecurityBaselineSettingState) SetState(value *SecurityBaselineComplianceState)() {
    m.state = value
}
