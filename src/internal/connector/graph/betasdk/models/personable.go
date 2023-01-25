package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Personable 
type Personable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBirthday()(*string)
    GetCompanyName()(*string)
    GetDepartment()(*string)
    GetDisplayName()(*string)
    GetEmailAddresses()([]RankedEmailAddressable)
    GetGivenName()(*string)
    GetIsFavorite()(*bool)
    GetMailboxType()(*string)
    GetOfficeLocation()(*string)
    GetPersonNotes()(*string)
    GetPersonType()(*string)
    GetPhones()([]Phoneable)
    GetPostalAddresses()([]Locationable)
    GetProfession()(*string)
    GetSources()([]PersonDataSourceable)
    GetSurname()(*string)
    GetTitle()(*string)
    GetUserPrincipalName()(*string)
    GetWebsites()([]Websiteable)
    GetYomiCompany()(*string)
    SetBirthday(value *string)()
    SetCompanyName(value *string)()
    SetDepartment(value *string)()
    SetDisplayName(value *string)()
    SetEmailAddresses(value []RankedEmailAddressable)()
    SetGivenName(value *string)()
    SetIsFavorite(value *bool)()
    SetMailboxType(value *string)()
    SetOfficeLocation(value *string)()
    SetPersonNotes(value *string)()
    SetPersonType(value *string)()
    SetPhones(value []Phoneable)()
    SetPostalAddresses(value []Locationable)()
    SetProfession(value *string)()
    SetSources(value []PersonDataSourceable)()
    SetSurname(value *string)()
    SetTitle(value *string)()
    SetUserPrincipalName(value *string)()
    SetWebsites(value []Websiteable)()
    SetYomiCompany(value *string)()
}
