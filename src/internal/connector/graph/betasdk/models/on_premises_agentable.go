package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesAgentable 
type OnPremisesAgentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgentGroups()([]OnPremisesAgentGroupable)
    GetExternalIp()(*string)
    GetMachineName()(*string)
    GetStatus()(*AgentStatus)
    GetSupportedPublishingTypes()([]OnPremisesPublishingType)
    SetAgentGroups(value []OnPremisesAgentGroupable)()
    SetExternalIp(value *string)()
    SetMachineName(value *string)()
    SetStatus(value *AgentStatus)()
    SetSupportedPublishingTypes(value []OnPremisesPublishingType)()
}
