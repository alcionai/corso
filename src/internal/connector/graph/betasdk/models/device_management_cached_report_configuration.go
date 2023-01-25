package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCachedReportConfiguration entity representing the configuration of a cached report
type DeviceManagementCachedReportConfiguration struct {
    Entity
    // Time that the cached report expires
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Filters applied on report creation.
    filter *string
    // Time that the cached report was last refreshed
    lastRefreshDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Caller-managed metadata associated with the report
    metadata *string
    // Ordering of columns in the report
    orderBy []string
    // Name of the report
    reportName *string
    // Columns selected from the report
    select_escaped []string
    // Possible statuses associated with a generated report
    status *DeviceManagementReportStatus
}
// NewDeviceManagementCachedReportConfiguration instantiates a new deviceManagementCachedReportConfiguration and sets the default values.
func NewDeviceManagementCachedReportConfiguration()(*DeviceManagementCachedReportConfiguration) {
    m := &DeviceManagementCachedReportConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementCachedReportConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCachedReportConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCachedReportConfiguration(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. Time that the cached report expires
func (m *DeviceManagementCachedReportConfiguration) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCachedReportConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["lastRefreshDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshDateTime(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
        }
        return nil
    }
    res["orderBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrderBy(res)
        }
        return nil
    }
    res["reportName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportName(val)
        }
        return nil
    }
    res["select"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceManagementReportStatus))
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. Filters applied on report creation.
func (m *DeviceManagementCachedReportConfiguration) GetFilter()(*string) {
    return m.filter
}
// GetLastRefreshDateTime gets the lastRefreshDateTime property value. Time that the cached report was last refreshed
func (m *DeviceManagementCachedReportConfiguration) GetLastRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshDateTime
}
// GetMetadata gets the metadata property value. Caller-managed metadata associated with the report
func (m *DeviceManagementCachedReportConfiguration) GetMetadata()(*string) {
    return m.metadata
}
// GetOrderBy gets the orderBy property value. Ordering of columns in the report
func (m *DeviceManagementCachedReportConfiguration) GetOrderBy()([]string) {
    return m.orderBy
}
// GetReportName gets the reportName property value. Name of the report
func (m *DeviceManagementCachedReportConfiguration) GetReportName()(*string) {
    return m.reportName
}
// GetSelect gets the select property value. Columns selected from the report
func (m *DeviceManagementCachedReportConfiguration) GetSelect()([]string) {
    return m.select_escaped
}
// GetStatus gets the status property value. Possible statuses associated with a generated report
func (m *DeviceManagementCachedReportConfiguration) GetStatus()(*DeviceManagementReportStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DeviceManagementCachedReportConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshDateTime", m.GetLastRefreshDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    if m.GetOrderBy() != nil {
        err = writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportName", m.GetReportName())
        if err != nil {
            return err
        }
    }
    if m.GetSelect() != nil {
        err = writer.WriteCollectionOfStringValues("select", m.GetSelect())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. Time that the cached report expires
func (m *DeviceManagementCachedReportConfiguration) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetFilter sets the filter property value. Filters applied on report creation.
func (m *DeviceManagementCachedReportConfiguration) SetFilter(value *string)() {
    m.filter = value
}
// SetLastRefreshDateTime sets the lastRefreshDateTime property value. Time that the cached report was last refreshed
func (m *DeviceManagementCachedReportConfiguration) SetLastRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshDateTime = value
}
// SetMetadata sets the metadata property value. Caller-managed metadata associated with the report
func (m *DeviceManagementCachedReportConfiguration) SetMetadata(value *string)() {
    m.metadata = value
}
// SetOrderBy sets the orderBy property value. Ordering of columns in the report
func (m *DeviceManagementCachedReportConfiguration) SetOrderBy(value []string)() {
    m.orderBy = value
}
// SetReportName sets the reportName property value. Name of the report
func (m *DeviceManagementCachedReportConfiguration) SetReportName(value *string)() {
    m.reportName = value
}
// SetSelect sets the select property value. Columns selected from the report
func (m *DeviceManagementCachedReportConfiguration) SetSelect(value []string)() {
    m.select_escaped = value
}
// SetStatus sets the status property value. Possible statuses associated with a generated report
func (m *DeviceManagementCachedReportConfiguration) SetStatus(value *DeviceManagementReportStatus)() {
    m.status = value
}
