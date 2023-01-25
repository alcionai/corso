package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEnterpriseWiFiConfiguration 
type IosEnterpriseWiFiConfiguration struct {
    IosWiFiConfiguration
    // Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
    authenticationMethod *WiFiAuthenticationMethod
    // Tenant level settings for the Derived Credentials to be used for authentication.
    derivedCredentialSettings DeviceManagementDerivedCredentialSettingsable
    // EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
    eapFastConfiguration *EapFastConfiguration
    // Extensible Authentication Protocol (EAP) configuration types.
    eapType *EapType
    // Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
    identityCertificateForClientAuthentication IosCertificateProfileBaseable
    // Non-EAP Method for Authentication when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
    innerAuthenticationProtocolForEapTtls *NonEapAuthenticationMethodForEapTtlsType
    // Enable identity privacy (Outer Identity) when EAP Type is configured to EAP - TTLS, EAP - FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this Wi-Fi connection using their real username is displayed as 'anonymous'.
    outerIdentityPrivacyTemporaryValue *string
    // Password format string used to build the password to connect to wifi
    passwordFormatString *string
    // Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can contain a maximum of 500 elements.
    rootCertificatesForServerValidation []IosTrustedRootCertificateable
    // Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
    trustedServerCertificateNames []string
    // Username format string used to build the username to connect to wifi
    usernameFormatString *string
}
// NewIosEnterpriseWiFiConfiguration instantiates a new IosEnterpriseWiFiConfiguration and sets the default values.
func NewIosEnterpriseWiFiConfiguration()(*IosEnterpriseWiFiConfiguration) {
    m := &IosEnterpriseWiFiConfiguration{
        IosWiFiConfiguration: *NewIosWiFiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosEnterpriseWiFiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosEnterpriseWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEnterpriseWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEnterpriseWiFiConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *IosEnterpriseWiFiConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    return m.authenticationMethod
}
// GetDerivedCredentialSettings gets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosEnterpriseWiFiConfiguration) GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable) {
    return m.derivedCredentialSettings
}
// GetEapFastConfiguration gets the eapFastConfiguration property value. EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
func (m *IosEnterpriseWiFiConfiguration) GetEapFastConfiguration()(*EapFastConfiguration) {
    return m.eapFastConfiguration
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *IosEnterpriseWiFiConfiguration) GetEapType()(*EapType) {
    return m.eapType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosEnterpriseWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosWiFiConfiguration.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*WiFiAuthenticationMethod))
        }
        return nil
    }
    res["derivedCredentialSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialSettings(val.(DeviceManagementDerivedCredentialSettingsable))
        }
        return nil
    }
    res["eapFastConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEapFastConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapFastConfiguration(val.(*EapFastConfiguration))
        }
        return nil
    }
    res["eapType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEapType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapType(val.(*EapType))
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(IosCertificateProfileBaseable))
        }
        return nil
    }
    res["innerAuthenticationProtocolForEapTtls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForEapTtlsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForEapTtls(val.(*NonEapAuthenticationMethodForEapTtlsType))
        }
        return nil
    }
    res["outerIdentityPrivacyTemporaryValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuterIdentityPrivacyTemporaryValue(val)
        }
        return nil
    }
    res["passwordFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordFormatString(val)
        }
        return nil
    }
    res["rootCertificatesForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosTrustedRootCertificateable, len(val))
            for i, v := range val {
                res[i] = v.(IosTrustedRootCertificateable)
            }
            m.SetRootCertificatesForServerValidation(res)
        }
        return nil
    }
    res["trustedServerCertificateNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTrustedServerCertificateNames(res)
        }
        return nil
    }
    res["usernameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameFormatString(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
func (m *IosEnterpriseWiFiConfiguration) GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable) {
    return m.identityCertificateForClientAuthentication
}
// GetInnerAuthenticationProtocolForEapTtls gets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *IosEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType) {
    return m.innerAuthenticationProtocolForEapTtls
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP - TTLS, EAP - FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this Wi-Fi connection using their real username is displayed as 'anonymous'.
func (m *IosEnterpriseWiFiConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    return m.outerIdentityPrivacyTemporaryValue
}
// GetPasswordFormatString gets the passwordFormatString property value. Password format string used to build the password to connect to wifi
func (m *IosEnterpriseWiFiConfiguration) GetPasswordFormatString()(*string) {
    return m.passwordFormatString
}
// GetRootCertificatesForServerValidation gets the rootCertificatesForServerValidation property value. Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can contain a maximum of 500 elements.
func (m *IosEnterpriseWiFiConfiguration) GetRootCertificatesForServerValidation()([]IosTrustedRootCertificateable) {
    return m.rootCertificatesForServerValidation
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *IosEnterpriseWiFiConfiguration) GetTrustedServerCertificateNames()([]string) {
    return m.trustedServerCertificateNames
}
// GetUsernameFormatString gets the usernameFormatString property value. Username format string used to build the username to connect to wifi
func (m *IosEnterpriseWiFiConfiguration) GetUsernameFormatString()(*string) {
    return m.usernameFormatString
}
// Serialize serializes information the current object
func (m *IosEnterpriseWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosWiFiConfiguration.Serialize(writer)
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
        err = writer.WriteObjectValue("derivedCredentialSettings", m.GetDerivedCredentialSettings())
        if err != nil {
            return err
        }
    }
    if m.GetEapFastConfiguration() != nil {
        cast := (*m.GetEapFastConfiguration()).String()
        err = writer.WriteStringValue("eapFastConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEapType() != nil {
        cast := (*m.GetEapType()).String()
        err = writer.WriteStringValue("eapType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificateForClientAuthentication", m.GetIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    if m.GetInnerAuthenticationProtocolForEapTtls() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForEapTtls()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForEapTtls", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outerIdentityPrivacyTemporaryValue", m.GetOuterIdentityPrivacyTemporaryValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordFormatString", m.GetPasswordFormatString())
        if err != nil {
            return err
        }
    }
    if m.GetRootCertificatesForServerValidation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRootCertificatesForServerValidation()))
        for i, v := range m.GetRootCertificatesForServerValidation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rootCertificatesForServerValidation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrustedServerCertificateNames() != nil {
        err = writer.WriteCollectionOfStringValues("trustedServerCertificateNames", m.GetTrustedServerCertificateNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameFormatString", m.GetUsernameFormatString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Authentication Method when EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *IosEnterpriseWiFiConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetDerivedCredentialSettings sets the derivedCredentialSettings property value. Tenant level settings for the Derived Credentials to be used for authentication.
func (m *IosEnterpriseWiFiConfiguration) SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)() {
    m.derivedCredentialSettings = value
}
// SetEapFastConfiguration sets the eapFastConfiguration property value. EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type. Possible values are: noProtectedAccessCredential, useProtectedAccessCredential, useProtectedAccessCredentialAndProvision, useProtectedAccessCredentialAndProvisionAnonymously.
func (m *IosEnterpriseWiFiConfiguration) SetEapFastConfiguration(value *EapFastConfiguration)() {
    m.eapFastConfiguration = value
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *IosEnterpriseWiFiConfiguration) SetEapType(value *EapType)() {
    m.eapType = value
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication).
func (m *IosEnterpriseWiFiConfiguration) SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)() {
    m.identityCertificateForClientAuthentication = value
}
// SetInnerAuthenticationProtocolForEapTtls sets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *IosEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)() {
    m.innerAuthenticationProtocolForEapTtls = value
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP - TTLS, EAP - FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this Wi-Fi connection using their real username is displayed as 'anonymous'.
func (m *IosEnterpriseWiFiConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    m.outerIdentityPrivacyTemporaryValue = value
}
// SetPasswordFormatString sets the passwordFormatString property value. Password format string used to build the password to connect to wifi
func (m *IosEnterpriseWiFiConfiguration) SetPasswordFormatString(value *string)() {
    m.passwordFormatString = value
}
// SetRootCertificatesForServerValidation sets the rootCertificatesForServerValidation property value. Trusted Root Certificates for Server Validation when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. If you provide this value you do not need to provide trustedServerCertificateNames, and vice versa. This collection can contain a maximum of 500 elements.
func (m *IosEnterpriseWiFiConfiguration) SetRootCertificatesForServerValidation(value []IosTrustedRootCertificateable)() {
    m.rootCertificatesForServerValidation = value
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *IosEnterpriseWiFiConfiguration) SetTrustedServerCertificateNames(value []string)() {
    m.trustedServerCertificateNames = value
}
// SetUsernameFormatString sets the usernameFormatString property value. Username format string used to build the username to connect to wifi
func (m *IosEnterpriseWiFiConfiguration) SetUsernameFormatString(value *string)() {
    m.usernameFormatString = value
}
