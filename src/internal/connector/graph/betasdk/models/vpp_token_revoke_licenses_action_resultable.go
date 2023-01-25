package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppTokenRevokeLicensesActionResultable 
type VppTokenRevokeLicensesActionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VppTokenActionResultable
    GetActionFailureReason()(*VppTokenActionFailureReason)
    GetFailedLicensesCount()(*int32)
    GetTotalLicensesCount()(*int32)
    SetActionFailureReason(value *VppTokenActionFailureReason)()
    SetFailedLicensesCount(value *int32)()
    SetTotalLicensesCount(value *int32)()
}
