package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeApp 
type AndroidDeviceOwnerKioskModeApp struct {
    AndroidDeviceOwnerKioskModeFolderItem
    // Class name of application
    className *string
    // Package name of application
    package_escaped *string
}
// NewAndroidDeviceOwnerKioskModeApp instantiates a new AndroidDeviceOwnerKioskModeApp and sets the default values.
func NewAndroidDeviceOwnerKioskModeApp()(*AndroidDeviceOwnerKioskModeApp) {
    m := &AndroidDeviceOwnerKioskModeApp{
        AndroidDeviceOwnerKioskModeFolderItem: *NewAndroidDeviceOwnerKioskModeFolderItem(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerKioskModeApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerKioskModeAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeApp(), nil
}
// GetClassName gets the className property value. Class name of application
func (m *AndroidDeviceOwnerKioskModeApp) GetClassName()(*string) {
    return m.className
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerKioskModeFolderItem.GetFieldDeserializers()
    res["className"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassName(val)
        }
        return nil
    }
    res["package"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackage(val)
        }
        return nil
    }
    return res
}
// GetPackage gets the package property value. Package name of application
func (m *AndroidDeviceOwnerKioskModeApp) GetPackage()(*string) {
    return m.package_escaped
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerKioskModeFolderItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("className", m.GetClassName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("package", m.GetPackage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassName sets the className property value. Class name of application
func (m *AndroidDeviceOwnerKioskModeApp) SetClassName(value *string)() {
    m.className = value
}
// SetPackage sets the package property value. Package name of application
func (m *AndroidDeviceOwnerKioskModeApp) SetPackage(value *string)() {
    m.package_escaped = value
}
