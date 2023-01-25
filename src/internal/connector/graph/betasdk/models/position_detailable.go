package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PositionDetailable 
type PositionDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompany()(CompanyDetailable)
    GetDescription()(*string)
    GetEndMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetJobTitle()(*string)
    GetOdataType()(*string)
    GetRole()(*string)
    GetStartMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSummary()(*string)
    SetCompany(value CompanyDetailable)()
    SetDescription(value *string)()
    SetEndMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetJobTitle(value *string)()
    SetOdataType(value *string)()
    SetRole(value *string)()
    SetStartMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSummary(value *string)()
}
