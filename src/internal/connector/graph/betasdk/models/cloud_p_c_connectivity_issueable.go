package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPCConnectivityIssueable 
type CloudPCConnectivityIssueable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetErrorCode()(*string)
    GetErrorDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetErrorDescription()(*string)
    GetRecommendedAction()(*string)
    GetUserId()(*string)
    SetDeviceId(value *string)()
    SetErrorCode(value *string)()
    SetErrorDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetErrorDescription(value *string)()
    SetRecommendedAction(value *string)()
    SetUserId(value *string)()
}
