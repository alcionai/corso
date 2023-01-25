package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleScopeTagable 
type RoleScopeTagable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]RoleScopeTagAutoAssignmentable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsBuiltIn()(*bool)
    SetAssignments(value []RoleScopeTagAutoAssignmentable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsBuiltIn(value *bool)()
}
