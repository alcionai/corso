package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppAppAssignedLicenseable 
type MacOsVppAppAssignedLicenseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUserEmailAddress()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetUserEmailAddress(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
