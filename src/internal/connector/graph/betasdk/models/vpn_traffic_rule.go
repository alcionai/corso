package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnTrafficRule vPN Traffic Rule definition.
type VpnTrafficRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // App identifier, if this traffic rule is triggered by an app.
    appId *string
    // Indicates the type of app that a VPN traffic rule is associated with.
    appType *VpnTrafficRuleAppType
    // Claims associated with this traffic rule.
    claims *string
    // Local address range. This collection can contain a maximum of 500 elements.
    localAddressRanges []IPv4Rangeable
    // Local port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
    localPortRanges []NumberRangeable
    // Name.
    name *string
    // The OdataType property
    odataType *string
    // Protocols (0-255). Valid values 0 to 255
    protocols *int32
    // Remote address range. This collection can contain a maximum of 500 elements.
    remoteAddressRanges []IPv4Rangeable
    // Remote port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
    remotePortRanges []NumberRangeable
    // Specifies the routing policy for a VPN traffic rule.
    routingPolicyType *VpnTrafficRuleRoutingPolicyType
}
// NewVpnTrafficRule instantiates a new vpnTrafficRule and sets the default values.
func NewVpnTrafficRule()(*VpnTrafficRule) {
    m := &VpnTrafficRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVpnTrafficRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnTrafficRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnTrafficRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnTrafficRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppId gets the appId property value. App identifier, if this traffic rule is triggered by an app.
func (m *VpnTrafficRule) GetAppId()(*string) {
    return m.appId
}
// GetAppType gets the appType property value. Indicates the type of app that a VPN traffic rule is associated with.
func (m *VpnTrafficRule) GetAppType()(*VpnTrafficRuleAppType) {
    return m.appType
}
// GetClaims gets the claims property value. Claims associated with this traffic rule.
func (m *VpnTrafficRule) GetClaims()(*string) {
    return m.claims
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VpnTrafficRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnTrafficRuleAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppType(val.(*VpnTrafficRuleAppType))
        }
        return nil
    }
    res["claims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaims(val)
        }
        return nil
    }
    res["localAddressRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIPv4RangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IPv4Rangeable, len(val))
            for i, v := range val {
                res[i] = v.(IPv4Rangeable)
            }
            m.SetLocalAddressRanges(res)
        }
        return nil
    }
    res["localPortRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNumberRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NumberRangeable, len(val))
            for i, v := range val {
                res[i] = v.(NumberRangeable)
            }
            m.SetLocalPortRanges(res)
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
    res["protocols"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocols(val)
        }
        return nil
    }
    res["remoteAddressRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIPv4RangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IPv4Rangeable, len(val))
            for i, v := range val {
                res[i] = v.(IPv4Rangeable)
            }
            m.SetRemoteAddressRanges(res)
        }
        return nil
    }
    res["remotePortRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNumberRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NumberRangeable, len(val))
            for i, v := range val {
                res[i] = v.(NumberRangeable)
            }
            m.SetRemotePortRanges(res)
        }
        return nil
    }
    res["routingPolicyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnTrafficRuleRoutingPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoutingPolicyType(val.(*VpnTrafficRuleRoutingPolicyType))
        }
        return nil
    }
    return res
}
// GetLocalAddressRanges gets the localAddressRanges property value. Local address range. This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) GetLocalAddressRanges()([]IPv4Rangeable) {
    return m.localAddressRanges
}
// GetLocalPortRanges gets the localPortRanges property value. Local port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) GetLocalPortRanges()([]NumberRangeable) {
    return m.localPortRanges
}
// GetName gets the name property value. Name.
func (m *VpnTrafficRule) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VpnTrafficRule) GetOdataType()(*string) {
    return m.odataType
}
// GetProtocols gets the protocols property value. Protocols (0-255). Valid values 0 to 255
func (m *VpnTrafficRule) GetProtocols()(*int32) {
    return m.protocols
}
// GetRemoteAddressRanges gets the remoteAddressRanges property value. Remote address range. This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) GetRemoteAddressRanges()([]IPv4Rangeable) {
    return m.remoteAddressRanges
}
// GetRemotePortRanges gets the remotePortRanges property value. Remote port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) GetRemotePortRanges()([]NumberRangeable) {
    return m.remotePortRanges
}
// GetRoutingPolicyType gets the routingPolicyType property value. Specifies the routing policy for a VPN traffic rule.
func (m *VpnTrafficRule) GetRoutingPolicyType()(*VpnTrafficRuleRoutingPolicyType) {
    return m.routingPolicyType
}
// Serialize serializes information the current object
func (m *VpnTrafficRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetAppType() != nil {
        cast := (*m.GetAppType()).String()
        err := writer.WriteStringValue("appType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("claims", m.GetClaims())
        if err != nil {
            return err
        }
    }
    if m.GetLocalAddressRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalAddressRanges()))
        for i, v := range m.GetLocalAddressRanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("localAddressRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalPortRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalPortRanges()))
        for i, v := range m.GetLocalPortRanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("localPortRanges", cast)
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
        err := writer.WriteInt32Value("protocols", m.GetProtocols())
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAddressRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteAddressRanges()))
        for i, v := range m.GetRemoteAddressRanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("remoteAddressRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemotePortRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemotePortRanges()))
        for i, v := range m.GetRemotePortRanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("remotePortRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoutingPolicyType() != nil {
        cast := (*m.GetRoutingPolicyType()).String()
        err := writer.WriteStringValue("routingPolicyType", &cast)
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
func (m *VpnTrafficRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppId sets the appId property value. App identifier, if this traffic rule is triggered by an app.
func (m *VpnTrafficRule) SetAppId(value *string)() {
    m.appId = value
}
// SetAppType sets the appType property value. Indicates the type of app that a VPN traffic rule is associated with.
func (m *VpnTrafficRule) SetAppType(value *VpnTrafficRuleAppType)() {
    m.appType = value
}
// SetClaims sets the claims property value. Claims associated with this traffic rule.
func (m *VpnTrafficRule) SetClaims(value *string)() {
    m.claims = value
}
// SetLocalAddressRanges sets the localAddressRanges property value. Local address range. This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetLocalAddressRanges(value []IPv4Rangeable)() {
    m.localAddressRanges = value
}
// SetLocalPortRanges sets the localPortRanges property value. Local port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetLocalPortRanges(value []NumberRangeable)() {
    m.localPortRanges = value
}
// SetName sets the name property value. Name.
func (m *VpnTrafficRule) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VpnTrafficRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProtocols sets the protocols property value. Protocols (0-255). Valid values 0 to 255
func (m *VpnTrafficRule) SetProtocols(value *int32)() {
    m.protocols = value
}
// SetRemoteAddressRanges sets the remoteAddressRanges property value. Remote address range. This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetRemoteAddressRanges(value []IPv4Rangeable)() {
    m.remoteAddressRanges = value
}
// SetRemotePortRanges sets the remotePortRanges property value. Remote port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetRemotePortRanges(value []NumberRangeable)() {
    m.remotePortRanges = value
}
// SetRoutingPolicyType sets the routingPolicyType property value. Specifies the routing policy for a VPN traffic rule.
func (m *VpnTrafficRule) SetRoutingPolicyType(value *VpnTrafficRuleRoutingPolicyType)() {
    m.routingPolicyType = value
}
