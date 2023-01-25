package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleSettingsable 
type PrivilegedRoleSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalOnElevation()(*bool)
    GetApproverIds()([]string)
    GetElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetIsMfaOnElevationConfigurable()(*bool)
    GetLastGlobalAdmin()(*bool)
    GetMaxElavationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMfaOnElevation()(*bool)
    GetMinElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetNotificationToUserOnElevation()(*bool)
    GetTicketingInfoOnElevation()(*bool)
    SetApprovalOnElevation(value *bool)()
    SetApproverIds(value []string)()
    SetElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetIsMfaOnElevationConfigurable(value *bool)()
    SetLastGlobalAdmin(value *bool)()
    SetMaxElavationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMfaOnElevation(value *bool)()
    SetMinElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetNotificationToUserOnElevation(value *bool)()
    SetTicketingInfoOnElevation(value *bool)()
}
