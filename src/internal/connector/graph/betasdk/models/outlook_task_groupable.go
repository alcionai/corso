package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTaskGroupable 
type OutlookTaskGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeKey()(*string)
    GetGroupKey()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIsDefaultGroup()(*bool)
    GetName()(*string)
    GetTaskFolders()([]OutlookTaskFolderable)
    SetChangeKey(value *string)()
    SetGroupKey(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIsDefaultGroup(value *bool)()
    SetName(value *string)()
    SetTaskFolders(value []OutlookTaskFolderable)()
}
