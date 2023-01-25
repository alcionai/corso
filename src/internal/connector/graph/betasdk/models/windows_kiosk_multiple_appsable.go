package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskMultipleAppsable 
type WindowsKioskMultipleAppsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsKioskAppConfigurationable
    GetAllowAccessToDownloadsFolder()(*bool)
    GetApps()([]WindowsKioskAppBaseable)
    GetDisallowDesktopApps()(*bool)
    GetShowTaskBar()(*bool)
    GetStartMenuLayoutXml()([]byte)
    SetAllowAccessToDownloadsFolder(value *bool)()
    SetApps(value []WindowsKioskAppBaseable)()
    SetDisallowDesktopApps(value *bool)()
    SetShowTaskBar(value *bool)()
    SetStartMenuLayoutXml(value []byte)()
}
