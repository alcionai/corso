package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetedManagedAppConfigurationPolicySetItem 
type TargetedManagedAppConfigurationPolicySetItem struct {
    PolicySetItem
}
// NewTargetedManagedAppConfigurationPolicySetItem instantiates a new TargetedManagedAppConfigurationPolicySetItem and sets the default values.
func NewTargetedManagedAppConfigurationPolicySetItem()(*TargetedManagedAppConfigurationPolicySetItem) {
    m := &TargetedManagedAppConfigurationPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.targetedManagedAppConfigurationPolicySetItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTargetedManagedAppConfigurationPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetedManagedAppConfigurationPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetedManagedAppConfigurationPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetedManagedAppConfigurationPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TargetedManagedAppConfigurationPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
