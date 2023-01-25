package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingGroupInstance 
type DeviceManagementConfigurationSettingGroupInstance struct {
    DeviceManagementConfigurationSettingInstance
}
// NewDeviceManagementConfigurationSettingGroupInstance instantiates a new DeviceManagementConfigurationSettingGroupInstance and sets the default values.
func NewDeviceManagementConfigurationSettingGroupInstance()(*DeviceManagementConfigurationSettingGroupInstance) {
    m := &DeviceManagementConfigurationSettingGroupInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSettingGroupInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSettingGroupInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingGroupInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingGroupInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingGroupInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingGroupInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
