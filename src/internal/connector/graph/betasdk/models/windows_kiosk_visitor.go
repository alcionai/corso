package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskVisitor 
type WindowsKioskVisitor struct {
    WindowsKioskUser
}
// NewWindowsKioskVisitor instantiates a new WindowsKioskVisitor and sets the default values.
func NewWindowsKioskVisitor()(*WindowsKioskVisitor) {
    m := &WindowsKioskVisitor{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskVisitor";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskVisitorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskVisitorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskVisitor(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskVisitor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsKioskVisitor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
