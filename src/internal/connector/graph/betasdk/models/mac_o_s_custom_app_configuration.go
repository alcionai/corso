package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSCustomAppConfiguration 
type MacOSCustomAppConfiguration struct {
    DeviceConfiguration
    // Bundle id for targeting.
    bundleId *string
    // Configuration xml. (UTF8 encoded byte array)
    configurationXml []byte
    // Configuration file name (.plist
    fileName *string
}
// NewMacOSCustomAppConfiguration instantiates a new MacOSCustomAppConfiguration and sets the default values.
func NewMacOSCustomAppConfiguration()(*MacOSCustomAppConfiguration) {
    m := &MacOSCustomAppConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSCustomAppConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSCustomAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSCustomAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSCustomAppConfiguration(), nil
}
// GetBundleId gets the bundleId property value. Bundle id for targeting.
func (m *MacOSCustomAppConfiguration) GetBundleId()(*string) {
    return m.bundleId
}
// GetConfigurationXml gets the configurationXml property value. Configuration xml. (UTF8 encoded byte array)
func (m *MacOSCustomAppConfiguration) GetConfigurationXml()([]byte) {
    return m.configurationXml
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSCustomAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
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
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. Configuration file name (.plist
func (m *MacOSCustomAppConfiguration) GetFileName()(*string) {
    return m.fileName
}
// Serialize serializes information the current object
func (m *MacOSCustomAppConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("bundleId", m.GetBundleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("configurationXml", m.GetConfigurationXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBundleId sets the bundleId property value. Bundle id for targeting.
func (m *MacOSCustomAppConfiguration) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetConfigurationXml sets the configurationXml property value. Configuration xml. (UTF8 encoded byte array)
func (m *MacOSCustomAppConfiguration) SetConfigurationXml(value []byte)() {
    m.configurationXml = value
}
// SetFileName sets the fileName property value. Configuration file name (.plist
func (m *MacOSCustomAppConfiguration) SetFileName(value *string)() {
    m.fileName = value
}
