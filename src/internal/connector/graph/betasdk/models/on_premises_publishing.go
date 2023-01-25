package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishing 
type OnPremisesPublishing struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
    alternateUrl *string
    // The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
    applicationServerTimeout *string
    // Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
    applicationType *string
    // Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
    externalAuthenticationType *ExternalAuthenticationType
    // The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
    externalUrl *string
    // The internal url of the application. For example, https://intranet/.
    internalUrl *string
    // Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
    isBackendCertificateValidationEnabled *bool
    // Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
    isHttpOnlyCookieEnabled *bool
    // Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
    isOnPremPublishingEnabled *bool
    // Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
    isPersistentCookieEnabled *bool
    // Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
    isSecureCookieEnabled *bool
    // Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
    isStateSessionEnabled *bool
    // Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
    isTranslateHostHeaderEnabled *bool
    // Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
    isTranslateLinksInBodyEnabled *bool
    // The OdataType property
    odataType *string
    // Represents the application segment collection for an on-premises wildcard application.
    onPremisesApplicationSegments []OnPremisesApplicationSegmentable
    // The segmentsConfiguration property
    segmentsConfiguration SegmentConfigurationable
    // Represents the single sign-on configuration for the on-premises application.
    singleSignOnSettings OnPremisesPublishingSingleSignOnable
    // The useAlternateUrlForTranslationAndRedirect property
    useAlternateUrlForTranslationAndRedirect *bool
    // Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
    verifiedCustomDomainCertificatesMetadata VerifiedCustomDomainCertificatesMetadataable
    // The associated key credential for the custom domain used.
    verifiedCustomDomainKeyCredential KeyCredentialable
    // The associated password credential for the custom domain used.
    verifiedCustomDomainPasswordCredential PasswordCredentialable
}
// NewOnPremisesPublishing instantiates a new onPremisesPublishing and sets the default values.
func NewOnPremisesPublishing()(*OnPremisesPublishing) {
    m := &OnPremisesPublishing{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesPublishingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesPublishingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesPublishing(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishing) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAlternateUrl gets the alternateUrl property value. If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
func (m *OnPremisesPublishing) GetAlternateUrl()(*string) {
    return m.alternateUrl
}
// GetApplicationServerTimeout gets the applicationServerTimeout property value. The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
func (m *OnPremisesPublishing) GetApplicationServerTimeout()(*string) {
    return m.applicationServerTimeout
}
// GetApplicationType gets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) GetApplicationType()(*string) {
    return m.applicationType
}
// GetExternalAuthenticationType gets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
func (m *OnPremisesPublishing) GetExternalAuthenticationType()(*ExternalAuthenticationType) {
    return m.externalAuthenticationType
}
// GetExternalUrl gets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
func (m *OnPremisesPublishing) GetExternalUrl()(*string) {
    return m.externalUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishing) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alternateUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateUrl(val)
        }
        return nil
    }
    res["applicationServerTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationServerTimeout(val)
        }
        return nil
    }
    res["applicationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationType(val)
        }
        return nil
    }
    res["externalAuthenticationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAuthenticationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAuthenticationType(val.(*ExternalAuthenticationType))
        }
        return nil
    }
    res["externalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["internalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalUrl(val)
        }
        return nil
    }
    res["isBackendCertificateValidationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBackendCertificateValidationEnabled(val)
        }
        return nil
    }
    res["isHttpOnlyCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpOnlyCookieEnabled(val)
        }
        return nil
    }
    res["isOnPremPublishingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnPremPublishingEnabled(val)
        }
        return nil
    }
    res["isPersistentCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPersistentCookieEnabled(val)
        }
        return nil
    }
    res["isSecureCookieEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSecureCookieEnabled(val)
        }
        return nil
    }
    res["isStateSessionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsStateSessionEnabled(val)
        }
        return nil
    }
    res["isTranslateHostHeaderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateHostHeaderEnabled(val)
        }
        return nil
    }
    res["isTranslateLinksInBodyEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranslateLinksInBodyEnabled(val)
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
    res["onPremisesApplicationSegments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesApplicationSegmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesApplicationSegmentable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesApplicationSegmentable)
            }
            m.SetOnPremisesApplicationSegments(res)
        }
        return nil
    }
    res["segmentsConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSegmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSegmentsConfiguration(val.(SegmentConfigurationable))
        }
        return nil
    }
    res["singleSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnSettings(val.(OnPremisesPublishingSingleSignOnable))
        }
        return nil
    }
    res["useAlternateUrlForTranslationAndRedirect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseAlternateUrlForTranslationAndRedirect(val)
        }
        return nil
    }
    res["verifiedCustomDomainCertificatesMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerifiedCustomDomainCertificatesMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainCertificatesMetadata(val.(VerifiedCustomDomainCertificatesMetadataable))
        }
        return nil
    }
    res["verifiedCustomDomainKeyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainKeyCredential(val.(KeyCredentialable))
        }
        return nil
    }
    res["verifiedCustomDomainPasswordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedCustomDomainPasswordCredential(val.(PasswordCredentialable))
        }
        return nil
    }
    return res
}
// GetInternalUrl gets the internalUrl property value. The internal url of the application. For example, https://intranet/.
func (m *OnPremisesPublishing) GetInternalUrl()(*string) {
    return m.internalUrl
}
// GetIsBackendCertificateValidationEnabled gets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
func (m *OnPremisesPublishing) GetIsBackendCertificateValidationEnabled()(*bool) {
    return m.isBackendCertificateValidationEnabled
}
// GetIsHttpOnlyCookieEnabled gets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
func (m *OnPremisesPublishing) GetIsHttpOnlyCookieEnabled()(*bool) {
    return m.isHttpOnlyCookieEnabled
}
// GetIsOnPremPublishingEnabled gets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) GetIsOnPremPublishingEnabled()(*bool) {
    return m.isOnPremPublishingEnabled
}
// GetIsPersistentCookieEnabled gets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
func (m *OnPremisesPublishing) GetIsPersistentCookieEnabled()(*bool) {
    return m.isPersistentCookieEnabled
}
// GetIsSecureCookieEnabled gets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
func (m *OnPremisesPublishing) GetIsSecureCookieEnabled()(*bool) {
    return m.isSecureCookieEnabled
}
// GetIsStateSessionEnabled gets the isStateSessionEnabled property value. Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
func (m *OnPremisesPublishing) GetIsStateSessionEnabled()(*bool) {
    return m.isStateSessionEnabled
}
// GetIsTranslateHostHeaderEnabled gets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
func (m *OnPremisesPublishing) GetIsTranslateHostHeaderEnabled()(*bool) {
    return m.isTranslateHostHeaderEnabled
}
// GetIsTranslateLinksInBodyEnabled gets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
func (m *OnPremisesPublishing) GetIsTranslateLinksInBodyEnabled()(*bool) {
    return m.isTranslateLinksInBodyEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishing) GetOdataType()(*string) {
    return m.odataType
}
// GetOnPremisesApplicationSegments gets the onPremisesApplicationSegments property value. Represents the application segment collection for an on-premises wildcard application.
func (m *OnPremisesPublishing) GetOnPremisesApplicationSegments()([]OnPremisesApplicationSegmentable) {
    return m.onPremisesApplicationSegments
}
// GetSegmentsConfiguration gets the segmentsConfiguration property value. The segmentsConfiguration property
func (m *OnPremisesPublishing) GetSegmentsConfiguration()(SegmentConfigurationable) {
    return m.segmentsConfiguration
}
// GetSingleSignOnSettings gets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
func (m *OnPremisesPublishing) GetSingleSignOnSettings()(OnPremisesPublishingSingleSignOnable) {
    return m.singleSignOnSettings
}
// GetUseAlternateUrlForTranslationAndRedirect gets the useAlternateUrlForTranslationAndRedirect property value. The useAlternateUrlForTranslationAndRedirect property
func (m *OnPremisesPublishing) GetUseAlternateUrlForTranslationAndRedirect()(*bool) {
    return m.useAlternateUrlForTranslationAndRedirect
}
// GetVerifiedCustomDomainCertificatesMetadata gets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainCertificatesMetadata()(VerifiedCustomDomainCertificatesMetadataable) {
    return m.verifiedCustomDomainCertificatesMetadata
}
// GetVerifiedCustomDomainKeyCredential gets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainKeyCredential()(KeyCredentialable) {
    return m.verifiedCustomDomainKeyCredential
}
// GetVerifiedCustomDomainPasswordCredential gets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
func (m *OnPremisesPublishing) GetVerifiedCustomDomainPasswordCredential()(PasswordCredentialable) {
    return m.verifiedCustomDomainPasswordCredential
}
// Serialize serializes information the current object
func (m *OnPremisesPublishing) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alternateUrl", m.GetAlternateUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationServerTimeout", m.GetApplicationServerTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationType", m.GetApplicationType())
        if err != nil {
            return err
        }
    }
    if m.GetExternalAuthenticationType() != nil {
        cast := (*m.GetExternalAuthenticationType()).String()
        err := writer.WriteStringValue("externalAuthenticationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalUrl", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalUrl", m.GetInternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBackendCertificateValidationEnabled", m.GetIsBackendCertificateValidationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHttpOnlyCookieEnabled", m.GetIsHttpOnlyCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOnPremPublishingEnabled", m.GetIsOnPremPublishingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPersistentCookieEnabled", m.GetIsPersistentCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSecureCookieEnabled", m.GetIsSecureCookieEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isStateSessionEnabled", m.GetIsStateSessionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isTranslateHostHeaderEnabled", m.GetIsTranslateHostHeaderEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isTranslateLinksInBodyEnabled", m.GetIsTranslateLinksInBodyEnabled())
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
    if m.GetOnPremisesApplicationSegments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesApplicationSegments()))
        for i, v := range m.GetOnPremisesApplicationSegments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("onPremisesApplicationSegments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("segmentsConfiguration", m.GetSegmentsConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("singleSignOnSettings", m.GetSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useAlternateUrlForTranslationAndRedirect", m.GetUseAlternateUrlForTranslationAndRedirect())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainCertificatesMetadata", m.GetVerifiedCustomDomainCertificatesMetadata())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainKeyCredential", m.GetVerifiedCustomDomainKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verifiedCustomDomainPasswordCredential", m.GetVerifiedCustomDomainPasswordCredential())
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
func (m *OnPremisesPublishing) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAlternateUrl sets the alternateUrl property value. If you are configuring a traffic manager in front of multiple App Proxy applications, the alternateUrl is the user-friendly URL that will point to the traffic manager.
func (m *OnPremisesPublishing) SetAlternateUrl(value *string)() {
    m.alternateUrl = value
}
// SetApplicationServerTimeout sets the applicationServerTimeout property value. The duration the connector will wait for a response from the backend application before closing the connection. Possible values are default, long. When set to default, the backend application timeout has a length of 85 seconds. When set to long, the backend timeout is increased to 180 seconds. Use long if your server takes more than 85 seconds to respond to requests or if you are unable to access the application and the error status is 'Backend Timeout'. Default value is default.
func (m *OnPremisesPublishing) SetApplicationServerTimeout(value *string)() {
    m.applicationServerTimeout = value
}
// SetApplicationType sets the applicationType property value. Indicates if this application is an Application Proxy configured application. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) SetApplicationType(value *string)() {
    m.applicationType = value
}
// SetExternalAuthenticationType sets the externalAuthenticationType property value. Details the pre-authentication setting for the application. Pre-authentication enforces that users must authenticate before accessing the app. Passthru does not require authentication. Possible values are: passthru, aadPreAuthentication.
func (m *OnPremisesPublishing) SetExternalAuthenticationType(value *ExternalAuthenticationType)() {
    m.externalAuthenticationType = value
}
// SetExternalUrl sets the externalUrl property value. The published external url for the application. For example, https://intranet-contoso.msappproxy.net/.
func (m *OnPremisesPublishing) SetExternalUrl(value *string)() {
    m.externalUrl = value
}
// SetInternalUrl sets the internalUrl property value. The internal url of the application. For example, https://intranet/.
func (m *OnPremisesPublishing) SetInternalUrl(value *string)() {
    m.internalUrl = value
}
// SetIsBackendCertificateValidationEnabled sets the isBackendCertificateValidationEnabled property value. Indicates whether backend SSL certificate validation is enabled for the application. For all new Application Proxy apps, the property will be set to true by default. For all existing apps, the property will be set to false.
func (m *OnPremisesPublishing) SetIsBackendCertificateValidationEnabled(value *bool)() {
    m.isBackendCertificateValidationEnabled = value
}
// SetIsHttpOnlyCookieEnabled sets the isHttpOnlyCookieEnabled property value. Indicates if the HTTPOnly cookie flag should be set in the HTTP response headers. Set this value to true to have Application Proxy cookies include the HTTPOnly flag in the HTTP response headers. If using Remote Desktop Services, set this value to False. Default value is false.
func (m *OnPremisesPublishing) SetIsHttpOnlyCookieEnabled(value *bool)() {
    m.isHttpOnlyCookieEnabled = value
}
// SetIsOnPremPublishingEnabled sets the isOnPremPublishingEnabled property value. Indicates if the application is currently being published via Application Proxy or not. This is pre-set by the system. Read-only.
func (m *OnPremisesPublishing) SetIsOnPremPublishingEnabled(value *bool)() {
    m.isOnPremPublishingEnabled = value
}
// SetIsPersistentCookieEnabled sets the isPersistentCookieEnabled property value. Indicates if the Persistent cookie flag should be set in the HTTP response headers. Keep this value set to false. Only use this setting for applications that can't share cookies between processes. For more information about cookie settings, see Cookie settings for accessing on-premises applications in Azure Active Directory. Default value is false.
func (m *OnPremisesPublishing) SetIsPersistentCookieEnabled(value *bool)() {
    m.isPersistentCookieEnabled = value
}
// SetIsSecureCookieEnabled sets the isSecureCookieEnabled property value. Indicates if the Secure cookie flag should be set in the HTTP response headers. Set this value to true to transmit cookies over a secure channel such as an encrypted HTTPS request. Default value is true.
func (m *OnPremisesPublishing) SetIsSecureCookieEnabled(value *bool)() {
    m.isSecureCookieEnabled = value
}
// SetIsStateSessionEnabled sets the isStateSessionEnabled property value. Indicates whether validation of the state parameter when the client uses the OAuth 2.0 authorization code grant flow is enabled. This setting allows admins to specify whether they want to enable CSRF protection for their apps.
func (m *OnPremisesPublishing) SetIsStateSessionEnabled(value *bool)() {
    m.isStateSessionEnabled = value
}
// SetIsTranslateHostHeaderEnabled sets the isTranslateHostHeaderEnabled property value. Indicates if the application should translate urls in the reponse headers. Keep this value as true unless your application required the original host header in the authentication request. Default value is true.
func (m *OnPremisesPublishing) SetIsTranslateHostHeaderEnabled(value *bool)() {
    m.isTranslateHostHeaderEnabled = value
}
// SetIsTranslateLinksInBodyEnabled sets the isTranslateLinksInBodyEnabled property value. Indicates if the application should translate urls in the application body. Keep this value as false unless you have hardcoded HTML links to other on-premises applications and don't use custom domains. For more information, see Link translation with Application Proxy. Default value is false.
func (m *OnPremisesPublishing) SetIsTranslateLinksInBodyEnabled(value *bool)() {
    m.isTranslateLinksInBodyEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishing) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnPremisesApplicationSegments sets the onPremisesApplicationSegments property value. Represents the application segment collection for an on-premises wildcard application.
func (m *OnPremisesPublishing) SetOnPremisesApplicationSegments(value []OnPremisesApplicationSegmentable)() {
    m.onPremisesApplicationSegments = value
}
// SetSegmentsConfiguration sets the segmentsConfiguration property value. The segmentsConfiguration property
func (m *OnPremisesPublishing) SetSegmentsConfiguration(value SegmentConfigurationable)() {
    m.segmentsConfiguration = value
}
// SetSingleSignOnSettings sets the singleSignOnSettings property value. Represents the single sign-on configuration for the on-premises application.
func (m *OnPremisesPublishing) SetSingleSignOnSettings(value OnPremisesPublishingSingleSignOnable)() {
    m.singleSignOnSettings = value
}
// SetUseAlternateUrlForTranslationAndRedirect sets the useAlternateUrlForTranslationAndRedirect property value. The useAlternateUrlForTranslationAndRedirect property
func (m *OnPremisesPublishing) SetUseAlternateUrlForTranslationAndRedirect(value *bool)() {
    m.useAlternateUrlForTranslationAndRedirect = value
}
// SetVerifiedCustomDomainCertificatesMetadata sets the verifiedCustomDomainCertificatesMetadata property value. Details of the certificate associated with the application when a custom domain is in use. null when using the default domain. Read-only.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainCertificatesMetadata(value VerifiedCustomDomainCertificatesMetadataable)() {
    m.verifiedCustomDomainCertificatesMetadata = value
}
// SetVerifiedCustomDomainKeyCredential sets the verifiedCustomDomainKeyCredential property value. The associated key credential for the custom domain used.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainKeyCredential(value KeyCredentialable)() {
    m.verifiedCustomDomainKeyCredential = value
}
// SetVerifiedCustomDomainPasswordCredential sets the verifiedCustomDomainPasswordCredential property value. The associated password credential for the custom domain used.
func (m *OnPremisesPublishing) SetVerifiedCustomDomainPasswordCredential(value PasswordCredentialable)() {
    m.verifiedCustomDomainPasswordCredential = value
}
