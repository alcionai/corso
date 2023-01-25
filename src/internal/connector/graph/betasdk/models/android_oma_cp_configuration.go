package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidOmaCpConfiguration 
type AndroidOmaCpConfiguration struct {
    DeviceConfiguration
    // Configuration XML that will be applied to the device. When it is read, it only provides a placeholder string since the original data is encrypted and stored.
    configurationXml []byte
}
// NewAndroidOmaCpConfiguration instantiates a new AndroidOmaCpConfiguration and sets the default values.
func NewAndroidOmaCpConfiguration()(*AndroidOmaCpConfiguration) {
    m := &AndroidOmaCpConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidOmaCpConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidOmaCpConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidOmaCpConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidOmaCpConfiguration(), nil
}
// GetConfigurationXml gets the configurationXml property value. Configuration XML that will be applied to the device. When it is read, it only provides a placeholder string since the original data is encrypted and stored.
func (m *AndroidOmaCpConfiguration) GetConfigurationXml()([]byte) {
    return m.configurationXml
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidOmaCpConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["configurationXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationXml(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AndroidOmaCpConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("configurationXml", m.GetConfigurationXml())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationXml sets the configurationXml property value. Configuration XML that will be applied to the device. When it is read, it only provides a placeholder string since the original data is encrypted and stored.
func (m *AndroidOmaCpConfiguration) SetConfigurationXml(value []byte)() {
    m.configurationXml = value
}
