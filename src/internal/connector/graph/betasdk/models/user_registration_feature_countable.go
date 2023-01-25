package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationFeatureCountable 
type UserRegistrationFeatureCountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFeature()(*AuthenticationMethodFeature)
    GetOdataType()(*string)
    GetUserCount()(*int64)
    SetFeature(value *AuthenticationMethodFeature)()
    SetOdataType(value *string)()
    SetUserCount(value *int64)()
}
