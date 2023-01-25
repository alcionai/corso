package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonCertificationable 
type PersonCertificationable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificationId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetIssuedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetIssuingAuthority()(*string)
    GetIssuingCompany()(*string)
    GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetThumbnailUrl()(*string)
    GetWebUrl()(*string)
    SetCertificationId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetIssuedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetIssuingAuthority(value *string)()
    SetIssuingCompany(value *string)()
    SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetThumbnailUrl(value *string)()
    SetWebUrl(value *string)()
}
