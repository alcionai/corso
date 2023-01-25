package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingInstance setting instance within policy
type DeviceManagementConfigurationSettingInstance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Setting Definition Id
    settingDefinitionId *string
    // Setting Instance Template Reference
    settingInstanceTemplateReference DeviceManagementConfigurationSettingInstanceTemplateReferenceable
}
// NewDeviceManagementConfigurationSettingInstance instantiates a new deviceManagementConfigurationSettingInstance and sets the default values.
func NewDeviceManagementConfigurationSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    m := &DeviceManagementConfigurationSettingInstance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstance":
                        return NewDeviceManagementConfigurationChoiceSettingCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance":
                        return NewDeviceManagementConfigurationChoiceSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstance":
                        return NewDeviceManagementConfigurationGroupSettingCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationGroupSettingInstance":
                        return NewDeviceManagementConfigurationGroupSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionInstance":
                        return NewDeviceManagementConfigurationSettingGroupCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingGroupInstance":
                        return NewDeviceManagementConfigurationSettingGroupInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstance":
                        return NewDeviceManagementConfigurationSimpleSettingCollectionInstance(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance":
                        return NewDeviceManagementConfigurationSimpleSettingInstance(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingInstance(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstance) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["settingInstanceTemplateReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingInstanceTemplateReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplateReference(val.(DeviceManagementConfigurationSettingInstanceTemplateReferenceable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingInstance) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingDefinitionId gets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstance) GetSettingDefinitionId()(*string) {
    return m.settingDefinitionId
}
// GetSettingInstanceTemplateReference gets the settingInstanceTemplateReference property value. Setting Instance Template Reference
func (m *DeviceManagementConfigurationSettingInstance) GetSettingInstanceTemplateReference()(DeviceManagementConfigurationSettingInstanceTemplateReferenceable) {
    return m.settingInstanceTemplateReference
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("settingInstanceTemplateReference", m.GetSettingInstanceTemplateReference())
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
func (m *DeviceManagementConfigurationSettingInstance) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingInstance) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingDefinitionId sets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstance) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
// SetSettingInstanceTemplateReference sets the settingInstanceTemplateReference property value. Setting Instance Template Reference
func (m *DeviceManagementConfigurationSettingInstance) SetSettingInstanceTemplateReference(value DeviceManagementConfigurationSettingInstanceTemplateReferenceable)() {
    m.settingInstanceTemplateReference = value
}
