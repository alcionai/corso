package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishingProfileable 
type OnPremisesPublishingProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgentGroups()([]OnPremisesAgentGroupable)
    GetAgents()([]OnPremisesAgentable)
    GetConnectorGroups()([]ConnectorGroupable)
    GetConnectors()([]Connectorable)
    GetHybridAgentUpdaterConfiguration()(HybridAgentUpdaterConfigurationable)
    GetIsDefaultAccessEnabled()(*bool)
    GetIsEnabled()(*bool)
    GetPublishedResources()([]PublishedResourceable)
    SetAgentGroups(value []OnPremisesAgentGroupable)()
    SetAgents(value []OnPremisesAgentable)()
    SetConnectorGroups(value []ConnectorGroupable)()
    SetConnectors(value []Connectorable)()
    SetHybridAgentUpdaterConfiguration(value HybridAgentUpdaterConfigurationable)()
    SetIsDefaultAccessEnabled(value *bool)()
    SetIsEnabled(value *bool)()
    SetPublishedResources(value []PublishedResourceable)()
}
