package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10VpnConfiguration 
type Windows10VpnConfiguration struct {
    WindowsVpnConfiguration
    // Associated Apps. This collection can contain a maximum of 10000 elements.
    associatedApps []Windows10AssociatedAppsable
    // Windows 10 VPN connection types.
    authenticationMethod *Windows10VpnAuthenticationMethod
    // VPN connection types.
    connectionType *Windows10VpnConnectionType
    // Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
    cryptographySuite CryptographySuiteable
    // DNS rules. This collection can contain a maximum of 1000 elements.
    dnsRules []VpnDnsRuleable
    // Specify DNS suffixes to add to the DNS search list to properly route short names.
    dnsSuffixes []string
    // Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
    eapXml []byte
    // Enable Always On mode.
    enableAlwaysOn *bool
    // Enable conditional access.
    enableConditionalAccess *bool
    // Enable device tunnel.
    enableDeviceTunnel *bool
    // Enable IP address registration with internal DNS.
    enableDnsRegistration *bool
    // Enable single sign-on (SSO) with alternate certificate.
    enableSingleSignOnWithAlternateCertificate *bool
    // Enable split tunneling.
    enableSplitTunneling *bool
    // Identity certificate for client authentication when authentication method is certificate.
    identityCertificate WindowsCertificateProfileBaseable
    // ID of the Microsoft Tunnel site associated with the VPN profile.
    microsoftTunnelSiteId *string
    // Only associated Apps can use connection (per-app VPN).
    onlyAssociatedAppsCanUseConnection *bool
    // Profile target type. Possible values are: user, device, autoPilotDevice.
    profileTarget *Windows10VpnProfileTarget
    // Proxy Server.
    proxyServer Windows10VpnProxyServerable
    // Remember user credentials.
    rememberUserCredentials *bool
    // Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
    routes []VpnRouteable
    // Single sign-on Extended Key Usage (EKU).
    singleSignOnEku ExtendedKeyUsageable
    // Single sign-on issuer hash.
    singleSignOnIssuerHash *string
    // Traffic rules. This collection can contain a maximum of 1000 elements.
    trafficRules []VpnTrafficRuleable
    // Trusted Network Domains
    trustedNetworkDomains []string
    // Windows Information Protection (WIP) domain to associate with this connection.
    windowsInformationProtectionDomain *string
}
// NewWindows10VpnConfiguration instantiates a new Windows10VpnConfiguration and sets the default values.
func NewWindows10VpnConfiguration()(*Windows10VpnConfiguration) {
    m := &Windows10VpnConfiguration{
        WindowsVpnConfiguration: *NewWindowsVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10VpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10VpnConfiguration(), nil
}
// GetAssociatedApps gets the associatedApps property value. Associated Apps. This collection can contain a maximum of 10000 elements.
func (m *Windows10VpnConfiguration) GetAssociatedApps()([]Windows10AssociatedAppsable) {
    return m.associatedApps
}
// GetAuthenticationMethod gets the authenticationMethod property value. Windows 10 VPN connection types.
func (m *Windows10VpnConfiguration) GetAuthenticationMethod()(*Windows10VpnAuthenticationMethod) {
    return m.authenticationMethod
}
// GetConnectionType gets the connectionType property value. VPN connection types.
func (m *Windows10VpnConfiguration) GetConnectionType()(*Windows10VpnConnectionType) {
    return m.connectionType
}
// GetCryptographySuite gets the cryptographySuite property value. Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
func (m *Windows10VpnConfiguration) GetCryptographySuite()(CryptographySuiteable) {
    return m.cryptographySuite
}
// GetDnsRules gets the dnsRules property value. DNS rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetDnsRules()([]VpnDnsRuleable) {
    return m.dnsRules
}
// GetDnsSuffixes gets the dnsSuffixes property value. Specify DNS suffixes to add to the DNS search list to properly route short names.
func (m *Windows10VpnConfiguration) GetDnsSuffixes()([]string) {
    return m.dnsSuffixes
}
// GetEapXml gets the eapXml property value. Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
func (m *Windows10VpnConfiguration) GetEapXml()([]byte) {
    return m.eapXml
}
// GetEnableAlwaysOn gets the enableAlwaysOn property value. Enable Always On mode.
func (m *Windows10VpnConfiguration) GetEnableAlwaysOn()(*bool) {
    return m.enableAlwaysOn
}
// GetEnableConditionalAccess gets the enableConditionalAccess property value. Enable conditional access.
func (m *Windows10VpnConfiguration) GetEnableConditionalAccess()(*bool) {
    return m.enableConditionalAccess
}
// GetEnableDeviceTunnel gets the enableDeviceTunnel property value. Enable device tunnel.
func (m *Windows10VpnConfiguration) GetEnableDeviceTunnel()(*bool) {
    return m.enableDeviceTunnel
}
// GetEnableDnsRegistration gets the enableDnsRegistration property value. Enable IP address registration with internal DNS.
func (m *Windows10VpnConfiguration) GetEnableDnsRegistration()(*bool) {
    return m.enableDnsRegistration
}
// GetEnableSingleSignOnWithAlternateCertificate gets the enableSingleSignOnWithAlternateCertificate property value. Enable single sign-on (SSO) with alternate certificate.
func (m *Windows10VpnConfiguration) GetEnableSingleSignOnWithAlternateCertificate()(*bool) {
    return m.enableSingleSignOnWithAlternateCertificate
}
// GetEnableSplitTunneling gets the enableSplitTunneling property value. Enable split tunneling.
func (m *Windows10VpnConfiguration) GetEnableSplitTunneling()(*bool) {
    return m.enableSplitTunneling
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsVpnConfiguration.GetFieldDeserializers()
    res["associatedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10AssociatedAppsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10AssociatedAppsable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10AssociatedAppsable)
            }
            m.SetAssociatedApps(res)
        }
        return nil
    }
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10VpnAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*Windows10VpnAuthenticationMethod))
        }
        return nil
    }
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10VpnConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*Windows10VpnConnectionType))
        }
        return nil
    }
    res["cryptographySuite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCryptographySuiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCryptographySuite(val.(CryptographySuiteable))
        }
        return nil
    }
    res["dnsRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnDnsRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnDnsRuleable, len(val))
            for i, v := range val {
                res[i] = v.(VpnDnsRuleable)
            }
            m.SetDnsRules(res)
        }
        return nil
    }
    res["dnsSuffixes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDnsSuffixes(res)
        }
        return nil
    }
    res["eapXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapXml(val)
        }
        return nil
    }
    res["enableAlwaysOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAlwaysOn(val)
        }
        return nil
    }
    res["enableConditionalAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableConditionalAccess(val)
        }
        return nil
    }
    res["enableDeviceTunnel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDeviceTunnel(val)
        }
        return nil
    }
    res["enableDnsRegistration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDnsRegistration(val)
        }
        return nil
    }
    res["enableSingleSignOnWithAlternateCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleSignOnWithAlternateCertificate(val)
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
    res["identityCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificate(val.(WindowsCertificateProfileBaseable))
        }
        return nil
    }
    res["microsoftTunnelSiteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftTunnelSiteId(val)
        }
        return nil
    }
    res["onlyAssociatedAppsCanUseConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlyAssociatedAppsCanUseConnection(val)
        }
        return nil
    }
    res["profileTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10VpnProfileTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileTarget(val.(*Windows10VpnProfileTarget))
        }
        return nil
    }
    res["proxyServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows10VpnProxyServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyServer(val.(Windows10VpnProxyServerable))
        }
        return nil
    }
    res["rememberUserCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRememberUserCredentials(val)
        }
        return nil
    }
    res["routes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnRouteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnRouteable, len(val))
            for i, v := range val {
                res[i] = v.(VpnRouteable)
            }
            m.SetRoutes(res)
        }
        return nil
    }
    res["singleSignOnEku"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExtendedKeyUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnEku(val.(ExtendedKeyUsageable))
        }
        return nil
    }
    res["singleSignOnIssuerHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnIssuerHash(val)
        }
        return nil
    }
    res["trafficRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVpnTrafficRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VpnTrafficRuleable, len(val))
            for i, v := range val {
                res[i] = v.(VpnTrafficRuleable)
            }
            m.SetTrafficRules(res)
        }
        return nil
    }
    res["trustedNetworkDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTrustedNetworkDomains(res)
        }
        return nil
    }
    res["windowsInformationProtectionDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsInformationProtectionDomain(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificate gets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *Windows10VpnConfiguration) GetIdentityCertificate()(WindowsCertificateProfileBaseable) {
    return m.identityCertificate
}
// GetMicrosoftTunnelSiteId gets the microsoftTunnelSiteId property value. ID of the Microsoft Tunnel site associated with the VPN profile.
func (m *Windows10VpnConfiguration) GetMicrosoftTunnelSiteId()(*string) {
    return m.microsoftTunnelSiteId
}
// GetOnlyAssociatedAppsCanUseConnection gets the onlyAssociatedAppsCanUseConnection property value. Only associated Apps can use connection (per-app VPN).
func (m *Windows10VpnConfiguration) GetOnlyAssociatedAppsCanUseConnection()(*bool) {
    return m.onlyAssociatedAppsCanUseConnection
}
// GetProfileTarget gets the profileTarget property value. Profile target type. Possible values are: user, device, autoPilotDevice.
func (m *Windows10VpnConfiguration) GetProfileTarget()(*Windows10VpnProfileTarget) {
    return m.profileTarget
}
// GetProxyServer gets the proxyServer property value. Proxy Server.
func (m *Windows10VpnConfiguration) GetProxyServer()(Windows10VpnProxyServerable) {
    return m.proxyServer
}
// GetRememberUserCredentials gets the rememberUserCredentials property value. Remember user credentials.
func (m *Windows10VpnConfiguration) GetRememberUserCredentials()(*bool) {
    return m.rememberUserCredentials
}
// GetRoutes gets the routes property value. Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetRoutes()([]VpnRouteable) {
    return m.routes
}
// GetSingleSignOnEku gets the singleSignOnEku property value. Single sign-on Extended Key Usage (EKU).
func (m *Windows10VpnConfiguration) GetSingleSignOnEku()(ExtendedKeyUsageable) {
    return m.singleSignOnEku
}
// GetSingleSignOnIssuerHash gets the singleSignOnIssuerHash property value. Single sign-on issuer hash.
func (m *Windows10VpnConfiguration) GetSingleSignOnIssuerHash()(*string) {
    return m.singleSignOnIssuerHash
}
// GetTrafficRules gets the trafficRules property value. Traffic rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) GetTrafficRules()([]VpnTrafficRuleable) {
    return m.trafficRules
}
// GetTrustedNetworkDomains gets the trustedNetworkDomains property value. Trusted Network Domains
func (m *Windows10VpnConfiguration) GetTrustedNetworkDomains()([]string) {
    return m.trustedNetworkDomains
}
// GetWindowsInformationProtectionDomain gets the windowsInformationProtectionDomain property value. Windows Information Protection (WIP) domain to associate with this connection.
func (m *Windows10VpnConfiguration) GetWindowsInformationProtectionDomain()(*string) {
    return m.windowsInformationProtectionDomain
}
// Serialize serializes information the current object
func (m *Windows10VpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociatedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociatedApps()))
        for i, v := range m.GetAssociatedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("associatedApps", cast)
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
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cryptographySuite", m.GetCryptographySuite())
        if err != nil {
            return err
        }
    }
    if m.GetDnsRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDnsRules()))
        for i, v := range m.GetDnsRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dnsRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDnsSuffixes() != nil {
        err = writer.WriteCollectionOfStringValues("dnsSuffixes", m.GetDnsSuffixes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("eapXml", m.GetEapXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableAlwaysOn", m.GetEnableAlwaysOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableConditionalAccess", m.GetEnableConditionalAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableDeviceTunnel", m.GetEnableDeviceTunnel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableDnsRegistration", m.GetEnableDnsRegistration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSingleSignOnWithAlternateCertificate", m.GetEnableSingleSignOnWithAlternateCertificate())
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
    {
        err = writer.WriteObjectValue("identityCertificate", m.GetIdentityCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftTunnelSiteId", m.GetMicrosoftTunnelSiteId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onlyAssociatedAppsCanUseConnection", m.GetOnlyAssociatedAppsCanUseConnection())
        if err != nil {
            return err
        }
    }
    if m.GetProfileTarget() != nil {
        cast := (*m.GetProfileTarget()).String()
        err = writer.WriteStringValue("profileTarget", &cast)
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
        err = writer.WriteBoolValue("rememberUserCredentials", m.GetRememberUserCredentials())
        if err != nil {
            return err
        }
    }
    if m.GetRoutes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoutes()))
        for i, v := range m.GetRoutes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("routes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnEku", m.GetSingleSignOnEku())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("singleSignOnIssuerHash", m.GetSingleSignOnIssuerHash())
        if err != nil {
            return err
        }
    }
    if m.GetTrafficRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrafficRules()))
        for i, v := range m.GetTrafficRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("trafficRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrustedNetworkDomains() != nil {
        err = writer.WriteCollectionOfStringValues("trustedNetworkDomains", m.GetTrustedNetworkDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windowsInformationProtectionDomain", m.GetWindowsInformationProtectionDomain())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociatedApps sets the associatedApps property value. Associated Apps. This collection can contain a maximum of 10000 elements.
func (m *Windows10VpnConfiguration) SetAssociatedApps(value []Windows10AssociatedAppsable)() {
    m.associatedApps = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. Windows 10 VPN connection types.
func (m *Windows10VpnConfiguration) SetAuthenticationMethod(value *Windows10VpnAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetConnectionType sets the connectionType property value. VPN connection types.
func (m *Windows10VpnConfiguration) SetConnectionType(value *Windows10VpnConnectionType)() {
    m.connectionType = value
}
// SetCryptographySuite sets the cryptographySuite property value. Cryptography Suite security settings for IKEv2 VPN in Windows10 and above
func (m *Windows10VpnConfiguration) SetCryptographySuite(value CryptographySuiteable)() {
    m.cryptographySuite = value
}
// SetDnsRules sets the dnsRules property value. DNS rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetDnsRules(value []VpnDnsRuleable)() {
    m.dnsRules = value
}
// SetDnsSuffixes sets the dnsSuffixes property value. Specify DNS suffixes to add to the DNS search list to properly route short names.
func (m *Windows10VpnConfiguration) SetDnsSuffixes(value []string)() {
    m.dnsSuffixes = value
}
// SetEapXml sets the eapXml property value. Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
func (m *Windows10VpnConfiguration) SetEapXml(value []byte)() {
    m.eapXml = value
}
// SetEnableAlwaysOn sets the enableAlwaysOn property value. Enable Always On mode.
func (m *Windows10VpnConfiguration) SetEnableAlwaysOn(value *bool)() {
    m.enableAlwaysOn = value
}
// SetEnableConditionalAccess sets the enableConditionalAccess property value. Enable conditional access.
func (m *Windows10VpnConfiguration) SetEnableConditionalAccess(value *bool)() {
    m.enableConditionalAccess = value
}
// SetEnableDeviceTunnel sets the enableDeviceTunnel property value. Enable device tunnel.
func (m *Windows10VpnConfiguration) SetEnableDeviceTunnel(value *bool)() {
    m.enableDeviceTunnel = value
}
// SetEnableDnsRegistration sets the enableDnsRegistration property value. Enable IP address registration with internal DNS.
func (m *Windows10VpnConfiguration) SetEnableDnsRegistration(value *bool)() {
    m.enableDnsRegistration = value
}
// SetEnableSingleSignOnWithAlternateCertificate sets the enableSingleSignOnWithAlternateCertificate property value. Enable single sign-on (SSO) with alternate certificate.
func (m *Windows10VpnConfiguration) SetEnableSingleSignOnWithAlternateCertificate(value *bool)() {
    m.enableSingleSignOnWithAlternateCertificate = value
}
// SetEnableSplitTunneling sets the enableSplitTunneling property value. Enable split tunneling.
func (m *Windows10VpnConfiguration) SetEnableSplitTunneling(value *bool)() {
    m.enableSplitTunneling = value
}
// SetIdentityCertificate sets the identityCertificate property value. Identity certificate for client authentication when authentication method is certificate.
func (m *Windows10VpnConfiguration) SetIdentityCertificate(value WindowsCertificateProfileBaseable)() {
    m.identityCertificate = value
}
// SetMicrosoftTunnelSiteId sets the microsoftTunnelSiteId property value. ID of the Microsoft Tunnel site associated with the VPN profile.
func (m *Windows10VpnConfiguration) SetMicrosoftTunnelSiteId(value *string)() {
    m.microsoftTunnelSiteId = value
}
// SetOnlyAssociatedAppsCanUseConnection sets the onlyAssociatedAppsCanUseConnection property value. Only associated Apps can use connection (per-app VPN).
func (m *Windows10VpnConfiguration) SetOnlyAssociatedAppsCanUseConnection(value *bool)() {
    m.onlyAssociatedAppsCanUseConnection = value
}
// SetProfileTarget sets the profileTarget property value. Profile target type. Possible values are: user, device, autoPilotDevice.
func (m *Windows10VpnConfiguration) SetProfileTarget(value *Windows10VpnProfileTarget)() {
    m.profileTarget = value
}
// SetProxyServer sets the proxyServer property value. Proxy Server.
func (m *Windows10VpnConfiguration) SetProxyServer(value Windows10VpnProxyServerable)() {
    m.proxyServer = value
}
// SetRememberUserCredentials sets the rememberUserCredentials property value. Remember user credentials.
func (m *Windows10VpnConfiguration) SetRememberUserCredentials(value *bool)() {
    m.rememberUserCredentials = value
}
// SetRoutes sets the routes property value. Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetRoutes(value []VpnRouteable)() {
    m.routes = value
}
// SetSingleSignOnEku sets the singleSignOnEku property value. Single sign-on Extended Key Usage (EKU).
func (m *Windows10VpnConfiguration) SetSingleSignOnEku(value ExtendedKeyUsageable)() {
    m.singleSignOnEku = value
}
// SetSingleSignOnIssuerHash sets the singleSignOnIssuerHash property value. Single sign-on issuer hash.
func (m *Windows10VpnConfiguration) SetSingleSignOnIssuerHash(value *string)() {
    m.singleSignOnIssuerHash = value
}
// SetTrafficRules sets the trafficRules property value. Traffic rules. This collection can contain a maximum of 1000 elements.
func (m *Windows10VpnConfiguration) SetTrafficRules(value []VpnTrafficRuleable)() {
    m.trafficRules = value
}
// SetTrustedNetworkDomains sets the trustedNetworkDomains property value. Trusted Network Domains
func (m *Windows10VpnConfiguration) SetTrustedNetworkDomains(value []string)() {
    m.trustedNetworkDomains = value
}
// SetWindowsInformationProtectionDomain sets the windowsInformationProtectionDomain property value. Windows Information Protection (WIP) domain to associate with this connection.
func (m *Windows10VpnConfiguration) SetWindowsInformationProtectionDomain(value *string)() {
    m.windowsInformationProtectionDomain = value
}
