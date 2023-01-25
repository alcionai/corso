package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUserRegistrationCountable 
type CredentialUserRegistrationCountable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalUserCount()(*int64)
    GetUserRegistrationCounts()([]UserRegistrationCountable)
    SetTotalUserCount(value *int64)()
    SetUserRegistrationCounts(value []UserRegistrationCountable)()
}
