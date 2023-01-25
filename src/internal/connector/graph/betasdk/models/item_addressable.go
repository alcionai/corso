package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemAddressable 
type ItemAddressable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetail()(PhysicalAddressable)
    GetDisplayName()(*string)
    GetGeoCoordinates()(GeoCoordinatesable)
    SetDetail(value PhysicalAddressable)()
    SetDisplayName(value *string)()
    SetGeoCoordinates(value GeoCoordinatesable)()
}
