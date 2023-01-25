package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCustomAttributeShellScriptCollectionResponse 
type DeviceCustomAttributeShellScriptCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceCustomAttributeShellScriptable
}
// NewDeviceCustomAttributeShellScriptCollectionResponse instantiates a new DeviceCustomAttributeShellScriptCollectionResponse and sets the default values.
func NewDeviceCustomAttributeShellScriptCollectionResponse()(*DeviceCustomAttributeShellScriptCollectionResponse) {
    m := &DeviceCustomAttributeShellScriptCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceCustomAttributeShellScriptCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCustomAttributeShellScriptCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCustomAttributeShellScriptCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCustomAttributeShellScriptCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCustomAttributeShellScriptable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCustomAttributeShellScriptable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceCustomAttributeShellScriptCollectionResponse) GetValue()([]DeviceCustomAttributeShellScriptable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceCustomAttributeShellScriptCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceCustomAttributeShellScriptCollectionResponse) SetValue(value []DeviceCustomAttributeShellScriptable)() {
    m.value = value
}
