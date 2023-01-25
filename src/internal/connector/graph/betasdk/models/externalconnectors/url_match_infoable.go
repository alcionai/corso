package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UrlMatchInfoable 
type UrlMatchInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseUrls()([]string)
    GetOdataType()(*string)
    GetUrlPattern()(*string)
    SetBaseUrls(value []string)()
    SetOdataType(value *string)()
    SetUrlPattern(value *string)()
}
