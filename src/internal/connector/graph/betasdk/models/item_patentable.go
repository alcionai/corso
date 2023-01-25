package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPatentable 
type ItemPatentable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsPending()(*bool)
    GetIssuedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetIssuingAuthority()(*string)
    GetNumber()(*string)
    GetWebUrl()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsPending(value *bool)()
    SetIssuedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetIssuingAuthority(value *string)()
    SetNumber(value *string)()
    SetWebUrl(value *string)()
}
