package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleAssignmentable 
type PrivilegedRoleAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsElevated()(*bool)
    GetResultMessage()(*string)
    GetRoleId()(*string)
    GetRoleInfo()(PrivilegedRoleable)
    GetUserId()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsElevated(value *bool)()
    SetResultMessage(value *string)()
    SetRoleId(value *string)()
    SetRoleInfo(value PrivilegedRoleable)()
    SetUserId(value *string)()
}
