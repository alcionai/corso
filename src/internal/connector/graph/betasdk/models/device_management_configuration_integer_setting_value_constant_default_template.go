package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate 
type DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate struct {
    DeviceManagementConfigurationIntegerSettingValueDefaultTemplate
    // Default Constant Value. Valid values -2147483648 to 2147483647
    constantValue *int32
}
// NewDeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate instantiates a new DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate and sets the default values.
func NewDeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate()(*DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) {
    m := &DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate{
        DeviceManagementConfigurationIntegerSettingValueDefaultTemplate: *NewDeviceManagementConfigurationIntegerSettingValueDefaultTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate(), nil
}
// GetConstantValue gets the constantValue property value. Default Constant Value. Valid values -2147483648 to 2147483647
func (m *DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) GetConstantValue()(*int32) {
    return m.constantValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationIntegerSettingValueDefaultTemplate.GetFieldDeserializers()
    res["constantValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConstantValue(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationIntegerSettingValueDefaultTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("constantValue", m.GetConstantValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConstantValue sets the constantValue property value. Default Constant Value. Valid values -2147483648 to 2147483647
func (m *DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) SetConstantValue(value *int32)() {
    m.constantValue = value
}
