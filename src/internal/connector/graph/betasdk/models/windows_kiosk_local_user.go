package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskLocalUser 
type WindowsKioskLocalUser struct {
    WindowsKioskUser
    // The local user that will be locked to this kiosk configuration
    userName *string
}
// NewWindowsKioskLocalUser instantiates a new WindowsKioskLocalUser and sets the default values.
func NewWindowsKioskLocalUser()(*WindowsKioskLocalUser) {
    m := &WindowsKioskLocalUser{
        WindowsKioskUser: *NewWindowsKioskUser(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskLocalUser";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskLocalUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskLocalUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskLocalUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskLocalUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsKioskUser.GetFieldDeserializers()
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetUserName gets the userName property value. The local user that will be locked to this kiosk configuration
func (m *WindowsKioskLocalUser) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *WindowsKioskLocalUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsKioskUser.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserName sets the userName property value. The local user that will be locked to this kiosk configuration
func (m *WindowsKioskLocalUser) SetUserName(value *string)() {
    m.userName = value
}
