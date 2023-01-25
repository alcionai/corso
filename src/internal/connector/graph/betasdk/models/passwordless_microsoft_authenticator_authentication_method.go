package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordlessMicrosoftAuthenticatorAuthenticationMethod 
type PasswordlessMicrosoftAuthenticatorAuthenticationMethod struct {
    AuthenticationMethod
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The timestamp when this method was registered to the user.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The device property
    device Deviceable
    // The display name of the mobile device as given by the user.
    displayName *string
}
// NewPasswordlessMicrosoftAuthenticatorAuthenticationMethod instantiates a new PasswordlessMicrosoftAuthenticatorAuthenticationMethod and sets the default values.
func NewPasswordlessMicrosoftAuthenticatorAuthenticationMethod()(*PasswordlessMicrosoftAuthenticatorAuthenticationMethod) {
    m := &PasswordlessMicrosoftAuthenticatorAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    odataTypeValue := "#microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethod";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePasswordlessMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordlessMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordlessMicrosoftAuthenticatorAuthenticationMethod(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreationDateTime gets the creationDateTime property value. The timestamp when this method was registered to the user.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetDevice gets the device property value. The device property
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetDevice()(Deviceable) {
    return m.device
}
// GetDisplayName gets the displayName property value. The display name of the mobile device as given by the user.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["creationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(Deviceable))
        }
        return nil
    }
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
    return res
}
// Serialize serializes information the current object
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreationDateTime sets the creationDateTime property value. The timestamp when this method was registered to the user.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetDevice sets the device property value. The device property
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetDevice(value Deviceable)() {
    m.device = value
}
// SetDisplayName sets the displayName property value. The display name of the mobile device as given by the user.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetDisplayName(value *string)() {
    m.displayName = value
}
