package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenIdConnectProvider 
type OpenIdConnectProvider struct {
    IdentityProvider
    // After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. It is a required property.
    claimsMapping ClaimsMappingable
    // The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
    domainHint *string
    // The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration . For the OpenID Connect identity provider you are looking to add, you will need to provide the metadata URL. It is a required property and is read only after creation.
    metadataUrl *string
    // The responseMode property
    responseMode *OpenIdConnectResponseMode
    // The responseType property
    responseType *OpenIdConnectResponseTypes
    // Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended separated by space. For more details about the scope limitations see RFC6749 Section 3.3. It is a required property.
    scope *string
}
// NewOpenIdConnectProvider instantiates a new OpenIdConnectProvider and sets the default values.
func NewOpenIdConnectProvider()(*OpenIdConnectProvider) {
    m := &OpenIdConnectProvider{
        IdentityProvider: *NewIdentityProvider(),
    }
    return m
}
// CreateOpenIdConnectProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenIdConnectProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenIdConnectProvider(), nil
}
// GetClaimsMapping gets the claimsMapping property value. After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. It is a required property.
func (m *OpenIdConnectProvider) GetClaimsMapping()(ClaimsMappingable) {
    return m.claimsMapping
}
// GetDomainHint gets the domainHint property value. The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
func (m *OpenIdConnectProvider) GetDomainHint()(*string) {
    return m.domainHint
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenIdConnectProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityProvider.GetFieldDeserializers()
    res["claimsMapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClaimsMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaimsMapping(val.(ClaimsMappingable))
        }
        return nil
    }
    res["domainHint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainHint(val)
        }
        return nil
    }
    res["metadataUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataUrl(val)
        }
        return nil
    }
    res["responseMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOpenIdConnectResponseMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseMode(val.(*OpenIdConnectResponseMode))
        }
        return nil
    }
    res["responseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOpenIdConnectResponseTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseType(val.(*OpenIdConnectResponseTypes))
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    return res
}
// GetMetadataUrl gets the metadataUrl property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration . For the OpenID Connect identity provider you are looking to add, you will need to provide the metadata URL. It is a required property and is read only after creation.
func (m *OpenIdConnectProvider) GetMetadataUrl()(*string) {
    return m.metadataUrl
}
// GetResponseMode gets the responseMode property value. The responseMode property
func (m *OpenIdConnectProvider) GetResponseMode()(*OpenIdConnectResponseMode) {
    return m.responseMode
}
// GetResponseType gets the responseType property value. The responseType property
func (m *OpenIdConnectProvider) GetResponseType()(*OpenIdConnectResponseTypes) {
    return m.responseType
}
// GetScope gets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended separated by space. For more details about the scope limitations see RFC6749 Section 3.3. It is a required property.
func (m *OpenIdConnectProvider) GetScope()(*string) {
    return m.scope
}
// Serialize serializes information the current object
func (m *OpenIdConnectProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityProvider.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("claimsMapping", m.GetClaimsMapping())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainHint", m.GetDomainHint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadataUrl", m.GetMetadataUrl())
        if err != nil {
            return err
        }
    }
    if m.GetResponseMode() != nil {
        cast := (*m.GetResponseMode()).String()
        err = writer.WriteStringValue("responseMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetResponseType() != nil {
        cast := (*m.GetResponseType()).String()
        err = writer.WriteStringValue("responseType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClaimsMapping sets the claimsMapping property value. After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. It is a required property.
func (m *OpenIdConnectProvider) SetClaimsMapping(value ClaimsMappingable)() {
    m.claimsMapping = value
}
// SetDomainHint sets the domainHint property value. The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
func (m *OpenIdConnectProvider) SetDomainHint(value *string)() {
    m.domainHint = value
}
// SetMetadataUrl sets the metadataUrl property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration . For the OpenID Connect identity provider you are looking to add, you will need to provide the metadata URL. It is a required property and is read only after creation.
func (m *OpenIdConnectProvider) SetMetadataUrl(value *string)() {
    m.metadataUrl = value
}
// SetResponseMode sets the responseMode property value. The responseMode property
func (m *OpenIdConnectProvider) SetResponseMode(value *OpenIdConnectResponseMode)() {
    m.responseMode = value
}
// SetResponseType sets the responseType property value. The responseType property
func (m *OpenIdConnectProvider) SetResponseType(value *OpenIdConnectResponseTypes)() {
    m.responseType = value
}
// SetScope sets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended separated by space. For more details about the scope limitations see RFC6749 Section 3.3. It is a required property.
func (m *OpenIdConnectProvider) SetScope(value *string)() {
    m.scope = value
}
