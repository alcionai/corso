package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineStateSummaryable 
type SecurityBaselineStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflictCount()(*int32)
    GetErrorCount()(*int32)
    GetNotApplicableCount()(*int32)
    GetNotSecureCount()(*int32)
    GetSecureCount()(*int32)
    GetUnknownCount()(*int32)
    SetConflictCount(value *int32)()
    SetErrorCount(value *int32)()
    SetNotApplicableCount(value *int32)()
    SetNotSecureCount(value *int32)()
    SetSecureCount(value *int32)()
    SetUnknownCount(value *int32)()
}
