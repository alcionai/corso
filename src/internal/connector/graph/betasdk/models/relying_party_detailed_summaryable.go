package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RelyingPartyDetailedSummaryable 
type RelyingPartyDetailedSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedSignInCount()(*int64)
    GetMigrationStatus()(*MigrationStatus)
    GetMigrationValidationDetails()([]KeyValuePairable)
    GetRelyingPartyId()(*string)
    GetRelyingPartyName()(*string)
    GetReplyUrls()([]string)
    GetServiceId()(*string)
    GetSignInSuccessRate()(*float64)
    GetSuccessfulSignInCount()(*int64)
    GetTotalSignInCount()(*int64)
    GetUniqueUserCount()(*int64)
    SetFailedSignInCount(value *int64)()
    SetMigrationStatus(value *MigrationStatus)()
    SetMigrationValidationDetails(value []KeyValuePairable)()
    SetRelyingPartyId(value *string)()
    SetRelyingPartyName(value *string)()
    SetReplyUrls(value []string)()
    SetServiceId(value *string)()
    SetSignInSuccessRate(value *float64)()
    SetSuccessfulSignInCount(value *int64)()
    SetTotalSignInCount(value *int64)()
    SetUniqueUserCount(value *int64)()
}
