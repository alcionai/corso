package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XWifiConfiguration 
type Windows10XWifiConfiguration struct {
    DeviceManagementResourceAccessProfileBase
    // ID to the Authentication Certificate
    authenticationCertificateId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
    customXml []byte
    // Custom Xml file name.
    customXmlFileName *string
}
// NewWindows10XWifiConfiguration instantiates a new Windows10XWifiConfiguration and sets the default values.
func NewWindows10XWifiConfiguration()(*Windows10XWifiConfiguration) {
    m := &Windows10XWifiConfiguration{
        DeviceManagementResourceAccessProfileBase: *NewDeviceManagementResourceAccessProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10XWifiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10XWifiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XWifiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XWifiConfiguration(), nil
}
// GetAuthenticationCertificateId gets the authenticationCertificateId property value. ID to the Authentication Certificate
func (m *Windows10XWifiConfiguration) GetAuthenticationCertificateId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.authenticationCertificateId
}
// GetCustomXml gets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
func (m *Windows10XWifiConfiguration) GetCustomXml()([]byte) {
    return m.customXml
}
// GetCustomXmlFileName gets the customXmlFileName property value. Custom Xml file name.
func (m *Windows10XWifiConfiguration) GetCustomXmlFileName()(*string) {
    return m.customXmlFileName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XWifiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementResourceAccessProfileBase.GetFieldDeserializers()
    res["authenticationCertificateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationCertificateId(val)
        }
        return nil
    }
    res["customXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomXml(val)
        }
        return nil
    }
    res["customXmlFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomXmlFileName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Windows10XWifiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementResourceAccessProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteUUIDValue("authenticationCertificateId", m.GetAuthenticationCertificateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("customXml", m.GetCustomXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customXmlFileName", m.GetCustomXmlFileName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationCertificateId sets the authenticationCertificateId property value. ID to the Authentication Certificate
func (m *Windows10XWifiConfiguration) SetAuthenticationCertificateId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.authenticationCertificateId = value
}
// SetCustomXml sets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
func (m *Windows10XWifiConfiguration) SetCustomXml(value []byte)() {
    m.customXml = value
}
// SetCustomXmlFileName sets the customXmlFileName property value. Custom Xml file name.
func (m *Windows10XWifiConfiguration) SetCustomXmlFileName(value *string)() {
    m.customXmlFileName = value
}
