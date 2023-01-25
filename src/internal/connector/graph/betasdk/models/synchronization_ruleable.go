package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationRuleable 
type SynchronizationRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEditable()(*bool)
    GetId()(*string)
    GetMetadata()([]StringKeyStringValuePairable)
    GetName()(*string)
    GetObjectMappings()([]ObjectMappingable)
    GetOdataType()(*string)
    GetPriority()(*int32)
    GetSourceDirectoryName()(*string)
    GetTargetDirectoryName()(*string)
    SetEditable(value *bool)()
    SetId(value *string)()
    SetMetadata(value []StringKeyStringValuePairable)()
    SetName(value *string)()
    SetObjectMappings(value []ObjectMappingable)()
    SetOdataType(value *string)()
    SetPriority(value *int32)()
    SetSourceDirectoryName(value *string)()
    SetTargetDirectoryName(value *string)()
}
