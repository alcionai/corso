package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppProductCodeDetectionable 
type Win32LobAppProductCodeDetectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppDetectionable
    GetProductCode()(*string)
    GetProductVersion()(*string)
    GetProductVersionOperator()(*Win32LobAppDetectionOperator)
    SetProductCode(value *string)()
    SetProductVersion(value *string)()
    SetProductVersionOperator(value *Win32LobAppDetectionOperator)()
}
