package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingCategory entity representing a setting category
type DeviceManagementSettingCategory struct {
    Entity
    // The category name
    displayName *string
    // The category contains top level required setting
    hasRequiredSetting *bool
    // The setting definitions this category contains
    settingDefinitions []DeviceManagementSettingDefinitionable
}
// NewDeviceManagementSettingCategory instantiates a new deviceManagementSettingCategory and sets the default values.
func NewDeviceManagementSettingCategory()(*DeviceManagementSettingCategory) {
    m := &DeviceManagementSettingCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementSettingCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementIntentSettingCategory":
                        return NewDeviceManagementIntentSettingCategory(), nil
                    case "#microsoft.graph.deviceManagementTemplateSettingCategory":
                        return NewDeviceManagementTemplateSettingCategory(), nil
                }
            }
        }
    }
    return NewDeviceManagementSettingCategory(), nil
}
// GetDisplayName gets the displayName property value. The category name
func (m *DeviceManagementSettingCategory) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["hasRequiredSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasRequiredSetting(val)
        }
        return nil
    }
    res["settingDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingDefinitionable)
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    return res
}
// GetHasRequiredSetting gets the hasRequiredSetting property value. The category contains top level required setting
func (m *DeviceManagementSettingCategory) GetHasRequiredSetting()(*bool) {
    return m.hasRequiredSetting
}
// GetSettingDefinitions gets the settingDefinitions property value. The setting definitions this category contains
func (m *DeviceManagementSettingCategory) GetSettingDefinitions()([]DeviceManagementSettingDefinitionable) {
    return m.settingDefinitions
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasRequiredSetting", m.GetHasRequiredSetting())
        if err != nil {
            return err
        }
    }
    if m.GetSettingDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingDefinitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The category name
func (m *DeviceManagementSettingCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHasRequiredSetting sets the hasRequiredSetting property value. The category contains top level required setting
func (m *DeviceManagementSettingCategory) SetHasRequiredSetting(value *bool)() {
    m.hasRequiredSetting = value
}
// SetSettingDefinitions sets the settingDefinitions property value. The setting definitions this category contains
func (m *DeviceManagementSettingCategory) SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)() {
    m.settingDefinitions = value
}
