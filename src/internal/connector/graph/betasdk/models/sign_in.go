package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SignIn 
type SignIn struct {
    Entity
    // The application name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
    appDisplayName *string
    // The application identifier in Azure Active Directory. Supports $filter (eq operator only).
    appId *string
    // A list of conditional access policies that are triggered by the corresponding sign-in activity.
    appliedConditionalAccessPolicies []AppliedConditionalAccessPolicyable
    // The appliedEventListeners property
    appliedEventListeners []AppliedAuthenticationEventListenerable
    // Provides details about the app and device used during an Azure AD authentication step.
    authenticationAppDeviceDetails AuthenticationAppDeviceDetailsable
    // Provides details of the Azure AD policies applied to a user and client authentication app during an authentication step.
    authenticationAppPolicyEvaluationDetails []AuthenticationAppPolicyDetailsable
    // Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
    authenticationContextClassReferences []AuthenticationContextable
    // The result of the authentication attempt and additional details on the authentication method.
    authenticationDetails []AuthenticationDetailable
    // The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
    authenticationMethodsUsed []string
    // Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
    authenticationProcessingDetails []KeyValueable
    // Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
    authenticationProtocol *ProtocolType
    // This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
    authenticationRequirement *string
    // Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
    authenticationRequirementPolicies []AuthenticationRequirementPolicyable
    // The Autonomous System Number (ASN) of the network used by the actor.
    autonomousSystemNumber *int32
    // Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
    azureResourceId *string
    // The legacy client used for sign-in activity. For example: Browser, Exchange ActiveSync, Modern clients, IMAP, MAPI, SMTP, or POP. Supports $filter (eq operator only).
    clientAppUsed *string
    // Describes the credential type that a user client or service principal provided to Azure AD to authenticate itself. You may wish to review clientCredentialType to track and eliminate less secure credential types or to watch for clients and service principals using anomalous credential types. The possible values are: none, clientSecret, clientAssertion, federatedIdentityCredential, managedIdentity, certificate, unknownFutureValue.
    clientCredentialType *ClientCredentialType
    // The status of the conditional access policy triggered. Possible values: success, failure, notApplied, or unknownFutureValue. Supports $filter (eq operator only).
    conditionalAccessStatus *ConditionalAccessStatus
    // The identifier that's sent from the client when sign-in is initiated. This is used for troubleshooting the corresponding sign-in activity when calling for support. Supports $filter (eq operator only).
    correlationId *string
    // The date and time the sign-in was initiated. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
    crossTenantAccessType *SignInAccessType
    // The device information from where the sign-in occurred. Includes information such as deviceId, OS, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSystem properties.
    deviceDetail DeviceDetailable
    // Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
    federatedCredentialId *string
    // During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
    flaggedForReview *bool
    // The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
    homeTenantId *string
    // For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
    homeTenantName *string
    // Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue, remoteDesktopToken.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed. Also, please note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: remoteDesktopToken.
    incomingTokenType *IncomingTokenType
    // The IP address of the client from where the sign-in occurred. Supports $filter (eq and startsWith operators only).
    ipAddress *string
    // The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
    ipAddressFromResourceProvider *string
    // Indicates whether a user sign in is interactive. In interactive sign in, the user provides an authentication factor to Azure AD. These factors include passwords, responses to MFA challenges, biometric factors, or QR codes that a user provides to Azure AD or an associated app. In non-interactive sign in, the user doesn't provide an authentication factor. Instead, the client app uses a token or code to authenticate or access a resource on behalf of a user. Non-interactive sign ins are commonly used for a client to sign in on a user's behalf in a process transparent to the user.
    isInteractive *bool
    // Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
    isTenantRestricted *bool
    // The city, state, and 2 letter country code from where the sign-in occurred. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
    location SignInLocationable
    // The mfaDetail property
    mfaDetail MfaDetailable
    // The network location details including the type of network used and its names.
    networkLocationDetails []NetworkLocationDetailable
    // The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
    originalRequestId *string
    // Contains information about the Azure AD Private Link policy that is associated with the sign in event.
    privateLinkDetails PrivateLinkDetailsable
    // The request processing time in milliseconds in AD STS.
    processingTimeInMilliseconds *int32
    // The name of the resource that the user signed in to. Supports $filter (eq operator only).
    resourceDisplayName *string
    // The identifier of the resource that the user signed in to. Supports $filter (eq operator only).
    resourceId *string
    // The identifier of the service principal representing the target resource in the sign-in event.
    resourceServicePrincipalId *string
    // The tenant identifier of the resource referenced in the sign in.
    resourceTenantId *string
    // The reason behind a specific state of a risky user, sign-in, or a risk event. Possible values: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, or unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
    riskDetail *RiskDetail
    // The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
    riskEventTypes_v2 []string
    // The aggregated risk level. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
    riskLevelAggregated *RiskLevel
    // The risk level during sign-in. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
    riskLevelDuringSignIn *RiskLevel
    // The risk state of a risky user, sign-in, or a risk event. Possible values: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, or unknownFutureValue. Supports $filter (eq operator only).
    riskState *RiskState
    // The unique identifier of the key credential used by the service principal to authenticate.
    servicePrincipalCredentialKeyId *string
    // The certificate thumbprint of the certificate used by the service principal to authenticate.
    servicePrincipalCredentialThumbprint *string
    // The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
    servicePrincipalId *string
    // The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
    servicePrincipalName *string
    // Any conditional access session management policies that were applied during the sign-in event.
    sessionLifetimePolicies []SessionLifetimePolicyable
    // Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue. Supports $filter (eq, ne).
    signInEventTypes []string
    // The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
    signInIdentifier *string
    // The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
    signInIdentifierType *SignInIdentifierType
    // The sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
    status SignInStatusable
    // The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
    tokenIssuerName *string
    // The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
    tokenIssuerType *TokenIssuerType
    // A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
    uniqueTokenIdentifier *string
    // The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
    userAgent *string
    // The display name of the user. Supports $filter (eq and startsWith operators only).
    userDisplayName *string
    // The identifier of the user. Supports $filter (eq operator only).
    userId *string
    // The UPN of the user. Supports $filter (eq and startsWith operators only).
    userPrincipalName *string
    // Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
    userType *SignInUserType
}
// NewSignIn instantiates a new SignIn and sets the default values.
func NewSignIn()(*SignIn) {
    m := &SignIn{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSignInFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSignInFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignIn(), nil
}
// GetAppDisplayName gets the appDisplayName property value. The application name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppId gets the appId property value. The application identifier in Azure Active Directory. Supports $filter (eq operator only).
func (m *SignIn) GetAppId()(*string) {
    return m.appId
}
// GetAppliedConditionalAccessPolicies gets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicyable) {
    return m.appliedConditionalAccessPolicies
}
// GetAppliedEventListeners gets the appliedEventListeners property value. The appliedEventListeners property
func (m *SignIn) GetAppliedEventListeners()([]AppliedAuthenticationEventListenerable) {
    return m.appliedEventListeners
}
// GetAuthenticationAppDeviceDetails gets the authenticationAppDeviceDetails property value. Provides details about the app and device used during an Azure AD authentication step.
func (m *SignIn) GetAuthenticationAppDeviceDetails()(AuthenticationAppDeviceDetailsable) {
    return m.authenticationAppDeviceDetails
}
// GetAuthenticationAppPolicyEvaluationDetails gets the authenticationAppPolicyEvaluationDetails property value. Provides details of the Azure AD policies applied to a user and client authentication app during an authentication step.
func (m *SignIn) GetAuthenticationAppPolicyEvaluationDetails()([]AuthenticationAppPolicyDetailsable) {
    return m.authenticationAppPolicyEvaluationDetails
}
// GetAuthenticationContextClassReferences gets the authenticationContextClassReferences property value. Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
func (m *SignIn) GetAuthenticationContextClassReferences()([]AuthenticationContextable) {
    return m.authenticationContextClassReferences
}
// GetAuthenticationDetails gets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
func (m *SignIn) GetAuthenticationDetails()([]AuthenticationDetailable) {
    return m.authenticationDetails
}
// GetAuthenticationMethodsUsed gets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
func (m *SignIn) GetAuthenticationMethodsUsed()([]string) {
    return m.authenticationMethodsUsed
}
// GetAuthenticationProcessingDetails gets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
func (m *SignIn) GetAuthenticationProcessingDetails()([]KeyValueable) {
    return m.authenticationProcessingDetails
}
// GetAuthenticationProtocol gets the authenticationProtocol property value. Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
func (m *SignIn) GetAuthenticationProtocol()(*ProtocolType) {
    return m.authenticationProtocol
}
// GetAuthenticationRequirement gets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetAuthenticationRequirement()(*string) {
    return m.authenticationRequirement
}
// GetAuthenticationRequirementPolicies gets the authenticationRequirementPolicies property value. Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
func (m *SignIn) GetAuthenticationRequirementPolicies()([]AuthenticationRequirementPolicyable) {
    return m.authenticationRequirementPolicies
}
// GetAutonomousSystemNumber gets the autonomousSystemNumber property value. The Autonomous System Number (ASN) of the network used by the actor.
func (m *SignIn) GetAutonomousSystemNumber()(*int32) {
    return m.autonomousSystemNumber
}
// GetAzureResourceId gets the azureResourceId property value. Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
func (m *SignIn) GetAzureResourceId()(*string) {
    return m.azureResourceId
}
// GetClientAppUsed gets the clientAppUsed property value. The legacy client used for sign-in activity. For example: Browser, Exchange ActiveSync, Modern clients, IMAP, MAPI, SMTP, or POP. Supports $filter (eq operator only).
func (m *SignIn) GetClientAppUsed()(*string) {
    return m.clientAppUsed
}
// GetClientCredentialType gets the clientCredentialType property value. Describes the credential type that a user client or service principal provided to Azure AD to authenticate itself. You may wish to review clientCredentialType to track and eliminate less secure credential types or to watch for clients and service principals using anomalous credential types. The possible values are: none, clientSecret, clientAssertion, federatedIdentityCredential, managedIdentity, certificate, unknownFutureValue.
func (m *SignIn) GetClientCredentialType()(*ClientCredentialType) {
    return m.clientCredentialType
}
// GetConditionalAccessStatus gets the conditionalAccessStatus property value. The status of the conditional access policy triggered. Possible values: success, failure, notApplied, or unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetConditionalAccessStatus()(*ConditionalAccessStatus) {
    return m.conditionalAccessStatus
}
// GetCorrelationId gets the correlationId property value. The identifier that's sent from the client when sign-in is initiated. This is used for troubleshooting the corresponding sign-in activity when calling for support. Supports $filter (eq operator only).
func (m *SignIn) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the sign-in was initiated. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
func (m *SignIn) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCrossTenantAccessType gets the crossTenantAccessType property value. Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
func (m *SignIn) GetCrossTenantAccessType()(*SignInAccessType) {
    return m.crossTenantAccessType
}
// GetDeviceDetail gets the deviceDetail property value. The device information from where the sign-in occurred. Includes information such as deviceId, OS, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSystem properties.
func (m *SignIn) GetDeviceDetail()(DeviceDetailable) {
    return m.deviceDetail
}
// GetFederatedCredentialId gets the federatedCredentialId property value. Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
func (m *SignIn) GetFederatedCredentialId()(*string) {
    return m.federatedCredentialId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignIn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
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
    res["appliedConditionalAccessPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppliedConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppliedConditionalAccessPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AppliedConditionalAccessPolicyable)
            }
            m.SetAppliedConditionalAccessPolicies(res)
        }
        return nil
    }
    res["appliedEventListeners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppliedAuthenticationEventListenerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppliedAuthenticationEventListenerable, len(val))
            for i, v := range val {
                res[i] = v.(AppliedAuthenticationEventListenerable)
            }
            m.SetAppliedEventListeners(res)
        }
        return nil
    }
    res["authenticationAppDeviceDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationAppDeviceDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationAppDeviceDetails(val.(AuthenticationAppDeviceDetailsable))
        }
        return nil
    }
    res["authenticationAppPolicyEvaluationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationAppPolicyDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationAppPolicyDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationAppPolicyDetailsable)
            }
            m.SetAuthenticationAppPolicyEvaluationDetails(res)
        }
        return nil
    }
    res["authenticationContextClassReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationContextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationContextable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationContextable)
            }
            m.SetAuthenticationContextClassReferences(res)
        }
        return nil
    }
    res["authenticationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationDetailable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationDetailable)
            }
            m.SetAuthenticationDetails(res)
        }
        return nil
    }
    res["authenticationMethodsUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAuthenticationMethodsUsed(res)
        }
        return nil
    }
    res["authenticationProcessingDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValueable)
            }
            m.SetAuthenticationProcessingDetails(res)
        }
        return nil
    }
    res["authenticationProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtocolType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationProtocol(val.(*ProtocolType))
        }
        return nil
    }
    res["authenticationRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationRequirement(val)
        }
        return nil
    }
    res["authenticationRequirementPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationRequirementPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationRequirementPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationRequirementPolicyable)
            }
            m.SetAuthenticationRequirementPolicies(res)
        }
        return nil
    }
    res["autonomousSystemNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutonomousSystemNumber(val)
        }
        return nil
    }
    res["azureResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureResourceId(val)
        }
        return nil
    }
    res["clientAppUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppUsed(val)
        }
        return nil
    }
    res["clientCredentialType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientCredentialType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientCredentialType(val.(*ClientCredentialType))
        }
        return nil
    }
    res["conditionalAccessStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessStatus(val.(*ConditionalAccessStatus))
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["crossTenantAccessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossTenantAccessType(val.(*SignInAccessType))
        }
        return nil
    }
    res["deviceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDetail(val.(DeviceDetailable))
        }
        return nil
    }
    res["federatedCredentialId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedCredentialId(val)
        }
        return nil
    }
    res["flaggedForReview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlaggedForReview(val)
        }
        return nil
    }
    res["homeTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
        }
        return nil
    }
    res["homeTenantName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantName(val)
        }
        return nil
    }
    res["incomingTokenType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncomingTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncomingTokenType(val.(*IncomingTokenType))
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["ipAddressFromResourceProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressFromResourceProvider(val)
        }
        return nil
    }
    res["isInteractive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInteractive(val)
        }
        return nil
    }
    res["isTenantRestricted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTenantRestricted(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(SignInLocationable))
        }
        return nil
    }
    res["mfaDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMfaDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaDetail(val.(MfaDetailable))
        }
        return nil
    }
    res["networkLocationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNetworkLocationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkLocationDetailable, len(val))
            for i, v := range val {
                res[i] = v.(NetworkLocationDetailable)
            }
            m.SetNetworkLocationDetails(res)
        }
        return nil
    }
    res["originalRequestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalRequestId(val)
        }
        return nil
    }
    res["privateLinkDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivateLinkDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateLinkDetails(val.(PrivateLinkDetailsable))
        }
        return nil
    }
    res["processingTimeInMilliseconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingTimeInMilliseconds(val)
        }
        return nil
    }
    res["resourceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceDisplayName(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["resourceServicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceServicePrincipalId(val)
        }
        return nil
    }
    res["resourceTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceTenantId(val)
        }
        return nil
    }
    res["riskDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskDetail(val.(*RiskDetail))
        }
        return nil
    }
    res["riskEventTypes_v2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRiskEventTypes_v2(res)
        }
        return nil
    }
    res["riskLevelAggregated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevelAggregated(val.(*RiskLevel))
        }
        return nil
    }
    res["riskLevelDuringSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevelDuringSignIn(val.(*RiskLevel))
        }
        return nil
    }
    res["riskState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskState(val.(*RiskState))
        }
        return nil
    }
    res["servicePrincipalCredentialKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalCredentialKeyId(val)
        }
        return nil
    }
    res["servicePrincipalCredentialThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalCredentialThumbprint(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    res["sessionLifetimePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSessionLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SessionLifetimePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(SessionLifetimePolicyable)
            }
            m.SetSessionLifetimePolicies(res)
        }
        return nil
    }
    res["signInEventTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSignInEventTypes(res)
        }
        return nil
    }
    res["signInIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInIdentifier(val)
        }
        return nil
    }
    res["signInIdentifierType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInIdentifierType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInIdentifierType(val.(*SignInIdentifierType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(SignInStatusable))
        }
        return nil
    }
    res["tokenIssuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenIssuerName(val)
        }
        return nil
    }
    res["tokenIssuerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTokenIssuerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenIssuerType(val.(*TokenIssuerType))
        }
        return nil
    }
    res["uniqueTokenIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueTokenIdentifier(val)
        }
        return nil
    }
    res["userAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAgent(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*SignInUserType))
        }
        return nil
    }
    return res
}
// GetFlaggedForReview gets the flaggedForReview property value. During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
func (m *SignIn) GetFlaggedForReview()(*bool) {
    return m.flaggedForReview
}
// GetHomeTenantId gets the homeTenantId property value. The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
func (m *SignIn) GetHomeTenantId()(*string) {
    return m.homeTenantId
}
// GetHomeTenantName gets the homeTenantName property value. For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *SignIn) GetHomeTenantName()(*string) {
    return m.homeTenantName
}
// GetIncomingTokenType gets the incomingTokenType property value. Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue, remoteDesktopToken.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed. Also, please note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: remoteDesktopToken.
func (m *SignIn) GetIncomingTokenType()(*IncomingTokenType) {
    return m.incomingTokenType
}
// GetIpAddress gets the ipAddress property value. The IP address of the client from where the sign-in occurred. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetIpAddressFromResourceProvider gets the ipAddressFromResourceProvider property value. The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
func (m *SignIn) GetIpAddressFromResourceProvider()(*string) {
    return m.ipAddressFromResourceProvider
}
// GetIsInteractive gets the isInteractive property value. Indicates whether a user sign in is interactive. In interactive sign in, the user provides an authentication factor to Azure AD. These factors include passwords, responses to MFA challenges, biometric factors, or QR codes that a user provides to Azure AD or an associated app. In non-interactive sign in, the user doesn't provide an authentication factor. Instead, the client app uses a token or code to authenticate or access a resource on behalf of a user. Non-interactive sign ins are commonly used for a client to sign in on a user's behalf in a process transparent to the user.
func (m *SignIn) GetIsInteractive()(*bool) {
    return m.isInteractive
}
// GetIsTenantRestricted gets the isTenantRestricted property value. Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
func (m *SignIn) GetIsTenantRestricted()(*bool) {
    return m.isTenantRestricted
}
// GetLocation gets the location property value. The city, state, and 2 letter country code from where the sign-in occurred. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
func (m *SignIn) GetLocation()(SignInLocationable) {
    return m.location
}
// GetMfaDetail gets the mfaDetail property value. The mfaDetail property
func (m *SignIn) GetMfaDetail()(MfaDetailable) {
    return m.mfaDetail
}
// GetNetworkLocationDetails gets the networkLocationDetails property value. The network location details including the type of network used and its names.
func (m *SignIn) GetNetworkLocationDetails()([]NetworkLocationDetailable) {
    return m.networkLocationDetails
}
// GetOriginalRequestId gets the originalRequestId property value. The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
func (m *SignIn) GetOriginalRequestId()(*string) {
    return m.originalRequestId
}
// GetPrivateLinkDetails gets the privateLinkDetails property value. Contains information about the Azure AD Private Link policy that is associated with the sign in event.
func (m *SignIn) GetPrivateLinkDetails()(PrivateLinkDetailsable) {
    return m.privateLinkDetails
}
// GetProcessingTimeInMilliseconds gets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
func (m *SignIn) GetProcessingTimeInMilliseconds()(*int32) {
    return m.processingTimeInMilliseconds
}
// GetResourceDisplayName gets the resourceDisplayName property value. The name of the resource that the user signed in to. Supports $filter (eq operator only).
func (m *SignIn) GetResourceDisplayName()(*string) {
    return m.resourceDisplayName
}
// GetResourceId gets the resourceId property value. The identifier of the resource that the user signed in to. Supports $filter (eq operator only).
func (m *SignIn) GetResourceId()(*string) {
    return m.resourceId
}
// GetResourceServicePrincipalId gets the resourceServicePrincipalId property value. The identifier of the service principal representing the target resource in the sign-in event.
func (m *SignIn) GetResourceServicePrincipalId()(*string) {
    return m.resourceServicePrincipalId
}
// GetResourceTenantId gets the resourceTenantId property value. The tenant identifier of the resource referenced in the sign in.
func (m *SignIn) GetResourceTenantId()(*string) {
    return m.resourceTenantId
}
// GetRiskDetail gets the riskDetail property value. The reason behind a specific state of a risky user, sign-in, or a risk event. Possible values: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, or unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) GetRiskDetail()(*RiskDetail) {
    return m.riskDetail
}
// GetRiskEventTypes_v2 gets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetRiskEventTypes_v2()([]string) {
    return m.riskEventTypes_v2
}
// GetRiskLevelAggregated gets the riskLevelAggregated property value. The aggregated risk level. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) GetRiskLevelAggregated()(*RiskLevel) {
    return m.riskLevelAggregated
}
// GetRiskLevelDuringSignIn gets the riskLevelDuringSignIn property value. The risk level during sign-in. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) GetRiskLevelDuringSignIn()(*RiskLevel) {
    return m.riskLevelDuringSignIn
}
// GetRiskState gets the riskState property value. The risk state of a risky user, sign-in, or a risk event. Possible values: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, or unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetRiskState()(*RiskState) {
    return m.riskState
}
// GetServicePrincipalCredentialKeyId gets the servicePrincipalCredentialKeyId property value. The unique identifier of the key credential used by the service principal to authenticate.
func (m *SignIn) GetServicePrincipalCredentialKeyId()(*string) {
    return m.servicePrincipalCredentialKeyId
}
// GetServicePrincipalCredentialThumbprint gets the servicePrincipalCredentialThumbprint property value. The certificate thumbprint of the certificate used by the service principal to authenticate.
func (m *SignIn) GetServicePrincipalCredentialThumbprint()(*string) {
    return m.servicePrincipalCredentialThumbprint
}
// GetServicePrincipalId gets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetServicePrincipalId()(*string) {
    return m.servicePrincipalId
}
// GetServicePrincipalName gets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetServicePrincipalName()(*string) {
    return m.servicePrincipalName
}
// GetSessionLifetimePolicies gets the sessionLifetimePolicies property value. Any conditional access session management policies that were applied during the sign-in event.
func (m *SignIn) GetSessionLifetimePolicies()([]SessionLifetimePolicyable) {
    return m.sessionLifetimePolicies
}
// GetSignInEventTypes gets the signInEventTypes property value. Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue. Supports $filter (eq, ne).
func (m *SignIn) GetSignInEventTypes()([]string) {
    return m.signInEventTypes
}
// GetSignInIdentifier gets the signInIdentifier property value. The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
func (m *SignIn) GetSignInIdentifier()(*string) {
    return m.signInIdentifier
}
// GetSignInIdentifierType gets the signInIdentifierType property value. The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
func (m *SignIn) GetSignInIdentifierType()(*SignInIdentifierType) {
    return m.signInIdentifierType
}
// GetStatus gets the status property value. The sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
func (m *SignIn) GetStatus()(SignInStatusable) {
    return m.status
}
// GetTokenIssuerName gets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
func (m *SignIn) GetTokenIssuerName()(*string) {
    return m.tokenIssuerName
}
// GetTokenIssuerType gets the tokenIssuerType property value. The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
func (m *SignIn) GetTokenIssuerType()(*TokenIssuerType) {
    return m.tokenIssuerType
}
// GetUniqueTokenIdentifier gets the uniqueTokenIdentifier property value. A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
func (m *SignIn) GetUniqueTokenIdentifier()(*string) {
    return m.uniqueTokenIdentifier
}
// GetUserAgent gets the userAgent property value. The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserAgent()(*string) {
    return m.userAgent
}
// GetUserDisplayName gets the userDisplayName property value. The display name of the user. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. The identifier of the user. Supports $filter (eq operator only).
func (m *SignIn) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. The UPN of the user. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserType gets the userType property value. Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
func (m *SignIn) GetUserType()(*SignInUserType) {
    return m.userType
}
// Serialize serializes information the current object
func (m *SignIn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetAppliedConditionalAccessPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliedConditionalAccessPolicies()))
        for i, v := range m.GetAppliedConditionalAccessPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliedConditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppliedEventListeners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliedEventListeners()))
        for i, v := range m.GetAppliedEventListeners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliedEventListeners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationAppDeviceDetails", m.GetAuthenticationAppDeviceDetails())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationAppPolicyEvaluationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationAppPolicyEvaluationDetails()))
        for i, v := range m.GetAuthenticationAppPolicyEvaluationDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationAppPolicyEvaluationDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationContextClassReferences() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationContextClassReferences()))
        for i, v := range m.GetAuthenticationContextClassReferences() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationContextClassReferences", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationDetails()))
        for i, v := range m.GetAuthenticationDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationMethodsUsed() != nil {
        err = writer.WriteCollectionOfStringValues("authenticationMethodsUsed", m.GetAuthenticationMethodsUsed())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationProcessingDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationProcessingDetails()))
        for i, v := range m.GetAuthenticationProcessingDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationProcessingDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationProtocol() != nil {
        cast := (*m.GetAuthenticationProtocol()).String()
        err = writer.WriteStringValue("authenticationProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authenticationRequirement", m.GetAuthenticationRequirement())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationRequirementPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationRequirementPolicies()))
        for i, v := range m.GetAuthenticationRequirementPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationRequirementPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("autonomousSystemNumber", m.GetAutonomousSystemNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureResourceId", m.GetAzureResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientAppUsed", m.GetClientAppUsed())
        if err != nil {
            return err
        }
    }
    if m.GetClientCredentialType() != nil {
        cast := (*m.GetClientCredentialType()).String()
        err = writer.WriteStringValue("clientCredentialType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessStatus() != nil {
        cast := (*m.GetConditionalAccessStatus()).String()
        err = writer.WriteStringValue("conditionalAccessStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCrossTenantAccessType() != nil {
        cast := (*m.GetCrossTenantAccessType()).String()
        err = writer.WriteStringValue("crossTenantAccessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceDetail", m.GetDeviceDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("federatedCredentialId", m.GetFederatedCredentialId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("flaggedForReview", m.GetFlaggedForReview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homeTenantId", m.GetHomeTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homeTenantName", m.GetHomeTenantName())
        if err != nil {
            return err
        }
    }
    if m.GetIncomingTokenType() != nil {
        cast := (*m.GetIncomingTokenType()).String()
        err = writer.WriteStringValue("incomingTokenType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddressFromResourceProvider", m.GetIpAddressFromResourceProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInteractive", m.GetIsInteractive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTenantRestricted", m.GetIsTenantRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mfaDetail", m.GetMfaDetail())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkLocationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkLocationDetails()))
        for i, v := range m.GetNetworkLocationDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("networkLocationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalRequestId", m.GetOriginalRequestId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("privateLinkDetails", m.GetPrivateLinkDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("processingTimeInMilliseconds", m.GetProcessingTimeInMilliseconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceDisplayName", m.GetResourceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceServicePrincipalId", m.GetResourceServicePrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceTenantId", m.GetResourceTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetRiskDetail() != nil {
        cast := (*m.GetRiskDetail()).String()
        err = writer.WriteStringValue("riskDetail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskEventTypes_v2() != nil {
        err = writer.WriteCollectionOfStringValues("riskEventTypes_v2", m.GetRiskEventTypes_v2())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelAggregated() != nil {
        cast := (*m.GetRiskLevelAggregated()).String()
        err = writer.WriteStringValue("riskLevelAggregated", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelDuringSignIn() != nil {
        cast := (*m.GetRiskLevelDuringSignIn()).String()
        err = writer.WriteStringValue("riskLevelDuringSignIn", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskState() != nil {
        cast := (*m.GetRiskState()).String()
        err = writer.WriteStringValue("riskState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalCredentialKeyId", m.GetServicePrincipalCredentialKeyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalCredentialThumbprint", m.GetServicePrincipalCredentialThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetSessionLifetimePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessionLifetimePolicies()))
        for i, v := range m.GetSessionLifetimePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sessionLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignInEventTypes() != nil {
        err = writer.WriteCollectionOfStringValues("signInEventTypes", m.GetSignInEventTypes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInIdentifier", m.GetSignInIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetSignInIdentifierType() != nil {
        cast := (*m.GetSignInIdentifierType()).String()
        err = writer.WriteStringValue("signInIdentifierType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenIssuerName", m.GetTokenIssuerName())
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuerType() != nil {
        cast := (*m.GetTokenIssuerType()).String()
        err = writer.WriteStringValue("tokenIssuerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uniqueTokenIdentifier", m.GetUniqueTokenIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userAgent", m.GetUserAgent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err = writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. The application name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppId sets the appId property value. The application identifier in Azure Active Directory. Supports $filter (eq operator only).
func (m *SignIn) SetAppId(value *string)() {
    m.appId = value
}
// SetAppliedConditionalAccessPolicies sets the appliedConditionalAccessPolicies property value. A list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicyable)() {
    m.appliedConditionalAccessPolicies = value
}
// SetAppliedEventListeners sets the appliedEventListeners property value. The appliedEventListeners property
func (m *SignIn) SetAppliedEventListeners(value []AppliedAuthenticationEventListenerable)() {
    m.appliedEventListeners = value
}
// SetAuthenticationAppDeviceDetails sets the authenticationAppDeviceDetails property value. Provides details about the app and device used during an Azure AD authentication step.
func (m *SignIn) SetAuthenticationAppDeviceDetails(value AuthenticationAppDeviceDetailsable)() {
    m.authenticationAppDeviceDetails = value
}
// SetAuthenticationAppPolicyEvaluationDetails sets the authenticationAppPolicyEvaluationDetails property value. Provides details of the Azure AD policies applied to a user and client authentication app during an authentication step.
func (m *SignIn) SetAuthenticationAppPolicyEvaluationDetails(value []AuthenticationAppPolicyDetailsable)() {
    m.authenticationAppPolicyEvaluationDetails = value
}
// SetAuthenticationContextClassReferences sets the authenticationContextClassReferences property value. Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
func (m *SignIn) SetAuthenticationContextClassReferences(value []AuthenticationContextable)() {
    m.authenticationContextClassReferences = value
}
// SetAuthenticationDetails sets the authenticationDetails property value. The result of the authentication attempt and additional details on the authentication method.
func (m *SignIn) SetAuthenticationDetails(value []AuthenticationDetailable)() {
    m.authenticationDetails = value
}
// SetAuthenticationMethodsUsed sets the authenticationMethodsUsed property value. The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA, or PHS.
func (m *SignIn) SetAuthenticationMethodsUsed(value []string)() {
    m.authenticationMethodsUsed = value
}
// SetAuthenticationProcessingDetails sets the authenticationProcessingDetails property value. Additional authentication processing details, such as the agent name in case of PTA/PHS or Server/farm name in case of federated authentication.
func (m *SignIn) SetAuthenticationProcessingDetails(value []KeyValueable)() {
    m.authenticationProcessingDetails = value
}
// SetAuthenticationProtocol sets the authenticationProtocol property value. Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc, wsFederation, saml20, deviceCode, unknownFutureValue. For authentications that use protocols other than the possible values listed, the protocol type is listed as none.
func (m *SignIn) SetAuthenticationProtocol(value *ProtocolType)() {
    m.authenticationProtocol = value
}
// SetAuthenticationRequirement sets the authenticationRequirement property value. This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetAuthenticationRequirement(value *string)() {
    m.authenticationRequirement = value
}
// SetAuthenticationRequirementPolicies sets the authenticationRequirementPolicies property value. Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security defaults.
func (m *SignIn) SetAuthenticationRequirementPolicies(value []AuthenticationRequirementPolicyable)() {
    m.authenticationRequirementPolicies = value
}
// SetAutonomousSystemNumber sets the autonomousSystemNumber property value. The Autonomous System Number (ASN) of the network used by the actor.
func (m *SignIn) SetAutonomousSystemNumber(value *int32)() {
    m.autonomousSystemNumber = value
}
// SetAzureResourceId sets the azureResourceId property value. Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
func (m *SignIn) SetAzureResourceId(value *string)() {
    m.azureResourceId = value
}
// SetClientAppUsed sets the clientAppUsed property value. The legacy client used for sign-in activity. For example: Browser, Exchange ActiveSync, Modern clients, IMAP, MAPI, SMTP, or POP. Supports $filter (eq operator only).
func (m *SignIn) SetClientAppUsed(value *string)() {
    m.clientAppUsed = value
}
// SetClientCredentialType sets the clientCredentialType property value. Describes the credential type that a user client or service principal provided to Azure AD to authenticate itself. You may wish to review clientCredentialType to track and eliminate less secure credential types or to watch for clients and service principals using anomalous credential types. The possible values are: none, clientSecret, clientAssertion, federatedIdentityCredential, managedIdentity, certificate, unknownFutureValue.
func (m *SignIn) SetClientCredentialType(value *ClientCredentialType)() {
    m.clientCredentialType = value
}
// SetConditionalAccessStatus sets the conditionalAccessStatus property value. The status of the conditional access policy triggered. Possible values: success, failure, notApplied, or unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetConditionalAccessStatus(value *ConditionalAccessStatus)() {
    m.conditionalAccessStatus = value
}
// SetCorrelationId sets the correlationId property value. The identifier that's sent from the client when sign-in is initiated. This is used for troubleshooting the corresponding sign-in activity when calling for support. Supports $filter (eq operator only).
func (m *SignIn) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the sign-in was initiated. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
func (m *SignIn) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCrossTenantAccessType sets the crossTenantAccessType property value. Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none, b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue. If the sign in did not cross tenant boundaries, the value is none.
func (m *SignIn) SetCrossTenantAccessType(value *SignInAccessType)() {
    m.crossTenantAccessType = value
}
// SetDeviceDetail sets the deviceDetail property value. The device information from where the sign-in occurred. Includes information such as deviceId, OS, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSystem properties.
func (m *SignIn) SetDeviceDetail(value DeviceDetailable)() {
    m.deviceDetail = value
}
// SetFederatedCredentialId sets the federatedCredentialId property value. Contains the identifier of an application's federated identity credential, if a federated identity credential was used to sign in.
func (m *SignIn) SetFederatedCredentialId(value *string)() {
    m.federatedCredentialId = value
}
// SetFlaggedForReview sets the flaggedForReview property value. During a failed sign in, a user may click a button in the Azure portal to mark the failed event for tenant admins. If a user clicked the button to flag the failed sign in, this value is true.
func (m *SignIn) SetFlaggedForReview(value *bool)() {
    m.flaggedForReview = value
}
// SetHomeTenantId sets the homeTenantId property value. The tenant identifier of the user initiating the sign in. Not applicable in Managed Identity or service principal sign ins.
func (m *SignIn) SetHomeTenantId(value *string)() {
    m.homeTenantId = value
}
// SetHomeTenantName sets the homeTenantName property value. For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *SignIn) SetHomeTenantName(value *string)() {
    m.homeTenantName = value
}
// SetIncomingTokenType sets the incomingTokenType property value. Indicates the token types that were presented to Azure AD to authenticate the actor in the sign in. The possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue, remoteDesktopToken.  NOTE Azure AD may have also used token types not listed in this Enum type to authenticate the actor. Do not infer the lack of a token if it is not one of the types listed. Also, please note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: remoteDesktopToken.
func (m *SignIn) SetIncomingTokenType(value *IncomingTokenType)() {
    m.incomingTokenType = value
}
// SetIpAddress sets the ipAddress property value. The IP address of the client from where the sign-in occurred. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetIpAddressFromResourceProvider sets the ipAddressFromResourceProvider property value. The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some policies. For example, when a user interacts with Exchange Online, the IP address Exchange receives from the user may be recorded here. This value is often null.
func (m *SignIn) SetIpAddressFromResourceProvider(value *string)() {
    m.ipAddressFromResourceProvider = value
}
// SetIsInteractive sets the isInteractive property value. Indicates whether a user sign in is interactive. In interactive sign in, the user provides an authentication factor to Azure AD. These factors include passwords, responses to MFA challenges, biometric factors, or QR codes that a user provides to Azure AD or an associated app. In non-interactive sign in, the user doesn't provide an authentication factor. Instead, the client app uses a token or code to authenticate or access a resource on behalf of a user. Non-interactive sign ins are commonly used for a client to sign in on a user's behalf in a process transparent to the user.
func (m *SignIn) SetIsInteractive(value *bool)() {
    m.isInteractive = value
}
// SetIsTenantRestricted sets the isTenantRestricted property value. Shows whether the sign in event was subject to an Azure AD tenant restriction policy.
func (m *SignIn) SetIsTenantRestricted(value *bool)() {
    m.isTenantRestricted = value
}
// SetLocation sets the location property value. The city, state, and 2 letter country code from where the sign-in occurred. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
func (m *SignIn) SetLocation(value SignInLocationable)() {
    m.location = value
}
// SetMfaDetail sets the mfaDetail property value. The mfaDetail property
func (m *SignIn) SetMfaDetail(value MfaDetailable)() {
    m.mfaDetail = value
}
// SetNetworkLocationDetails sets the networkLocationDetails property value. The network location details including the type of network used and its names.
func (m *SignIn) SetNetworkLocationDetails(value []NetworkLocationDetailable)() {
    m.networkLocationDetails = value
}
// SetOriginalRequestId sets the originalRequestId property value. The request identifier of the first request in the authentication sequence. Supports $filter (eq operator only).
func (m *SignIn) SetOriginalRequestId(value *string)() {
    m.originalRequestId = value
}
// SetPrivateLinkDetails sets the privateLinkDetails property value. Contains information about the Azure AD Private Link policy that is associated with the sign in event.
func (m *SignIn) SetPrivateLinkDetails(value PrivateLinkDetailsable)() {
    m.privateLinkDetails = value
}
// SetProcessingTimeInMilliseconds sets the processingTimeInMilliseconds property value. The request processing time in milliseconds in AD STS.
func (m *SignIn) SetProcessingTimeInMilliseconds(value *int32)() {
    m.processingTimeInMilliseconds = value
}
// SetResourceDisplayName sets the resourceDisplayName property value. The name of the resource that the user signed in to. Supports $filter (eq operator only).
func (m *SignIn) SetResourceDisplayName(value *string)() {
    m.resourceDisplayName = value
}
// SetResourceId sets the resourceId property value. The identifier of the resource that the user signed in to. Supports $filter (eq operator only).
func (m *SignIn) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetResourceServicePrincipalId sets the resourceServicePrincipalId property value. The identifier of the service principal representing the target resource in the sign-in event.
func (m *SignIn) SetResourceServicePrincipalId(value *string)() {
    m.resourceServicePrincipalId = value
}
// SetResourceTenantId sets the resourceTenantId property value. The tenant identifier of the resource referenced in the sign in.
func (m *SignIn) SetResourceTenantId(value *string)() {
    m.resourceTenantId = value
}
// SetRiskDetail sets the riskDetail property value. The reason behind a specific state of a risky user, sign-in, or a risk event. Possible values: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, or unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// SetRiskEventTypes_v2 sets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetRiskEventTypes_v2(value []string)() {
    m.riskEventTypes_v2 = value
}
// SetRiskLevelAggregated sets the riskLevelAggregated property value. The aggregated risk level. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) SetRiskLevelAggregated(value *RiskLevel)() {
    m.riskLevelAggregated = value
}
// SetRiskLevelDuringSignIn sets the riskLevelDuringSignIn property value. The risk level during sign-in. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers are returned hidden.
func (m *SignIn) SetRiskLevelDuringSignIn(value *RiskLevel)() {
    m.riskLevelDuringSignIn = value
}
// SetRiskState sets the riskState property value. The risk state of a risky user, sign-in, or a risk event. Possible values: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, or unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// SetServicePrincipalCredentialKeyId sets the servicePrincipalCredentialKeyId property value. The unique identifier of the key credential used by the service principal to authenticate.
func (m *SignIn) SetServicePrincipalCredentialKeyId(value *string)() {
    m.servicePrincipalCredentialKeyId = value
}
// SetServicePrincipalCredentialThumbprint sets the servicePrincipalCredentialThumbprint property value. The certificate thumbprint of the certificate used by the service principal to authenticate.
func (m *SignIn) SetServicePrincipalCredentialThumbprint(value *string)() {
    m.servicePrincipalCredentialThumbprint = value
}
// SetServicePrincipalId sets the servicePrincipalId property value. The application identifier used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
// SetServicePrincipalName sets the servicePrincipalName property value. The application name used for sign-in. This field is populated when you are signing in using an application. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// SetSessionLifetimePolicies sets the sessionLifetimePolicies property value. Any conditional access session management policies that were applied during the sign-in event.
func (m *SignIn) SetSessionLifetimePolicies(value []SessionLifetimePolicyable)() {
    m.sessionLifetimePolicies = value
}
// SetSignInEventTypes sets the signInEventTypes property value. Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed identity sign ins, the category is managedIdentity. For service principal sign ins, the category is servicePrincipal. Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue. Supports $filter (eq, ne).
func (m *SignIn) SetSignInEventTypes(value []string)() {
    m.signInEventTypes = value
}
// SetSignInIdentifier sets the signInIdentifier property value. The identification that the user provided to sign in. It may be the userPrincipalName but it's also populated when a user signs in using other identifiers.
func (m *SignIn) SetSignInIdentifier(value *string)() {
    m.signInIdentifier = value
}
// SetSignInIdentifierType sets the signInIdentifierType property value. The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode, onPremisesUserPrincipalName, unknownFutureValue.
func (m *SignIn) SetSignInIdentifierType(value *SignInIdentifierType)() {
    m.signInIdentifierType = value
}
// SetStatus sets the status property value. The sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
func (m *SignIn) SetStatus(value SignInStatusable)() {
    m.status = value
}
// SetTokenIssuerName sets the tokenIssuerName property value. The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq operator only).
func (m *SignIn) SetTokenIssuerName(value *string)() {
    m.tokenIssuerName = value
}
// SetTokenIssuerType sets the tokenIssuerType property value. The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue, AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter , NPSExtension.
func (m *SignIn) SetTokenIssuerType(value *TokenIssuerType)() {
    m.tokenIssuerType = value
}
// SetUniqueTokenIdentifier sets the uniqueTokenIdentifier property value. A unique base64 encoded request identifier used to track tokens issued by Azure AD as they are redeemed at resource providers.
func (m *SignIn) SetUniqueTokenIdentifier(value *string)() {
    m.uniqueTokenIdentifier = value
}
// SetUserAgent sets the userAgent property value. The user agent information related to sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserAgent(value *string)() {
    m.userAgent = value
}
// SetUserDisplayName sets the userDisplayName property value. The display name of the user. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. The identifier of the user. Supports $filter (eq operator only).
func (m *SignIn) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The UPN of the user. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserType sets the userType property value. Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest, unknownFutureValue.
func (m *SignIn) SetUserType(value *SignInUserType)() {
    m.userType = value
}
