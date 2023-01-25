package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingCollectionDefinition 
type DeviceManagementConfigurationChoiceSettingCollectionDefinition struct {
    DeviceManagementConfigurationChoiceSettingDefinition
    // Maximum number of choices in the collection. Valid values 1 to 100
    maximumCount *int32
    // Minimum number of choices in the collection. Valid values 1 to 100
    minimumCount *int32
}
// NewDeviceManagementConfigurationChoiceSettingCollectionDefinition instantiates a new DeviceManagementConfigurationChoiceSettingCollectionDefinition and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingCollectionDefinition()(*DeviceManagementConfigurationChoiceSettingCollectionDefinition) {
    m := &DeviceManagementConfigurationChoiceSettingCollectionDefinition{
        DeviceManagementConfigurationChoiceSettingDefinition: *NewDeviceManagementConfigurationChoiceSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingCollectionDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingCollectionDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingCollectionDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingCollectionDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationChoiceSettingDefinition.GetFieldDeserializers()
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
// GetMaximumCount gets the maximumCount property value. Maximum number of choices in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationChoiceSettingCollectionDefinition) GetMaximumCount()(*int32) {
    return m.maximumCount
}
// GetMinimumCount gets the minimumCount property value. Minimum number of choices in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationChoiceSettingCollectionDefinition) GetMinimumCount()(*int32) {
    return m.minimumCount
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingCollectionDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationChoiceSettingDefinition.Serialize(writer)
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
// SetMaximumCount sets the maximumCount property value. Maximum number of choices in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationChoiceSettingCollectionDefinition) SetMaximumCount(value *int32)() {
    m.maximumCount = value
}
// SetMinimumCount sets the minimumCount property value. Minimum number of choices in the collection. Valid values 1 to 100
func (m *DeviceManagementConfigurationChoiceSettingCollectionDefinition) SetMinimumCount(value *int32)() {
    m.minimumCount = value
}
