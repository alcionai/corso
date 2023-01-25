package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityDetailable 
type Office365GroupsActivityDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExchangeMailboxStorageUsedInBytes()(*int64)
    GetExchangeMailboxTotalItemCount()(*int64)
    GetExchangeReceivedEmailCount()(*int64)
    GetExternalMemberCount()(*int64)
    GetGroupDisplayName()(*string)
    GetGroupId()(*string)
    GetGroupType()(*string)
    GetIsDeleted()(*bool)
    GetLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetMemberCount()(*int64)
    GetOwnerPrincipalName()(*string)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSharePointActiveFileCount()(*int64)
    GetSharePointSiteStorageUsedInBytes()(*int64)
    GetSharePointTotalFileCount()(*int64)
    GetTeamsChannelMessagesCount()(*int64)
    GetTeamsMeetingsOrganizedCount()(*int64)
    GetYammerLikedMessageCount()(*int64)
    GetYammerPostedMessageCount()(*int64)
    GetYammerReadMessageCount()(*int64)
    SetExchangeMailboxStorageUsedInBytes(value *int64)()
    SetExchangeMailboxTotalItemCount(value *int64)()
    SetExchangeReceivedEmailCount(value *int64)()
    SetExternalMemberCount(value *int64)()
    SetGroupDisplayName(value *string)()
    SetGroupId(value *string)()
    SetGroupType(value *string)()
    SetIsDeleted(value *bool)()
    SetLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetMemberCount(value *int64)()
    SetOwnerPrincipalName(value *string)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSharePointActiveFileCount(value *int64)()
    SetSharePointSiteStorageUsedInBytes(value *int64)()
    SetSharePointTotalFileCount(value *int64)()
    SetTeamsChannelMessagesCount(value *int64)()
    SetTeamsMeetingsOrganizedCount(value *int64)()
    SetYammerLikedMessageCount(value *int64)()
    SetYammerPostedMessageCount(value *int64)()
    SetYammerReadMessageCount(value *int64)()
}
