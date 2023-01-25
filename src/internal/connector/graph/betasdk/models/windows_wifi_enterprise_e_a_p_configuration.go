package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWifiEnterpriseEAPConfiguration 
type WindowsWifiEnterpriseEAPConfiguration struct {
    WindowsWifiConfiguration
    // Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential.
    authenticationMethod *WiFiAuthenticationMethod
    // Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
    authenticationPeriodInSeconds *int32
    // Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
    authenticationRetryDelayPeriodInSeconds *int32
    // Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you’re using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest.
    authenticationType *WifiAuthenticationType
    // Specify whether to cache user credentials on the device so that users don’t need to keep entering them each time they connect.
    cacheCredentials *bool
    // Specify whether to prevent the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP.
    disableUserPromptForServerValidation *bool
    // Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
    eapolStartPeriodInSeconds *int32
    // Extensible Authentication Protocol (EAP) configuration types.
    eapType *EapType
    // Specify whether the wifi connection should enable pairwise master key caching.
    enablePairwiseMasterKeyCaching *bool
    // Specify whether pre-authentication should be enabled.
    enablePreAuthentication *bool
    // Specify identity certificate for client authentication.
    identityCertificateForClientAuthentication WindowsCertificateProfileBaseable
    // Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
    innerAuthenticationProtocolForEAPTTLS *NonEapAuthenticationMethodForEapTtlsType
    // Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
    maximumAuthenticationFailures *int32
    // Specify maximum authentication timeout (in seconds).  Valid range: 1-120
    maximumAuthenticationTimeoutInSeconds *int32
    // Specifiy the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
    maximumEAPOLStartMessages *int32
    // Specify maximum number of pairwise master keys in cache.  Valid range: 1-255
    maximumNumberOfPairwiseMasterKeysInCache *int32
    // Specify maximum pairwise master key cache time (in minutes).  Valid range: 5-1440
    maximumPairwiseMasterKeyCacheTimeInMinutes *int32
    // Specify maximum pre-authentication attempts.  Valid range: 1-16
    maximumPreAuthenticationAttempts *int32
    // Specify the network single sign on type. Possible values are: disabled, prelogon, postlogon.
    networkSingleSignOn *NetworkSingleSignOnType
    // Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
    outerIdentityPrivacyTemporaryValue *string
    // Specify whether to enable verification of server's identity by validating the certificate when EAP type is selected as PEAP.
    performServerValidation *bool
    // Specify whether the wifi connection should prompt for additional authentication credentials.
    promptForAdditionalAuthenticationCredentials *bool
    // Specify whether to enable cryptographic binding when EAP type is selected as PEAP.
    requireCryptographicBinding *bool
    // Specify root certificate for client validation.
    rootCertificateForClientValidation Windows81TrustedRootCertificateable
    // Specify root certificate for server validation. This collection can contain a maximum of 500 elements.
    rootCertificatesForServerValidation []Windows81TrustedRootCertificateable
    // Specify trusted server certificate names.
    trustedServerCertificateNames []string
    // Specifiy whether to change the virtual LAN used by the device based on the user’s credentials. Cannot be used when NetworkSingleSignOnType is set to ​Disabled.
    userBasedVirtualLan *bool
}
// NewWindowsWifiEnterpriseEAPConfiguration instantiates a new WindowsWifiEnterpriseEAPConfiguration and sets the default values.
func NewWindowsWifiEnterpriseEAPConfiguration()(*WindowsWifiEnterpriseEAPConfiguration) {
    m := &WindowsWifiEnterpriseEAPConfiguration{
        WindowsWifiConfiguration: *NewWindowsWifiConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsWifiEnterpriseEAPConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsWifiEnterpriseEAPConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsWifiEnterpriseEAPConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsWifiEnterpriseEAPConfiguration(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationMethod()(*WiFiAuthenticationMethod) {
    return m.authenticationMethod
}
// GetAuthenticationPeriodInSeconds gets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationPeriodInSeconds()(*int32) {
    return m.authenticationPeriodInSeconds
}
// GetAuthenticationRetryDelayPeriodInSeconds gets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationRetryDelayPeriodInSeconds()(*int32) {
    return m.authenticationRetryDelayPeriodInSeconds
}
// GetAuthenticationType gets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you’re using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetAuthenticationType()(*WifiAuthenticationType) {
    return m.authenticationType
}
// GetCacheCredentials gets the cacheCredentials property value. Specify whether to cache user credentials on the device so that users don’t need to keep entering them each time they connect.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetCacheCredentials()(*bool) {
    return m.cacheCredentials
}
// GetDisableUserPromptForServerValidation gets the disableUserPromptForServerValidation property value. Specify whether to prevent the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetDisableUserPromptForServerValidation()(*bool) {
    return m.disableUserPromptForServerValidation
}
// GetEapolStartPeriodInSeconds gets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEapolStartPeriodInSeconds()(*int32) {
    return m.eapolStartPeriodInSeconds
}
// GetEapType gets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEapType()(*EapType) {
    return m.eapType
}
// GetEnablePairwiseMasterKeyCaching gets the enablePairwiseMasterKeyCaching property value. Specify whether the wifi connection should enable pairwise master key caching.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEnablePairwiseMasterKeyCaching()(*bool) {
    return m.enablePairwiseMasterKeyCaching
}
// GetEnablePreAuthentication gets the enablePreAuthentication property value. Specify whether pre-authentication should be enabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetEnablePreAuthentication()(*bool) {
    return m.enablePreAuthentication
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsWifiEnterpriseEAPConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsWifiConfiguration.GetFieldDeserializers()
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
        val, err := n.GetEnumValue(ParseWifiAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationType(val.(*WifiAuthenticationType))
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
    res["enablePairwiseMasterKeyCaching"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePairwiseMasterKeyCaching(val)
        }
        return nil
    }
    res["enablePreAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnablePreAuthentication(val)
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
    res["maximumAuthenticationTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAuthenticationTimeoutInSeconds(val)
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
    res["maximumNumberOfPairwiseMasterKeysInCache"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumNumberOfPairwiseMasterKeysInCache(val)
        }
        return nil
    }
    res["maximumPairwiseMasterKeyCacheTimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPairwiseMasterKeyCacheTimeInMinutes(val)
        }
        return nil
    }
    res["maximumPreAuthenticationAttempts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPreAuthenticationAttempts(val)
        }
        return nil
    }
    res["networkSingleSignOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkSingleSignOnType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkSingleSignOn(val.(*NetworkSingleSignOnType))
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
    res["promptForAdditionalAuthenticationCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPromptForAdditionalAuthenticationCredentials(val)
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
    res["userBasedVirtualLan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserBasedVirtualLan(val)
        }
        return nil
    }
    return res
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetIdentityCertificateForClientAuthentication()(WindowsCertificateProfileBaseable) {
    return m.identityCertificateForClientAuthentication
}
// GetInnerAuthenticationProtocolForEAPTTLS gets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetInnerAuthenticationProtocolForEAPTTLS()(*NonEapAuthenticationMethodForEapTtlsType) {
    return m.innerAuthenticationProtocolForEAPTTLS
}
// GetMaximumAuthenticationFailures gets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumAuthenticationFailures()(*int32) {
    return m.maximumAuthenticationFailures
}
// GetMaximumAuthenticationTimeoutInSeconds gets the maximumAuthenticationTimeoutInSeconds property value. Specify maximum authentication timeout (in seconds).  Valid range: 1-120
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumAuthenticationTimeoutInSeconds()(*int32) {
    return m.maximumAuthenticationTimeoutInSeconds
}
// GetMaximumEAPOLStartMessages gets the maximumEAPOLStartMessages property value. Specifiy the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumEAPOLStartMessages()(*int32) {
    return m.maximumEAPOLStartMessages
}
// GetMaximumNumberOfPairwiseMasterKeysInCache gets the maximumNumberOfPairwiseMasterKeysInCache property value. Specify maximum number of pairwise master keys in cache.  Valid range: 1-255
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumNumberOfPairwiseMasterKeysInCache()(*int32) {
    return m.maximumNumberOfPairwiseMasterKeysInCache
}
// GetMaximumPairwiseMasterKeyCacheTimeInMinutes gets the maximumPairwiseMasterKeyCacheTimeInMinutes property value. Specify maximum pairwise master key cache time (in minutes).  Valid range: 5-1440
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumPairwiseMasterKeyCacheTimeInMinutes()(*int32) {
    return m.maximumPairwiseMasterKeyCacheTimeInMinutes
}
// GetMaximumPreAuthenticationAttempts gets the maximumPreAuthenticationAttempts property value. Specify maximum pre-authentication attempts.  Valid range: 1-16
func (m *WindowsWifiEnterpriseEAPConfiguration) GetMaximumPreAuthenticationAttempts()(*int32) {
    return m.maximumPreAuthenticationAttempts
}
// GetNetworkSingleSignOn gets the networkSingleSignOn property value. Specify the network single sign on type. Possible values are: disabled, prelogon, postlogon.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetNetworkSingleSignOn()(*NetworkSingleSignOnType) {
    return m.networkSingleSignOn
}
// GetOuterIdentityPrivacyTemporaryValue gets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetOuterIdentityPrivacyTemporaryValue()(*string) {
    return m.outerIdentityPrivacyTemporaryValue
}
// GetPerformServerValidation gets the performServerValidation property value. Specify whether to enable verification of server's identity by validating the certificate when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetPerformServerValidation()(*bool) {
    return m.performServerValidation
}
// GetPromptForAdditionalAuthenticationCredentials gets the promptForAdditionalAuthenticationCredentials property value. Specify whether the wifi connection should prompt for additional authentication credentials.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetPromptForAdditionalAuthenticationCredentials()(*bool) {
    return m.promptForAdditionalAuthenticationCredentials
}
// GetRequireCryptographicBinding gets the requireCryptographicBinding property value. Specify whether to enable cryptographic binding when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetRequireCryptographicBinding()(*bool) {
    return m.requireCryptographicBinding
}
// GetRootCertificateForClientValidation gets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetRootCertificateForClientValidation()(Windows81TrustedRootCertificateable) {
    return m.rootCertificateForClientValidation
}
// GetRootCertificatesForServerValidation gets the rootCertificatesForServerValidation property value. Specify root certificate for server validation. This collection can contain a maximum of 500 elements.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetRootCertificatesForServerValidation()([]Windows81TrustedRootCertificateable) {
    return m.rootCertificatesForServerValidation
}
// GetTrustedServerCertificateNames gets the trustedServerCertificateNames property value. Specify trusted server certificate names.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetTrustedServerCertificateNames()([]string) {
    return m.trustedServerCertificateNames
}
// GetUserBasedVirtualLan gets the userBasedVirtualLan property value. Specifiy whether to change the virtual LAN used by the device based on the user’s credentials. Cannot be used when NetworkSingleSignOnType is set to ​Disabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) GetUserBasedVirtualLan()(*bool) {
    return m.userBasedVirtualLan
}
// Serialize serializes information the current object
func (m *WindowsWifiEnterpriseEAPConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsWifiConfiguration.Serialize(writer)
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
        err = writer.WriteBoolValue("enablePairwiseMasterKeyCaching", m.GetEnablePairwiseMasterKeyCaching())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enablePreAuthentication", m.GetEnablePreAuthentication())
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
        err = writer.WriteInt32Value("maximumAuthenticationTimeoutInSeconds", m.GetMaximumAuthenticationTimeoutInSeconds())
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
        err = writer.WriteInt32Value("maximumNumberOfPairwiseMasterKeysInCache", m.GetMaximumNumberOfPairwiseMasterKeysInCache())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumPairwiseMasterKeyCacheTimeInMinutes", m.GetMaximumPairwiseMasterKeyCacheTimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumPreAuthenticationAttempts", m.GetMaximumPreAuthenticationAttempts())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkSingleSignOn() != nil {
        cast := (*m.GetNetworkSingleSignOn()).String()
        err = writer.WriteStringValue("networkSingleSignOn", &cast)
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
        err = writer.WriteBoolValue("promptForAdditionalAuthenticationCredentials", m.GetPromptForAdditionalAuthenticationCredentials())
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
    if m.GetTrustedServerCertificateNames() != nil {
        err = writer.WriteCollectionOfStringValues("trustedServerCertificateNames", m.GetTrustedServerCertificateNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("userBasedVirtualLan", m.GetUserBasedVirtualLan())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Specify the authentication method. Possible values are: certificate, usernameAndPassword, derivedCredential.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationMethod(value *WiFiAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetAuthenticationPeriodInSeconds sets the authenticationPeriodInSeconds property value. Specify the number of seconds for the client to wait after an authentication attempt before failing. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationPeriodInSeconds(value *int32)() {
    m.authenticationPeriodInSeconds = value
}
// SetAuthenticationRetryDelayPeriodInSeconds sets the authenticationRetryDelayPeriodInSeconds property value. Specify the number of seconds between a failed authentication and the next authentication attempt. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationRetryDelayPeriodInSeconds(value *int32)() {
    m.authenticationRetryDelayPeriodInSeconds = value
}
// SetAuthenticationType sets the authenticationType property value. Specify whether to authenticate the user, the device, either, or to use guest authentication (none). If you’re using certificate authentication, make sure the certificate type matches the authentication type. Possible values are: none, user, machine, machineOrUser, guest.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetAuthenticationType(value *WifiAuthenticationType)() {
    m.authenticationType = value
}
// SetCacheCredentials sets the cacheCredentials property value. Specify whether to cache user credentials on the device so that users don’t need to keep entering them each time they connect.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetCacheCredentials(value *bool)() {
    m.cacheCredentials = value
}
// SetDisableUserPromptForServerValidation sets the disableUserPromptForServerValidation property value. Specify whether to prevent the user from being prompted to authorize new servers for trusted certification authorities when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetDisableUserPromptForServerValidation(value *bool)() {
    m.disableUserPromptForServerValidation = value
}
// SetEapolStartPeriodInSeconds sets the eapolStartPeriodInSeconds property value. Specify the number of seconds to wait before sending an EAPOL (Extensible Authentication Protocol over LAN) Start message. Valid range 1-3600.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEapolStartPeriodInSeconds(value *int32)() {
    m.eapolStartPeriodInSeconds = value
}
// SetEapType sets the eapType property value. Extensible Authentication Protocol (EAP) configuration types.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEapType(value *EapType)() {
    m.eapType = value
}
// SetEnablePairwiseMasterKeyCaching sets the enablePairwiseMasterKeyCaching property value. Specify whether the wifi connection should enable pairwise master key caching.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEnablePairwiseMasterKeyCaching(value *bool)() {
    m.enablePairwiseMasterKeyCaching = value
}
// SetEnablePreAuthentication sets the enablePreAuthentication property value. Specify whether pre-authentication should be enabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetEnablePreAuthentication(value *bool)() {
    m.enablePreAuthentication = value
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Specify identity certificate for client authentication.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetIdentityCertificateForClientAuthentication(value WindowsCertificateProfileBaseable)() {
    m.identityCertificateForClientAuthentication = value
}
// SetInnerAuthenticationProtocolForEAPTTLS sets the innerAuthenticationProtocolForEAPTTLS property value. Specify inner authentication protocol for EAP TTLS. Possible values are: unencryptedPassword, challengeHandshakeAuthenticationProtocol, microsoftChap, microsoftChapVersionTwo.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetInnerAuthenticationProtocolForEAPTTLS(value *NonEapAuthenticationMethodForEapTtlsType)() {
    m.innerAuthenticationProtocolForEAPTTLS = value
}
// SetMaximumAuthenticationFailures sets the maximumAuthenticationFailures property value. Specify the maximum authentication failures allowed for a set of credentials. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumAuthenticationFailures(value *int32)() {
    m.maximumAuthenticationFailures = value
}
// SetMaximumAuthenticationTimeoutInSeconds sets the maximumAuthenticationTimeoutInSeconds property value. Specify maximum authentication timeout (in seconds).  Valid range: 1-120
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumAuthenticationTimeoutInSeconds(value *int32)() {
    m.maximumAuthenticationTimeoutInSeconds = value
}
// SetMaximumEAPOLStartMessages sets the maximumEAPOLStartMessages property value. Specifiy the maximum number of EAPOL (Extensible Authentication Protocol over LAN) Start messages to be sent before returning failure. Valid range 1-100.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumEAPOLStartMessages(value *int32)() {
    m.maximumEAPOLStartMessages = value
}
// SetMaximumNumberOfPairwiseMasterKeysInCache sets the maximumNumberOfPairwiseMasterKeysInCache property value. Specify maximum number of pairwise master keys in cache.  Valid range: 1-255
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumNumberOfPairwiseMasterKeysInCache(value *int32)() {
    m.maximumNumberOfPairwiseMasterKeysInCache = value
}
// SetMaximumPairwiseMasterKeyCacheTimeInMinutes sets the maximumPairwiseMasterKeyCacheTimeInMinutes property value. Specify maximum pairwise master key cache time (in minutes).  Valid range: 5-1440
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumPairwiseMasterKeyCacheTimeInMinutes(value *int32)() {
    m.maximumPairwiseMasterKeyCacheTimeInMinutes = value
}
// SetMaximumPreAuthenticationAttempts sets the maximumPreAuthenticationAttempts property value. Specify maximum pre-authentication attempts.  Valid range: 1-16
func (m *WindowsWifiEnterpriseEAPConfiguration) SetMaximumPreAuthenticationAttempts(value *int32)() {
    m.maximumPreAuthenticationAttempts = value
}
// SetNetworkSingleSignOn sets the networkSingleSignOn property value. Specify the network single sign on type. Possible values are: disabled, prelogon, postlogon.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetNetworkSingleSignOn(value *NetworkSingleSignOnType)() {
    m.networkSingleSignOn = value
}
// SetOuterIdentityPrivacyTemporaryValue sets the outerIdentityPrivacyTemporaryValue property value. Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetOuterIdentityPrivacyTemporaryValue(value *string)() {
    m.outerIdentityPrivacyTemporaryValue = value
}
// SetPerformServerValidation sets the performServerValidation property value. Specify whether to enable verification of server's identity by validating the certificate when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetPerformServerValidation(value *bool)() {
    m.performServerValidation = value
}
// SetPromptForAdditionalAuthenticationCredentials sets the promptForAdditionalAuthenticationCredentials property value. Specify whether the wifi connection should prompt for additional authentication credentials.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetPromptForAdditionalAuthenticationCredentials(value *bool)() {
    m.promptForAdditionalAuthenticationCredentials = value
}
// SetRequireCryptographicBinding sets the requireCryptographicBinding property value. Specify whether to enable cryptographic binding when EAP type is selected as PEAP.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetRequireCryptographicBinding(value *bool)() {
    m.requireCryptographicBinding = value
}
// SetRootCertificateForClientValidation sets the rootCertificateForClientValidation property value. Specify root certificate for client validation.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetRootCertificateForClientValidation(value Windows81TrustedRootCertificateable)() {
    m.rootCertificateForClientValidation = value
}
// SetRootCertificatesForServerValidation sets the rootCertificatesForServerValidation property value. Specify root certificate for server validation. This collection can contain a maximum of 500 elements.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetRootCertificatesForServerValidation(value []Windows81TrustedRootCertificateable)() {
    m.rootCertificatesForServerValidation = value
}
// SetTrustedServerCertificateNames sets the trustedServerCertificateNames property value. Specify trusted server certificate names.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetTrustedServerCertificateNames(value []string)() {
    m.trustedServerCertificateNames = value
}
// SetUserBasedVirtualLan sets the userBasedVirtualLan property value. Specifiy whether to change the virtual LAN used by the device based on the user’s credentials. Cannot be used when NetworkSingleSignOnType is set to ​Disabled.
func (m *WindowsWifiEnterpriseEAPConfiguration) SetUserBasedVirtualLan(value *bool)() {
    m.userBasedVirtualLan = value
}
