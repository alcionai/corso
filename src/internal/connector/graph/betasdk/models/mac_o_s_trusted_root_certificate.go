package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSTrustedRootCertificate 
type MacOSTrustedRootCertificate struct {
    DeviceConfiguration
    // File name to display in UI.
    certFileName *string
    // Trusted Root Certificate.
    trustedRootCertificate []byte
}
// NewMacOSTrustedRootCertificate instantiates a new macOSTrustedRootCertificate and sets the default values.
func NewMacOSTrustedRootCertificate()(*MacOSTrustedRootCertificate) {
    m := &MacOSTrustedRootCertificate{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSTrustedRootCertificate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSTrustedRootCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSTrustedRootCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSTrustedRootCertificate(), nil
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
func (m *MacOSTrustedRootCertificate) GetCertFileName()(*string) {
    return m.certFileName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSTrustedRootCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["certFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertFileName(val)
        }
        return nil
    }
    res["trustedRootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustedRootCertificate(val)
        }
        return nil
    }
    return res
}
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate.
func (m *MacOSTrustedRootCertificate) GetTrustedRootCertificate()([]byte) {
    return m.trustedRootCertificate
}
// Serialize serializes information the current object
func (m *MacOSTrustedRootCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certFileName", m.GetCertFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("trustedRootCertificate", m.GetTrustedRootCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertFileName sets the certFileName property value. File name to display in UI.
func (m *MacOSTrustedRootCertificate) SetCertFileName(value *string)() {
    m.certFileName = value
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate.
func (m *MacOSTrustedRootCertificate) SetTrustedRootCertificate(value []byte)() {
    m.trustedRootCertificate = value
}
