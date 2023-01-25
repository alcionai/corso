package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NdesConnectorable 
type NdesConnectorable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectorVersion()(*string)
    GetDisplayName()(*string)
    GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMachineName()(*string)
    GetRoleScopeTagIds()([]string)
    GetState()(*NdesConnectorState)
    SetConnectorVersion(value *string)()
    SetDisplayName(value *string)()
    SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMachineName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetState(value *NdesConnectorState)()
}
