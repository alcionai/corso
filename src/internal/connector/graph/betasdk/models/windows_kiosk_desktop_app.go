package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskDesktopApp 
type WindowsKioskDesktopApp struct {
    WindowsKioskAppBase
    // Define the DesktopApplicationID of the app
    desktopApplicationId *string
    // Define the DesktopApplicationLinkPath of the app
    desktopApplicationLinkPath *string
    // Define the path of a desktop app
    path *string
}
// NewWindowsKioskDesktopApp instantiates a new WindowsKioskDesktopApp and sets the default values.
func NewWindowsKioskDesktopApp()(*WindowsKioskDesktopApp) {
    m := &WindowsKioskDesktopApp{
        WindowsKioskAppBase: *NewWindowsKioskAppBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskDesktopApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskDesktopAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskDesktopAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskDesktopApp(), nil
}
// GetDesktopApplicationId gets the desktopApplicationId property value. Define the DesktopApplicationID of the app
func (m *WindowsKioskDesktopApp) GetDesktopApplicationId()(*string) {
    return m.desktopApplicationId
}
// GetDesktopApplicationLinkPath gets the desktopApplicationLinkPath property value. Define the DesktopApplicationLinkPath of the app
func (m *WindowsKioskDesktopApp) GetDesktopApplicationLinkPath()(*string) {
    return m.desktopApplicationLinkPath
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskDesktopApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppBase.GetFieldDeserializers()
    res["desktopApplicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDesktopApplicationId(val)
        }
        return nil
    }
    res["desktopApplicationLinkPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDesktopApplicationLinkPath(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    return res
}
// GetPath gets the path property value. Define the path of a desktop app
func (m *WindowsKioskDesktopApp) GetPath()(*string) {
    return m.path
}
// Serialize serializes information the current object
func (m *WindowsKioskDesktopApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("desktopApplicationId", m.GetDesktopApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("desktopApplicationLinkPath", m.GetDesktopApplicationLinkPath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDesktopApplicationId sets the desktopApplicationId property value. Define the DesktopApplicationID of the app
func (m *WindowsKioskDesktopApp) SetDesktopApplicationId(value *string)() {
    m.desktopApplicationId = value
}
// SetDesktopApplicationLinkPath sets the desktopApplicationLinkPath property value. Define the DesktopApplicationLinkPath of the app
func (m *WindowsKioskDesktopApp) SetDesktopApplicationLinkPath(value *string)() {
    m.desktopApplicationLinkPath = value
}
// SetPath sets the path property value. Define the path of a desktop app
func (m *WindowsKioskDesktopApp) SetPath(value *string)() {
    m.path = value
}
