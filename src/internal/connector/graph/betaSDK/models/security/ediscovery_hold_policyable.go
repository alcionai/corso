package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryHoldPolicyable 
type EdiscoveryHoldPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
    GetContentQuery()(*string)
    GetErrors()([]string)
    GetIsEnabled()(*bool)
    GetSiteSources()([]SiteSourceable)
    GetUserSources()([]UserSourceable)
    SetContentQuery(value *string)()
    SetErrors(value []string)()
    SetIsEnabled(value *bool)()
    SetSiteSources(value []SiteSourceable)()
    SetUserSources(value []UserSourceable)()
}
