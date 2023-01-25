package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingable 
type AttributeMappingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetExportMissingReferences()(*bool)
    GetFlowBehavior()(*AttributeFlowBehavior)
    GetFlowType()(*AttributeFlowType)
    GetMatchingPriority()(*int32)
    GetOdataType()(*string)
    GetSource()(AttributeMappingSourceable)
    GetTargetAttributeName()(*string)
    SetDefaultValue(value *string)()
    SetExportMissingReferences(value *bool)()
    SetFlowBehavior(value *AttributeFlowBehavior)()
    SetFlowType(value *AttributeFlowType)()
    SetMatchingPriority(value *int32)()
    SetOdataType(value *string)()
    SetSource(value AttributeMappingSourceable)()
    SetTargetAttributeName(value *string)()
}
