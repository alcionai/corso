package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessUsersable 
type ConditionalAccessUsersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeGroups()([]string)
    GetExcludeGuestsOrExternalUsers()(ConditionalAccessGuestsOrExternalUsersable)
    GetExcludeRoles()([]string)
    GetExcludeUsers()([]string)
    GetIncludeGroups()([]string)
    GetIncludeGuestsOrExternalUsers()(ConditionalAccessGuestsOrExternalUsersable)
    GetIncludeRoles()([]string)
    GetIncludeUsers()([]string)
    GetOdataType()(*string)
    SetExcludeGroups(value []string)()
    SetExcludeGuestsOrExternalUsers(value ConditionalAccessGuestsOrExternalUsersable)()
    SetExcludeRoles(value []string)()
    SetExcludeUsers(value []string)()
    SetIncludeGroups(value []string)()
    SetIncludeGuestsOrExternalUsers(value ConditionalAccessGuestsOrExternalUsersable)()
    SetIncludeRoles(value []string)()
    SetIncludeUsers(value []string)()
    SetOdataType(value *string)()
}
