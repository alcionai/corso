package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationSignInSummaryable 
type ApplicationSignInSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetFailedSignInCount()(*int64)
    GetSuccessfulSignInCount()(*int64)
    GetSuccessPercentage()(*float64)
    SetAppDisplayName(value *string)()
    SetFailedSignInCount(value *int64)()
    SetSuccessfulSignInCount(value *int64)()
    SetSuccessPercentage(value *float64)()
}
