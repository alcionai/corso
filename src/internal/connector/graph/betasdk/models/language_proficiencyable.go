package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LanguageProficiencyable 
type LanguageProficiencyable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetProficiency()(*LanguageProficiencyLevel)
    GetReading()(*LanguageProficiencyLevel)
    GetSpoken()(*LanguageProficiencyLevel)
    GetTag()(*string)
    GetThumbnailUrl()(*string)
    GetWritten()(*LanguageProficiencyLevel)
    SetDisplayName(value *string)()
    SetProficiency(value *LanguageProficiencyLevel)()
    SetReading(value *LanguageProficiencyLevel)()
    SetSpoken(value *LanguageProficiencyLevel)()
    SetTag(value *string)()
    SetThumbnailUrl(value *string)()
    SetWritten(value *LanguageProficiencyLevel)()
}
