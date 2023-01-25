package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingValueCollectionResponse 
type DeviceManagementConfigurationSimpleSettingValueCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceManagementConfigurationSimpleSettingValueable
}
// NewDeviceManagementConfigurationSimpleSettingValueCollectionResponse instantiates a new DeviceManagementConfigurationSimpleSettingValueCollectionResponse and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingValueCollectionResponse()(*DeviceManagementConfigurationSimpleSettingValueCollectionResponse) {
    m := &DeviceManagementConfigurationSimpleSettingValueCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingValueCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingValueCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingValueCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingValueCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSimpleSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSimpleSettingValueable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSimpleSettingValueable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceManagementConfigurationSimpleSettingValueCollectionResponse) GetValue()([]DeviceManagementConfigurationSimpleSettingValueable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingValueCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DeviceManagementConfigurationSimpleSettingValueCollectionResponse) SetValue(value []DeviceManagementConfigurationSimpleSettingValueable)() {
    m.value = value
}
