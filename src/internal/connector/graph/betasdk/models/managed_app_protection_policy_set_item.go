package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppProtectionPolicySetItem 
type ManagedAppProtectionPolicySetItem struct {
    PolicySetItem
    // TargetedAppManagementLevels of the ManagedAppPolicySetItem.
    targetedAppManagementLevels *string
}
// NewManagedAppProtectionPolicySetItem instantiates a new ManagedAppProtectionPolicySetItem and sets the default values.
func NewManagedAppProtectionPolicySetItem()(*ManagedAppProtectionPolicySetItem) {
    m := &ManagedAppProtectionPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.managedAppProtectionPolicySetItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedAppProtectionPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppProtectionPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppProtectionPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppProtectionPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    res["targetedAppManagementLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedAppManagementLevels(val)
        }
        return nil
    }
    return res
}
// GetTargetedAppManagementLevels gets the targetedAppManagementLevels property value. TargetedAppManagementLevels of the ManagedAppPolicySetItem.
func (m *ManagedAppProtectionPolicySetItem) GetTargetedAppManagementLevels()(*string) {
    return m.targetedAppManagementLevels
}
// Serialize serializes information the current object
func (m *ManagedAppProtectionPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetedAppManagementLevels", m.GetTargetedAppManagementLevels())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetedAppManagementLevels sets the targetedAppManagementLevels property value. TargetedAppManagementLevels of the ManagedAppPolicySetItem.
func (m *ManagedAppProtectionPolicySetItem) SetTargetedAppManagementLevels(value *string)() {
    m.targetedAppManagementLevels = value
}
