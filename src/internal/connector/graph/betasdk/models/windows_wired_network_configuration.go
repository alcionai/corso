package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWiredNetworkConfiguration 
type WindowsWiredNetworkConfiguration struct {
    DeviceConfiguration
    // Specify the duration for which automatic authentication attempts will be blocked from occuring after a failed authentication attempt.
    authenticationBlockPeriodInMinutes *int32
    // Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
    authenticationMethod *WiredNetworkAuthenticationMethod
    // Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
    authenticationPeriodInSeconds *int32
    // Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
    authenticationRetryDelayPeriodInSeconds *int32
    // Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you're using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest. Possible values are: none, user, machine, machineOrUser, guest, unknownFutureValue.
    authenticationType *WiredNetworkAuthenticationType
    // When TRUE, caches user credentials on the device so that users don't need to keep entering them each time they connect. When FALSE, do not cache credentials. Default value is FALSE.
    cacheCredentials *bool
    // When TRUE, prevents the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP. When FALSE, does not prevent the user from being prompted. Default value is FALSE.
    disableUserPromptForServerValidation *bool
    // Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
    eapolStartPeriodInSeconds *int32
    // Extensible Authentication Protocol (EAP) configuration types.
    eapType *EapType
    // When TRUE, the automatic configuration service for wired networks requires the use of 802.1X for port authentication. When FALSE, 802.1X is not required. Default value is FALSE.
    enforce8021X *bool
    // When TRUE, forces FIPS compliance. When FALSE, does not enable FIPS compliance. Default value is FALSE.
    forceFIPSCompliance *bool
    // Specify identity certificate for client authentication.
    identityCertificateForClientAuthentication WindowsCertificateProfileBaseable
    // Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
    innerAuthenticationProtocolForEAPTTLS *NonEapAuthenticationMethodForEapTtlsType
    // Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
    maximumAuthenticationFailures *int32
    // Specify the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
    maximumEAPOLStartMessages *int32
    // Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
    outerIdentityPrivacyTemporaryValue *string
    // When TRUE, enables verification of server's identity by validating the certificate when EAP type is selected as PEAP. When FALSE, the certificate is not validated. Default value is TRUE.
    performServerValidation *bool
    // When TRUE, enables cryptographic binding when EAP type is selected as PEAP. When FALSE, does not enable cryptogrpahic binding. Default value is TRUE.
    requireCryptographicBinding *bool
    // Specify root certificate for client validation.
    rootCertificateForClientValidation Windows81TrustedRootCertificateable
    // Specify root certificates for server validation. This collection can contain a maximum of 500 elements.
    rootCertificatesForServerValidation []Windows81TrustedRootCertificateable
    // Specify the secondary authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
    secondaryAuthenticationMethod *WiredNetworkAuthenticationMethod
    // Specify secondary identity certificate for client authentication.
    secondaryIdentityCertificateForClientAuthentication WindowsCertificateProfileBaseable
    // Specify secondary root certificate for client validation.
    secondaryRootCertificateForClientValidation Windows81TrustedRootCertificateable
    // Specify trusted server certificate names.
    trustedServerCertificateNames []string
}
// NewWindowsWiredNetworkConfiguration instantiates a new WindowsWiredNetworkConfiguration and sets the default values.
func NewWindowsWiredNetworkConfiguration()(*WindowsWiredNetworkConfiguration) {
    m := &WindowsWiredNetworkConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsWiredNetworkConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsWiredNetworkConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsWiredNetworkConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsWiredNetworkConfiguration(), nil
}
// GetAuthenticationBlockPeriodInMinutes gets the authenticationBlockPeriodInMinutes property value. Specify the duration for which automatic authentication attempts will be blocked from occuring after a failed authentication attempt.
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationBlockPeriodInMinutes()(*int32) {
    return m.authenticationBlockPeriodInMinutes
}
// GetAuthenticationMethod gets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationMethod()(*WiredNetworkAuthenticationMethod) {
    return m.authenticationMethod
}
// GetAuthenticationPeriodInSeconds gets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationPeriodInSeconds()(*int32) {
    return m.authenticationPeriodInSeconds
}
// GetAuthenticationRetryDelayPeriodInSeconds gets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationRetryDelayPeriodInSeconds()(*int32) {
    return m.authenticationRetryDelayPeriodInSeconds
}
// GetAuthenticationType gets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you're using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest. Possible values are: none, user, machine, machineOrUser, guest, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) GetAuthenticationType()(*WiredNetworkAuthenticationType) {
    return m.authenticationType
}
// GetCacheCredentials gets the cacheCredentials property value. When TRUE, caches user credentials on the device so that users don't need to keep entering them each time they connect. When FALSE, do not cache credentials. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) GetCacheCredentials()(*bool) {
    return m.cacheCredentials
}
// GetDisableUserPromptForServerValidation gets the disableUserPromptForServerValidation property value. When TRUE, prevents the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP. When FALSE, does not prevent the user from being prompted. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) GetDisableUserPromptForServerValidation()(*bool) {
    return m.disableUserPromptForServerValidation
}
// GetEapolStartPeriodInSeconds gets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) GetEapolStartPeriodInSeconds()(*int32) {
    return m.eapolStartPeriodInSeconds
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *WindowsWiredNetworkConfiguration) GetEapType()(*EapType) {
    return m.eapType
}
// GetEnforce8021X gets the enforce8021X property value. When TRUE, the automatic configuration service for wired networks requires the use of 802.1X for port authentication. When FALSE, 802.1X is not required. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) GetEnforce8021X()(*bool) {
    return m.enforce8021X
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsWiredNetworkConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["authenticationBlockPeriodInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationBlockPeriodInMinutes(val)
        }
        return nil
    }
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiredNetworkAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(*WiredNetworkAuthenticationMethod))
        }
        return nil
    }
    res["authenticationPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationPeriodInSeconds(val)
        }
        return nil
    }
    res["authenticationRetryDelayPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationRetryDelayPeriodInSeconds(val)
        }
        return nil
    }
    res["authenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiredNetworkAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val.(*WiredNetworkAuthenticationType))
        }
        return nil
    }
    res["cacheCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheCredentials(val)
        }
        return nil
    }
    res["disableUserPromptForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUserPromptForServerValidation(val)
        }
        return nil
    }
    res["eapolStartPeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEapolStartPeriodInSeconds(val)
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
    res["enforce8021X"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforce8021X(val)
        }
        return nil
    }
    res["forceFIPSCompliance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceFIPSCompliance(val)
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(WindowsCertificateProfileBaseable))
        }
        return nil
    }
    res["innerAuthenticationProtocolForEAPTTLS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonEapAuthenticationMethodForEapTtlsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerAuthenticationProtocolForEAPTTLS(val.(*NonEapAuthenticationMethodForEapTtlsType))
        }
        return nil
    }
    res["maximumAuthenticationFailures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAuthenticationFailures(val)
        }
        return nil
    }
    res["maximumEAPOLStartMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumEAPOLStartMessages(val)
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
    res["performServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerformServerValidation(val)
        }
        return nil
    }
    res["requireCryptographicBinding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCryptographicBinding(val)
        }
        return nil
    }
    res["rootCertificateForClientValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateForClientValidation(val.(Windows81TrustedRootCertificateable))
        }
        return nil
    }
    res["rootCertificatesForServerValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows81TrustedRootCertificateable, len(val))
            for i, v := range val {
                res[i] = v.(Windows81TrustedRootCertificateable)
            }
            m.SetRootCertificatesForServerValidation(res)
        }
        return nil
    }
    res["secondaryAuthenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiredNetworkAuthenticationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryAuthenticationMethod(val.(*WiredNetworkAuthenticationMethod))
        }
        return nil
    }
    res["secondaryIdentityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryIdentityCertificateForClientAuthentication(val.(WindowsCertificateProfileBaseable))
        }
        return nil
    }
    res["secondaryRootCertificateForClientValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryRootCertificateForClientValidation(val.(Windows81TrustedRootCertificateable))
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
    return res
}
// GetForceFIPSCompliance gets the forceFIPSCompliance property value. When TRUE, forces FIPS compliance. When FALSE, does not enable FIPS compliance. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) GetForceFIPSCompliance()(*bool) {
    return m.forceFIPSCompliance
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
func (m *WindowsWiredNetworkConfiguration) GetIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable) {
    return m.identityCertificateForClientAuthentication
}
// GetInnerAuthenticationProtocolForEAPTTLS gets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *WindowsWiredNetworkConfiguration) GetInnerAuthenticationProtocolForEAPTTLS()(*NonEapAuthenticationMethodForEapTtlsType) {
    return m.innerAuthenticationProtocolForEAPTTLS
}
// GetMaximumAuthenticationFailures gets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
func (m *WindowsWiredNetworkConfiguration) GetMaximumAuthenticationFailures()(*int32) {
    return m.maximumAuthenticationFailures
}
// GetMaximumEAPOLStartMessages gets the maximumEAPOLStartMessages property value. Specify the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
func (m *WindowsWiredNetworkConfiguration) GetMaximumEAPOLStartMessages()(*int32) {
    return m.maximumEAPOLStartMessages
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
func (m *WindowsWiredNetworkConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    return m.outerIdentityPrivacyTemporaryValue
}
// GetPerformServerValidation gets the performServerValidation property value. When TRUE, enables verification of server's identity by validating the certificate when EAP type is selected as PEAP. When FALSE, the certificate is not validated. Default value is TRUE.
func (m *WindowsWiredNetworkConfiguration) GetPerformServerValidation()(*bool) {
    return m.performServerValidation
}
// GetRequireCryptographicBinding gets the requireCryptographicBinding property value. When TRUE, enables cryptographic binding when EAP type is selected as PEAP. When FALSE, does not enable cryptogrpahic binding. Default value is TRUE.
func (m *WindowsWiredNetworkConfiguration) GetRequireCryptographicBinding()(*bool) {
    return m.requireCryptographicBinding
}
// GetRootCertificateForClientValidation gets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
func (m *WindowsWiredNetworkConfiguration) GetRootCertificateForClientValidation()(Windows81TrustedRootCertificateable) {
    return m.rootCertificateForClientValidation
}
// GetRootCertificatesForServerValidation gets the rootCertificatesForServerValidation property value. Specify root certificates for server validation. This collection can contain a maximum of 500 elements.
func (m *WindowsWiredNetworkConfiguration) GetRootCertificatesForServerValidation()([]Windows81TrustedRootCertificateable) {
    return m.rootCertificatesForServerValidation
}
// GetSecondaryAuthenticationMethod gets the secondaryAuthenticationMethod property value. Specify the secondary authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) GetSecondaryAuthenticationMethod()(*WiredNetworkAuthenticationMethod) {
    return m.secondaryAuthenticationMethod
}
// GetSecondaryIdentityCertificateForClientAuthentication gets the secondaryIdentityCertificateForClientAuthentication property value. Specify secondary identity certificate for client authentication.
func (m *WindowsWiredNetworkConfiguration) GetSecondaryIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable) {
    return m.secondaryIdentityCertificateForClientAuthentication
}
// GetSecondaryRootCertificateForClientValidation gets the secondaryRootCertificateForClientValidation property value. Specify secondary root certificate for client validation.
func (m *WindowsWiredNetworkConfiguration) GetSecondaryRootCertificateForClientValidation()(Windows81TrustedRootCertificateable) {
    return m.secondaryRootCertificateForClientValidation
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Specify trusted server certificate names.
func (m *WindowsWiredNetworkConfiguration) GetTrustedServerCertificateNames()([]string) {
    return m.trustedServerCertificateNames
}
// Serialize serializes information the current object
func (m *WindowsWiredNetworkConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("authenticationBlockPeriodInMinutes", m.GetAuthenticationBlockPeriodInMinutes())
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
        err = writer.WriteInt32Value("authenticationPeriodInSeconds", m.GetAuthenticationPeriodInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("authenticationRetryDelayPeriodInSeconds", m.GetAuthenticationRetryDelayPeriodInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationType() != nil {
        cast := (*m.GetAuthenticationType()).String()
        err = writer.WriteStringValue("authenticationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cacheCredentials", m.GetCacheCredentials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableUserPromptForServerValidation", m.GetDisableUserPromptForServerValidation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("eapolStartPeriodInSeconds", m.GetEapolStartPeriodInSeconds())
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
        err = writer.WriteBoolValue("enforce8021X", m.GetEnforce8021X())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("forceFIPSCompliance", m.GetForceFIPSCompliance())
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
    if m.GetInnerAuthenticationProtocolForEAPTTLS() != nil {
        cast := (*m.GetInnerAuthenticationProtocolForEAPTTLS()).String()
        err = writer.WriteStringValue("innerAuthenticationProtocolForEAPTTLS", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAuthenticationFailures", m.GetMaximumAuthenticationFailures())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumEAPOLStartMessages", m.GetMaximumEAPOLStartMessages())
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
        err = writer.WriteBoolValue("performServerValidation", m.GetPerformServerValidation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireCryptographicBinding", m.GetRequireCryptographicBinding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificateForClientValidation", m.GetRootCertificateForClientValidation())
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
    if m.GetSecondaryAuthenticationMethod() != nil {
        cast := (*m.GetSecondaryAuthenticationMethod()).String()
        err = writer.WriteStringValue("secondaryAuthenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("secondaryIdentityCertificateForClientAuthentication", m.GetSecondaryIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("secondaryRootCertificateForClientValidation", m.GetSecondaryRootCertificateForClientValidation())
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
    return nil
}
// SetAuthenticationBlockPeriodInMinutes sets the authenticationBlockPeriodInMinutes property value. Specify the duration for which automatic authentication attempts will be blocked from occuring after a failed authentication attempt.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationBlockPeriodInMinutes(value *int32)() {
    m.authenticationBlockPeriodInMinutes = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationMethod(value *WiredNetworkAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetAuthenticationPeriodInSeconds sets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationPeriodInSeconds(value *int32)() {
    m.authenticationPeriodInSeconds = value
}
// SetAuthenticationRetryDelayPeriodInSeconds sets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationRetryDelayPeriodInSeconds(value *int32)() {
    m.authenticationRetryDelayPeriodInSeconds = value
}
// SetAuthenticationType sets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you're using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest. Possible values are: none, user, machine, machineOrUser, guest, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) SetAuthenticationType(value *WiredNetworkAuthenticationType)() {
    m.authenticationType = value
}
// SetCacheCredentials sets the cacheCredentials property value. When TRUE, caches user credentials on the device so that users don't need to keep entering them each time they connect. When FALSE, do not cache credentials. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetCacheCredentials(value *bool)() {
    m.cacheCredentials = value
}
// SetDisableUserPromptForServerValidation sets the disableUserPromptForServerValidation property value. When TRUE, prevents the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP. When FALSE, does not prevent the user from being prompted. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetDisableUserPromptForServerValidation(value *bool)() {
    m.disableUserPromptForServerValidation = value
}
// SetEapolStartPeriodInSeconds sets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
func (m *WindowsWiredNetworkConfiguration) SetEapolStartPeriodInSeconds(value *int32)() {
    m.eapolStartPeriodInSeconds = value
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *WindowsWiredNetworkConfiguration) SetEapType(value *EapType)() {
    m.eapType = value
}
// SetEnforce8021X sets the enforce8021X property value. When TRUE, the automatic configuration service for wired networks requires the use of 802.1X for port authentication. When FALSE, 802.1X is not required. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetEnforce8021X(value *bool)() {
    m.enforce8021X = value
}
// SetForceFIPSCompliance sets the forceFIPSCompliance property value. When TRUE, forces FIPS compliance. When FALSE, does not enable FIPS compliance. Default value is FALSE.
func (m *WindowsWiredNetworkConfiguration) SetForceFIPSCompliance(value *bool)() {
    m.forceFIPSCompliance = value
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
func (m *WindowsWiredNetworkConfiguration) SetIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)() {
    m.identityCertificateForClientAuthentication = value
}
// SetInnerAuthenticationProtocolForEAPTTLS sets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *WindowsWiredNetworkConfiguration) SetInnerAuthenticationProtocolForEAPTTLS(value *NonEapAuthenticationMethodForEapTtlsType)() {
    m.innerAuthenticationProtocolForEAPTTLS = value
}
// SetMaximumAuthenticationFailures sets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
func (m *WindowsWiredNetworkConfiguration) SetMaximumAuthenticationFailures(value *int32)() {
    m.maximumAuthenticationFailures = value
}
// SetMaximumEAPOLStartMessages sets the maximumEAPOLStartMessages property value. Specify the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
func (m *WindowsWiredNetworkConfiguration) SetMaximumEAPOLStartMessages(value *int32)() {
    m.maximumEAPOLStartMessages = value
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
func (m *WindowsWiredNetworkConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    m.outerIdentityPrivacyTemporaryValue = value
}
// SetPerformServerValidation sets the performServerValidation property value. When TRUE, enables verification of server's identity by validating the certificate when EAP type is selected as PEAP. When FALSE, the certificate is not validated. Default value is TRUE.
func (m *WindowsWiredNetworkConfiguration) SetPerformServerValidation(value *bool)() {
    m.performServerValidation = value
}
// SetRequireCryptographicBinding sets the requireCryptographicBinding property value. When TRUE, enables cryptographic binding when EAP type is selected as PEAP. When FALSE, does not enable cryptogrpahic binding. Default value is TRUE.
func (m *WindowsWiredNetworkConfiguration) SetRequireCryptographicBinding(value *bool)() {
    m.requireCryptographicBinding = value
}
// SetRootCertificateForClientValidation sets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
func (m *WindowsWiredNetworkConfiguration) SetRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)() {
    m.rootCertificateForClientValidation = value
}
// SetRootCertificatesForServerValidation sets the rootCertificatesForServerValidation property value. Specify root certificates for server validation. This collection can contain a maximum of 500 elements.
func (m *WindowsWiredNetworkConfiguration) SetRootCertificatesForServerValidation(value []Windows81TrustedRootCertificateable)() {
    m.rootCertificatesForServerValidation = value
}
// SetSecondaryAuthenticationMethod sets the secondaryAuthenticationMethod property value. Specify the secondary authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential. Possible values are: certificate, usernameAndPassword, derivedCredential, unknownFutureValue.
func (m *WindowsWiredNetworkConfiguration) SetSecondaryAuthenticationMethod(value *WiredNetworkAuthenticationMethod)() {
    m.secondaryAuthenticationMethod = value
}
// SetSecondaryIdentityCertificateForClientAuthentication sets the secondaryIdentityCertificateForClientAuthentication property value. Specify secondary identity certificate for client authentication.
func (m *WindowsWiredNetworkConfiguration) SetSecondaryIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)() {
    m.secondaryIdentityCertificateForClientAuthentication = value
}
// SetSecondaryRootCertificateForClientValidation sets the secondaryRootCertificateForClientValidation property value. Specify secondary root certificate for client validation.
func (m *WindowsWiredNetworkConfiguration) SetSecondaryRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)() {
    m.secondaryRootCertificateForClientValidation = value
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Specify trusted server certificate names.
func (m *WindowsWiredNetworkConfiguration) SetTrustedServerCertificateNames(value []string)() {
    m.trustedServerCertificateNames = value
}
