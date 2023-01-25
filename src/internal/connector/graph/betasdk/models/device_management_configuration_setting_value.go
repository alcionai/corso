package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingValue setting value
type DeviceManagementConfigurationSettingValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Setting value template reference
    settingValueTemplateReference DeviceManagementConfigurationSettingValueTemplateReferenceable
}
// NewDeviceManagementConfigurationSettingValue instantiates a new deviceManagementConfigurationSettingValue and sets the default values.
func NewDeviceManagementConfigurationSettingValue()(*DeviceManagementConfigurationSettingValue) {
    m := &DeviceManagementConfigurationSettingValue{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingValue":
                        return NewDeviceManagementConfigurationChoiceSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationGroupSettingValue":
                        return NewDeviceManagementConfigurationGroupSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationIntegerSettingValue":
                        return NewDeviceManagementConfigurationIntegerSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationReferenceSettingValue":
                        return NewDeviceManagementConfigurationReferenceSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSecretSettingValue":
                        return NewDeviceManagementConfigurationSecretSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingValue":
                        return NewDeviceManagementConfigurationSimpleSettingValue(), nil
                    case "#microsoft.graph.deviceManagementConfigurationStringSettingValue":
                        return NewDeviceManagementConfigurationStringSettingValue(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingValue(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingValue) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["settingValueTemplateReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingValueTemplateReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueTemplateReference(val.(DeviceManagementConfigurationSettingValueTemplateReferenceable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingValue) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingValueTemplateReference gets the settingValueTemplateReference property value. Setting value template reference
func (m *DeviceManagementConfigurationSettingValue) GetSettingValueTemplateReference()(DeviceManagementConfigurationSettingValueTemplateReferenceable) {
    return m.settingValueTemplateReference
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settingValueTemplateReference", m.GetSettingValueTemplateReference())
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
func (m *DeviceManagementConfigurationSettingValue) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingValue) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingValueTemplateReference sets the settingValueTemplateReference property value. Setting value template reference
func (m *DeviceManagementConfigurationSettingValue) SetSettingValueTemplateReference(value DeviceManagementConfigurationSettingValueTemplateReferenceable)() {
    m.settingValueTemplateReference = value
}
