package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcRemoteActionResultable 
type CloudPcRemoteActionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionName()(*string)
    GetActionState()(*ActionState)
    GetCloudPcId()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceId()(*string)
    GetOdataType()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatusDetails()(CloudPcStatusDetailsable)
    SetActionName(value *string)()
    SetActionState(value *ActionState)()
    SetCloudPcId(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceId(value *string)()
    SetOdataType(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatusDetails(value CloudPcStatusDetailsable)()
}
