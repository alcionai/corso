package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcExportJob 
type CloudPcExportJob struct {
    Entity
    // The date and time when the export job expires.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status of the export job. The possible values are: notStarted, inProgress, completed, unknownFutureValue. Read-only.
    exportJobStatus *CloudPcExportJobStatus
    // The storage account URL of the exported report. It can be used to download the file.
    exportUrl *string
    // The filter applied on the report.
    filter *string
    // The format of the exported report.
    format *string
    // The report name. The possible values are: remoteConnectionHistoricalReports, dailyAggregatedRemoteConnectionReports, totalAggregatedRemoteConnectionReports, unknownFutureValue.
    reportName *CloudPcReportName
    // The date and time when the export job was requested.
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The selected columns of the report.
    select_escaped []string
}
// NewCloudPcExportJob instantiates a new CloudPcExportJob and sets the default values.
func NewCloudPcExportJob()(*CloudPcExportJob) {
    m := &CloudPcExportJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcExportJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcExportJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcExportJob(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The date and time when the export job expires.
func (m *CloudPcExportJob) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetExportJobStatus gets the exportJobStatus property value. The status of the export job. The possible values are: notStarted, inProgress, completed, unknownFutureValue. Read-only.
func (m *CloudPcExportJob) GetExportJobStatus()(*CloudPcExportJobStatus) {
    return m.exportJobStatus
}
// GetExportUrl gets the exportUrl property value. The storage account URL of the exported report. It can be used to download the file.
func (m *CloudPcExportJob) GetExportUrl()(*string) {
    return m.exportUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcExportJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["exportJobStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcExportJobStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportJobStatus(val.(*CloudPcExportJobStatus))
        }
        return nil
    }
    res["exportUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportUrl(val)
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
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["reportName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcReportName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportName(val.(*CloudPcReportName))
        }
        return nil
    }
    res["requestDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDateTime(val)
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
    return res
}
// GetFilter gets the filter property value. The filter applied on the report.
func (m *CloudPcExportJob) GetFilter()(*string) {
    return m.filter
}
// GetFormat gets the format property value. The format of the exported report.
func (m *CloudPcExportJob) GetFormat()(*string) {
    return m.format
}
// GetReportName gets the reportName property value. The report name. The possible values are: remoteConnectionHistoricalReports, dailyAggregatedRemoteConnectionReports, totalAggregatedRemoteConnectionReports, unknownFutureValue.
func (m *CloudPcExportJob) GetReportName()(*CloudPcReportName) {
    return m.reportName
}
// GetRequestDateTime gets the requestDateTime property value. The date and time when the export job was requested.
func (m *CloudPcExportJob) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.requestDateTime
}
// GetSelect gets the select property value. The selected columns of the report.
func (m *CloudPcExportJob) GetSelect()([]string) {
    return m.select_escaped
}
// Serialize serializes information the current object
func (m *CloudPcExportJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetExportJobStatus() != nil {
        cast := (*m.GetExportJobStatus()).String()
        err = writer.WriteStringValue("exportJobStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("exportUrl", m.GetExportUrl())
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
        err = writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    if m.GetReportName() != nil {
        cast := (*m.GetReportName()).String()
        err = writer.WriteStringValue("reportName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
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
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. The date and time when the export job expires.
func (m *CloudPcExportJob) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetExportJobStatus sets the exportJobStatus property value. The status of the export job. The possible values are: notStarted, inProgress, completed, unknownFutureValue. Read-only.
func (m *CloudPcExportJob) SetExportJobStatus(value *CloudPcExportJobStatus)() {
    m.exportJobStatus = value
}
// SetExportUrl sets the exportUrl property value. The storage account URL of the exported report. It can be used to download the file.
func (m *CloudPcExportJob) SetExportUrl(value *string)() {
    m.exportUrl = value
}
// SetFilter sets the filter property value. The filter applied on the report.
func (m *CloudPcExportJob) SetFilter(value *string)() {
    m.filter = value
}
// SetFormat sets the format property value. The format of the exported report.
func (m *CloudPcExportJob) SetFormat(value *string)() {
    m.format = value
}
// SetReportName sets the reportName property value. The report name. The possible values are: remoteConnectionHistoricalReports, dailyAggregatedRemoteConnectionReports, totalAggregatedRemoteConnectionReports, unknownFutureValue.
func (m *CloudPcExportJob) SetReportName(value *CloudPcReportName)() {
    m.reportName = value
}
// SetRequestDateTime sets the requestDateTime property value. The date and time when the export job was requested.
func (m *CloudPcExportJob) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestDateTime = value
}
// SetSelect sets the select property value. The selected columns of the report.
func (m *CloudPcExportJob) SetSelect(value []string)() {
    m.select_escaped = value
}
