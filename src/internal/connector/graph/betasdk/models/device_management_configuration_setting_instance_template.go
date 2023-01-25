package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingInstanceTemplate setting Instance Template
type DeviceManagementConfigurationSettingInstanceTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates if a policy must specify this setting.
    isRequired *bool
    // The OdataType property
    odataType *string
    // Setting Definition Id
    settingDefinitionId *string
    // Setting Instance Template Id
    settingInstanceTemplateId *string
}
// NewDeviceManagementConfigurationSettingInstanceTemplate instantiates a new deviceManagementConfigurationSettingInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationSettingInstanceTemplate()(*DeviceManagementConfigurationSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationSettingInstanceTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstanceTemplate":
                        return NewDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate(), nil
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstanceTemplate":
                        return NewDeviceManagementConfigurationChoiceSettingInstanceTemplate(), nil
                    case "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstanceTemplate":
                        return NewDeviceManagementConfigurationGroupSettingCollectionInstanceTemplate(), nil
                    case "#microsoft.graph.deviceManagementConfigurationGroupSettingInstanceTemplate":
                        return NewDeviceManagementConfigurationGroupSettingInstanceTemplate(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstanceTemplate":
                        return NewDeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstanceTemplate":
                        return NewDeviceManagementConfigurationSimpleSettingInstanceTemplate(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingInstanceTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
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
    res["settingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionId(val)
        }
        return nil
    }
    res["settingInstanceTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplateId(val)
        }
        return nil
    }
    return res
}
// GetIsRequired gets the isRequired property value. Indicates if a policy must specify this setting.
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetIsRequired()(*bool) {
    return m.isRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingDefinitionId gets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetSettingDefinitionId()(*string) {
    return m.settingDefinitionId
}
// GetSettingInstanceTemplateId gets the settingInstanceTemplateId property value. Setting Instance Template Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetSettingInstanceTemplateId()(*string) {
    return m.settingInstanceTemplateId
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
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
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingInstanceTemplateId", m.GetSettingInstanceTemplateId())
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
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsRequired sets the isRequired property value. Indicates if a policy must specify this setting.
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingDefinitionId sets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
// SetSettingInstanceTemplateId sets the settingInstanceTemplateId property value. Setting Instance Template Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetSettingInstanceTemplateId(value *string)() {
    m.settingInstanceTemplateId = value
}
