package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleVpnConfiguration 
type AppleVpnConfiguration struct {
    DeviceConfiguration
    // Associated Domains
    associatedDomains []string
    // VPN Authentication Method.
    authenticationMethod *VpnAuthenticationMethod
    // Connection name displayed to the user.
    connectionName *string
    // Apple VPN connection type.
    connectionType *AppleVpnConnectionType
    // Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
    customData []KeyValueable
    // Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
    customKeyValueData []KeyValuePairable
    // Toggle to prevent user from disabling automatic VPN in the Settings app
    disableOnDemandUserOverride *bool
    // Whether to disconnect after on-demand connection idles
    disconnectOnIdle *bool
    // The length of time in seconds to wait before disconnecting an on-demand connection. Valid values 0 to 65535
    disconnectOnIdleTimerInSeconds *int32
    // Setting this to true creates Per-App VPN payload which can later be associated with Apps that can trigger this VPN conneciton on the end user's iOS device.
    enablePerApp *bool
    // Send all network traffic through VPN.
    enableSplitTunneling *bool
    // Domains that are accessed through the public internet instead of through VPN, even when per-app VPN is activated
    excludedDomains []string
    // Identifier provided by VPN vendor when connection type is set to Custom VPN. For example: Cisco AnyConnect uses an identifier of the form com.cisco.anyconnect.applevpn.plugin
    identifier *string
    // Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
    loginGroupOrDomain *string
    // On-Demand Rules. This collection can contain a maximum of 500 elements.
    onDemandRules []VpnOnDemandRuleable
    // Opt-In to sharing the device's Id to third-party vpn clients for use during network access control validation.
    optInToDeviceIdSharing *bool
    // Provider type for per-app VPN. Possible values are: notConfigured, appProxy, packetTunnel.
    providerType *VpnProviderType
    // Proxy Server.
    proxyServer VpnProxyServerable
    // Realm when connection type is set to Pulse Secure.
    realm *string
    // Role when connection type is set to Pulse Secure.
    role *string
    // Safari domains when this VPN per App setting is enabled. In addition to the apps associated with this VPN, Safari domains specified here will also be able to trigger this VPN connection.
    safariDomains []string
    // VPN Server definition.
    server VpnServerable
}
// NewAppleVpnConfiguration instantiates a new AppleVpnConfiguration and sets the default values.
func NewAppleVpnConfiguration()(*AppleVpnConfiguration) {
    m := &AppleVpnConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.appleVpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAppleVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosikEv2VpnConfiguration":
                        return NewIosikEv2VpnConfiguration(), nil
                    case "#microsoft.graph.iosVpnConfiguration":
                        return NewIosVpnConfiguration(), nil
                    case "#microsoft.graph.macOSVpnConfiguration":
                        return NewMacOSVpnConfiguration(), nil
                }
            }
        }
    }
    return NewAppleVpnConfiguration(), nil
}
// GetAssociatedDomains gets the associatedDomains property value. Associated Domains
func (m *AppleVpnConfiguration) GetAssociatedDomains()([]string) {
    return m.associatedDomains
}
// GetAuthenticationMethod gets the authenticationMethod property value. VPN Authentication Method.
func (m *AppleVpnConfiguration) GetAuthenticationMethod()(*VpnAuthenticationMethod) {
    return m.authenticationMethod
}
// GetConnectionName gets the connectionName property value. Connection name displayed to the user.
func (m *AppleVpnConfiguration) GetConnectionName()(*string) {
    return m.connectionName
}
// GetConnectionType gets the connectionType property value. Apple VPN connection type.
func (m *AppleVpnConfiguration) GetConnectionType()(*AppleVpnConnectionType) {
    return m.connectionType
}
// GetCustomData gets the customData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) GetCustomData()([]KeyValueable) {
    return m.customData
}
// GetCustomKeyValueData gets the customKeyValueData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) GetCustomKeyValueData()([]KeyValuePairable) {
    return m.customKeyValueData
}
// GetDisableOnDemandUserOverride gets the disableOnDemandUserOverride property value. Toggle to prevent user from disabling automatic VPN in the Settings app
func (m *AppleVpnConfiguration) GetDisableOnDemandUserOverride()(*bool) {
    return m.disableOnDemandUserOverride
}
// GetDisconnectOnIdle gets the disconnectOnIdle property value. Whether to disconnect after on-demand connection idles
func (m *AppleVpnConfiguration) GetDisconnectOnIdle()(*bool) {
    return m.disconnectOnIdle
}
// GetDisconnectOnIdleTimerInSeconds gets the disconnectOnIdleTimerInSeconds property value. The length of time in seconds to wait before disconnecting an on-demand connection. Valid values 0 to 65535
func (m *AppleVpnConfiguration) GetDisconnectOnIdleTimerInSeconds()(*int32) {
    return m.disconnectOnIdleTimerInSeconds
}
// GetEnablePerApp gets the enablePerApp property value. Setting this to true creates Per-App VPN payload which can later be associated with Apps that can trigger this VPN conneciton on the end user's iOS device.
func (m *AppleVpnConfiguration) GetEnablePerApp()(*bool) {
    return m.enablePerApp
}
// GetEnableSplitTunneling gets the enableSplitTunneling property value. Send all network traffic through VPN.
func (m *AppleVpnConfiguration) GetEnableSplitTunneling()(*bool) {
    return m.enableSplitTunneling
}
// GetExcludedDomains gets the excludedDomains property value. Domains that are accessed through the public internet instead of through VPN, even when per-app VPN is activated
func (m *AppleVpnConfiguration) GetExcludedDomains()([]string) {
    return m.excludedDomains
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["associatedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssociatedDomains(res)
        }
        return nil
    }
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
        val, err := n.GetEnumValue(ParseAppleVpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*AppleVpnConnectionType))
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
    res["disableOnDemandUserOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableOnDemandUserOverride(val)
        }
        return nil
    }
    res["disconnectOnIdle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisconnectOnIdle(val)
        }
        return nil
    }
    res["disconnectOnIdleTimerInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisconnectOnIdleTimerInSeconds(val)
        }
        return nil
    }
    res["enablePerApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePerApp(val)
        }
        return nil
    }
    res["enableSplitTunneling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSplitTunneling(val)
        }
        return nil
    }
    res["excludedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludedDomains(res)
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["loginGroupOrDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginGroupOrDomain(val)
        }
        return nil
    }
    res["onDemandRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnOnDemandRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnOnDemandRuleable, len(val))
            for i, v := range val {
                res[i] = v.(VpnOnDemandRuleable)
            }
            m.SetOnDemandRules(res)
        }
        return nil
    }
    res["optInToDeviceIdSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptInToDeviceIdSharing(val)
        }
        return nil
    }
    res["providerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderType(val.(*VpnProviderType))
        }
        return nil
    }
    res["proxyServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVpnProxyServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyServer(val.(VpnProxyServerable))
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
    res["safariDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSafariDomains(res)
        }
        return nil
    }
    res["server"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVpnServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServer(val.(VpnServerable))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. Identifier provided by VPN vendor when connection type is set to Custom VPN. For example: Cisco AnyConnect uses an identifier of the form com.cisco.anyconnect.applevpn.plugin
func (m *AppleVpnConfiguration) GetIdentifier()(*string) {
    return m.identifier
}
// GetLoginGroupOrDomain gets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
func (m *AppleVpnConfiguration) GetLoginGroupOrDomain()(*string) {
    return m.loginGroupOrDomain
}
// GetOnDemandRules gets the onDemandRules property value. On-Demand Rules. This collection can contain a maximum of 500 elements.
func (m *AppleVpnConfiguration) GetOnDemandRules()([]VpnOnDemandRuleable) {
    return m.onDemandRules
}
// GetOptInToDeviceIdSharing gets the optInToDeviceIdSharing property value. Opt-In to sharing the device's Id to third-party vpn clients for use during network access control validation.
func (m *AppleVpnConfiguration) GetOptInToDeviceIdSharing()(*bool) {
    return m.optInToDeviceIdSharing
}
// GetProviderType gets the providerType property value. Provider type for per-app VPN. Possible values are: notConfigured, appProxy, packetTunnel.
func (m *AppleVpnConfiguration) GetProviderType()(*VpnProviderType) {
    return m.providerType
}
// GetProxyServer gets the proxyServer property value. Proxy Server.
func (m *AppleVpnConfiguration) GetProxyServer()(VpnProxyServerable) {
    return m.proxyServer
}
// GetRealm gets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AppleVpnConfiguration) GetRealm()(*string) {
    return m.realm
}
// GetRole gets the role property value. Role when connection type is set to Pulse Secure.
func (m *AppleVpnConfiguration) GetRole()(*string) {
    return m.role
}
// GetSafariDomains gets the safariDomains property value. Safari domains when this VPN per App setting is enabled. In addition to the apps associated with this VPN, Safari domains specified here will also be able to trigger this VPN connection.
func (m *AppleVpnConfiguration) GetSafariDomains()([]string) {
    return m.safariDomains
}
// GetServer gets the server property value. VPN Server definition.
func (m *AppleVpnConfiguration) GetServer()(VpnServerable) {
    return m.server
}
// Serialize serializes information the current object
func (m *AppleVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociatedDomains() != nil {
        err = writer.WriteCollectionOfStringValues("associatedDomains", m.GetAssociatedDomains())
        if err != nil {
            return err
        }
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
        err = writer.WriteBoolValue("disableOnDemandUserOverride", m.GetDisableOnDemandUserOverride())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disconnectOnIdle", m.GetDisconnectOnIdle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("disconnectOnIdleTimerInSeconds", m.GetDisconnectOnIdleTimerInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enablePerApp", m.GetEnablePerApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSplitTunneling", m.GetEnableSplitTunneling())
        if err != nil {
            return err
        }
    }
    if m.GetExcludedDomains() != nil {
        err = writer.WriteCollectionOfStringValues("excludedDomains", m.GetExcludedDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginGroupOrDomain", m.GetLoginGroupOrDomain())
        if err != nil {
            return err
        }
    }
    if m.GetOnDemandRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnDemandRules()))
        for i, v := range m.GetOnDemandRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("onDemandRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("optInToDeviceIdSharing", m.GetOptInToDeviceIdSharing())
        if err != nil {
            return err
        }
    }
    if m.GetProviderType() != nil {
        cast := (*m.GetProviderType()).String()
        err = writer.WriteStringValue("providerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("proxyServer", m.GetProxyServer())
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
    if m.GetSafariDomains() != nil {
        err = writer.WriteCollectionOfStringValues("safariDomains", m.GetSafariDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("server", m.GetServer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociatedDomains sets the associatedDomains property value. Associated Domains
func (m *AppleVpnConfiguration) SetAssociatedDomains(value []string)() {
    m.associatedDomains = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. VPN Authentication Method.
func (m *AppleVpnConfiguration) SetAuthenticationMethod(value *VpnAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetConnectionName sets the connectionName property value. Connection name displayed to the user.
func (m *AppleVpnConfiguration) SetConnectionName(value *string)() {
    m.connectionName = value
}
// SetConnectionType sets the connectionType property value. Apple VPN connection type.
func (m *AppleVpnConfiguration) SetConnectionType(value *AppleVpnConnectionType)() {
    m.connectionType = value
}
// SetCustomData sets the customData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) SetCustomData(value []KeyValueable)() {
    m.customData = value
}
// SetCustomKeyValueData sets the customKeyValueData property value. Custom data when connection type is set to Custom VPN. Use this field to enable functionality not supported by Intune, but available in your VPN solution. Contact your VPN vendor to learn how to add these key/value pairs. This collection can contain a maximum of 25 elements.
func (m *AppleVpnConfiguration) SetCustomKeyValueData(value []KeyValuePairable)() {
    m.customKeyValueData = value
}
// SetDisableOnDemandUserOverride sets the disableOnDemandUserOverride property value. Toggle to prevent user from disabling automatic VPN in the Settings app
func (m *AppleVpnConfiguration) SetDisableOnDemandUserOverride(value *bool)() {
    m.disableOnDemandUserOverride = value
}
// SetDisconnectOnIdle sets the disconnectOnIdle property value. Whether to disconnect after on-demand connection idles
func (m *AppleVpnConfiguration) SetDisconnectOnIdle(value *bool)() {
    m.disconnectOnIdle = value
}
// SetDisconnectOnIdleTimerInSeconds sets the disconnectOnIdleTimerInSeconds property value. The length of time in seconds to wait before disconnecting an on-demand connection. Valid values 0 to 65535
func (m *AppleVpnConfiguration) SetDisconnectOnIdleTimerInSeconds(value *int32)() {
    m.disconnectOnIdleTimerInSeconds = value
}
// SetEnablePerApp sets the enablePerApp property value. Setting this to true creates Per-App VPN payload which can later be associated with Apps that can trigger this VPN conneciton on the end user's iOS device.
func (m *AppleVpnConfiguration) SetEnablePerApp(value *bool)() {
    m.enablePerApp = value
}
// SetEnableSplitTunneling sets the enableSplitTunneling property value. Send all network traffic through VPN.
func (m *AppleVpnConfiguration) SetEnableSplitTunneling(value *bool)() {
    m.enableSplitTunneling = value
}
// SetExcludedDomains sets the excludedDomains property value. Domains that are accessed through the public internet instead of through VPN, even when per-app VPN is activated
func (m *AppleVpnConfiguration) SetExcludedDomains(value []string)() {
    m.excludedDomains = value
}
// SetIdentifier sets the identifier property value. Identifier provided by VPN vendor when connection type is set to Custom VPN. For example: Cisco AnyConnect uses an identifier of the form com.cisco.anyconnect.applevpn.plugin
func (m *AppleVpnConfiguration) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetLoginGroupOrDomain sets the loginGroupOrDomain property value. Login group or domain when connection type is set to Dell SonicWALL Mobile Connection.
func (m *AppleVpnConfiguration) SetLoginGroupOrDomain(value *string)() {
    m.loginGroupOrDomain = value
}
// SetOnDemandRules sets the onDemandRules property value. On-Demand Rules. This collection can contain a maximum of 500 elements.
func (m *AppleVpnConfiguration) SetOnDemandRules(value []VpnOnDemandRuleable)() {
    m.onDemandRules = value
}
// SetOptInToDeviceIdSharing sets the optInToDeviceIdSharing property value. Opt-In to sharing the device's Id to third-party vpn clients for use during network access control validation.
func (m *AppleVpnConfiguration) SetOptInToDeviceIdSharing(value *bool)() {
    m.optInToDeviceIdSharing = value
}
// SetProviderType sets the providerType property value. Provider type for per-app VPN. Possible values are: notConfigured, appProxy, packetTunnel.
func (m *AppleVpnConfiguration) SetProviderType(value *VpnProviderType)() {
    m.providerType = value
}
// SetProxyServer sets the proxyServer property value. Proxy Server.
func (m *AppleVpnConfiguration) SetProxyServer(value VpnProxyServerable)() {
    m.proxyServer = value
}
// SetRealm sets the realm property value. Realm when connection type is set to Pulse Secure.
func (m *AppleVpnConfiguration) SetRealm(value *string)() {
    m.realm = value
}
// SetRole sets the role property value. Role when connection type is set to Pulse Secure.
func (m *AppleVpnConfiguration) SetRole(value *string)() {
    m.role = value
}
// SetSafariDomains sets the safariDomains property value. Safari domains when this VPN per App setting is enabled. In addition to the apps associated with this VPN, Safari domains specified here will also be able to trigger this VPN connection.
func (m *AppleVpnConfiguration) SetSafariDomains(value []string)() {
    m.safariDomains = value
}
// SetServer sets the server property value. VPN Server definition.
func (m *AppleVpnConfiguration) SetServer(value VpnServerable)() {
    m.server = value
}
