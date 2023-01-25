package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10VpnProxyServer 
type Windows10VpnProxyServer struct {
    VpnProxyServer
    // Bypass proxy server for local address.
    bypassProxyServerForLocalAddress *bool
}
// NewWindows10VpnProxyServer instantiates a new Windows10VpnProxyServer and sets the default values.
func NewWindows10VpnProxyServer()(*Windows10VpnProxyServer) {
    m := &Windows10VpnProxyServer{
        VpnProxyServer: *NewVpnProxyServer(),
    }
    odataTypeValue := "#microsoft.graph.windows10VpnProxyServer";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10VpnProxyServerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10VpnProxyServerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10VpnProxyServer(), nil
}
// GetBypassProxyServerForLocalAddress gets the bypassProxyServerForLocalAddress property value. Bypass proxy server for local address.
func (m *Windows10VpnProxyServer) GetBypassProxyServerForLocalAddress()(*bool) {
    return m.bypassProxyServerForLocalAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10VpnProxyServer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VpnProxyServer.GetFieldDeserializers()
    res["bypassProxyServerForLocalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassProxyServerForLocalAddress(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Windows10VpnProxyServer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VpnProxyServer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("bypassProxyServerForLocalAddress", m.GetBypassProxyServerForLocalAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBypassProxyServerForLocalAddress sets the bypassProxyServerForLocalAddress property value. Bypass proxy server for local address.
func (m *Windows10VpnProxyServer) SetBypassProxyServerForLocalAddress(value *bool)() {
    m.bypassProxyServerForLocalAddress = value
}
