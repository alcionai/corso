package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationLinkedObjectsable 
type SynchronizationLinkedObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManager()(SynchronizationJobSubjectable)
    GetMembers()([]SynchronizationJobSubjectable)
    GetOdataType()(*string)
    GetOwners()([]SynchronizationJobSubjectable)
    SetManager(value SynchronizationJobSubjectable)()
    SetMembers(value []SynchronizationJobSubjectable)()
    SetOdataType(value *string)()
    SetOwners(value []SynchronizationJobSubjectable)()
}
