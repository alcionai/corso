package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintJobable 
type PrintJobable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConfiguration()(PrintJobConfigurationable)
    GetCreatedBy()(UserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetDocuments()([]PrintDocumentable)
    GetErrorCode()(*int32)
    GetIsFetchable()(*bool)
    GetRedirectedFrom()(*string)
    GetRedirectedTo()(*string)
    GetStatus()(PrintJobStatusable)
    GetTasks()([]PrintTaskable)
    SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConfiguration(value PrintJobConfigurationable)()
    SetCreatedBy(value UserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetDocuments(value []PrintDocumentable)()
    SetErrorCode(value *int32)()
    SetIsFetchable(value *bool)()
    SetRedirectedFrom(value *string)()
    SetRedirectedTo(value *string)()
    SetStatus(value PrintJobStatusable)()
    SetTasks(value []PrintTaskable)()
}
