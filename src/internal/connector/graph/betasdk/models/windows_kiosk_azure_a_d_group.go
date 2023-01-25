package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskAzureADGroup 
type WindowsKioskAzureADGroup struct {
    WindowsKioskUser
    // The display name of the AzureAD group that will be locked to this kiosk configuration
    displayName *string
    // The ID of the AzureAD group that will be locked to this kiosk configuration
    groupId *string
}
// NewWindowsKioskAzureADGroup instantiates a new WindowsKioskAzureADGroup and sets the default values.
func NewWindowsKioskAzureADGroup()(*WindowsKioskAzureADGroup) {
    m := &WindowsKioskAzureADGroup{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskAzureADGroup";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskAzureADGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskAzureADGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskAzureADGroup(), nil
}
// GetDisplayName gets the displayName property value. The display name of the AzureAD group that will be locked to this kiosk configuration
func (m *WindowsKioskAzureADGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskAzureADGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The ID of the AzureAD group that will be locked to this kiosk configuration
func (m *WindowsKioskAzureADGroup) GetGroupId()(*string) {
    return m.groupId
}
// Serialize serializes information the current object
func (m *WindowsKioskAzureADGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the AzureAD group that will be locked to this kiosk configuration
func (m *WindowsKioskAzureADGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupId sets the groupId property value. The ID of the AzureAD group that will be locked to this kiosk configuration
func (m *WindowsKioskAzureADGroup) SetGroupId(value *string)() {
    m.groupId = value
}
