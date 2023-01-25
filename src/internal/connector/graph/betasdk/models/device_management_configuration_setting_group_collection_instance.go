package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingGroupCollectionInstance 
type DeviceManagementConfigurationSettingGroupCollectionInstance struct {
    DeviceManagementConfigurationSettingInstance
}
// NewDeviceManagementConfigurationSettingGroupCollectionInstance instantiates a new DeviceManagementConfigurationSettingGroupCollectionInstance and sets the default values.
func NewDeviceManagementConfigurationSettingGroupCollectionInstance()(*DeviceManagementConfigurationSettingGroupCollectionInstance) {
    m := &DeviceManagementConfigurationSettingGroupCollectionInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSettingGroupCollectionInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingGroupCollectionInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingGroupCollectionInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingGroupCollectionInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingGroupCollectionInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
