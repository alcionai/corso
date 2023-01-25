package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeDefinitionable 
type AttributeDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnchor()(*bool)
    GetApiExpressions()([]StringKeyStringValuePairable)
    GetCaseExact()(*bool)
    GetDefaultValue()(*string)
    GetFlowNullValues()(*bool)
    GetMetadata()([]MetadataEntryable)
    GetMultivalued()(*bool)
    GetMutability()(*Mutability)
    GetName()(*string)
    GetOdataType()(*string)
    GetReferencedObjects()([]ReferencedObjectable)
    GetRequired()(*bool)
    GetType()(*AttributeType)
    SetAnchor(value *bool)()
    SetApiExpressions(value []StringKeyStringValuePairable)()
    SetCaseExact(value *bool)()
    SetDefaultValue(value *string)()
    SetFlowNullValues(value *bool)()
    SetMetadata(value []MetadataEntryable)()
    SetMultivalued(value *bool)()
    SetMutability(value *Mutability)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetReferencedObjects(value []ReferencedObjectable)()
    SetRequired(value *bool)()
    SetType(value *AttributeType)()
}
