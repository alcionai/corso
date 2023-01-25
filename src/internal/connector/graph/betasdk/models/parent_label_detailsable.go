package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParentLabelDetailsable 
type ParentLabelDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetDescription()(*string)
    GetId()(*string)
    GetIsActive()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetParent()(ParentLabelDetailsable)
    GetSensitivity()(*int32)
    GetTooltip()(*string)
    SetColor(value *string)()
    SetDescription(value *string)()
    SetId(value *string)()
    SetIsActive(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetParent(value ParentLabelDetailsable)()
    SetSensitivity(value *int32)()
    SetTooltip(value *string)()
}
