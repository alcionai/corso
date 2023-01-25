package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComplianceScheduledActionForRule scheduled Action for Rule
type DeviceManagementComplianceScheduledActionForRule struct {
    Entity
    // Name of the rule which this scheduled action applies to.
    ruleName *string
    // The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
    scheduledActionConfigurations []DeviceManagementComplianceActionItemable
}
// NewDeviceManagementComplianceScheduledActionForRule instantiates a new deviceManagementComplianceScheduledActionForRule and sets the default values.
func NewDeviceManagementComplianceScheduledActionForRule()(*DeviceManagementComplianceScheduledActionForRule) {
    m := &DeviceManagementComplianceScheduledActionForRule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComplianceScheduledActionForRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplianceScheduledActionForRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ruleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    res["scheduledActionConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementComplianceActionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementComplianceActionItemable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementComplianceActionItemable)
            }
            m.SetScheduledActionConfigurations(res)
        }
        return nil
    }
    return res
}
// GetRuleName gets the ruleName property value. Name of the rule which this scheduled action applies to.
func (m *DeviceManagementComplianceScheduledActionForRule) GetRuleName()(*string) {
    return m.ruleName
}
// GetScheduledActionConfigurations gets the scheduledActionConfigurations property value. The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceScheduledActionForRule) GetScheduledActionConfigurations()([]DeviceManagementComplianceActionItemable) {
    return m.scheduledActionConfigurations
}
// Serialize serializes information the current object
func (m *DeviceManagementComplianceScheduledActionForRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    if m.GetScheduledActionConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScheduledActionConfigurations()))
        for i, v := range m.GetScheduledActionConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("scheduledActionConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRuleName sets the ruleName property value. Name of the rule which this scheduled action applies to.
func (m *DeviceManagementComplianceScheduledActionForRule) SetRuleName(value *string)() {
    m.ruleName = value
}
// SetScheduledActionConfigurations sets the scheduledActionConfigurations property value. The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceScheduledActionForRule) SetScheduledActionConfigurations(value []DeviceManagementComplianceActionItemable)() {
    m.scheduledActionConfigurations = value
}
