package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUsageSummaryable 
type CredentialUsageSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthMethod()(*UsageAuthMethod)
    GetFailureActivityCount()(*int64)
    GetFeature()(*FeatureType)
    GetSuccessfulActivityCount()(*int64)
    SetAuthMethod(value *UsageAuthMethod)()
    SetFailureActivityCount(value *int64)()
    SetFeature(value *FeatureType)()
    SetSuccessfulActivityCount(value *int64)()
}
