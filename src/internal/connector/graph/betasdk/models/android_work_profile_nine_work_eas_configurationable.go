package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileNineWorkEasConfigurationable 
type AndroidWorkProfileNineWorkEasConfigurationable interface {
    AndroidWorkProfileEasEmailProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSyncCalendar()(*bool)
    GetSyncContacts()(*bool)
    GetSyncTasks()(*bool)
    SetSyncCalendar(value *bool)()
    SetSyncContacts(value *bool)()
    SetSyncTasks(value *bool)()
}
