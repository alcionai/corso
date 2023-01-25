package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsNetworkIsolationPolicy windows Network Isolation Policy
type WindowsNetworkIsolationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Contains a list of enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy. This collection can contain a maximum of 500 elements.
    enterpriseCloudResources []ProxiedDomainable
    // This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseCloudResources policy to force traffic to the matched cloud resources through these proxies.
    enterpriseInternalProxyServers []string
    // Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to. This collection can contain a maximum of 500 elements.
    enterpriseIPRanges []IpRangeable
    // Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false.
    enterpriseIPRangesAreAuthoritative *bool
    // This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected. These locations will be considered a safe destination for enterprise data to be shared to.
    enterpriseNetworkDomainNames []string
    // This is a list of proxy servers. Any server not on this list is considered non-enterprise.
    enterpriseProxyServers []string
    // Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
    enterpriseProxyServersAreAuthoritative *bool
    // List of domain names that can used for work or personal resource.
    neutralDomainResources []string
    // The OdataType property
    odataType *string
}
// NewWindowsNetworkIsolationPolicy instantiates a new windowsNetworkIsolationPolicy and sets the default values.
func NewWindowsNetworkIsolationPolicy()(*WindowsNetworkIsolationPolicy) {
    m := &WindowsNetworkIsolationPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsNetworkIsolationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsNetworkIsolationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsNetworkIsolationPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsNetworkIsolationPolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnterpriseCloudResources gets the enterpriseCloudResources property value. Contains a list of enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy. This collection can contain a maximum of 500 elements.
func (m *WindowsNetworkIsolationPolicy) GetEnterpriseCloudResources()([]ProxiedDomainable) {
    return m.enterpriseCloudResources
}
// GetEnterpriseInternalProxyServers gets the enterpriseInternalProxyServers property value. This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseCloudResources policy to force traffic to the matched cloud resources through these proxies.
func (m *WindowsNetworkIsolationPolicy) GetEnterpriseInternalProxyServers()([]string) {
    return m.enterpriseInternalProxyServers
}
// GetEnterpriseIPRanges gets the enterpriseIPRanges property value. Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to. This collection can contain a maximum of 500 elements.
func (m *WindowsNetworkIsolationPolicy) GetEnterpriseIPRanges()([]IpRangeable) {
    return m.enterpriseIPRanges
}
// GetEnterpriseIPRangesAreAuthoritative gets the enterpriseIPRangesAreAuthoritative property value. Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false.
func (m *WindowsNetworkIsolationPolicy) GetEnterpriseIPRangesAreAuthoritative()(*bool) {
    return m.enterpriseIPRangesAreAuthoritative
}
// GetEnterpriseNetworkDomainNames gets the enterpriseNetworkDomainNames property value. This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected. These locations will be considered a safe destination for enterprise data to be shared to.
func (m *WindowsNetworkIsolationPolicy) GetEnterpriseNetworkDomainNames()([]string) {
    return m.enterpriseNetworkDomainNames
}
// GetEnterpriseProxyServers gets the enterpriseProxyServers property value. This is a list of proxy servers. Any server not on this list is considered non-enterprise.
func (m *WindowsNetworkIsolationPolicy) GetEnterpriseProxyServers()([]string) {
    return m.enterpriseProxyServers
}
// GetEnterpriseProxyServersAreAuthoritative gets the enterpriseProxyServersAreAuthoritative property value. Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
func (m *WindowsNetworkIsolationPolicy) GetEnterpriseProxyServersAreAuthoritative()(*bool) {
    return m.enterpriseProxyServersAreAuthoritative
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsNetworkIsolationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enterpriseCloudResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProxiedDomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProxiedDomainable, len(val))
            for i, v := range val {
                res[i] = v.(ProxiedDomainable)
            }
            m.SetEnterpriseCloudResources(res)
        }
        return nil
    }
    res["enterpriseInternalProxyServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnterpriseInternalProxyServers(res)
        }
        return nil
    }
    res["enterpriseIPRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpRangeable, len(val))
            for i, v := range val {
                res[i] = v.(IpRangeable)
            }
            m.SetEnterpriseIPRanges(res)
        }
        return nil
    }
    res["enterpriseIPRangesAreAuthoritative"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseIPRangesAreAuthoritative(val)
        }
        return nil
    }
    res["enterpriseNetworkDomainNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnterpriseNetworkDomainNames(res)
        }
        return nil
    }
    res["enterpriseProxyServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnterpriseProxyServers(res)
        }
        return nil
    }
    res["enterpriseProxyServersAreAuthoritative"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseProxyServersAreAuthoritative(val)
        }
        return nil
    }
    res["neutralDomainResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNeutralDomainResources(res)
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
// GetNeutralDomainResources gets the neutralDomainResources property value. List of domain names that can used for work or personal resource.
func (m *WindowsNetworkIsolationPolicy) GetNeutralDomainResources()([]string) {
    return m.neutralDomainResources
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsNetworkIsolationPolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *WindowsNetworkIsolationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnterpriseCloudResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnterpriseCloudResources()))
        for i, v := range m.GetEnterpriseCloudResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("enterpriseCloudResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseInternalProxyServers() != nil {
        err := writer.WriteCollectionOfStringValues("enterpriseInternalProxyServers", m.GetEnterpriseInternalProxyServers())
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseIPRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnterpriseIPRanges()))
        for i, v := range m.GetEnterpriseIPRanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("enterpriseIPRanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enterpriseIPRangesAreAuthoritative", m.GetEnterpriseIPRangesAreAuthoritative())
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseNetworkDomainNames() != nil {
        err := writer.WriteCollectionOfStringValues("enterpriseNetworkDomainNames", m.GetEnterpriseNetworkDomainNames())
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseProxyServers() != nil {
        err := writer.WriteCollectionOfStringValues("enterpriseProxyServers", m.GetEnterpriseProxyServers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enterpriseProxyServersAreAuthoritative", m.GetEnterpriseProxyServersAreAuthoritative())
        if err != nil {
            return err
        }
    }
    if m.GetNeutralDomainResources() != nil {
        err := writer.WriteCollectionOfStringValues("neutralDomainResources", m.GetNeutralDomainResources())
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
func (m *WindowsNetworkIsolationPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnterpriseCloudResources sets the enterpriseCloudResources property value. Contains a list of enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy. This collection can contain a maximum of 500 elements.
func (m *WindowsNetworkIsolationPolicy) SetEnterpriseCloudResources(value []ProxiedDomainable)() {
    m.enterpriseCloudResources = value
}
// SetEnterpriseInternalProxyServers sets the enterpriseInternalProxyServers property value. This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseCloudResources policy to force traffic to the matched cloud resources through these proxies.
func (m *WindowsNetworkIsolationPolicy) SetEnterpriseInternalProxyServers(value []string)() {
    m.enterpriseInternalProxyServers = value
}
// SetEnterpriseIPRanges sets the enterpriseIPRanges property value. Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to. This collection can contain a maximum of 500 elements.
func (m *WindowsNetworkIsolationPolicy) SetEnterpriseIPRanges(value []IpRangeable)() {
    m.enterpriseIPRanges = value
}
// SetEnterpriseIPRangesAreAuthoritative sets the enterpriseIPRangesAreAuthoritative property value. Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false.
func (m *WindowsNetworkIsolationPolicy) SetEnterpriseIPRangesAreAuthoritative(value *bool)() {
    m.enterpriseIPRangesAreAuthoritative = value
}
// SetEnterpriseNetworkDomainNames sets the enterpriseNetworkDomainNames property value. This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected. These locations will be considered a safe destination for enterprise data to be shared to.
func (m *WindowsNetworkIsolationPolicy) SetEnterpriseNetworkDomainNames(value []string)() {
    m.enterpriseNetworkDomainNames = value
}
// SetEnterpriseProxyServers sets the enterpriseProxyServers property value. This is a list of proxy servers. Any server not on this list is considered non-enterprise.
func (m *WindowsNetworkIsolationPolicy) SetEnterpriseProxyServers(value []string)() {
    m.enterpriseProxyServers = value
}
// SetEnterpriseProxyServersAreAuthoritative sets the enterpriseProxyServersAreAuthoritative property value. Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false
func (m *WindowsNetworkIsolationPolicy) SetEnterpriseProxyServersAreAuthoritative(value *bool)() {
    m.enterpriseProxyServersAreAuthoritative = value
}
// SetNeutralDomainResources sets the neutralDomainResources property value. List of domain names that can used for work or personal resource.
func (m *WindowsNetworkIsolationPolicy) SetNeutralDomainResources(value []string)() {
    m.neutralDomainResources = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsNetworkIsolationPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
