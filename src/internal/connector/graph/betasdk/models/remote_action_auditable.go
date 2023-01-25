package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteActionAuditable 
type RemoteActionAuditable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*RemoteAction)
    GetActionState()(*ActionState)
    GetDeviceDisplayName()(*string)
    GetDeviceIMEI()(*string)
    GetDeviceOwnerUserPrincipalName()(*string)
    GetInitiatedByUserPrincipalName()(*string)
    GetManagedDeviceId()(*string)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUserName()(*string)
    SetAction(value *RemoteAction)()
    SetActionState(value *ActionState)()
    SetDeviceDisplayName(value *string)()
    SetDeviceIMEI(value *string)()
    SetDeviceOwnerUserPrincipalName(value *string)()
    SetInitiatedByUserPrincipalName(value *string)()
    SetManagedDeviceId(value *string)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUserName(value *string)()
}
