package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectorGroupable 
type ConnectorGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplications()([]Applicationable)
    GetConnectorGroupType()(*ConnectorGroupType)
    GetIsDefault()(*bool)
    GetMembers()([]Connectorable)
    GetName()(*string)
    GetRegion()(*ConnectorGroupRegion)
    SetApplications(value []Applicationable)()
    SetConnectorGroupType(value *ConnectorGroupType)()
    SetIsDefault(value *bool)()
    SetMembers(value []Connectorable)()
    SetName(value *string)()
    SetRegion(value *ConnectorGroupRegion)()
}
