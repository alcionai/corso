package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskSingleWin32Appable 
type WindowsKioskSingleWin32Appable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppConfigurationable
    GetWin32App()(WindowsKioskWin32Appable)
    SetWin32App(value WindowsKioskWin32Appable)()
}
