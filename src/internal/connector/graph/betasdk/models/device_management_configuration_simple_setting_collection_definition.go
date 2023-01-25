package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingCollectionDefinition 
type DeviceManagementConfigurationSimpleSettingCollectionDefinition struct {
    DeviceManagementConfigurationSimpleSettingDefinition
    // Maximum number of simple settings in the collection. Valid values 1 to 100
    maximumCount *int32
    // Minimum number of simple settings in the collection. Valid values 1 to 100
    minimumCount *int32
}
// NewDeviceManagementConfigurationSimpleSettingCollectionDefinition instantiates a new DeviceManagementConfigurationSimpleSettingCollectionDefinition and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingCollectionDefinition()(*DeviceManagementConfigurationSimpleSettingCollectionDefinition) {
    m := &DeviceManagementConfigurationSimpleSettingCollectionDefinition{
        DeviceManagementConfigurationSimpleSettingDefinition: *NewDeviceManagementConfigurationSimpleSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingCollectionDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingCollectionDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingCollectionDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingDefinition.GetFieldDeserializers()
    res["maximumCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCount(val)
        }
        return nil
    }
    res["minimumCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCount(val)
        }
        return nil
    }
    return res
}
// GetMaximumCount gets the maximumCount property value. Maximum number of simple settings in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) GetMaximumCount()(*int32) {
    return m.maximumCount
}
// GetMinimumCount gets the minimumCount property value. Minimum number of simple settings in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) GetMinimumCount()(*int32) {
    return m.minimumCount
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSimpleSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumCount", m.GetMaximumCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumCount", m.GetMinimumCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaximumCount sets the maximumCount property value. Maximum number of simple settings in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) SetMaximumCount(value *int32)() {
    m.maximumCount = value
}
// SetMinimumCount sets the minimumCount property value. Minimum number of simple settings in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationSimpleSettingCollectionDefinition) SetMinimumCount(value *int32)() {
    m.minimumCount = value
}
