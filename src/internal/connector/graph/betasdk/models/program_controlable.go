package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProgramControlable 
type ProgramControlable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetControlId()(*string)
    GetControlTypeId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetOwner()(UserIdentityable)
    GetProgram()(Programable)
    GetProgramId()(*string)
    GetResource()(ProgramResourceable)
    GetStatus()(*string)
    SetControlId(value *string)()
    SetControlTypeId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetOwner(value UserIdentityable)()
    SetProgram(value Programable)()
    SetProgramId(value *string)()
    SetResource(value ProgramResourceable)()
    SetStatus(value *string)()
}
