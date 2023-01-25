package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityStorage 
type Office365GroupsActivityStorage struct {
    Entity
    // The storage used in group mailbox.
    mailboxStorageUsedInBytes *int64
    // The snapshot date for Exchange and SharePoint used storage.
    reportDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of days the report covers.
    reportPeriod *string
    // The latest date of the content.
    reportRefreshDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The storage used in SharePoint document library.
    siteStorageUsedInBytes *int64
}
// NewOffice365GroupsActivityStorage instantiates a new Office365GroupsActivityStorage and sets the default values.
func NewOffice365GroupsActivityStorage()(*Office365GroupsActivityStorage) {
    m := &Office365GroupsActivityStorage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365GroupsActivityStorageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365GroupsActivityStorageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365GroupsActivityStorage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityStorage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["mailboxStorageUsedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxStorageUsedInBytes(val)
        }
        return nil
    }
    res["reportDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
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
    res["siteStorageUsedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteStorageUsedInBytes(val)
        }
        return nil
    }
    return res
}
// GetMailboxStorageUsedInBytes gets the mailboxStorageUsedInBytes property value. The storage used in group mailbox.
func (m *Office365GroupsActivityStorage) GetMailboxStorageUsedInBytes()(*int64) {
    return m.mailboxStorageUsedInBytes
}
// GetReportDate gets the reportDate property value. The snapshot date for Exchange and SharePoint used storage.
func (m *Office365GroupsActivityStorage) GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportDate
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityStorage) GetReportPeriod()(*string) {
    return m.reportPeriod
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityStorage) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportRefreshDate
}
// GetSiteStorageUsedInBytes gets the siteStorageUsedInBytes property value. The storage used in SharePoint document library.
func (m *Office365GroupsActivityStorage) GetSiteStorageUsedInBytes()(*int64) {
    return m.siteStorageUsedInBytes
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityStorage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("mailboxStorageUsedInBytes", m.GetMailboxStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportDate", m.GetReportDate())
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
        err = writer.WriteInt64Value("siteStorageUsedInBytes", m.GetSiteStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMailboxStorageUsedInBytes sets the mailboxStorageUsedInBytes property value. The storage used in group mailbox.
func (m *Office365GroupsActivityStorage) SetMailboxStorageUsedInBytes(value *int64)() {
    m.mailboxStorageUsedInBytes = value
}
// SetReportDate sets the reportDate property value. The snapshot date for Exchange and SharePoint used storage.
func (m *Office365GroupsActivityStorage) SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityStorage) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityStorage) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportRefreshDate = value
}
// SetSiteStorageUsedInBytes sets the siteStorageUsedInBytes property value. The storage used in SharePoint document library.
func (m *Office365GroupsActivityStorage) SetSiteStorageUsedInBytes(value *int64)() {
    m.siteStorageUsedInBytes = value
}
