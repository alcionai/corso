package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSAppleEventReceiverable 
type MacOSAppleEventReceiverable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowed()(*bool)
    GetCodeRequirement()(*string)
    GetIdentifier()(*string)
    GetIdentifierType()(*MacOSProcessIdentifierType)
    GetOdataType()(*string)
    SetAllowed(value *bool)()
    SetCodeRequirement(value *string)()
    SetIdentifier(value *string)()
    SetIdentifierType(value *MacOSProcessIdentifierType)()
    SetOdataType(value *string)()
}
