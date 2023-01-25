package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddHeaderable 
type AddHeaderable interface {
    MarkContentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlignment()(*Alignment)
    GetMargin()(*int32)
    SetAlignment(value *Alignment)()
    SetMargin(value *int32)()
}
