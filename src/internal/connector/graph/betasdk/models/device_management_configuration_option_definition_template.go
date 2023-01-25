package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationOptionDefinitionTemplate option Definition Template
type DeviceManagementConfigurationOptionDefinitionTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Option Children
    children []DeviceManagementConfigurationSettingInstanceTemplateable
    // Option ItemId
    itemId *string
    // The OdataType property
    odataType *string
}
// NewDeviceManagementConfigurationOptionDefinitionTemplate instantiates a new deviceManagementConfigurationOptionDefinitionTemplate and sets the default values.
func NewDeviceManagementConfigurationOptionDefinitionTemplate()(*DeviceManagementConfigurationOptionDefinitionTemplate) {
    m := &DeviceManagementConfigurationOptionDefinitionTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationOptionDefinitionTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationOptionDefinitionTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationOptionDefinitionTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetChildren gets the children property value. Option Children
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) GetChildren()([]DeviceManagementConfigurationSettingInstanceTemplateable) {
    return m.children
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
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
// GetItemId gets the itemId property value. Option ItemId
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) GetItemId()(*string) {
    return m.itemId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("itemId", m.GetItemId())
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
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetChildren sets the children property value. Option Children
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) SetChildren(value []DeviceManagementConfigurationSettingInstanceTemplateable)() {
    m.children = value
}
// SetItemId sets the itemId property value. Option ItemId
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) SetItemId(value *string)() {
    m.itemId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationOptionDefinitionTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
