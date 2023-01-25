package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsageRightable 
type UsageRightable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCatalogId()(*string)
    GetServiceIdentifier()(*string)
    GetState()(*UsageRightState)
    SetCatalogId(value *string)()
    SetServiceIdentifier(value *string)()
    SetState(value *UsageRightState)()
}
