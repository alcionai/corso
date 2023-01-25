package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InferenceDataable 
type InferenceDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfidenceScore()(*float64)
    GetOdataType()(*string)
    GetUserHasVerifiedAccuracy()(*bool)
    SetConfidenceScore(value *float64)()
    SetOdataType(value *string)()
    SetUserHasVerifiedAccuracy(value *bool)()
}
