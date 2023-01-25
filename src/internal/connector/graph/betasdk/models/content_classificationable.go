package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentClassificationable 
type ContentClassificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfidence()(*int32)
    GetMatches()([]MatchLocationable)
    GetOdataType()(*string)
    GetSensitiveTypeId()(*string)
    GetUniqueCount()(*int32)
    SetConfidence(value *int32)()
    SetMatches(value []MatchLocationable)()
    SetOdataType(value *string)()
    SetSensitiveTypeId(value *string)()
    SetUniqueCount(value *int32)()
}
