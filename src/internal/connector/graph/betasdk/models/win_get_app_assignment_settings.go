package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WinGetAppAssignmentSettings 
type WinGetAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // The install time settings to apply for this app assignment.
    installTimeSettings WinGetAppInstallTimeSettingsable
    // Contains value for notification status.
    notifications *WinGetAppNotification
    // The reboot settings to apply for this app assignment.
    restartSettings WinGetAppRestartSettingsable
}
// NewWinGetAppAssignmentSettings instantiates a new WinGetAppAssignmentSettings and sets the default values.
func NewWinGetAppAssignmentSettings()(*WinGetAppAssignmentSettings) {
    m := &WinGetAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.winGetAppAssignmentSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWinGetAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWinGetAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWinGetAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WinGetAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["installTimeSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWinGetAppInstallTimeSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallTimeSettings(val.(WinGetAppInstallTimeSettingsable))
        }
        return nil
    }
    res["notifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWinGetAppNotification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifications(val.(*WinGetAppNotification))
        }
        return nil
    }
    res["restartSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWinGetAppRestartSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartSettings(val.(WinGetAppRestartSettingsable))
        }
        return nil
    }
    return res
}
// GetInstallTimeSettings gets the installTimeSettings property value. The install time settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) GetInstallTimeSettings()(WinGetAppInstallTimeSettingsable) {
    return m.installTimeSettings
}
// GetNotifications gets the notifications property value. Contains value for notification status.
func (m *WinGetAppAssignmentSettings) GetNotifications()(*WinGetAppNotification) {
    return m.notifications
}
// GetRestartSettings gets the restartSettings property value. The reboot settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) GetRestartSettings()(WinGetAppRestartSettingsable) {
    return m.restartSettings
}
// Serialize serializes information the current object
func (m *WinGetAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("installTimeSettings", m.GetInstallTimeSettings())
        if err != nil {
            return err
        }
    }
    if m.GetNotifications() != nil {
        cast := (*m.GetNotifications()).String()
        err = writer.WriteStringValue("notifications", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("restartSettings", m.GetRestartSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInstallTimeSettings sets the installTimeSettings property value. The install time settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) SetInstallTimeSettings(value WinGetAppInstallTimeSettingsable)() {
    m.installTimeSettings = value
}
// SetNotifications sets the notifications property value. Contains value for notification status.
func (m *WinGetAppAssignmentSettings) SetNotifications(value *WinGetAppNotification)() {
    m.notifications = value
}
// SetRestartSettings sets the restartSettings property value. The reboot settings to apply for this app assignment.
func (m *WinGetAppAssignmentSettings) SetRestartSettings(value WinGetAppRestartSettingsable)() {
    m.restartSettings = value
}
