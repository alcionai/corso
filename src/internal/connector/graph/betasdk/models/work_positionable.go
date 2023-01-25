package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkPositionable 
type WorkPositionable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]string)
    GetColleagues()([]RelatedPersonable)
    GetDetail()(PositionDetailable)
    GetIsCurrent()(*bool)
    GetManager()(RelatedPersonable)
    SetCategories(value []string)()
    SetColleagues(value []RelatedPersonable)()
    SetDetail(value PositionDetailable)()
    SetIsCurrent(value *bool)()
    SetManager(value RelatedPersonable)()
}
