package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosikEv2VpnConfiguration 
type IosikEv2VpnConfiguration struct {
    IosVpnConfiguration
    // Allows the use of child security association parameters by setting all parameters to the device's default unless explicitly specified.
    allowDefaultChildSecurityAssociationParameters *bool
    // Allows the use of security association parameters by setting all parameters to the device's default unless explicitly specified.
    allowDefaultSecurityAssociationParameters *bool
    // AlwaysOn Configuration
    alwaysOnConfiguration AppleVpnAlwaysOnConfigurationable
    // Child Security Association Parameters
    childSecurityAssociationParameters IosVpnSecurityAssociationParametersable
    // The type of VPN client authentication type
    clientAuthenticationType *VpnClientAuthenticationType
    // Determine how often to check if a peer connection is still active. . Possible values are: medium, none, low, high.
    deadPeerDetectionRate *VpnDeadPeerDetectionRate
    // Disable MOBIKE
    disableMobilityAndMultihoming *bool
    // Disable Redirect
    disableRedirect *bool
    // Determines if Always on VPN is enabled
    enableAlwaysOnConfiguration *bool
    // Enables a best-effort revocation check; server response timeouts will not cause it to fail
    enableCertificateRevocationCheck *bool
    // Enables EAP only authentication
    enableEAP *bool
    // Enable Perfect Forward Secrecy (PFS).
    enablePerfectForwardSecrecy *bool
    // Enable Use Internal Subnet Attributes.
    enableUseInternalSubnetAttributes *bool
    // The type of VPN local identifier
    localIdentifier *VpnLocalIdentifier
    // Maximum transmission unit. Valid values 1280 to 1400
    mtuSizeInBytes *int32
    // Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
    remoteIdentifier *string
    // Security Association Parameters
    securityAssociationParameters IosVpnSecurityAssociationParametersable
    // Common name of the IKEv2 Server Certificate used in Server Authentication
    serverCertificateCommonName *string
    // Issuer Common name of the IKEv2 Server Certificate issuer used in Authentication
    serverCertificateIssuerCommonName *string
    // The type of certificate the VPN server will present to the VPN client for authentication. Possible values are: rsa, ecdsa256, ecdsa384, ecdsa521.
    serverCertificateType *VpnServerCertificateType
    // Used when Shared Secret Authentication is selected
    sharedSecret *string
    // The maximum TLS version to be used with EAP-TLS authentication
    tlsMaximumVersion *string
    // The minimum TLS version to be used with EAP-TLS authentication
    tlsMinimumVersion *string
}
// NewIosikEv2VpnConfiguration instantiates a new IosikEv2VpnConfiguration and sets the default values.
func NewIosikEv2VpnConfiguration()(*IosikEv2VpnConfiguration) {
    m := &IosikEv2VpnConfiguration{
        IosVpnConfiguration: *NewIosVpnConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosikEv2VpnConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosikEv2VpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosikEv2VpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosikEv2VpnConfiguration(), nil
}
// GetAllowDefaultChildSecurityAssociationParameters gets the allowDefaultChildSecurityAssociationParameters property value. Allows the use of child security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) GetAllowDefaultChildSecurityAssociationParameters()(*bool) {
    return m.allowDefaultChildSecurityAssociationParameters
}
// GetAllowDefaultSecurityAssociationParameters gets the allowDefaultSecurityAssociationParameters property value. Allows the use of security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) GetAllowDefaultSecurityAssociationParameters()(*bool) {
    return m.allowDefaultSecurityAssociationParameters
}
// GetAlwaysOnConfiguration gets the alwaysOnConfiguration property value. AlwaysOn Configuration
func (m *IosikEv2VpnConfiguration) GetAlwaysOnConfiguration()(AppleVpnAlwaysOnConfigurationable) {
    return m.alwaysOnConfiguration
}
// GetChildSecurityAssociationParameters gets the childSecurityAssociationParameters property value. Child Security Association Parameters
func (m *IosikEv2VpnConfiguration) GetChildSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable) {
    return m.childSecurityAssociationParameters
}
// GetClientAuthenticationType gets the clientAuthenticationType property value. The type of VPN client authentication type
func (m *IosikEv2VpnConfiguration) GetClientAuthenticationType()(*VpnClientAuthenticationType) {
    return m.clientAuthenticationType
}
// GetDeadPeerDetectionRate gets the deadPeerDetectionRate property value. Determine how often to check if a peer connection is still active. . Possible values are: medium, none, low, high.
func (m *IosikEv2VpnConfiguration) GetDeadPeerDetectionRate()(*VpnDeadPeerDetectionRate) {
    return m.deadPeerDetectionRate
}
// GetDisableMobilityAndMultihoming gets the disableMobilityAndMultihoming property value. Disable MOBIKE
func (m *IosikEv2VpnConfiguration) GetDisableMobilityAndMultihoming()(*bool) {
    return m.disableMobilityAndMultihoming
}
// GetDisableRedirect gets the disableRedirect property value. Disable Redirect
func (m *IosikEv2VpnConfiguration) GetDisableRedirect()(*bool) {
    return m.disableRedirect
}
// GetEnableAlwaysOnConfiguration gets the enableAlwaysOnConfiguration property value. Determines if Always on VPN is enabled
func (m *IosikEv2VpnConfiguration) GetEnableAlwaysOnConfiguration()(*bool) {
    return m.enableAlwaysOnConfiguration
}
// GetEnableCertificateRevocationCheck gets the enableCertificateRevocationCheck property value. Enables a best-effort revocation check; server response timeouts will not cause it to fail
func (m *IosikEv2VpnConfiguration) GetEnableCertificateRevocationCheck()(*bool) {
    return m.enableCertificateRevocationCheck
}
// GetEnableEAP gets the enableEAP property value. Enables EAP only authentication
func (m *IosikEv2VpnConfiguration) GetEnableEAP()(*bool) {
    return m.enableEAP
}
// GetEnablePerfectForwardSecrecy gets the enablePerfectForwardSecrecy property value. Enable Perfect Forward Secrecy (PFS).
func (m *IosikEv2VpnConfiguration) GetEnablePerfectForwardSecrecy()(*bool) {
    return m.enablePerfectForwardSecrecy
}
// GetEnableUseInternalSubnetAttributes gets the enableUseInternalSubnetAttributes property value. Enable Use Internal Subnet Attributes.
func (m *IosikEv2VpnConfiguration) GetEnableUseInternalSubnetAttributes()(*bool) {
    return m.enableUseInternalSubnetAttributes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosikEv2VpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosVpnConfiguration.GetFieldDeserializers()
    res["allowDefaultChildSecurityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDefaultChildSecurityAssociationParameters(val)
        }
        return nil
    }
    res["allowDefaultSecurityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDefaultSecurityAssociationParameters(val)
        }
        return nil
    }
    res["alwaysOnConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppleVpnAlwaysOnConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlwaysOnConfiguration(val.(AppleVpnAlwaysOnConfigurationable))
        }
        return nil
    }
    res["childSecurityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildSecurityAssociationParameters(val.(IosVpnSecurityAssociationParametersable))
        }
        return nil
    }
    res["clientAuthenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnClientAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAuthenticationType(val.(*VpnClientAuthenticationType))
        }
        return nil
    }
    res["deadPeerDetectionRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnDeadPeerDetectionRate)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadPeerDetectionRate(val.(*VpnDeadPeerDetectionRate))
        }
        return nil
    }
    res["disableMobilityAndMultihoming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableMobilityAndMultihoming(val)
        }
        return nil
    }
    res["disableRedirect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableRedirect(val)
        }
        return nil
    }
    res["enableAlwaysOnConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAlwaysOnConfiguration(val)
        }
        return nil
    }
    res["enableCertificateRevocationCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableCertificateRevocationCheck(val)
        }
        return nil
    }
    res["enableEAP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableEAP(val)
        }
        return nil
    }
    res["enablePerfectForwardSecrecy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePerfectForwardSecrecy(val)
        }
        return nil
    }
    res["enableUseInternalSubnetAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableUseInternalSubnetAttributes(val)
        }
        return nil
    }
    res["localIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnLocalIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalIdentifier(val.(*VpnLocalIdentifier))
        }
        return nil
    }
    res["mtuSizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMtuSizeInBytes(val)
        }
        return nil
    }
    res["remoteIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteIdentifier(val)
        }
        return nil
    }
    res["securityAssociationParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAssociationParameters(val.(IosVpnSecurityAssociationParametersable))
        }
        return nil
    }
    res["serverCertificateCommonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCertificateCommonName(val)
        }
        return nil
    }
    res["serverCertificateIssuerCommonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCertificateIssuerCommonName(val)
        }
        return nil
    }
    res["serverCertificateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnServerCertificateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCertificateType(val.(*VpnServerCertificateType))
        }
        return nil
    }
    res["sharedSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedSecret(val)
        }
        return nil
    }
    res["tlsMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlsMaximumVersion(val)
        }
        return nil
    }
    res["tlsMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlsMinimumVersion(val)
        }
        return nil
    }
    return res
}
// GetLocalIdentifier gets the localIdentifier property value. The type of VPN local identifier
func (m *IosikEv2VpnConfiguration) GetLocalIdentifier()(*VpnLocalIdentifier) {
    return m.localIdentifier
}
// GetMtuSizeInBytes gets the mtuSizeInBytes property value. Maximum transmission unit. Valid values 1280 to 1400
func (m *IosikEv2VpnConfiguration) GetMtuSizeInBytes()(*int32) {
    return m.mtuSizeInBytes
}
// GetRemoteIdentifier gets the remoteIdentifier property value. Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
func (m *IosikEv2VpnConfiguration) GetRemoteIdentifier()(*string) {
    return m.remoteIdentifier
}
// GetSecurityAssociationParameters gets the securityAssociationParameters property value. Security Association Parameters
func (m *IosikEv2VpnConfiguration) GetSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable) {
    return m.securityAssociationParameters
}
// GetServerCertificateCommonName gets the serverCertificateCommonName property value. Common name of the IKEv2 Server Certificate used in Server Authentication
func (m *IosikEv2VpnConfiguration) GetServerCertificateCommonName()(*string) {
    return m.serverCertificateCommonName
}
// GetServerCertificateIssuerCommonName gets the serverCertificateIssuerCommonName property value. Issuer Common name of the IKEv2 Server Certificate issuer used in Authentication
func (m *IosikEv2VpnConfiguration) GetServerCertificateIssuerCommonName()(*string) {
    return m.serverCertificateIssuerCommonName
}
// GetServerCertificateType gets the serverCertificateType property value. The type of certificate the VPN server will present to the VPN client for authentication. Possible values are: rsa, ecdsa256, ecdsa384, ecdsa521.
func (m *IosikEv2VpnConfiguration) GetServerCertificateType()(*VpnServerCertificateType) {
    return m.serverCertificateType
}
// GetSharedSecret gets the sharedSecret property value. Used when Shared Secret Authentication is selected
func (m *IosikEv2VpnConfiguration) GetSharedSecret()(*string) {
    return m.sharedSecret
}
// GetTlsMaximumVersion gets the tlsMaximumVersion property value. The maximum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) GetTlsMaximumVersion()(*string) {
    return m.tlsMaximumVersion
}
// GetTlsMinimumVersion gets the tlsMinimumVersion property value. The minimum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) GetTlsMinimumVersion()(*string) {
    return m.tlsMinimumVersion
}
// Serialize serializes information the current object
func (m *IosikEv2VpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosVpnConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowDefaultChildSecurityAssociationParameters", m.GetAllowDefaultChildSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowDefaultSecurityAssociationParameters", m.GetAllowDefaultSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("alwaysOnConfiguration", m.GetAlwaysOnConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("childSecurityAssociationParameters", m.GetChildSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    if m.GetClientAuthenticationType() != nil {
        cast := (*m.GetClientAuthenticationType()).String()
        err = writer.WriteStringValue("clientAuthenticationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeadPeerDetectionRate() != nil {
        cast := (*m.GetDeadPeerDetectionRate()).String()
        err = writer.WriteStringValue("deadPeerDetectionRate", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableMobilityAndMultihoming", m.GetDisableMobilityAndMultihoming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableRedirect", m.GetDisableRedirect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableAlwaysOnConfiguration", m.GetEnableAlwaysOnConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableCertificateRevocationCheck", m.GetEnableCertificateRevocationCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableEAP", m.GetEnableEAP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enablePerfectForwardSecrecy", m.GetEnablePerfectForwardSecrecy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableUseInternalSubnetAttributes", m.GetEnableUseInternalSubnetAttributes())
        if err != nil {
            return err
        }
    }
    if m.GetLocalIdentifier() != nil {
        cast := (*m.GetLocalIdentifier()).String()
        err = writer.WriteStringValue("localIdentifier", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mtuSizeInBytes", m.GetMtuSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remoteIdentifier", m.GetRemoteIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("securityAssociationParameters", m.GetSecurityAssociationParameters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverCertificateCommonName", m.GetServerCertificateCommonName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverCertificateIssuerCommonName", m.GetServerCertificateIssuerCommonName())
        if err != nil {
            return err
        }
    }
    if m.GetServerCertificateType() != nil {
        cast := (*m.GetServerCertificateType()).String()
        err = writer.WriteStringValue("serverCertificateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sharedSecret", m.GetSharedSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tlsMaximumVersion", m.GetTlsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tlsMinimumVersion", m.GetTlsMinimumVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowDefaultChildSecurityAssociationParameters sets the allowDefaultChildSecurityAssociationParameters property value. Allows the use of child security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) SetAllowDefaultChildSecurityAssociationParameters(value *bool)() {
    m.allowDefaultChildSecurityAssociationParameters = value
}
// SetAllowDefaultSecurityAssociationParameters sets the allowDefaultSecurityAssociationParameters property value. Allows the use of security association parameters by setting all parameters to the device's default unless explicitly specified.
func (m *IosikEv2VpnConfiguration) SetAllowDefaultSecurityAssociationParameters(value *bool)() {
    m.allowDefaultSecurityAssociationParameters = value
}
// SetAlwaysOnConfiguration sets the alwaysOnConfiguration property value. AlwaysOn Configuration
func (m *IosikEv2VpnConfiguration) SetAlwaysOnConfiguration(value AppleVpnAlwaysOnConfigurationable)() {
    m.alwaysOnConfiguration = value
}
// SetChildSecurityAssociationParameters sets the childSecurityAssociationParameters property value. Child Security Association Parameters
func (m *IosikEv2VpnConfiguration) SetChildSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)() {
    m.childSecurityAssociationParameters = value
}
// SetClientAuthenticationType sets the clientAuthenticationType property value. The type of VPN client authentication type
func (m *IosikEv2VpnConfiguration) SetClientAuthenticationType(value *VpnClientAuthenticationType)() {
    m.clientAuthenticationType = value
}
// SetDeadPeerDetectionRate sets the deadPeerDetectionRate property value. Determine how often to check if a peer connection is still active. . Possible values are: medium, none, low, high.
func (m *IosikEv2VpnConfiguration) SetDeadPeerDetectionRate(value *VpnDeadPeerDetectionRate)() {
    m.deadPeerDetectionRate = value
}
// SetDisableMobilityAndMultihoming sets the disableMobilityAndMultihoming property value. Disable MOBIKE
func (m *IosikEv2VpnConfiguration) SetDisableMobilityAndMultihoming(value *bool)() {
    m.disableMobilityAndMultihoming = value
}
// SetDisableRedirect sets the disableRedirect property value. Disable Redirect
func (m *IosikEv2VpnConfiguration) SetDisableRedirect(value *bool)() {
    m.disableRedirect = value
}
// SetEnableAlwaysOnConfiguration sets the enableAlwaysOnConfiguration property value. Determines if Always on VPN is enabled
func (m *IosikEv2VpnConfiguration) SetEnableAlwaysOnConfiguration(value *bool)() {
    m.enableAlwaysOnConfiguration = value
}
// SetEnableCertificateRevocationCheck sets the enableCertificateRevocationCheck property value. Enables a best-effort revocation check; server response timeouts will not cause it to fail
func (m *IosikEv2VpnConfiguration) SetEnableCertificateRevocationCheck(value *bool)() {
    m.enableCertificateRevocationCheck = value
}
// SetEnableEAP sets the enableEAP property value. Enables EAP only authentication
func (m *IosikEv2VpnConfiguration) SetEnableEAP(value *bool)() {
    m.enableEAP = value
}
// SetEnablePerfectForwardSecrecy sets the enablePerfectForwardSecrecy property value. Enable Perfect Forward Secrecy (PFS).
func (m *IosikEv2VpnConfiguration) SetEnablePerfectForwardSecrecy(value *bool)() {
    m.enablePerfectForwardSecrecy = value
}
// SetEnableUseInternalSubnetAttributes sets the enableUseInternalSubnetAttributes property value. Enable Use Internal Subnet Attributes.
func (m *IosikEv2VpnConfiguration) SetEnableUseInternalSubnetAttributes(value *bool)() {
    m.enableUseInternalSubnetAttributes = value
}
// SetLocalIdentifier sets the localIdentifier property value. The type of VPN local identifier
func (m *IosikEv2VpnConfiguration) SetLocalIdentifier(value *VpnLocalIdentifier)() {
    m.localIdentifier = value
}
// SetMtuSizeInBytes sets the mtuSizeInBytes property value. Maximum transmission unit. Valid values 1280 to 1400
func (m *IosikEv2VpnConfiguration) SetMtuSizeInBytes(value *int32)() {
    m.mtuSizeInBytes = value
}
// SetRemoteIdentifier sets the remoteIdentifier property value. Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
func (m *IosikEv2VpnConfiguration) SetRemoteIdentifier(value *string)() {
    m.remoteIdentifier = value
}
// SetSecurityAssociationParameters sets the securityAssociationParameters property value. Security Association Parameters
func (m *IosikEv2VpnConfiguration) SetSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)() {
    m.securityAssociationParameters = value
}
// SetServerCertificateCommonName sets the serverCertificateCommonName property value. Common name of the IKEv2 Server Certificate used in Server Authentication
func (m *IosikEv2VpnConfiguration) SetServerCertificateCommonName(value *string)() {
    m.serverCertificateCommonName = value
}
// SetServerCertificateIssuerCommonName sets the serverCertificateIssuerCommonName property value. Issuer Common name of the IKEv2 Server Certificate issuer used in Authentication
func (m *IosikEv2VpnConfiguration) SetServerCertificateIssuerCommonName(value *string)() {
    m.serverCertificateIssuerCommonName = value
}
// SetServerCertificateType sets the serverCertificateType property value. The type of certificate the VPN server will present to the VPN client for authentication. Possible values are: rsa, ecdsa256, ecdsa384, ecdsa521.
func (m *IosikEv2VpnConfiguration) SetServerCertificateType(value *VpnServerCertificateType)() {
    m.serverCertificateType = value
}
// SetSharedSecret sets the sharedSecret property value. Used when Shared Secret Authentication is selected
func (m *IosikEv2VpnConfiguration) SetSharedSecret(value *string)() {
    m.sharedSecret = value
}
// SetTlsMaximumVersion sets the tlsMaximumVersion property value. The maximum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) SetTlsMaximumVersion(value *string)() {
    m.tlsMaximumVersion = value
}
// SetTlsMinimumVersion sets the tlsMinimumVersion property value. The minimum TLS version to be used with EAP-TLS authentication
func (m *IosikEv2VpnConfiguration) SetTlsMinimumVersion(value *string)() {
    m.tlsMinimumVersion = value
}
