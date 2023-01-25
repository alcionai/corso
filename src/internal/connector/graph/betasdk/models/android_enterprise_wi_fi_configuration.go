package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidEnterpriseWiFiConfiguration 
type AndroidEnterpriseWiFiConfiguration struct {
    AndroidWiFiConfiguration
    // Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
    authenticationMethod *WiFiAuthenticationMethod
    // Extensible Authentication Protocol (EAP) Configuration Types.
    eapType *AndroidEapType
    // Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
    identityCertificateForClientAuthentication AndroidCertificateProfileBaseable
    // Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
    innerAuthenticationProtocolForEapTtls *NonEapAuthenticationMethodForEapTtlsType
    // Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
    innerAuthenticationProtocolForPeap *NonEapAuthenticationMethodForPeap
    // Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
    outerIdentityPrivacyTemporaryValue *string
    // Password format string used to build the password to connect to wifi
    passwordFormatString *string
    // PreSharedKey used to build the password to connect to wifi
    preSharedKey *string
    // Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
    rootCertificateForServerValidation AndroidTrustedRootCertificateable
    // Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
    trustedServerCertificateNames []string
    // Username format string used to build the username to connect to wifi
    usernameFormatString *string
}
// NewAndroidEnterpriseWiFiConfiguration instantiates a new AndroidEnterpriseWiFiConfiguration and sets the default values.
func NewAndroidEnterpriseWiFiConfiguration()(*AndroidEnterpriseWiFiConfiguration) {
    m := &AndroidEnterpriseWiFiConfiguration{
        AndroidWiFiConfiguration: *NewAndroidWiFiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidEnterpriseWiFiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidEnterpriseWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidEnterpriseWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidEnterpriseWiFiConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *AndroidEnterpriseWiFiConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    return m.authenticationMethod
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
func (m *AndroidEnterpriseWiFiConfiguration) GetEapType()(*AndroidEapType) {
    return m.eapType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidEnterpriseWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidWiFiConfiguration.GetFieldDeserializers()
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
    res["eapType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidEapType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapType(val.(*AndroidEapType))
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(AndroidCertificateProfileBaseable))
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
    res["innerAuthenticationProtocolForPeap"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForPeap)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForPeap(val.(*NonEapAuthenticationMethodForPeap))
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
    res["preSharedKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreSharedKey(val)
        }
        return nil
    }
    res["rootCertificateForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForServerValidation(val.(AndroidTrustedRootCertificateable))
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
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
func (m *AndroidEnterpriseWiFiConfiguration) GetIdentityCertificateForClientAuthentication()(AndroidCertificateProfileBaseable) {
    return m.identityCertificateForClientAuthentication
}
// GetInnerAuthenticationProtocolForEapTtls gets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *AndroidEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType) {
    return m.innerAuthenticationProtocolForEapTtls
}
// GetInnerAuthenticationProtocolForPeap gets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
func (m *AndroidEnterpriseWiFiConfiguration) GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap) {
    return m.innerAuthenticationProtocolForPeap
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
func (m *AndroidEnterpriseWiFiConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    return m.outerIdentityPrivacyTemporaryValue
}
// GetPasswordFormatString gets the passwordFormatString property value. Password format string used to build the password to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) GetPasswordFormatString()(*string) {
    return m.passwordFormatString
}
// GetPreSharedKey gets the preSharedKey property value. PreSharedKey used to build the password to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) GetPreSharedKey()(*string) {
    return m.preSharedKey
}
// GetRootCertificateForServerValidation gets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
func (m *AndroidEnterpriseWiFiConfiguration) GetRootCertificateForServerValidation()(AndroidTrustedRootCertificateable) {
    return m.rootCertificateForServerValidation
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *AndroidEnterpriseWiFiConfiguration) GetTrustedServerCertificateNames()([]string) {
    return m.trustedServerCertificateNames
}
// GetUsernameFormatString gets the usernameFormatString property value. Username format string used to build the username to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) GetUsernameFormatString()(*string) {
    return m.usernameFormatString
}
// Serialize serializes information the current object
func (m *AndroidEnterpriseWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidWiFiConfiguration.Serialize(writer)
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
    if m.GetInnerAuthenticationProtocolForPeap() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForPeap()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForPeap", &cast)
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
    {
        err = writer.WriteStringValue("preSharedKey", m.GetPreSharedKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificateForServerValidation", m.GetRootCertificateForServerValidation())
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
// SetAuthenticationMethod sets the authenticationMethod property value. Indicates the Authentication Method the client (device) needs to use when the EAP Type is configured to PEAP or EAP-TTLS. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *AndroidEnterpriseWiFiConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) Configuration Types.
func (m *AndroidEnterpriseWiFiConfiguration) SetEapType(value *AndroidEapType)() {
    m.eapType = value
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for client authentication when EAP Type is configured to EAP-TLS, EAP-TTLS (with Certificate Authentication), or PEAP (with Certificate Authentication). This is the certificate presented by client to the Wi-Fi endpoint. The authentication server sitting behind the Wi-Fi endpoint must accept this certificate to successfully establish a Wi-Fi connection.
func (m *AndroidEnterpriseWiFiConfiguration) SetIdentityCertificateForClientAuthentication(value AndroidCertificateProfileBaseable)() {
    m.identityCertificateForClientAuthentication = value
}
// SetInnerAuthenticationProtocolForEapTtls sets the innerAuthenticationProtocolForEapTtls property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *AndroidEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)() {
    m.innerAuthenticationProtocolForEapTtls = value
}
// SetInnerAuthenticationProtocolForPeap sets the innerAuthenticationProtocolForPeap property value. Non-EAP Method for Authentication (Inner Identity) when EAP Type is PEAP and Authenticationmethod is Username and Password. Possible values are: none, microsoftChapVersionTwo.
func (m *AndroidEnterpriseWiFiConfiguration) SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)() {
    m.innerAuthenticationProtocolForPeap = value
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS or PEAP. The String provided here is used to mask the username of individual users when they attempt to connect to Wi-Fi network.
func (m *AndroidEnterpriseWiFiConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    m.outerIdentityPrivacyTemporaryValue = value
}
// SetPasswordFormatString sets the passwordFormatString property value. Password format string used to build the password to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) SetPasswordFormatString(value *string)() {
    m.passwordFormatString = value
}
// SetPreSharedKey sets the preSharedKey property value. PreSharedKey used to build the password to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) SetPreSharedKey(value *string)() {
    m.preSharedKey = value
}
// SetRootCertificateForServerValidation sets the rootCertificateForServerValidation property value. Trusted Root Certificate for Server Validation when EAP Type is configured to EAP-TLS, EAP-TTLS or PEAP. This is the certificate presented by the Wi-Fi endpoint when the device attempts to connect to Wi-Fi endpoint. The device (or user) must accept this certificate to continue the connection attempt.
func (m *AndroidEnterpriseWiFiConfiguration) SetRootCertificateForServerValidation(value AndroidTrustedRootCertificateable)() {
    m.rootCertificateForServerValidation = value
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users' devices when they connect to this Wi-Fi network.
func (m *AndroidEnterpriseWiFiConfiguration) SetTrustedServerCertificateNames(value []string)() {
    m.trustedServerCertificateNames = value
}
// SetUsernameFormatString sets the usernameFormatString property value. Username format string used to build the username to connect to wifi
func (m *AndroidEnterpriseWiFiConfiguration) SetUsernameFormatString(value *string)() {
    m.usernameFormatString = value
}
