package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchHitable 
type SearchHitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Get_id()(*string)
    Get_score()(*int32)
    Get_source()(Entityable)
    Get_summary()(*string)
    GetContentSource()(*string)
    GetHitId()(*string)
    GetIsCollapsed()(*bool)
    GetOdataType()(*string)
    GetRank()(*int32)
    GetResource()(Entityable)
    GetResultTemplateId()(*string)
    GetSummary()(*string)
    Set_id(value *string)()
    Set_score(value *int32)()
    Set_source(value Entityable)()
    Set_summary(value *string)()
    SetContentSource(value *string)()
    SetHitId(value *string)()
    SetIsCollapsed(value *bool)()
    SetOdataType(value *string)()
    SetRank(value *int32)()
    SetResource(value Entityable)()
    SetResultTemplateId(value *string)()
    SetSummary(value *string)()
}
