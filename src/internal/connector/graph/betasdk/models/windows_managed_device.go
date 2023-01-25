package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagedDevice 
type WindowsManagedDevice struct {
    ManagedDevice
}
// NewWindowsManagedDevice instantiates a new WindowsManagedDevice and sets the default values.
func NewWindowsManagedDevice()(*WindowsManagedDevice) {
    m := &WindowsManagedDevice{
        ManagedDevice: *NewManagedDevice(),
    }
    odataTypeValue := "#microsoft.graph.windowsManagedDevice";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagedDevice(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedDevice.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedDevice.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
