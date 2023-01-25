package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleSummaryable 
type PrivilegedRoleSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetElevatedCount()(*int32)
    GetManagedCount()(*int32)
    GetMfaEnabled()(*bool)
    GetStatus()(*RoleSummaryStatus)
    GetUsersCount()(*int32)
    SetElevatedCount(value *int32)()
    SetManagedCount(value *int32)()
    SetMfaEnabled(value *bool)()
    SetStatus(value *RoleSummaryStatus)()
    SetUsersCount(value *int32)()
}
