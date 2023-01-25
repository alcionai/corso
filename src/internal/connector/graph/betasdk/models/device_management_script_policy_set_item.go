package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptPolicySetItem 
type DeviceManagementScriptPolicySetItem struct {
    PolicySetItem
}
// NewDeviceManagementScriptPolicySetItem instantiates a new DeviceManagementScriptPolicySetItem and sets the default values.
func NewDeviceManagementScriptPolicySetItem()(*DeviceManagementScriptPolicySetItem) {
    m := &DeviceManagementScriptPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementScriptPolicySetItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementScriptPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
