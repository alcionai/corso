package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintJob provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PrintJob struct {
    Entity
    // The acknowledgedDateTime property
    acknowledgedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The completedDateTime property
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A group of settings that a printer should use to print a job.
    configuration PrintJobConfigurationable
    // The createdBy property
    createdBy UserIdentityable
    // The DateTimeOffset when the job was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The displayName property
    displayName *string
    // The documents property
    documents []PrintDocumentable
    // The errorCode property
    errorCode *int32
    // If true, document can be fetched by printer.
    isFetchable *bool
    // Contains the source job URL, if the job has been redirected from another printer.
    redirectedFrom *string
    // Contains the destination job URL, if the job has been redirected to another printer.
    redirectedTo *string
    // The status of the print job. Read-only.
    status PrintJobStatusable
    // A list of printTasks that were triggered by this print job.
    tasks []PrintTaskable
}
// NewPrintJob instantiates a new printJob and sets the default values.
func NewPrintJob()(*PrintJob) {
    m := &PrintJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintJob(), nil
}
// GetAcknowledgedDateTime gets the acknowledgedDateTime property value. The acknowledgedDateTime property
func (m *PrintJob) GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.acknowledgedDateTime
}
// GetCompletedDateTime gets the completedDateTime property value. The completedDateTime property
func (m *PrintJob) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetConfiguration gets the configuration property value. A group of settings that a printer should use to print a job.
func (m *PrintJob) GetConfiguration()(PrintJobConfigurationable) {
    return m.configuration
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *PrintJob) GetCreatedBy()(UserIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
func (m *PrintJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PrintJob) GetDisplayName()(*string) {
    return m.displayName
}
// GetDocuments gets the documents property value. The documents property
func (m *PrintJob) GetDocuments()([]PrintDocumentable) {
    return m.documents
}
// GetErrorCode gets the errorCode property value. The errorCode property
func (m *PrintJob) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acknowledgedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcknowledgedDateTime(val)
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintJobConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(PrintJobConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["documents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintDocumentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintDocumentable, len(val))
            for i, v := range val {
                res[i] = v.(PrintDocumentable)
            }
            m.SetDocuments(res)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["isFetchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFetchable(val)
        }
        return nil
    }
    res["redirectedFrom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectedFrom(val)
        }
        return nil
    }
    res["redirectedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectedTo(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintJobStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(PrintJobStatusable))
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskable, len(val))
            for i, v := range val {
                res[i] = v.(PrintTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetIsFetchable gets the isFetchable property value. If true, document can be fetched by printer.
func (m *PrintJob) GetIsFetchable()(*bool) {
    return m.isFetchable
}
// GetRedirectedFrom gets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
func (m *PrintJob) GetRedirectedFrom()(*string) {
    return m.redirectedFrom
}
// GetRedirectedTo gets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
func (m *PrintJob) GetRedirectedTo()(*string) {
    return m.redirectedTo
}
// GetStatus gets the status property value. The status of the print job. Read-only.
func (m *PrintJob) GetStatus()(PrintJobStatusable) {
    return m.status
}
// GetTasks gets the tasks property value. A list of printTasks that were triggered by this print job.
func (m *PrintJob) GetTasks()([]PrintTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *PrintJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acknowledgedDateTime", m.GetAcknowledgedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDocuments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDocuments()))
        for i, v := range m.GetDocuments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("documents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFetchable", m.GetIsFetchable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectedFrom", m.GetRedirectedFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectedTo", m.GetRedirectedTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcknowledgedDateTime sets the acknowledgedDateTime property value. The acknowledgedDateTime property
func (m *PrintJob) SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acknowledgedDateTime = value
}
// SetCompletedDateTime sets the completedDateTime property value. The completedDateTime property
func (m *PrintJob) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetConfiguration sets the configuration property value. A group of settings that a printer should use to print a job.
func (m *PrintJob) SetConfiguration(value PrintJobConfigurationable)() {
    m.configuration = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *PrintJob) SetCreatedBy(value UserIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
func (m *PrintJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PrintJob) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDocuments sets the documents property value. The documents property
func (m *PrintJob) SetDocuments(value []PrintDocumentable)() {
    m.documents = value
}
// SetErrorCode sets the errorCode property value. The errorCode property
func (m *PrintJob) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetIsFetchable sets the isFetchable property value. If true, document can be fetched by printer.
func (m *PrintJob) SetIsFetchable(value *bool)() {
    m.isFetchable = value
}
// SetRedirectedFrom sets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
func (m *PrintJob) SetRedirectedFrom(value *string)() {
    m.redirectedFrom = value
}
// SetRedirectedTo sets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
func (m *PrintJob) SetRedirectedTo(value *string)() {
    m.redirectedTo = value
}
// SetStatus sets the status property value. The status of the print job. Read-only.
func (m *PrintJob) SetStatus(value PrintJobStatusable)() {
    m.status = value
}
// SetTasks sets the tasks property value. A list of printTasks that were triggered by this print job.
func (m *PrintJob) SetTasks(value []PrintTaskable)() {
    m.tasks = value
}
