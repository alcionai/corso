package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkNetworkConfiguration 
type TeamworkNetworkConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The default gateway is the path used to pass information when the destination is unknown to the device.
    defaultGateway *string
    // The network domain of the device, for example, contoso.com.
    domainName *string
    // The device name on a network.
    hostName *string
    // The IP address is a numerical label that uniquely identifies every device connected to the internet.
    ipAddress *string
    // True if DHCP is enabled.
    isDhcpEnabled *bool
    // True if the PC port is enabled.
    isPCPortEnabled *bool
    // The OdataType property
    odataType *string
    // A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
    primaryDns *string
    // A secondary DNS is used when the primary DNS is not available.
    secondaryDns *string
    // A subnet mask is a number that distinguishes the network address and the host address within an IP address.
    subnetMask *string
}
// NewTeamworkNetworkConfiguration instantiates a new teamworkNetworkConfiguration and sets the default values.
func NewTeamworkNetworkConfiguration()(*TeamworkNetworkConfiguration) {
    m := &TeamworkNetworkConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkNetworkConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkNetworkConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkNetworkConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkNetworkConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultGateway gets the defaultGateway property value. The default gateway is the path used to pass information when the destination is unknown to the device.
func (m *TeamworkNetworkConfiguration) GetDefaultGateway()(*string) {
    return m.defaultGateway
}
// GetDomainName gets the domainName property value. The network domain of the device, for example, contoso.com.
func (m *TeamworkNetworkConfiguration) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkNetworkConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultGateway"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultGateway(val)
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["hostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["isDhcpEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDhcpEnabled(val)
        }
        return nil
    }
    res["isPCPortEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPCPortEnabled(val)
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
    res["primaryDns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryDns(val)
        }
        return nil
    }
    res["secondaryDns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryDns(val)
        }
        return nil
    }
    res["subnetMask"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetMask(val)
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. The device name on a network.
func (m *TeamworkNetworkConfiguration) GetHostName()(*string) {
    return m.hostName
}
// GetIpAddress gets the ipAddress property value. The IP address is a numerical label that uniquely identifies every device connected to the internet.
func (m *TeamworkNetworkConfiguration) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetIsDhcpEnabled gets the isDhcpEnabled property value. True if DHCP is enabled.
func (m *TeamworkNetworkConfiguration) GetIsDhcpEnabled()(*bool) {
    return m.isDhcpEnabled
}
// GetIsPCPortEnabled gets the isPCPortEnabled property value. True if the PC port is enabled.
func (m *TeamworkNetworkConfiguration) GetIsPCPortEnabled()(*bool) {
    return m.isPCPortEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkNetworkConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetPrimaryDns gets the primaryDns property value. A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
func (m *TeamworkNetworkConfiguration) GetPrimaryDns()(*string) {
    return m.primaryDns
}
// GetSecondaryDns gets the secondaryDns property value. A secondary DNS is used when the primary DNS is not available.
func (m *TeamworkNetworkConfiguration) GetSecondaryDns()(*string) {
    return m.secondaryDns
}
// GetSubnetMask gets the subnetMask property value. A subnet mask is a number that distinguishes the network address and the host address within an IP address.
func (m *TeamworkNetworkConfiguration) GetSubnetMask()(*string) {
    return m.subnetMask
}
// Serialize serializes information the current object
func (m *TeamworkNetworkConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultGateway", m.GetDefaultGateway())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostName", m.GetHostName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDhcpEnabled", m.GetIsDhcpEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPCPortEnabled", m.GetIsPCPortEnabled())
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
        err := writer.WriteStringValue("primaryDns", m.GetPrimaryDns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secondaryDns", m.GetSecondaryDns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetMask", m.GetSubnetMask())
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
func (m *TeamworkNetworkConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultGateway sets the defaultGateway property value. The default gateway is the path used to pass information when the destination is unknown to the device.
func (m *TeamworkNetworkConfiguration) SetDefaultGateway(value *string)() {
    m.defaultGateway = value
}
// SetDomainName sets the domainName property value. The network domain of the device, for example, contoso.com.
func (m *TeamworkNetworkConfiguration) SetDomainName(value *string)() {
    m.domainName = value
}
// SetHostName sets the hostName property value. The device name on a network.
func (m *TeamworkNetworkConfiguration) SetHostName(value *string)() {
    m.hostName = value
}
// SetIpAddress sets the ipAddress property value. The IP address is a numerical label that uniquely identifies every device connected to the internet.
func (m *TeamworkNetworkConfiguration) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetIsDhcpEnabled sets the isDhcpEnabled property value. True if DHCP is enabled.
func (m *TeamworkNetworkConfiguration) SetIsDhcpEnabled(value *bool)() {
    m.isDhcpEnabled = value
}
// SetIsPCPortEnabled sets the isPCPortEnabled property value. True if the PC port is enabled.
func (m *TeamworkNetworkConfiguration) SetIsPCPortEnabled(value *bool)() {
    m.isPCPortEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkNetworkConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrimaryDns sets the primaryDns property value. A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
func (m *TeamworkNetworkConfiguration) SetPrimaryDns(value *string)() {
    m.primaryDns = value
}
// SetSecondaryDns sets the secondaryDns property value. A secondary DNS is used when the primary DNS is not available.
func (m *TeamworkNetworkConfiguration) SetSecondaryDns(value *string)() {
    m.secondaryDns = value
}
// SetSubnetMask sets the subnetMask property value. A subnet mask is a number that distinguishes the network address and the host address within an IP address.
func (m *TeamworkNetworkConfiguration) SetSubnetMask(value *string)() {
    m.subnetMask = value
}
