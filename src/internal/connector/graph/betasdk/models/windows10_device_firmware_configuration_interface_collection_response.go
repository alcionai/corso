package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse 
type Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Windows10DeviceFirmwareConfigurationInterfaceable
}
// NewWindows10DeviceFirmwareConfigurationInterfaceCollectionResponse instantiates a new Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse and sets the default values.
func NewWindows10DeviceFirmwareConfigurationInterfaceCollectionResponse()(*Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse) {
    m := &Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindows10DeviceFirmwareConfigurationInterfaceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10DeviceFirmwareConfigurationInterfaceCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10DeviceFirmwareConfigurationInterfaceCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10DeviceFirmwareConfigurationInterfaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10DeviceFirmwareConfigurationInterfaceable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10DeviceFirmwareConfigurationInterfaceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse) GetValue()([]Windows10DeviceFirmwareConfigurationInterfaceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10DeviceFirmwareConfigurationInterfaceCollectionResponse) SetValue(value []Windows10DeviceFirmwareConfigurationInterfaceable)() {
    m.value = value
}
