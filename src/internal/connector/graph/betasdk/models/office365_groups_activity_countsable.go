package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityCountsable 
type Office365GroupsActivityCountsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExchangeEmailsReceived()(*int64)
    GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTeamsChannelMessages()(*int64)
    GetTeamsMeetingsOrganized()(*int64)
    GetYammerMessagesLiked()(*int64)
    GetYammerMessagesPosted()(*int64)
    GetYammerMessagesRead()(*int64)
    SetExchangeEmailsReceived(value *int64)()
    SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTeamsChannelMessages(value *int64)()
    SetTeamsMeetingsOrganized(value *int64)()
    SetYammerMessagesLiked(value *int64)()
    SetYammerMessagesPosted(value *int64)()
    SetYammerMessagesRead(value *int64)()
}
