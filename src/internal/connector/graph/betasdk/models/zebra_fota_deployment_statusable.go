package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaDeploymentStatusable 
type ZebraFotaDeploymentStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCancelRequested()(*bool)
    GetCompleteOrCanceledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetState()(*ZebraFotaDeploymentState)
    GetTotalAwaitingInstall()(*int32)
    GetTotalCanceled()(*int32)
    GetTotalCreated()(*int32)
    GetTotalDevices()(*int32)
    GetTotalDownloading()(*int32)
    GetTotalFailedDownload()(*int32)
    GetTotalFailedInstall()(*int32)
    GetTotalScheduled()(*int32)
    GetTotalSucceededInstall()(*int32)
    GetTotalUnknown()(*int32)
    SetCancelRequested(value *bool)()
    SetCompleteOrCanceledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetState(value *ZebraFotaDeploymentState)()
    SetTotalAwaitingInstall(value *int32)()
    SetTotalCanceled(value *int32)()
    SetTotalCreated(value *int32)()
    SetTotalDevices(value *int32)()
    SetTotalDownloading(value *int32)()
    SetTotalFailedDownload(value *int32)()
    SetTotalFailedInstall(value *int32)()
    SetTotalScheduled(value *int32)()
    SetTotalSucceededInstall(value *int32)()
    SetTotalUnknown(value *int32)()
}
