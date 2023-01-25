package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Settingsable 
type Settingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHasGraphMailbox()(*bool)
    GetHasLicense()(*bool)
    GetHasOptedOut()(*bool)
    GetOdataType()(*string)
    SetHasGraphMailbox(value *bool)()
    SetHasLicense(value *bool)()
    SetHasOptedOut(value *bool)()
    SetOdataType(value *string)()
}
