package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingInstance 
type DeviceManagementConfigurationSimpleSettingInstance struct {
    DeviceManagementConfigurationSettingInstance
    // The simpleSettingValue property
    simpleSettingValue DeviceManagementConfigurationSimpleSettingValueable
}
// NewDeviceManagementConfigurationSimpleSettingInstance instantiates a new DeviceManagementConfigurationSimpleSettingInstance and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingInstance()(*DeviceManagementConfigurationSimpleSettingInstance) {
    m := &DeviceManagementConfigurationSimpleSettingInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["simpleSettingValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSimpleSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimpleSettingValue(val.(DeviceManagementConfigurationSimpleSettingValueable))
        }
        return nil
    }
    return res
}
// GetSimpleSettingValue gets the simpleSettingValue property value. The simpleSettingValue property
func (m *DeviceManagementConfigurationSimpleSettingInstance) GetSimpleSettingValue()(DeviceManagementConfigurationSimpleSettingValueable) {
    return m.simpleSettingValue
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("simpleSettingValue", m.GetSimpleSettingValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSimpleSettingValue sets the simpleSettingValue property value. The simpleSettingValue property
func (m *DeviceManagementConfigurationSimpleSettingInstance) SetSimpleSettingValue(value DeviceManagementConfigurationSimpleSettingValueable)() {
    m.simpleSettingValue = value
}
