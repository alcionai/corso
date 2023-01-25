package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenIdConnectProviderable 
type OpenIdConnectProviderable interface {
    IdentityProviderable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClaimsMapping()(ClaimsMappingable)
    GetDomainHint()(*string)
    GetMetadataUrl()(*string)
    GetResponseMode()(*OpenIdConnectResponseMode)
    GetResponseType()(*OpenIdConnectResponseTypes)
    GetScope()(*string)
    SetClaimsMapping(value ClaimsMappingable)()
    SetDomainHint(value *string)()
    SetMetadataUrl(value *string)()
    SetResponseMode(value *OpenIdConnectResponseMode)()
    SetResponseType(value *OpenIdConnectResponseTypes)()
    SetScope(value *string)()
}
