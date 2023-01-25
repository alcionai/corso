package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityDetail 
type Office365GroupsActivityDetail struct {
    Entity
    // The storage used of the group mailbox.
    exchangeMailboxStorageUsedInBytes *int64
    // The number of items in the group mailbox.
    exchangeMailboxTotalItemCount *int64
    // The number of email that the group mailbox received.
    exchangeReceivedEmailCount *int64
    // The group external member count.
    externalMemberCount *int64
    // The display name of the group.
    groupDisplayName *string
    // The group id.
    groupId *string
    // The group type. Possible values are: Public or Private.
    groupType *string
    // Whether this user has been deleted or soft deleted.
    isDeleted *bool
    // The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
    lastActivityDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The group member count.
    memberCount *int64
    // The group owner principal name.
    ownerPrincipalName *string
    // The number of days the report covers.
    reportPeriod *string
    // The latest date of the content.
    reportRefreshDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of active files in SharePoint Group site.
    sharePointActiveFileCount *int64
    // The storage used by SharePoint Group site.
    sharePointSiteStorageUsedInBytes *int64
    // The total number of files in SharePoint Group site.
    sharePointTotalFileCount *int64
    // The teamsChannelMessagesCount property
    teamsChannelMessagesCount *int64
    // The teamsMeetingsOrganizedCount property
    teamsMeetingsOrganizedCount *int64
    // The number of messages liked in Yammer groups.
    yammerLikedMessageCount *int64
    // The number of messages posted to Yammer groups.
    yammerPostedMessageCount *int64
    // The number of messages read in Yammer groups.
    yammerReadMessageCount *int64
}
// NewOffice365GroupsActivityDetail instantiates a new Office365GroupsActivityDetail and sets the default values.
func NewOffice365GroupsActivityDetail()(*Office365GroupsActivityDetail) {
    m := &Office365GroupsActivityDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365GroupsActivityDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365GroupsActivityDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365GroupsActivityDetail(), nil
}
// GetExchangeMailboxStorageUsedInBytes gets the exchangeMailboxStorageUsedInBytes property value. The storage used of the group mailbox.
func (m *Office365GroupsActivityDetail) GetExchangeMailboxStorageUsedInBytes()(*int64) {
    return m.exchangeMailboxStorageUsedInBytes
}
// GetExchangeMailboxTotalItemCount gets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
func (m *Office365GroupsActivityDetail) GetExchangeMailboxTotalItemCount()(*int64) {
    return m.exchangeMailboxTotalItemCount
}
// GetExchangeReceivedEmailCount gets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
func (m *Office365GroupsActivityDetail) GetExchangeReceivedEmailCount()(*int64) {
    return m.exchangeReceivedEmailCount
}
// GetExternalMemberCount gets the externalMemberCount property value. The group external member count.
func (m *Office365GroupsActivityDetail) GetExternalMemberCount()(*int64) {
    return m.externalMemberCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeMailboxStorageUsedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeMailboxStorageUsedInBytes(val)
        }
        return nil
    }
    res["exchangeMailboxTotalItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeMailboxTotalItemCount(val)
        }
        return nil
    }
    res["exchangeReceivedEmailCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeReceivedEmailCount(val)
        }
        return nil
    }
    res["externalMemberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalMemberCount(val)
        }
        return nil
    }
    res["groupDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDisplayName(val)
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["groupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupType(val)
        }
        return nil
    }
    res["isDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["lastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["memberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberCount(val)
        }
        return nil
    }
    res["ownerPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerPrincipalName(val)
        }
        return nil
    }
    res["reportPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePointActiveFileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointActiveFileCount(val)
        }
        return nil
    }
    res["sharePointSiteStorageUsedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointSiteStorageUsedInBytes(val)
        }
        return nil
    }
    res["sharePointTotalFileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointTotalFileCount(val)
        }
        return nil
    }
    res["teamsChannelMessagesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsChannelMessagesCount(val)
        }
        return nil
    }
    res["teamsMeetingsOrganizedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsMeetingsOrganizedCount(val)
        }
        return nil
    }
    res["yammerLikedMessageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLikedMessageCount(val)
        }
        return nil
    }
    res["yammerPostedMessageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerPostedMessageCount(val)
        }
        return nil
    }
    res["yammerReadMessageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerReadMessageCount(val)
        }
        return nil
    }
    return res
}
// GetGroupDisplayName gets the groupDisplayName property value. The display name of the group.
func (m *Office365GroupsActivityDetail) GetGroupDisplayName()(*string) {
    return m.groupDisplayName
}
// GetGroupId gets the groupId property value. The group id.
func (m *Office365GroupsActivityDetail) GetGroupId()(*string) {
    return m.groupId
}
// GetGroupType gets the groupType property value. The group type. Possible values are: Public or Private.
func (m *Office365GroupsActivityDetail) GetGroupType()(*string) {
    return m.groupType
}
// GetIsDeleted gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365GroupsActivityDetail) GetIsDeleted()(*bool) {
    return m.isDeleted
}
// GetLastActivityDate gets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityDetail) GetLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.lastActivityDate
}
// GetMemberCount gets the memberCount property value. The group member count.
func (m *Office365GroupsActivityDetail) GetMemberCount()(*int64) {
    return m.memberCount
}
// GetOwnerPrincipalName gets the ownerPrincipalName property value. The group owner principal name.
func (m *Office365GroupsActivityDetail) GetOwnerPrincipalName()(*string) {
    return m.ownerPrincipalName
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityDetail) GetReportPeriod()(*string) {
    return m.reportPeriod
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityDetail) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportRefreshDate
}
// GetSharePointActiveFileCount gets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointActiveFileCount()(*int64) {
    return m.sharePointActiveFileCount
}
// GetSharePointSiteStorageUsedInBytes gets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointSiteStorageUsedInBytes()(*int64) {
    return m.sharePointSiteStorageUsedInBytes
}
// GetSharePointTotalFileCount gets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointTotalFileCount()(*int64) {
    return m.sharePointTotalFileCount
}
// GetTeamsChannelMessagesCount gets the teamsChannelMessagesCount property value. The teamsChannelMessagesCount property
func (m *Office365GroupsActivityDetail) GetTeamsChannelMessagesCount()(*int64) {
    return m.teamsChannelMessagesCount
}
// GetTeamsMeetingsOrganizedCount gets the teamsMeetingsOrganizedCount property value. The teamsMeetingsOrganizedCount property
func (m *Office365GroupsActivityDetail) GetTeamsMeetingsOrganizedCount()(*int64) {
    return m.teamsMeetingsOrganizedCount
}
// GetYammerLikedMessageCount gets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerLikedMessageCount()(*int64) {
    return m.yammerLikedMessageCount
}
// GetYammerPostedMessageCount gets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerPostedMessageCount()(*int64) {
    return m.yammerPostedMessageCount
}
// GetYammerReadMessageCount gets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerReadMessageCount()(*int64) {
    return m.yammerReadMessageCount
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("exchangeMailboxStorageUsedInBytes", m.GetExchangeMailboxStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("exchangeMailboxTotalItemCount", m.GetExchangeMailboxTotalItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("exchangeReceivedEmailCount", m.GetExchangeReceivedEmailCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("externalMemberCount", m.GetExternalMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupDisplayName", m.GetGroupDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupType", m.GetGroupType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("memberCount", m.GetMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerPrincipalName", m.GetOwnerPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharePointActiveFileCount", m.GetSharePointActiveFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharePointSiteStorageUsedInBytes", m.GetSharePointSiteStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharePointTotalFileCount", m.GetSharePointTotalFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamsChannelMessagesCount", m.GetTeamsChannelMessagesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamsMeetingsOrganizedCount", m.GetTeamsMeetingsOrganizedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerLikedMessageCount", m.GetYammerLikedMessageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerPostedMessageCount", m.GetYammerPostedMessageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerReadMessageCount", m.GetYammerReadMessageCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExchangeMailboxStorageUsedInBytes sets the exchangeMailboxStorageUsedInBytes property value. The storage used of the group mailbox.
func (m *Office365GroupsActivityDetail) SetExchangeMailboxStorageUsedInBytes(value *int64)() {
    m.exchangeMailboxStorageUsedInBytes = value
}
// SetExchangeMailboxTotalItemCount sets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
func (m *Office365GroupsActivityDetail) SetExchangeMailboxTotalItemCount(value *int64)() {
    m.exchangeMailboxTotalItemCount = value
}
// SetExchangeReceivedEmailCount sets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
func (m *Office365GroupsActivityDetail) SetExchangeReceivedEmailCount(value *int64)() {
    m.exchangeReceivedEmailCount = value
}
// SetExternalMemberCount sets the externalMemberCount property value. The group external member count.
func (m *Office365GroupsActivityDetail) SetExternalMemberCount(value *int64)() {
    m.externalMemberCount = value
}
// SetGroupDisplayName sets the groupDisplayName property value. The display name of the group.
func (m *Office365GroupsActivityDetail) SetGroupDisplayName(value *string)() {
    m.groupDisplayName = value
}
// SetGroupId sets the groupId property value. The group id.
func (m *Office365GroupsActivityDetail) SetGroupId(value *string)() {
    m.groupId = value
}
// SetGroupType sets the groupType property value. The group type. Possible values are: Public or Private.
func (m *Office365GroupsActivityDetail) SetGroupType(value *string)() {
    m.groupType = value
}
// SetIsDeleted sets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365GroupsActivityDetail) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetLastActivityDate sets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityDetail) SetLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.lastActivityDate = value
}
// SetMemberCount sets the memberCount property value. The group member count.
func (m *Office365GroupsActivityDetail) SetMemberCount(value *int64)() {
    m.memberCount = value
}
// SetOwnerPrincipalName sets the ownerPrincipalName property value. The group owner principal name.
func (m *Office365GroupsActivityDetail) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityDetail) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityDetail) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportRefreshDate = value
}
// SetSharePointActiveFileCount sets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointActiveFileCount(value *int64)() {
    m.sharePointActiveFileCount = value
}
// SetSharePointSiteStorageUsedInBytes sets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointSiteStorageUsedInBytes(value *int64)() {
    m.sharePointSiteStorageUsedInBytes = value
}
// SetSharePointTotalFileCount sets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointTotalFileCount(value *int64)() {
    m.sharePointTotalFileCount = value
}
// SetTeamsChannelMessagesCount sets the teamsChannelMessagesCount property value. The teamsChannelMessagesCount property
func (m *Office365GroupsActivityDetail) SetTeamsChannelMessagesCount(value *int64)() {
    m.teamsChannelMessagesCount = value
}
// SetTeamsMeetingsOrganizedCount sets the teamsMeetingsOrganizedCount property value. The teamsMeetingsOrganizedCount property
func (m *Office365GroupsActivityDetail) SetTeamsMeetingsOrganizedCount(value *int64)() {
    m.teamsMeetingsOrganizedCount = value
}
// SetYammerLikedMessageCount sets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerLikedMessageCount(value *int64)() {
    m.yammerLikedMessageCount = value
}
// SetYammerPostedMessageCount sets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerPostedMessageCount(value *int64)() {
    m.yammerPostedMessageCount = value
}
// SetYammerReadMessageCount sets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerReadMessageCount(value *int64)() {
    m.yammerReadMessageCount = value
}
