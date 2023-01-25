package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCachedReportConfigurationable 
type DeviceManagementCachedReportConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFilter()(*string)
    GetLastRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMetadata()(*string)
    GetOrderBy()([]string)
    GetReportName()(*string)
    GetSelect()([]string)
    GetStatus()(*DeviceManagementReportStatus)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFilter(value *string)()
    SetLastRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMetadata(value *string)()
    SetOrderBy(value []string)()
    SetReportName(value *string)()
    SetSelect(value []string)()
    SetStatus(value *DeviceManagementReportStatus)()
}
