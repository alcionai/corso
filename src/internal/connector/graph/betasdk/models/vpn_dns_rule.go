package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnDnsRule vPN DNS Rule definition.
type VpnDnsRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Automatically connect to the VPN when the device connects to this domain: Default False.
    autoTrigger *bool
    // Name.
    name *string
    // The OdataType property
    odataType *string
    // Keep this rule active even when the VPN is not connected: Default False
    persistent *bool
    // Proxy Server Uri.
    proxyServerUri *string
    // Servers.
    servers []string
}
// NewVpnDnsRule instantiates a new vpnDnsRule and sets the default values.
func NewVpnDnsRule()(*VpnDnsRule) {
    m := &VpnDnsRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVpnDnsRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnDnsRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnDnsRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnDnsRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAutoTrigger gets the autoTrigger property value. Automatically connect to the VPN when the device connects to this domain: Default False.
func (m *VpnDnsRule) GetAutoTrigger()(*bool) {
    return m.autoTrigger
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VpnDnsRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["autoTrigger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoTrigger(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["persistent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistent(val)
        }
        return nil
    }
    res["proxyServerUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyServerUri(val)
        }
        return nil
    }
    res["servers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetServers(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name.
func (m *VpnDnsRule) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VpnDnsRule) GetOdataType()(*string) {
    return m.odataType
}
// GetPersistent gets the persistent property value. Keep this rule active even when the VPN is not connected: Default False
func (m *VpnDnsRule) GetPersistent()(*bool) {
    return m.persistent
}
// GetProxyServerUri gets the proxyServerUri property value. Proxy Server Uri.
func (m *VpnDnsRule) GetProxyServerUri()(*string) {
    return m.proxyServerUri
}
// GetServers gets the servers property value. Servers.
func (m *VpnDnsRule) GetServers()([]string) {
    return m.servers
}
// Serialize serializes information the current object
func (m *VpnDnsRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("autoTrigger", m.GetAutoTrigger())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistent", m.GetPersistent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proxyServerUri", m.GetProxyServerUri())
        if err != nil {
            return err
        }
    }
    if m.GetServers() != nil {
        err := writer.WriteCollectionOfStringValues("servers", m.GetServers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnDnsRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAutoTrigger sets the autoTrigger property value. Automatically connect to the VPN when the device connects to this domain: Default False.
func (m *VpnDnsRule) SetAutoTrigger(value *bool)() {
    m.autoTrigger = value
}
// SetName sets the name property value. Name.
func (m *VpnDnsRule) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VpnDnsRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPersistent sets the persistent property value. Keep this rule active even when the VPN is not connected: Default False
func (m *VpnDnsRule) SetPersistent(value *bool)() {
    m.persistent = value
}
// SetProxyServerUri sets the proxyServerUri property value. Proxy Server Uri.
func (m *VpnDnsRule) SetProxyServerUri(value *string)() {
    m.proxyServerUri = value
}
// SetServers sets the servers property value. Servers.
func (m *VpnDnsRule) SetServers(value []string)() {
    m.servers = value
}
