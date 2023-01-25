package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81TrustedRootCertificate 
type Windows81TrustedRootCertificate struct {
    DeviceConfiguration
    // File name to display in UI.
    certFileName *string
    // Possible values for the Certificate Destination Store.
    destinationStore *CertificateDestinationStore
    // Trusted Root Certificate
    trustedRootCertificate []byte
}
// NewWindows81TrustedRootCertificate instantiates a new windows81TrustedRootCertificate and sets the default values.
func NewWindows81TrustedRootCertificate()(*Windows81TrustedRootCertificate) {
    m := &Windows81TrustedRootCertificate{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows81TrustedRootCertificate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows81TrustedRootCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81TrustedRootCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81TrustedRootCertificate(), nil
}
// GetCertFileName gets the certFileName property value. File name to display in UI.
func (m *Windows81TrustedRootCertificate) GetCertFileName()(*string) {
    return m.certFileName
}
// GetDestinationStore gets the destinationStore property value. Possible values for the Certificate Destination Store.
func (m *Windows81TrustedRootCertificate) GetDestinationStore()(*CertificateDestinationStore) {
    return m.destinationStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81TrustedRootCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["destinationStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateDestinationStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationStore(val.(*CertificateDestinationStore))
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
// GetTrustedRootCertificate gets the trustedRootCertificate property value. Trusted Root Certificate
func (m *Windows81TrustedRootCertificate) GetTrustedRootCertificate()([]byte) {
    return m.trustedRootCertificate
}
// Serialize serializes information the current object
func (m *Windows81TrustedRootCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDestinationStore() != nil {
        cast := (*m.GetDestinationStore()).String()
        err = writer.WriteStringValue("destinationStore", &cast)
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
func (m *Windows81TrustedRootCertificate) SetCertFileName(value *string)() {
    m.certFileName = value
}
// SetDestinationStore sets the destinationStore property value. Possible values for the Certificate Destination Store.
func (m *Windows81TrustedRootCertificate) SetDestinationStore(value *CertificateDestinationStore)() {
    m.destinationStore = value
}
// SetTrustedRootCertificate sets the trustedRootCertificate property value. Trusted Root Certificate
func (m *Windows81TrustedRootCertificate) SetTrustedRootCertificate(value []byte)() {
    m.trustedRootCertificate = value
}
