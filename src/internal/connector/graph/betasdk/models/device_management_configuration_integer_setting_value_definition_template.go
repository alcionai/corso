package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate integer Setting Value Definition Template
type DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Integer Setting Maximum Value. Valid values -2147483648 to 2147483647
    maxValue *int32
    // Integer Setting Minimum Value. Valid values -2147483648 to 2147483647
    minValue *int32
    // The OdataType property
    odataType *string
}
// NewDeviceManagementConfigurationIntegerSettingValueDefinitionTemplate instantiates a new deviceManagementConfigurationIntegerSettingValueDefinitionTemplate and sets the default values.
func NewDeviceManagementConfigurationIntegerSettingValueDefinitionTemplate()(*DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) {
    m := &DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationIntegerSettingValueDefinitionTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationIntegerSettingValueDefinitionTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationIntegerSettingValueDefinitionTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxValue(val)
        }
        return nil
    }
    res["minValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinValue(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetMaxValue gets the maxValue property value. Integer Setting Maximum Value. Valid values -2147483648 to 2147483647
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) GetMaxValue()(*int32) {
    return m.maxValue
}
// GetMinValue gets the minValue property value. Integer Setting Minimum Value. Valid values -2147483648 to 2147483647
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) GetMinValue()(*int32) {
    return m.minValue
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maxValue", m.GetMaxValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minValue", m.GetMinValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMaxValue sets the maxValue property value. Integer Setting Maximum Value. Valid values -2147483648 to 2147483647
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) SetMaxValue(value *int32)() {
    m.maxValue = value
}
// SetMinValue sets the minValue property value. Integer Setting Minimum Value. Valid values -2147483648 to 2147483647
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) SetMinValue(value *int32)() {
    m.minValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
