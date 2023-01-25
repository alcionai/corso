package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81VpnProxyServer 
type Windows81VpnProxyServer struct {
    VpnProxyServer
    // Automatically detect proxy settings.
    automaticallyDetectProxySettings *bool
    // Bypass proxy server for local address.
    bypassProxyServerForLocalAddress *bool
}
// NewWindows81VpnProxyServer instantiates a new Windows81VpnProxyServer and sets the default values.
func NewWindows81VpnProxyServer()(*Windows81VpnProxyServer) {
    m := &Windows81VpnProxyServer{
        VpnProxyServer: *NewVpnProxyServer(),
    }
    odataTypeValue := "#microsoft.graph.windows81VpnProxyServer";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows81VpnProxyServerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81VpnProxyServerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81VpnProxyServer(), nil
}
// GetAutomaticallyDetectProxySettings gets the automaticallyDetectProxySettings property value. Automatically detect proxy settings.
func (m *Windows81VpnProxyServer) GetAutomaticallyDetectProxySettings()(*bool) {
    return m.automaticallyDetectProxySettings
}
// GetBypassProxyServerForLocalAddress gets the bypassProxyServerForLocalAddress property value. Bypass proxy server for local address.
func (m *Windows81VpnProxyServer) GetBypassProxyServerForLocalAddress()(*bool) {
    return m.bypassProxyServerForLocalAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81VpnProxyServer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VpnProxyServer.GetFieldDeserializers()
    res["automaticallyDetectProxySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticallyDetectProxySettings(val)
        }
        return nil
    }
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
func (m *Windows81VpnProxyServer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VpnProxyServer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("automaticallyDetectProxySettings", m.GetAutomaticallyDetectProxySettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bypassProxyServerForLocalAddress", m.GetBypassProxyServerForLocalAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutomaticallyDetectProxySettings sets the automaticallyDetectProxySettings property value. Automatically detect proxy settings.
func (m *Windows81VpnProxyServer) SetAutomaticallyDetectProxySettings(value *bool)() {
    m.automaticallyDetectProxySettings = value
}
// SetBypassProxyServerForLocalAddress sets the bypassProxyServerForLocalAddress property value. Bypass proxy server for local address.
func (m *Windows81VpnProxyServer) SetBypassProxyServerForLocalAddress(value *bool)() {
    m.bypassProxyServerForLocalAddress = value
}
