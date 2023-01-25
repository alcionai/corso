package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSVpnConfiguration 
type MacOSVpnConfiguration struct {
    AppleVpnConfiguration
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate MacOSCertificateProfileBaseable
}
// NewMacOSVpnConfiguration instantiates a new MacOSVpnConfiguration and sets the default values.
func NewMacOSVpnConfiguration()(*MacOSVpnConfiguration) {
    m := &MacOSVpnConfiguration{
        AppleVpnConfiguration: *NewAppleVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSVpnConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleVpnConfiguration.GetFieldDeserializers()
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(MacOSCertificateProfileBaseable))
        }
        return nil
    }
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *MacOSVpnConfiguration) GetIdentityCertificate()(MacOSCertificateProfileBaseable) {
    return m.identityCertificate
}
// Serialize serializes information the current object
func (m *MacOSVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppleVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *MacOSVpnConfiguration) SetIdentityCertificate(value MacOSCertificateProfileBaseable)() {
    m.identityCertificate = value
}
