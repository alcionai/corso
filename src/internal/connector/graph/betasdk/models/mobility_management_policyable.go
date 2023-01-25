package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobilityManagementPolicyable 
type MobilityManagementPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppliesTo()(*PolicyScope)
    GetComplianceUrl()(*string)
    GetDescription()(*string)
    GetDiscoveryUrl()(*string)
    GetDisplayName()(*string)
    GetIncludedGroups()([]Groupable)
    GetIsValid()(*bool)
    GetTermsOfUseUrl()(*string)
    SetAppliesTo(value *PolicyScope)()
    SetComplianceUrl(value *string)()
    SetDescription(value *string)()
    SetDiscoveryUrl(value *string)()
    SetDisplayName(value *string)()
    SetIncludedGroups(value []Groupable)()
    SetIsValid(value *bool)()
    SetTermsOfUseUrl(value *string)()
}
