package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesAgentGroupable 
type OnPremisesAgentGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgents()([]OnPremisesAgentable)
    GetDisplayName()(*string)
    GetIsDefault()(*bool)
    GetPublishedResources()([]PublishedResourceable)
    GetPublishingType()(*OnPremisesPublishingType)
    SetAgents(value []OnPremisesAgentable)()
    SetDisplayName(value *string)()
    SetIsDefault(value *bool)()
    SetPublishedResources(value []PublishedResourceable)()
    SetPublishingType(value *OnPremisesPublishingType)()
}
