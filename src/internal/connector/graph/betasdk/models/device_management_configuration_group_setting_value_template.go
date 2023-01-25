package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingValueTemplate group Setting Value Template
type DeviceManagementConfigurationGroupSettingValueTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Group setting value children
    children []DeviceManagementConfigurationSettingInstanceTemplateable
    // The OdataType property
    odataType *string
    // Setting Value Template Id
    settingValueTemplateId *string
}
// NewDeviceManagementConfigurationGroupSettingValueTemplate instantiates a new deviceManagementConfigurationGroupSettingValueTemplate and sets the default values.
func NewDeviceManagementConfigurationGroupSettingValueTemplate()(*DeviceManagementConfigurationGroupSettingValueTemplate) {
    m := &DeviceManagementConfigurationGroupSettingValueTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationGroupSettingValueTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationGroupSettingValueTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingValueTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetChildren gets the children property value. Group setting value children
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) GetChildren()([]DeviceManagementConfigurationSettingInstanceTemplateable) {
    return m.children
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingInstanceTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingInstanceTemplateable)
            }
            m.SetChildren(res)
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
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingValueTemplateId gets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) GetSettingValueTemplateId()(*string) {
    return m.settingValueTemplateId
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("children", cast)
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
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetChildren sets the children property value. Group setting value children
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) SetChildren(value []DeviceManagementConfigurationSettingInstanceTemplateable)() {
    m.children = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingValueTemplateId sets the settingValueTemplateId property value. Setting Value Template Id
func (m *DeviceManagementConfigurationGroupSettingValueTemplate) SetSettingValueTemplateId(value *string)() {
    m.settingValueTemplateId = value
}
