package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosDerivedCredentialAuthenticationConfiguration 
type IosDerivedCredentialAuthenticationConfiguration struct {
    DeviceConfiguration
    // Tenant level settings for the Derived Credentials to be used for authentication.
    derivedCredentialSettings DeviceManagementDerivedCredentialSettingsable
}
// NewIosDerivedCredentialAuthenticationConfiguration instantiates a new IosDerivedCredentialAuthenticationConfiguration and sets the default values.
func NewIosDerivedCredentialAuthenticationConfiguration()(*IosDerivedCredentialAuthenticationConfiguration) {
    m := &IosDerivedCredentialAuthenticationConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosDerivedCredentialAuthenticationConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosDerivedCredentialAuthenticationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosDerivedCredentialAuthenticationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosDerivedCredentialAuthenticationConfiguration(), nil
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosDerivedCredentialAuthenticationConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    return m.derivedCredentialSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosDerivedCredentialAuthenticationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["derivedCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialSettings(val.(DeviceManagementDerivedCredentialSettingsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IosDerivedCredentialAuthenticationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("derivedCredentialSettings", m.GetDerivedCredentialSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosDerivedCredentialAuthenticationConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    m.derivedCredentialSettings = value
}
