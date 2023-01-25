package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingValueTemplate simple Setting Value Template
type DeviceManagementConfigurationSimpleSettingValueTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Setting Value Template Id
    settingValueTemplateId *string
}
// NewDeviceManagementConfigurationSimpleSettingValueTemplate instantiates a new deviceManagementConfigurationSimpleSettingValueTemplate and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingValueTemplate()(*DeviceManagementConfigurationSimpleSettingValueTemplate) {
    m := &DeviceManagementConfigurationSimpleSettingValueTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingValueTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingValueTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueTemplate":
                        return NewDeviceManagementConfigurationIntegerSettingValueTemplate(), nil
                    case "#microsoft.graph.deviceManagementConfigurationStringSettingValueTemplate":
                        return NewDeviceManagementConfigurationStringSettingValueTemplate(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSimpleSettingValueTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["settingValueTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueTemplateId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingValueTemplateId gets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) GetSettingValueTemplateId()(*string) {
    return m.settingValueTemplateId
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingValueTemplateId", m.GetSettingValueTemplateId())
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
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingValueTemplateId sets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationSimpleSettingValueTemplate) SetSettingValueTemplateId(value *string)() {
    m.settingValueTemplateId = value
}
