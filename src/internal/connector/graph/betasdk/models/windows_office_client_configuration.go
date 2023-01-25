package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsOfficeClientConfiguration 
type WindowsOfficeClientConfiguration struct {
    OfficeClientConfiguration
}
// NewWindowsOfficeClientConfiguration instantiates a new WindowsOfficeClientConfiguration and sets the default values.
func NewWindowsOfficeClientConfiguration()(*WindowsOfficeClientConfiguration) {
    m := &WindowsOfficeClientConfiguration{
        OfficeClientConfiguration: *NewOfficeClientConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsOfficeClientConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsOfficeClientConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsOfficeClientConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsOfficeClientConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsOfficeClientConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OfficeClientConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsOfficeClientConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OfficeClientConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
