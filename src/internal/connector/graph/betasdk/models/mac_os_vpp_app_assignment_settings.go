package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppAppAssignmentSettings 
type MacOsVppAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // Whether or not to uninstall the app when device is removed from Intune.
    uninstallOnDeviceRemoval *bool
    // Whether or not to use device licensing.
    useDeviceLicensing *bool
}
// NewMacOsVppAppAssignmentSettings instantiates a new MacOsVppAppAssignmentSettings and sets the default values.
func NewMacOsVppAppAssignmentSettings()(*MacOsVppAppAssignmentSettings) {
    m := &MacOsVppAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.macOsVppAppAssignmentSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOsVppAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOsVppAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOsVppAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOsVppAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["uninstallOnDeviceRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUninstallOnDeviceRemoval(val)
        }
        return nil
    }
    res["useDeviceLicensing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseDeviceLicensing(val)
        }
        return nil
    }
    return res
}
// GetUninstallOnDeviceRemoval gets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *MacOsVppAppAssignmentSettings) GetUninstallOnDeviceRemoval()(*bool) {
    return m.uninstallOnDeviceRemoval
}
// GetUseDeviceLicensing gets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *MacOsVppAppAssignmentSettings) GetUseDeviceLicensing()(*bool) {
    return m.useDeviceLicensing
}
// Serialize serializes information the current object
func (m *MacOsVppAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("uninstallOnDeviceRemoval", m.GetUninstallOnDeviceRemoval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useDeviceLicensing", m.GetUseDeviceLicensing())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUninstallOnDeviceRemoval sets the uninstallOnDeviceRemoval property value. Whether or not to uninstall the app when device is removed from Intune.
func (m *MacOsVppAppAssignmentSettings) SetUninstallOnDeviceRemoval(value *bool)() {
    m.uninstallOnDeviceRemoval = value
}
// SetUseDeviceLicensing sets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *MacOsVppAppAssignmentSettings) SetUseDeviceLicensing(value *bool)() {
    m.useDeviceLicensing = value
}
