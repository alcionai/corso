package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskSingleWin32App 
type WindowsKioskSingleWin32App struct {
    WindowsKioskAppConfiguration
    // The win32App property
    win32App WindowsKioskWin32Appable
}
// NewWindowsKioskSingleWin32App instantiates a new WindowsKioskSingleWin32App and sets the default values.
func NewWindowsKioskSingleWin32App()(*WindowsKioskSingleWin32App) {
    m := &WindowsKioskSingleWin32App{
        WindowsKioskAppConfiguration: *NewWindowsKioskAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskSingleWin32App";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskSingleWin32AppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskSingleWin32AppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskSingleWin32App(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskSingleWin32App) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppConfiguration.GetFieldDeserializers()
    res["win32App"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsKioskWin32AppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWin32App(val.(WindowsKioskWin32Appable))
        }
        return nil
    }
    return res
}
// GetWin32App gets the win32App property value. The win32App property
func (m *WindowsKioskSingleWin32App) GetWin32App()(WindowsKioskWin32Appable) {
    return m.win32App
}
// Serialize serializes information the current object
func (m *WindowsKioskSingleWin32App) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("win32App", m.GetWin32App())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWin32App sets the win32App property value. The win32App property
func (m *WindowsKioskSingleWin32App) SetWin32App(value WindowsKioskWin32Appable)() {
    m.win32App = value
}
