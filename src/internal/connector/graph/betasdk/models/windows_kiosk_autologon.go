package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskAutologon 
type WindowsKioskAutologon struct {
    WindowsKioskUser
}
// NewWindowsKioskAutologon instantiates a new WindowsKioskAutologon and sets the default values.
func NewWindowsKioskAutologon()(*WindowsKioskAutologon) {
    m := &WindowsKioskAutologon{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskAutologon";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskAutologonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskAutologonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskAutologon(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskAutologon) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsKioskAutologon) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
