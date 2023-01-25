package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityFileCounts 
type Office365GroupsActivityFileCounts struct {
    Entity
    // The number of files that were viewed, edited, shared, or synced in the group's SharePoint document library.
    active *int64
    // The date on which a number of files were active in the group's SharePoint site.
    reportDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of days the report covers.
    reportPeriod *string
    // The latest date of the content.
    reportRefreshDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The total number of files in the group's SharePoint document library.
    total *int64
}
// NewOffice365GroupsActivityFileCounts instantiates a new Office365GroupsActivityFileCounts and sets the default values.
func NewOffice365GroupsActivityFileCounts()(*Office365GroupsActivityFileCounts) {
    m := &Office365GroupsActivityFileCounts{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365GroupsActivityFileCountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365GroupsActivityFileCountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365GroupsActivityFileCounts(), nil
}
// GetActive gets the active property value. The number of files that were viewed, edited, shared, or synced in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) GetActive()(*int64) {
    return m.active
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityFileCounts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetReportDate gets the reportDate property value. The date on which a number of files were active in the group's SharePoint site.
func (m *Office365GroupsActivityFileCounts) GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportDate
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityFileCounts) GetReportPeriod()(*string) {
    return m.reportPeriod
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityFileCounts) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.reportRefreshDate
}
// GetTotal gets the total property value. The total number of files in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) GetTotal()(*int64) {
    return m.total
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityFileCounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetActive sets the active property value. The number of files that were viewed, edited, shared, or synced in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) SetActive(value *int64)() {
    m.active = value
}
// SetReportDate sets the reportDate property value. The date on which a number of files were active in the group's SharePoint site.
func (m *Office365GroupsActivityFileCounts) SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityFileCounts) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityFileCounts) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.reportRefreshDate = value
}
// SetTotal sets the total property value. The total number of files in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) SetTotal(value *int64)() {
    m.total = value
}
