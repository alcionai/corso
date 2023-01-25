package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Notificationable 
type Notificationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayTimeToLive()(*int32)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupName()(*string)
    GetPayload()(PayloadTypesable)
    GetPriority()(*Priority)
    GetTargetHostName()(*string)
    GetTargetPolicy()(TargetPolicyEndpointsable)
    SetDisplayTimeToLive(value *int32)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupName(value *string)()
    SetPayload(value PayloadTypesable)()
    SetPriority(value *Priority)()
    SetTargetHostName(value *string)()
    SetTargetPolicy(value TargetPolicyEndpointsable)()
}
