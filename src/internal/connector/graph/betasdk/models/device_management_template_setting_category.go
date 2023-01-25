package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTemplateSettingCategory 
type DeviceManagementTemplateSettingCategory struct {
    DeviceManagementSettingCategory
    // The settings this category contains
    recommendedSettings []DeviceManagementSettingInstanceable
}
// NewDeviceManagementTemplateSettingCategory instantiates a new DeviceManagementTemplateSettingCategory and sets the default values.
func NewDeviceManagementTemplateSettingCategory()(*DeviceManagementTemplateSettingCategory) {
    m := &DeviceManagementTemplateSettingCategory{
        DeviceManagementSettingCategory: *NewDeviceManagementSettingCategory(),
    }
    return m
}
// CreateDeviceManagementTemplateSettingCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTemplateSettingCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementTemplateSettingCategory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTemplateSettingCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingCategory.GetFieldDeserializers()
    res["recommendedSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingInstanceable)
            }
            m.SetRecommendedSettings(res)
        }
        return nil
    }
    return res
}
// GetRecommendedSettings gets the recommendedSettings property value. The settings this category contains
func (m *DeviceManagementTemplateSettingCategory) GetRecommendedSettings()([]DeviceManagementSettingInstanceable) {
    return m.recommendedSettings
}
// Serialize serializes information the current object
func (m *DeviceManagementTemplateSettingCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingCategory.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRecommendedSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecommendedSettings()))
        for i, v := range m.GetRecommendedSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("recommendedSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecommendedSettings sets the recommendedSettings property value. The settings this category contains
func (m *DeviceManagementTemplateSettingCategory) SetRecommendedSettings(value []DeviceManagementSettingInstanceable)() {
    m.recommendedSettings = value
}
