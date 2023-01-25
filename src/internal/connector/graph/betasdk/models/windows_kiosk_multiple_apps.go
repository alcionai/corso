package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskMultipleApps 
type WindowsKioskMultipleApps struct {
    WindowsKioskAppConfiguration
    // This setting allows access to Downloads folder in file explorer.
    allowAccessToDownloadsFolder *bool
    // These are the only Windows Store Apps that will be available to launch from the Start menu. This collection can contain a maximum of 128 elements.
    apps []WindowsKioskAppBaseable
    // This setting indicates that desktop apps are allowed. Default to true.
    disallowDesktopApps *bool
    // This setting allows the admin to specify whether the Task Bar is shown or not.
    showTaskBar *bool
    // Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
    startMenuLayoutXml []byte
}
// NewWindowsKioskMultipleApps instantiates a new WindowsKioskMultipleApps and sets the default values.
func NewWindowsKioskMultipleApps()(*WindowsKioskMultipleApps) {
    m := &WindowsKioskMultipleApps{
        WindowsKioskAppConfiguration: *NewWindowsKioskAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskMultipleApps";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskMultipleAppsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskMultipleAppsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskMultipleApps(), nil
}
// GetAllowAccessToDownloadsFolder gets the allowAccessToDownloadsFolder property value. This setting allows access to Downloads folder in file explorer.
func (m *WindowsKioskMultipleApps) GetAllowAccessToDownloadsFolder()(*bool) {
    return m.allowAccessToDownloadsFolder
}
// GetApps gets the apps property value. These are the only Windows Store Apps that will be available to launch from the Start menu. This collection can contain a maximum of 128 elements.
func (m *WindowsKioskMultipleApps) GetApps()([]WindowsKioskAppBaseable) {
    return m.apps
}
// GetDisallowDesktopApps gets the disallowDesktopApps property value. This setting indicates that desktop apps are allowed. Default to true.
func (m *WindowsKioskMultipleApps) GetDisallowDesktopApps()(*bool) {
    return m.disallowDesktopApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskMultipleApps) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppConfiguration.GetFieldDeserializers()
    res["allowAccessToDownloadsFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAccessToDownloadsFolder(val)
        }
        return nil
    }
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsKioskAppBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsKioskAppBaseable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsKioskAppBaseable)
            }
            m.SetApps(res)
        }
        return nil
    }
    res["disallowDesktopApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisallowDesktopApps(val)
        }
        return nil
    }
    res["showTaskBar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowTaskBar(val)
        }
        return nil
    }
    res["startMenuLayoutXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuLayoutXml(val)
        }
        return nil
    }
    return res
}
// GetShowTaskBar gets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsKioskMultipleApps) GetShowTaskBar()(*bool) {
    return m.showTaskBar
}
// GetStartMenuLayoutXml gets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsKioskMultipleApps) GetStartMenuLayoutXml()([]byte) {
    return m.startMenuLayoutXml
}
// Serialize serializes information the current object
func (m *WindowsKioskMultipleApps) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAccessToDownloadsFolder", m.GetAllowAccessToDownloadsFolder())
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disallowDesktopApps", m.GetDisallowDesktopApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showTaskBar", m.GetShowTaskBar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("startMenuLayoutXml", m.GetStartMenuLayoutXml())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAccessToDownloadsFolder sets the allowAccessToDownloadsFolder property value. This setting allows access to Downloads folder in file explorer.
func (m *WindowsKioskMultipleApps) SetAllowAccessToDownloadsFolder(value *bool)() {
    m.allowAccessToDownloadsFolder = value
}
// SetApps sets the apps property value. These are the only Windows Store Apps that will be available to launch from the Start menu. This collection can contain a maximum of 128 elements.
func (m *WindowsKioskMultipleApps) SetApps(value []WindowsKioskAppBaseable)() {
    m.apps = value
}
// SetDisallowDesktopApps sets the disallowDesktopApps property value. This setting indicates that desktop apps are allowed. Default to true.
func (m *WindowsKioskMultipleApps) SetDisallowDesktopApps(value *bool)() {
    m.disallowDesktopApps = value
}
// SetShowTaskBar sets the showTaskBar property value. This setting allows the admin to specify whether the Task Bar is shown or not.
func (m *WindowsKioskMultipleApps) SetShowTaskBar(value *bool)() {
    m.showTaskBar = value
}
// SetStartMenuLayoutXml sets the startMenuLayoutXml property value. Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
func (m *WindowsKioskMultipleApps) SetStartMenuLayoutXml(value []byte)() {
    m.startMenuLayoutXml = value
}
