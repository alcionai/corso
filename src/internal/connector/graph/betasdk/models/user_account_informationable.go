package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAccountInformationable 
type UserAccountInformationable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgeGroup()(*string)
    GetCountryCode()(*string)
    GetPreferredLanguageTag()(LocaleInfoable)
    GetUserPrincipalName()(*string)
    SetAgeGroup(value *string)()
    SetCountryCode(value *string)()
    SetPreferredLanguageTag(value LocaleInfoable)()
    SetUserPrincipalName(value *string)()
}
