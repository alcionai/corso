package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhoneXAP 
type WindowsPhoneXAP struct {
    MobileLobApp
    // The identity version.
    identityVersion *string
    // The minimum operating system required for a Windows mobile app.
    minimumSupportedOperatingSystem WindowsMinimumOperatingSystemable
    // The Product Identifier.
    productIdentifier *string
}
// NewWindowsPhoneXAP instantiates a new WindowsPhoneXAP and sets the default values.
func NewWindowsPhoneXAP()(*WindowsPhoneXAP) {
    m := &WindowsPhoneXAP{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhoneXAP";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsPhoneXAPFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhoneXAPFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhoneXAP(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhoneXAP) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    res["productIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductIdentifier(val)
        }
        return nil
    }
    return res
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *WindowsPhoneXAP) GetIdentityVersion()(*string) {
    return m.identityVersion
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsPhoneXAP) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetProductIdentifier gets the productIdentifier property value. The Product Identifier.
func (m *WindowsPhoneXAP) GetProductIdentifier()(*string) {
    return m.productIdentifier
}
// Serialize serializes information the current object
func (m *WindowsPhoneXAP) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productIdentifier", m.GetProductIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *WindowsPhoneXAP) SetIdentityVersion(value *string)() {
    m.identityVersion = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsPhoneXAP) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetProductIdentifier sets the productIdentifier property value. The Product Identifier.
func (m *WindowsPhoneXAP) SetProductIdentifier(value *string)() {
    m.productIdentifier = value
}
