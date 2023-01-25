package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AnswerKeywordable 
type AnswerKeywordable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeywords()([]string)
    GetMatchSimilarKeywords()(*bool)
    GetOdataType()(*string)
    GetReservedKeywords()([]string)
    SetKeywords(value []string)()
    SetMatchSimilarKeywords(value *bool)()
    SetOdataType(value *string)()
    SetReservedKeywords(value []string)()
}
