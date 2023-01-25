package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationalActivityDetailable 
type EducationalActivityDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAbbreviation()(*string)
    GetActivities()([]string)
    GetAwards()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFieldsOfStudy()([]string)
    GetGrade()(*string)
    GetNotes()(*string)
    GetOdataType()(*string)
    GetWebUrl()(*string)
    SetAbbreviation(value *string)()
    SetActivities(value []string)()
    SetAwards(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFieldsOfStudy(value []string)()
    SetGrade(value *string)()
    SetNotes(value *string)()
    SetOdataType(value *string)()
    SetWebUrl(value *string)()
}
