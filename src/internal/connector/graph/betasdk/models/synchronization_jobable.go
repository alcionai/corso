package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobable 
type SynchronizationJobable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSchedule()(SynchronizationScheduleable)
    GetSchema()(SynchronizationSchemaable)
    GetStatus()(SynchronizationStatusable)
    GetSynchronizationJobSettings()([]KeyValuePairable)
    GetTemplateId()(*string)
    SetSchedule(value SynchronizationScheduleable)()
    SetSchema(value SynchronizationSchemaable)()
    SetStatus(value SynchronizationStatusable)()
    SetSynchronizationJobSettings(value []KeyValuePairable)()
    SetTemplateId(value *string)()
}
