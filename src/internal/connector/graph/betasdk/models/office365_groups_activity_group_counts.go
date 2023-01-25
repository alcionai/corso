package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityGroupCounts 
type Office365GroupsActivityGroupCounts struct {
    Entity
    // The number of active groups. A group is considered active if any of the following occurred: group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
    active *int64
    // The date on which a number of groups were active.
    reportDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of days the report covers.
    reportPeriod *string
    // The latest date of the content.
    reportRefreshDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The total number of groups.
    total *int64
}
// NewOffice365GroupsActivityGroupCounts instantiates a new Office365GroupsActivityGroupCounts and sets the default values.
func NewOffice365GroupsActivityGroupCounts()(*Office365GroupsActivityGroupCounts) {
    m := &Office365GroupsActivityGroupCounts{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365GroupsActivityGroupCountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365GroupsActivityGroupCountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365GroupsActivityGroupCounts(), nil
}
// GetActive gets the active property value. The number of active groups. A group is considered active if any of the following occurred: group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityGroupCounts) GetActive()(*int64) {
    return m.active
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityGroupCounts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
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
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetReportDate gets the reportDate property value. The date on which a number of groups were active.
func (m *Office365GroupsActivityGroupCounts) GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportDate
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityGroupCounts) GetReportPeriod()(*string) {
    return m.reportPeriod
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityGroupCounts) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportRefreshDate
}
// GetTotal gets the total property value. The total number of groups.
func (m *Office365GroupsActivityGroupCounts) GetTotal()(*int64) {
    return m.total
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityGroupCounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("active", m.GetActive())
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
        err = writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActive sets the active property value. The number of active groups. A group is considered active if any of the following occurred: group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityGroupCounts) SetActive(value *int64)() {
    m.active = value
}
// SetReportDate sets the reportDate property value. The date on which a number of groups were active.
func (m *Office365GroupsActivityGroupCounts) SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityGroupCounts) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityGroupCounts) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportRefreshDate = value
}
// SetTotal sets the total property value. The total number of groups.
func (m *Office365GroupsActivityGroupCounts) SetTotal(value *int64)() {
    m.total = value
}
