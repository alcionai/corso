package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10PFXImportCertificateProfile 
type Windows10PFXImportCertificateProfile struct {
    DeviceConfiguration
    // Key Storage Provider (KSP) Import Options.
    keyStorageProvider *KeyStorageProviderOption
}
// NewWindows10PFXImportCertificateProfile instantiates a new Windows10PFXImportCertificateProfile and sets the default values.
func NewWindows10PFXImportCertificateProfile()(*Windows10PFXImportCertificateProfile) {
    m := &Windows10PFXImportCertificateProfile{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10PFXImportCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10PFXImportCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10PFXImportCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10PFXImportCertificateProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10PFXImportCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["keyStorageProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyStorageProviderOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyStorageProvider(val.(*KeyStorageProviderOption))
        }
        return nil
    }
    return res
}
// GetKeyStorageProvider gets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10PFXImportCertificateProfile) GetKeyStorageProvider()(*KeyStorageProviderOption) {
    return m.keyStorageProvider
}
// Serialize serializes information the current object
func (m *Windows10PFXImportCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetKeyStorageProvider() != nil {
        cast := (*m.GetKeyStorageProvider()).String()
        err = writer.WriteStringValue("keyStorageProvider", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeyStorageProvider sets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10PFXImportCertificateProfile) SetKeyStorageProvider(value *KeyStorageProviderOption)() {
    m.keyStorageProvider = value
}
