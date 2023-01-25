package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingCollectionInstance 
type DeviceManagementConfigurationSimpleSettingCollectionInstance struct {
    DeviceManagementConfigurationSettingInstance
    // Simple setting collection instance value
    simpleSettingCollectionValue []DeviceManagementConfigurationSimpleSettingValueable
}
// NewDeviceManagementConfigurationSimpleSettingCollectionInstance instantiates a new DeviceManagementConfigurationSimpleSettingCollectionInstance and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingCollectionInstance()(*DeviceManagementConfigurationSimpleSettingCollectionInstance) {
    m := &DeviceManagementConfigurationSimpleSettingCollectionInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingCollectionInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingCollectionInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingCollectionInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["simpleSettingCollectionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSimpleSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSimpleSettingValueable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSimpleSettingValueable)
            }
            m.SetSimpleSettingCollectionValue(res)
        }
        return nil
    }
    return res
}
// GetSimpleSettingCollectionValue gets the simpleSettingCollectionValue property value. Simple setting collection instance value
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstance) GetSimpleSettingCollectionValue()([]DeviceManagementConfigurationSimpleSettingValueable) {
    return m.simpleSettingCollectionValue
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSimpleSettingCollectionValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSimpleSettingCollectionValue()))
        for i, v := range m.GetSimpleSettingCollectionValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("simpleSettingCollectionValue", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSimpleSettingCollectionValue sets the simpleSettingCollectionValue property value. Simple setting collection instance value
func (m *DeviceManagementConfigurationSimpleSettingCollectionInstance) SetSimpleSettingCollectionValue(value []DeviceManagementConfigurationSimpleSettingValueable)() {
    m.simpleSettingCollectionValue = value
}
