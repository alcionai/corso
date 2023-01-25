package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonNameable 
type PersonNameable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetFirst()(*string)
    GetInitials()(*string)
    GetLanguageTag()(*string)
    GetLast()(*string)
    GetMaiden()(*string)
    GetMiddle()(*string)
    GetNickname()(*string)
    GetPronunciation()(PersonNamePronounciationable)
    GetSuffix()(*string)
    GetTitle()(*string)
    SetDisplayName(value *string)()
    SetFirst(value *string)()
    SetInitials(value *string)()
    SetLanguageTag(value *string)()
    SetLast(value *string)()
    SetMaiden(value *string)()
    SetMiddle(value *string)()
    SetNickname(value *string)()
    SetPronunciation(value PersonNamePronounciationable)()
    SetSuffix(value *string)()
    SetTitle(value *string)()
}
