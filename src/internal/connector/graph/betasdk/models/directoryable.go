package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Directoryable 
type Directoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnits()([]AdministrativeUnitable)
    GetAttributeSets()([]AttributeSetable)
    GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable)
    GetDeletedItems()([]DirectoryObjectable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetFederationConfigurations()([]IdentityProviderBaseable)
    GetImpactedResources()([]ImpactedResourceable)
    GetInboundSharedUserProfiles()([]InboundSharedUserProfileable)
    GetOnPremisesSynchronization()([]OnPremisesDirectorySynchronizationable)
    GetOutboundSharedUserProfiles()([]OutboundSharedUserProfileable)
    GetRecommendations()([]Recommendationable)
    GetSharedEmailDomains()([]SharedEmailDomainable)
    SetAdministrativeUnits(value []AdministrativeUnitable)()
    SetAttributeSets(value []AttributeSetable)()
    SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)()
    SetDeletedItems(value []DirectoryObjectable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetFederationConfigurations(value []IdentityProviderBaseable)()
    SetImpactedResources(value []ImpactedResourceable)()
    SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)()
    SetOnPremisesSynchronization(value []OnPremisesDirectorySynchronizationable)()
    SetOutboundSharedUserProfiles(value []OutboundSharedUserProfileable)()
    SetRecommendations(value []Recommendationable)()
    SetSharedEmailDomains(value []SharedEmailDomainable)()
}
