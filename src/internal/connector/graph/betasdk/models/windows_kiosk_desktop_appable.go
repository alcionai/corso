package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskDesktopAppable 
type WindowsKioskDesktopAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppBaseable
    GetDesktopApplicationId()(*string)
    GetDesktopApplicationLinkPath()(*string)
    GetPath()(*string)
    SetDesktopApplicationId(value *string)()
    SetDesktopApplicationLinkPath(value *string)()
    SetPath(value *string)()
}
