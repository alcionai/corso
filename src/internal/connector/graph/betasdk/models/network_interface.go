package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkInterface 
type NetworkInterface struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection, and so on).
    description *string
    // Last IPv4 address associated with this NIC.
    ipV4Address *string
    // Last Public (aka global) IPv6 address associated with this NIC.
    ipV6Address *string
    // Last local (link-local or site-local) IPv6 address associated with this NIC.
    localIpV6Address *string
    // MAC address of the NIC on this host.
    macAddress *string
    // The OdataType property
    odataType *string
}
// NewNetworkInterface instantiates a new networkInterface and sets the default values.
func NewNetworkInterface()(*NetworkInterface) {
    m := &NetworkInterface{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateNetworkInterfaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNetworkInterfaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkInterface(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInterface) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection, and so on).
func (m *NetworkInterface) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkInterface) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["ipV4Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpV4Address(val)
        }
        return nil
    }
    res["ipV6Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpV6Address(val)
        }
        return nil
    }
    res["localIpV6Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalIpV6Address(val)
        }
        return nil
    }
    res["macAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
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
    return res
}
// GetIpV4Address gets the ipV4Address property value. Last IPv4 address associated with this NIC.
func (m *NetworkInterface) GetIpV4Address()(*string) {
    return m.ipV4Address
}
// GetIpV6Address gets the ipV6Address property value. Last Public (aka global) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetIpV6Address()(*string) {
    return m.ipV6Address
}
// GetLocalIpV6Address gets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetLocalIpV6Address()(*string) {
    return m.localIpV6Address
}
// GetMacAddress gets the macAddress property value. MAC address of the NIC on this host.
func (m *NetworkInterface) GetMacAddress()(*string) {
    return m.macAddress
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NetworkInterface) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *NetworkInterface) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipV4Address", m.GetIpV4Address())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipV6Address", m.GetIpV6Address())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("localIpV6Address", m.GetLocalIpV6Address())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInterface) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection, and so on).
func (m *NetworkInterface) SetDescription(value *string)() {
    m.description = value
}
// SetIpV4Address sets the ipV4Address property value. Last IPv4 address associated with this NIC.
func (m *NetworkInterface) SetIpV4Address(value *string)() {
    m.ipV4Address = value
}
// SetIpV6Address sets the ipV6Address property value. Last Public (aka global) IPv6 address associated with this NIC.
func (m *NetworkInterface) SetIpV6Address(value *string)() {
    m.ipV6Address = value
}
// SetLocalIpV6Address sets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
func (m *NetworkInterface) SetLocalIpV6Address(value *string)() {
    m.localIpV6Address = value
}
// SetMacAddress sets the macAddress property value. MAC address of the NIC on this host.
func (m *NetworkInterface) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NetworkInterface) SetOdataType(value *string)() {
    m.odataType = value
}
