package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Pictureable 
type Pictureable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetContentType()(*string)
    GetHeight()(*int32)
    GetWidth()(*int32)
    SetContent(value []byte)()
    SetContentType(value *string)()
    SetHeight(value *int32)()
    SetWidth(value *int32)()
}
