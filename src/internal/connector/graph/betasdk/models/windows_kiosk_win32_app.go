package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskWin32App 
type WindowsKioskWin32App struct {
    WindowsKioskAppBase
    // This is the classicapppath to be used by v4 Win32 app while in Kiosk Mode
    classicAppPath *string
    // Edge kiosk (url) for Edge kiosk mode
    edgeKiosk *string
    // Edge kiosk idle timeout in minutes for Edge kiosk mode. Valid values 0 to 1440
    edgeKioskIdleTimeoutMinutes *int32
    // Edge kiosk type
    edgeKioskType *WindowsEdgeKioskType
    // Edge first run flag for Edge kiosk mode
    edgeNoFirstRun *bool
}
// NewWindowsKioskWin32App instantiates a new WindowsKioskWin32App and sets the default values.
func NewWindowsKioskWin32App()(*WindowsKioskWin32App) {
    m := &WindowsKioskWin32App{
        WindowsKioskAppBase: *NewWindowsKioskAppBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskWin32App";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskWin32AppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskWin32AppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskWin32App(), nil
}
// GetClassicAppPath gets the classicAppPath property value. This is the classicapppath to be used by v4 Win32 app while in Kiosk Mode
func (m *WindowsKioskWin32App) GetClassicAppPath()(*string) {
    return m.classicAppPath
}
// GetEdgeKiosk gets the edgeKiosk property value. Edge kiosk (url) for Edge kiosk mode
func (m *WindowsKioskWin32App) GetEdgeKiosk()(*string) {
    return m.edgeKiosk
}
// GetEdgeKioskIdleTimeoutMinutes gets the edgeKioskIdleTimeoutMinutes property value. Edge kiosk idle timeout in minutes for Edge kiosk mode. Valid values 0 to 1440
func (m *WindowsKioskWin32App) GetEdgeKioskIdleTimeoutMinutes()(*int32) {
    return m.edgeKioskIdleTimeoutMinutes
}
// GetEdgeKioskType gets the edgeKioskType property value. Edge kiosk type
func (m *WindowsKioskWin32App) GetEdgeKioskType()(*WindowsEdgeKioskType) {
    return m.edgeKioskType
}
// GetEdgeNoFirstRun gets the edgeNoFirstRun property value. Edge first run flag for Edge kiosk mode
func (m *WindowsKioskWin32App) GetEdgeNoFirstRun()(*bool) {
    return m.edgeNoFirstRun
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskWin32App) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskAppBase.GetFieldDeserializers()
    res["classicAppPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassicAppPath(val)
        }
        return nil
    }
    res["edgeKiosk"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeKiosk(val)
        }
        return nil
    }
    res["edgeKioskIdleTimeoutMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeKioskIdleTimeoutMinutes(val)
        }
        return nil
    }
    res["edgeKioskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsEdgeKioskType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeKioskType(val.(*WindowsEdgeKioskType))
        }
        return nil
    }
    res["edgeNoFirstRun"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeNoFirstRun(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsKioskWin32App) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskAppBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("classicAppPath", m.GetClassicAppPath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("edgeKiosk", m.GetEdgeKiosk())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("edgeKioskIdleTimeoutMinutes", m.GetEdgeKioskIdleTimeoutMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetEdgeKioskType() != nil {
        cast := (*m.GetEdgeKioskType()).String()
        err = writer.WriteStringValue("edgeKioskType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeNoFirstRun", m.GetEdgeNoFirstRun())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassicAppPath sets the classicAppPath property value. This is the classicapppath to be used by v4 Win32 app while in Kiosk Mode
func (m *WindowsKioskWin32App) SetClassicAppPath(value *string)() {
    m.classicAppPath = value
}
// SetEdgeKiosk sets the edgeKiosk property value. Edge kiosk (url) for Edge kiosk mode
func (m *WindowsKioskWin32App) SetEdgeKiosk(value *string)() {
    m.edgeKiosk = value
}
// SetEdgeKioskIdleTimeoutMinutes sets the edgeKioskIdleTimeoutMinutes property value. Edge kiosk idle timeout in minutes for Edge kiosk mode. Valid values 0 to 1440
func (m *WindowsKioskWin32App) SetEdgeKioskIdleTimeoutMinutes(value *int32)() {
    m.edgeKioskIdleTimeoutMinutes = value
}
// SetEdgeKioskType sets the edgeKioskType property value. Edge kiosk type
func (m *WindowsKioskWin32App) SetEdgeKioskType(value *WindowsEdgeKioskType)() {
    m.edgeKioskType = value
}
// SetEdgeNoFirstRun sets the edgeNoFirstRun property value. Edge first run flag for Edge kiosk mode
func (m *WindowsKioskWin32App) SetEdgeNoFirstRun(value *bool)() {
    m.edgeNoFirstRun = value
}
