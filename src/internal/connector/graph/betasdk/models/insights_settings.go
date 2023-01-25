package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsightsSettings 
type InsightsSettings struct {
    Entity
    // The ID of an Azure Active Directory group, of which the specified type of insights are disabled for its members. Default is empty. Optional.
    disabledForGroup *string
    // true if the specified type of insights are enabled for the organization; false if the specified type of insights are disabled for all users without exceptions. Default is true. Optional.
    isEnabledInOrganization *bool
}
// NewInsightsSettings instantiates a new insightsSettings and sets the default values.
func NewInsightsSettings()(*InsightsSettings) {
    m := &InsightsSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInsightsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInsightsSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsightsSettings(), nil
}
// GetDisabledForGroup gets the disabledForGroup property value. The ID of an Azure Active Directory group, of which the specified type of insights are disabled for its members. Default is empty. Optional.
func (m *InsightsSettings) GetDisabledForGroup()(*string) {
    return m.disabledForGroup
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightsSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["disabledForGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledForGroup(val)
        }
        return nil
    }
    res["isEnabledInOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledInOrganization(val)
        }
        return nil
    }
    return res
}
// GetIsEnabledInOrganization gets the isEnabledInOrganization property value. true if the specified type of insights are enabled for the organization; false if the specified type of insights are disabled for all users without exceptions. Default is true. Optional.
func (m *InsightsSettings) GetIsEnabledInOrganization()(*bool) {
    return m.isEnabledInOrganization
}
// Serialize serializes information the current object
func (m *InsightsSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("disabledForGroup", m.GetDisabledForGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledInOrganization", m.GetIsEnabledInOrganization())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisabledForGroup sets the disabledForGroup property value. The ID of an Azure Active Directory group, of which the specified type of insights are disabled for its members. Default is empty. Optional.
func (m *InsightsSettings) SetDisabledForGroup(value *string)() {
    m.disabledForGroup = value
}
// SetIsEnabledInOrganization sets the isEnabledInOrganization property value. true if the specified type of insights are enabled for the organization; false if the specified type of insights are disabled for all users without exceptions. Default is true. Optional.
func (m *InsightsSettings) SetIsEnabledInOrganization(value *bool)() {
    m.isEnabledInOrganization = value
}
