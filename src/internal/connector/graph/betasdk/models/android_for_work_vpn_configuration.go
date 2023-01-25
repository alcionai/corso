package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkVpnConfiguration 
type AndroidForWorkVpnConfiguration struct {
    DeviceConfiguration
    // VPN Authentication Method.
    authenticationMethod *VpnAuthenticationMethod
    // Connection name displayed to the user.
    connectionName *string
    // Android For Work VPN connection type.
    connectionType *AndroidForWorkVpnConnectionType
    // Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
    customData []KeyValueable
    // Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
    customKeyValueData []KeyValuePairable
    // Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
    fingerprint *string
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate AndroidForWorkCertificateProfileBaseable
    // Realm when connection type is set to Pulse Secure.
    realm *string
    // Role when connection type is set to Pulse Secure.
    role *string
    // List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
    servers []VpnServerable
}
// NewAndroidForWorkVpnConfiguration instantiates a new AndroidForWorkVpnConfiguration and sets the default values.
func NewAndroidForWorkVpnConfiguration()(*AndroidForWorkVpnConfiguration) {
    m := &AndroidForWorkVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidForWorkVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidForWorkVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkVpnConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidForWorkVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    return m.authenticationMethod
}
// GetConnectionName gets the connectionName property value. Connection name displayed to the user.
func (m *AndroidForWorkVpnConfiguration) GetConnectionName()(*string) {
    return m.connectionName
}
// GetConnectionType gets the connectionType property value. Android For Work VPN connection type.
func (m *AndroidForWorkVpnConfiguration) GetConnectionType()(*AndroidForWorkVpnConnectionType) {
    return m.connectionType
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) GetCustomData()([]KeyValueable) {
    return m.customData
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    return m.customKeyValueData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*VpnAuthenticationMethod))
        }
        return nil
    }
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
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*AndroidForWorkVpnConnectionType))
        }
        return nil
    }
    res["customData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValueable)
            }
            m.SetCustomData(res)
        }
        return nil
    }
    res["customKeyValueData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetCustomKeyValueData(res)
        }
        return nil
    }
    res["fingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprint(val)
        }
        return nil
    }
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidForWorkCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(AndroidForWorkCertificateProfileBaseable))
        }
        return nil
    }
    res["realm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealm(val)
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
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
// GetFingerprint gets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidForWorkVpnConfiguration) GetFingerprint()(*string) {
    return m.fingerprint
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidForWorkVpnConfiguration) GetIdentityCertificate()(AndroidForWorkCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) GetRealm()(*string) {
    return m.realm
}
// GetRole gets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) GetRole()(*string) {
    return m.role
}
// GetServers gets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkVpnConfiguration) GetServers()([]VpnServerable) {
    return m.servers
}
// Serialize serializes information the current object
func (m *AndroidForWorkVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectionName", m.GetConnectionName())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomData()))
        for i, v := range m.GetCustomData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customData", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomKeyValueData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomKeyValueData()))
        for i, v := range m.GetCustomKeyValueData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customKeyValueData", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fingerprint", m.GetFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("realm", m.GetRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("role", m.GetRole())
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
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AndroidForWorkVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AndroidForWorkVpnConfiguration) SetConnectionName(value *string)() {
    m.connectionName = value
}
// SetConnectionType sets the connectionType property value. Android For Work VPN connection type.
func (m *AndroidForWorkVpnConfiguration) SetConnectionType(value *AndroidForWorkVpnConnectionType)() {
    m.connectionType = value
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) SetCustomData(value []KeyValueable)() {
    m.customData = value
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
func (m *AndroidForWorkVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    m.customKeyValueData = value
}
// SetFingerprint sets the fingerprint property value. Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
func (m *AndroidForWorkVpnConfiguration) SetFingerprint(value *string)() {
    m.fingerprint = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *AndroidForWorkVpnConfiguration) SetIdentityCertificate(value AndroidForWorkCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) SetRealm(value *string)() {
    m.realm = value
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AndroidForWorkVpnConfiguration) SetRole(value *string)() {
    m.role = value
}
// SetServers sets the servers property value. List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkVpnConfiguration) SetServers(value []VpnServerable)() {
    m.servers = value
}
