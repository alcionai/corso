package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectDefinitionable 
type ObjectDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributes()([]AttributeDefinitionable)
    GetMetadata()([]MetadataEntryable)
    GetName()(*string)
    GetOdataType()(*string)
    GetSupportedApis()([]string)
    SetAttributes(value []AttributeDefinitionable)()
    SetMetadata(value []MetadataEntryable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetSupportedApis(value []string)()
}
