package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsVpnConfiguration 
type WindowsVpnConfiguration struct {
    DeviceConfiguration
    // Connection name displayed to the user.
    connectionName *string
    // Custom XML commands that configures the VPN connection. (UTF8 encoded byte array)
    customXml []byte
    // List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
    servers []VpnServerable
}
// NewWindowsVpnConfiguration instantiates a new WindowsVpnConfiguration and sets the default values.
func NewWindowsVpnConfiguration()(*WindowsVpnConfiguration) {
    m := &WindowsVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windows10VpnConfiguration":
                        return NewWindows10VpnConfiguration(), nil
                    case "#microsoft.graph.windows81VpnConfiguration":
                        return NewWindows81VpnConfiguration(), nil
                    case "#microsoft.graph.windowsPhone81VpnConfiguration":
                        return NewWindowsPhone81VpnConfiguration(), nil
                }
            }
        }
    }
    return NewWindowsVpnConfiguration(), nil
}
// GetConnectionName gets the connectionName property value. Connection name displayed to the user.
func (m *WindowsVpnConfiguration) GetConnectionName()(*string) {
    return m.connectionName
}
// GetCustomXml gets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 encoded byte array)
func (m *WindowsVpnConfiguration) GetCustomXml()([]byte) {
    return m.customXml
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["connectionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionName(val)
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
    res["servers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnServerable, len(val))
            for i, v := range val {
                res[i] = v.(VpnServerable)
            }
            m.SetServers(res)
        }
        return nil
    }
    return res
}
// GetServers gets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *WindowsVpnConfiguration) GetServers()([]VpnServerable) {
    return m.servers
}
// Serialize serializes information the current object
func (m *WindowsVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("connectionName", m.GetConnectionName())
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
    if m.GetServers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServers()))
        for i, v := range m.GetServers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("servers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *WindowsVpnConfiguration) SetConnectionName(value *string)() {
    m.connectionName = value
}
// SetCustomXml sets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 encoded byte array)
func (m *WindowsVpnConfiguration) SetCustomXml(value []byte)() {
    m.customXml = value
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *WindowsVpnConfiguration) SetServers(value []VpnServerable)() {
    m.servers = value
}
