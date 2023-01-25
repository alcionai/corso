package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppRelationshipable 
type MobileAppRelationshipable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTargetDisplayName()(*string)
    GetTargetDisplayVersion()(*string)
    GetTargetId()(*string)
    GetTargetPublisher()(*string)
    GetTargetType()(*MobileAppRelationshipType)
    SetTargetDisplayName(value *string)()
    SetTargetDisplayVersion(value *string)()
    SetTargetId(value *string)()
    SetTargetPublisher(value *string)()
    SetTargetType(value *MobileAppRelationshipType)()
}
