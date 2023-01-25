package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365ActiveUserCountsable 
type Office365ActiveUserCountsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExchange()(*int64)
    GetOffice365()(*int64)
    GetOneDrive()(*int64)
    GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSharePoint()(*int64)
    GetSkypeForBusiness()(*int64)
    GetTeams()(*int64)
    GetYammer()(*int64)
    SetExchange(value *int64)()
    SetOffice365(value *int64)()
    SetOneDrive(value *int64)()
    SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSharePoint(value *int64)()
    SetSkypeForBusiness(value *int64)()
    SetTeams(value *int64)()
    SetYammer(value *int64)()
}
