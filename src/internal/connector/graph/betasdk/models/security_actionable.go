package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityActionable 
type SecurityActionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionReason()(*string)
    GetAppId()(*string)
    GetAzureTenantId()(*string)
    GetClientContext()(*string)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetErrorInfo()(ResultInfoable)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetParameters()([]KeyValuePairable)
    GetStates()([]SecurityActionStateable)
    GetStatus()(*OperationStatus)
    GetUser()(*string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetActionReason(value *string)()
    SetAppId(value *string)()
    SetAzureTenantId(value *string)()
    SetClientContext(value *string)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetErrorInfo(value ResultInfoable)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetParameters(value []KeyValuePairable)()
    SetStates(value []SecurityActionStateable)()
    SetStatus(value *OperationStatus)()
    SetUser(value *string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
