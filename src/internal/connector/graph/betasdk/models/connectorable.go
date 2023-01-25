package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Connectorable 
type Connectorable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalIp()(*string)
    GetMachineName()(*string)
    GetMemberOf()([]ConnectorGroupable)
    GetStatus()(*ConnectorStatus)
    SetExternalIp(value *string)()
    SetMachineName(value *string)()
    SetMemberOf(value []ConnectorGroupable)()
    SetStatus(value *ConnectorStatus)()
}
